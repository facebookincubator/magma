/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <cassert>
#include <iostream>
#include <string>

#include <netinet/in.h>
#include <netinet/ip.h>
#include <net/if.h>
#include <net/ethernet.h>
#include <arpa/inet.h>

#include <tins/tins.h>
#include <tins/ip.h>

#include "PDUGenerator.h"

namespace {  // anonymous

/* TODO remove after testing */
void print_bytes(void* ptr, int size) {
  unsigned char* p = (unsigned char*) ptr;
  int i;
  for (i = 0; i < size; i++) {
    printf("%02hhX ", p[i]);
  }
  printf("\n");
}

void set_xid(struct pdu_info* pdu, std::string value) {
  memcpy(pdu->xid, value.c_str(), XID_LENGTH);
}

bool extract_ip_addr(
    const u_char* packet, std::string& src_ip, std::string& dst_ip) {
  const struct ether_header* ethernetHeader;
  const struct ip* ipHeader;
  char sourceIP[INET_ADDRSTRLEN];
  char destIP[INET_ADDRSTRLEN];

  ethernetHeader = (struct ether_header*) packet;
  if (ntohs(ethernetHeader->ether_type) == ETHERTYPE_IP) {
    ipHeader = (struct ip*) (packet + sizeof(struct ether_header));
    src_ip = inet_ntop(AF_INET, &(ipHeader->ip_src), sourceIP, INET_ADDRSTRLEN);
    dst_ip = inet_ntop(AF_INET, &(ipHeader->ip_dst), destIP, INET_ADDRSTRLEN);
    return true;
  }
  return false;
}

}  // namespace

namespace magma {

PDUGenerator::PDUGenerator(
    std::unique_ptr<ProxyConnector> proxy_connector,
    std::unique_ptr<DirectorydClient> directoryd_client,
    const std::string& pkt_dst_mac, const std::string& pkt_src_mac)
    : pkt_dst_mac_(pkt_dst_mac),
      pkt_src_mac_(pkt_src_mac),
      directoryd_client_(std::move(directoryd_client)),
      proxy_connector_(std::move(proxy_connector)) {}

void PDUGenerator::set_conditional_attr(
    const struct pcap_pkthdr* phdr, struct conditional_attributes* attributes) {
  attributes->timestamp = phdr->ts.tv_sec;
}

void* PDUGenerator::generate_pkt(
    const struct pcap_pkthdr* phdr, const u_char* pdata) {
  uint8_t* data        = static_cast<uint8_t*>(calloc(
      1, sizeof(struct pdu_info) + sizeof(struct conditional_attributes) +
             phdr->len));
  struct pdu_info* pdu = (struct pdu_info*) data;

  pdu->version  = PDU_VERSION;
  pdu->pdu_type = PDU_TYPE;
  pdu->header_length =
      sizeof(struct pdu_info) + sizeof(struct conditional_attributes);
  pdu->payload_length = phdr->len;
  pdu->payload_format = IP_PAYLOAD_FORMAT;

  set_conditional_attr(
      phdr, (struct conditional_attributes*) (data + sizeof(struct pdu_info)));

  memcpy(data + pdu->header_length, pdata, phdr->len);

  MLOG(MDEBUG) << "Generated packet length with length - "
               << pdu->header_length + pdu->payload_length;
  print_bytes(data, pdu->header_length + pdu->payload_length);

  return (void*) data;
}

bool PDUGenerator::send_packet(
    const struct pcap_pkthdr* phdr, const u_char* pdata) {
  std::string src_ip;
  std::string dst_ip;
  if (!extract_ip_addr(pdata, src_ip, dst_ip)) {
    MLOG(MERROR) << "Could not extract IP from the packet, skipping";
    return true;
  }

  MLOG(MDEBUG) << "Processing packet with src ip " << src_ip << " dst ip "
               << dst_ip;
  void* data           = generate_pkt(phdr, pdata);
  struct pdu_info* pdu = (struct pdu_info*) data;

  directoryd_client_->get_directoryd_xid_field(
      dst_ip, std::bind(
                  &PDUGenerator::handle_ip_lookup_callback, this, src_ip, data,
                  pdu, std::placeholders::_1, std::placeholders::_2));

  directoryd_client_->get_directoryd_xid_field(
      dst_ip, std::bind(
                  &PDUGenerator::handle_ip_lookup_callback, this, dst_ip, data,
                  pdu, std::placeholders::_1, std::placeholders::_2));
  return true;
}

void PDUGenerator::handle_ip_lookup_callback(
    std::string ip_addr, void* data, struct pdu_info* pdu, Status status,
    DirectoryField resp) {
  // TODO process the output of the directoryd lookup
  pdu->payload_direction = DIRECTION_TO_TARGET;
  set_xid(pdu, "tracking_123");
  proxy_connector_->send_data(data, pdu->header_length + pdu->payload_length);
  // free(data);

  if (!status.ok()) {
    MLOG(MDEBUG) << "Could not fetch subscriber with ip - " << ip_addr;
    return;
  }
  MLOG(MDEBUG) << "Got reply " << resp.value().c_str() << "for -" << ip_addr;

  // TODO create a cache that stores the IPs that were looked up successfully,
  // the amount of LI activated UEs is small and lookup should only occur on
  // first packet
}

}  // namespace magma