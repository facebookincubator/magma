/*
   Copyright 2020 The Magma Authors.
   This source code is licensed under the BSD-style license found in the
   LICENSE file in the root directory of this source tree.
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
 */

#pragma once
#include <sstream>
#include "M5gNasMessage.h"
#include "M5GExtendedProtocolDiscriminator.h"
#include "M5GPDUSessionIdentity.h"
#include "M5GPTI.h"
#include "M5GMessageType.h"
#include "M5GIntegrityProtMaxDataRate.h"
#include "M5GPDUSessionType.h"
#include "M5GSSCMode.h"

using namespace std;
namespace magma5g {
// PDUSessionEstablishmentRequest Message Class
class PDUSessionEstablishmentRequestMsg {
 public:
#define PDU_SESSION_ESTABLISH_REQ_MIN_LEN 4

#define REQUEST_5GSM_CAPABILITY_TYPE 0x28
#define REQUEST_MAXIMUM_NUMBER_OF_SUPPORTED_PACKET_FILTERS_TYPE 0x55
#define REQUEST_ALWAYS_ON_PDU_SESSION_REQUESTED_TYPE 0xB0
#define REQUEST_SM_PDU_DN_REQUEST_CONTAINER_TYPE 0x39
#define REQUEST_EXTENDED_PROTOCOL_CONFIGURATION_OPTIONS_TYPE 0x7B
#define REQUEST_HEADER_COMPRESSION_CONFIGURATION_TYPE 0x66
#define REQUEST_DS_TT_ETHERNET_PORT_MAC_ADDRESS_TYPE 0x6E
#define REQUEST_UE_DS_TT_RESIDENCE_TIME_TYPE 0x6F
#define REQUEST_PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE 0x7C

  ExtendedProtocolDiscriminatorMsg extended_protocol_discriminator;
  PDUSessionIdentityMsg pdu_session_identity;
  PTIMsg pti;
  MessageTypeMsg message_type;
  IntegrityProtMaxDataRateMsg integrity_prot_max_data_rate;
  PDUSessionTypeMsg pdu_session_type;
  SSCModeMsg ssc_mode;

  PDUSessionEstablishmentRequestMsg();
  ~PDUSessionEstablishmentRequestMsg();
  int DecodePDUSessionEstablishmentRequestMsg(
      PDUSessionEstablishmentRequestMsg* pdu_session_estab_request,
      uint8_t* buffer, uint32_t len);
  int EncodePDUSessionEstablishmentRequestMsg(
      PDUSessionEstablishmentRequestMsg* pdu_session_estab_request,
      uint8_t* buffer, uint32_t len);
};
}  // namespace magma5g
/******************************************************************************
  TS-24.501 Table 8.3.1.1.1: PDU SESSION ESTABLISHMENT Request message content
-------------------------------------------------------------------------------
|IEI|   Information Element  |    Type/Reference      |Presence|Format|Length |
|---|------------------------|------------------------|--------|------|-------|
|   |Extended protocol descr-|Extended Protocol descr-|    M   |  V   |  1    |
|   |-iminator               |-iminator 9.2           |        |      |       |
|---|------------------------|------------------------|--------|------|-------|
|   |PDU session ID          |PDU session ID 9.4      |    M   |  V   |  1    |
|---|------------------------|------------------------|--------|------|-------|
|   |PTI                     |Procedure transacti     |    M   |  V   |  1    |
|   |                        |identity 9.6            |        |      |       |
|---|------------------------|------------------------|--------|------|-------|
|   |PDU SESSION ESTABLISHME-|Message type 9.7        |    M   |  V   |  1    |
|   |-NT Request message ide-|                        |        |      |       |
|   |-ntity                  |                        |        |      |       |
|---|------------------------|------------------------|--------|------|-------|
|   |Integrity protection    |Integrity protection    |    M   |  V   |  2    |
|   |maximum data rate       |maximum data rate       |        |      |       |
|   |                        |9.11.4.7                |        |      |       |
|---|------------------------|------------------------|--------|------|-------|
|9- |PDU session type        |PDU session type        |    O   |  TV  |  1    |
|   |                        |9.11.4.14               |        |      |       |
|---|------------------------|------------------------|--------|------|-------|
|A- |SSC mode                |SSC mode 9.11.4,16      |    O   |  V   |  1/2  |
|---|------------------------|------------------------|--------|------|-------|
******************************************************************************/
