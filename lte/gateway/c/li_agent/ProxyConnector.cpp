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
#include <netinet/in.h>
#include <arpa/inet.h>
#include <iostream>

#include <libmnl/libmnl.h>
#include <linux/netfilter/nfnetlink.h>
#include <linux/netfilter/nfnetlink_conntrack.h>

#include <linux/if_packet.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <net/if.h>
#include <netinet/ether.h>
#include <linux/ip.h>
#include <memory>
#include <pcap.h>

#include <openssl/ssl.h>
#include <openssl/err.h>

#include "ProxyConnector.h"

#include "magma_logging.h"

namespace magma {
namespace lte {

ProxyConnector::ProxyConnector(
    const std::string& proxy_addr, const int proxy_port,
    const std::string& cert_file, const std::string& key_file)
    : proxy_addr_(proxy_addr),
      proxy_port_(proxy_port),
      cert_file_(cert_file),
      key_file_(key_file) {
  ssl_ = GetSSLSocket();
}

void ProxyConnector::LoadCertificates(SSL_CTX* ctx) {
  if (SSL_CTX_use_certificate_file(ctx, cert_file_.c_str(), SSL_FILETYPE_PEM) <=
      0) {
    ERR_print_errors_fp(stderr);
    abort();
  }
  if (SSL_CTX_use_PrivateKey_file(ctx, key_file_.c_str(), SSL_FILETYPE_PEM) <=
      0) {
    ERR_print_errors_fp(stderr);
    abort();
  }
  if (!SSL_CTX_check_private_key(ctx)) {
    MLOG(MERROR) << "Private key does not match the public certificate";
    abort();
  }
}

SSL_CTX* ProxyConnector::InitCTX(void) {
  SSL_CTX* ctx;

  OpenSSL_add_all_algorithms(); /* Load cryptos, et.al. */
  SSL_load_error_strings();     /* Bring in and register error messages */
  ctx = SSL_CTX_new(TLS_client_method()); /* Create new context */
  if (ctx == NULL) {
    ERR_print_errors_fp(stderr);
    abort();
  }
  return ctx;
}

int ProxyConnector::OpenConnection() {
  int sd;
  struct sockaddr_in serv_addr;

  sd                        = socket(AF_INET, SOCK_STREAM, 0);
  serv_addr.sin_family      = AF_INET;
  serv_addr.sin_addr.s_addr = INADDR_ANY;
  serv_addr.sin_port        = htons(proxy_port_);

  // TODO change to proxy addr
  if (inet_pton(AF_INET, "127.0.0.1", &serv_addr.sin_addr) <= 0) {
    MLOG(MERROR) << "Invalid address/ Address not supported";
    return -1;
  }
  if (connect(sd, (struct sockaddr*) &serv_addr, sizeof(struct sockaddr_in)) !=
      0) {
    MLOG(MERROR) << "Can't connect to the proxy, exiting";
    close(sd);
    abort();
  }
  return sd;
}

SSL* ProxyConnector::GetSSLSocket() {
  SSL* ssl;
  SSL_library_init();

  ctx_ = InitCTX();
  LoadCertificates(ctx_);
  proxy_ = OpenConnection();
  ssl    = SSL_new(ctx_);
  SSL_set_fd(ssl, proxy_);
  if (SSL_connect(ssl) == -1) {
    ERR_print_errors_fp(stderr);
    return NULL;
  }
  return ssl;
}

int ProxyConnector::SendData(void* data, uint32_t size) {
  // TODO we probably want to deal with write edge cases here
  SSL_write(ssl_, data, size);

  return 0;
}

void ProxyConnector::cleanup() {
  SSL_free(ssl_);
  close(proxy_);
  SSL_CTX_free(ctx_);
}

}  // namespace lte
}  // namespace magma
