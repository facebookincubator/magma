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
/****************************************************************************
  Source      ngap_messages_types.h
  Version     0.1
  Date        2020/07/28
  Product     NGAP stack
  Subsystem   Access and Mobility Management Function
  Author      Ashish Prajapati
  Description Defines NG Application Protocol Messages

*****************************************************************************/

#ifndef FILE_NGAP_MESSAGES_TYPES_SEEN
#define FILE_NGAP_MESSAGES_TYPES_SEEN

#include "3gpp_24.008.h"
#include "3gpp_38.401.h"
#include "3gpp_38.413.h"
#include "3gpp_38.331.h"
#include "3gpp_23.003.h"
#include "TrackingAreaIdentity.h"
#include "Ngap_Cause.h"

typedef uint16_t sctp_stream_id_t;
typedef uint32_t sctp_assoc_id_t;
typedef uint64_t gnb_ngap_id_key_t;
typedef uint64_t bitrate_t;

typedef char* APN_t;
typedef uint8_t APNRestriction_t;
typedef uint8_t DelayValue_t;
typedef uint8_t priority_level_t;
#define PRIORITY_LEVEL_FMT "0x%" PRIu8
#define PRIORITY_LEVEL_SCAN_FMT SCNu8
typedef uint32_t SequenceNumber_t;
typedef uint32_t access_restriction_t;
typedef uint32_t context_identifier_t;
typedef uint32_t rau_tau_timer_t;

typedef uint32_t ard_t;
typedef int pdn_cid_t;  // pdn connexion identity, related to esm protocol,
                        // sometimes type is mixed with int return code!!...
typedef uint8_t
    proc_tid_t;  // procedure transaction identity, related to esm protocol
typedef uint8_t qci_t;
typedef uint32_t teid_t;

#define NGAP_GNB_DEREGISTERED_IND(mSGpTR)                                      \
  (mSGpTR)->ittiMsg.ngap_gNB_deregistered_ind
#define NGAP_GNB_INITIATED_RESET_REQ(mSGpTR)                                   \
  (mSGpTR)->ittiMsg.ngap_gnb_initiated_reset_req
#define NGAP_GNB_INITIATED_RESET_ACK(mSGpTR)                                   \
  (mSGpTR)->ittiMsg.ngap_gnb_initiated_reset_ack
#define NGAP_UE_CONTEXT_RELEASE_REQ(mSGpTR)                                    \
  (mSGpTR)->ittiMsg.ngap_ue_context_release_req
#define NGAP_UE_CONTEXT_RELEASE_COMMAND(mSGpTR)                                \
  (mSGpTR)->ittiMsg.ngap_ue_context_release_command
#define NGAP_UE_CONTEXT_RELEASE_COMPLETE(mSGpTR)                               \
  (mSGpTR)->ittiMsg.ngap_ue_context_release_complete
#define NGAP_UE_CONTEXT_MODIFICATION_REQUEST(mSGpTR)                           \
  (mSGpTR)->ittiMsg.ngap_ue_context_mod_request
#define NGAP_UE_CONTEXT_MODIFICATION_RESPONSE(mSGpTR)                          \
  (mSGpTR)->ittiMsg.ngap_ue_context_mod_response
#define NGAP_UE_CONTEXT_MODIFICATION_FAILURE(mSGpTR)                           \
  (mSGpTR)->ittiMsg.ngap_ue_context_mod_failure

#define NGAP_INITIAL_UE_MESSAGE(mSGpTR)                                        \
  (mSGpTR)->ittiMsg.ngap_initial_ue_message
#define NGAP_NAS_DL_DATA_REQ(mSGpTR) (mSGpTR)->ittiMsg.ngap_nas_dl_data_req
#define NGAP_PAGING_REQUEST(mSGpTR) (mSGpTR)->ittiMsg.ngap_paging_request
#define NGAP_PATH_SWITCH_REQUEST(mSGpTR)                                       \
  (mSGpTR)->ittiMsg.ngap_path_switch_request
#define NGAP_PATH_SWITCH_REQUEST_ACK(mSGpTR)                                   \
  (mSGpTR)->ittiMsg.ngap_path_switch_request_ack
#define NGAP_PATH_SWITCH_REQUEST_FAILURE(mSGpTR)                               \
  (mSGpTR)->ittiMsg.ngap_path_switch_request_failure

// NOT a ITTI message
typedef struct ngap_initial_ue_message_s {
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  ecgi_t e_utran_cgi;
} ngap_initial_ue_message_t;

typedef struct itti_ngap_ue_context_mod_req_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
/* Use presence mask to identify presence of optional fields */
#define NGAP_UE_CONTEXT_MOD_LAI_PRESENT (1 << 0)
#define NGAP_UE_CONTEXT_MOD_CSFB_INDICATOR_PRESENT (1 << 1)
#define NGAP_UE_CONTEXT_MOD_UE_AMBR_INDICATOR_PRESENT (1 << 2)
  uint8_t presencemask;
  lai_t lai;
  int cs_fallback_indicator;
  ambr_t ue_ambr;
} itti_ngap_ue_context_mod_req_t;

typedef struct itti_ngap_ue_context_mod_resp_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
} itti_ngap_ue_context_mod_resp_t;

typedef struct itti_ngap_ue_context_mod_resp_fail_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  int64_t cause;
} itti_ngap_ue_context_mod_resp_fail_t;

typedef struct itti_ngap_initial_ctxt_setup_req_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;

  /* Key eNB */
  uint8_t kgnb[32];

  // ambr_t ambr;
  // ambr_t apn_ambr;

  /* EPS bearer ID */
  unsigned ebi : 4;

  /* QoS */
  qci_t qci;
  priority_level_t prio_level;
  // pre_emption_vulnerability_t pre_emp_vulnerability;
  // pre_emption_capability_t pre_emp_capability;

  /* S-GW TEID for user-plane */
  teid_t teid;
  /* S-GW IP address for User-Plane */
  // ip_address_t upf_address;
} itti_ngap_initial_ctxt_setup_req_t;

typedef struct itti_ngap_ue_cap_ind_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  uint8_t* radio_capabilities;
  size_t radio_capabilities_length;
} itti_ngap_ue_cap_ind_t;

#define NGAP_ITTI_UE_PER_DEREGISTER_MESSAGE 128
typedef struct itti_ngap_eNB_deregistered_ind_s {
  uint16_t nb_ue_to_deregister;
  gnb_ue_ngap_id_t gnb_ue_ngap_id[NGAP_ITTI_UE_PER_DEREGISTER_MESSAGE];
  amf_ue_ngap_id_t amf_ue_ngap_id[NGAP_ITTI_UE_PER_DEREGISTER_MESSAGE];
  uint32_t gnb_id;
} itti_ngap_eNB_deregistered_ind_t;

typedef enum ngap_reset_type_e {
  M5G_RESET_ALL = 0,
  M5G_RESET_PARTIAL
} ngap_reset_type_t;

typedef struct ng_sig_conn_id_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id;
} ng_sig_conn_id_t;

typedef struct itti_ngap_gnb_initiated_reset_req_s {
  uint32_t sctp_assoc_id;
  uint16_t sctp_stream_id;
  uint32_t gnb_id;
  ngap_reset_type_t ngap_reset_type;
  uint32_t num_ue;
  ng_sig_conn_id_t* ue_to_reset_list;
} itti_ngap_gnb_initiated_reset_req_t;

typedef struct itti_ngap_gnb_initiated_reset_ack_s {
  uint32_t sctp_assoc_id;
  uint16_t sctp_stream_id;
  ngap_reset_type_t ngap_reset_type;
  uint32_t num_ue;
  ng_sig_conn_id_t* ue_to_reset_list;
} itti_ngap_gnb_initiated_reset_ack_t;

// List of possible causes for AMF generated UE context release command towards
// eNB
enum Ngcause {
  NGAP_INVALID_CAUSE = 0,
  NGAP_NAS_NORMAL_RELEASE,
  NGAP_NAS_DEREGISTER,
  NGAP_RADIO_NR_GENERATED_REASON,
  NGAP_IMPLICIT_CONTEXT_RELEASE,
  NGAP_INITIAL_CONTEXT_SETUP_FAILED,
  NGAP_SCTP_SHUTDOWN_OR_RESET,
  NGAP_INITIAL_CONTEXT_SETUP_TMR_EXPRD,
  NGAP_INVALID_GNB_ID,
  NGAP_CSFB_TRIGGERED,
  NGAP_NAS_UE_NOT_AVAILABLE_FOR_PS
};
typedef struct itti_ngap_ue_context_release_command_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  enum Ngcause cause;
} itti_ngap_ue_context_release_command_t;

typedef struct itti_ngap_ue_context_release_req_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  uint32_t gnb_id;
  enum Ngcause relCause;
  Ngap_Cause_t cause;
} itti_ngap_ue_context_release_req_t;

typedef struct itti_ngap_dl_nas_data_req_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  bstring nas_msg; /* Downlink NAS message             */
} itti_ngap_nas_dl_data_req_t;

typedef struct itti_ngap_ue_context_release_complete_s {
  amf_ue_ngap_id_t amf_ue_ngap_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
} itti_ngap_ue_context_release_complete_t;

#if 0
typedef enum ngap_csfb_indicator_e {
  CSFB_REQUIRED,
  CSFB_HIGH_PRIORITY
} ngap_csfb_indicator_t;
#endif
typedef enum ngap_cn_domain_e {
  M5G_CN_DOMAIN_PS,
  M5G_CN_DOMAIN_CS
} ngap_cn_domain_t;

typedef struct itti_ngap_paging_request_s {
  char imsi[IMSI_BCD_DIGITS_MAX + 1];
  uint8_t imsi_length;
  tmsi_t m_tmsi;
  // amf_code_t amf_code;
  uint32_t amf_code;
  uint32_t sctp_assoc_id;
#define NGAP_PAGING_ID_IMSI 0X0
#define NGAP_PAGING_ID_STMSI 0X1
  uint8_t paging_id;
  ngap_cn_domain_t domain_indicator;
  uint8_t tai_list_count;
  paging_tai_list_t paging_tai_list[TRACKING_AREA_IDENTITY_MAX_NUM_OF_TAIS];
} itti_ngap_paging_request_t;

typedef struct itti_ngap_initial_ue_message_s {
  sctp_assoc_id_t sctp_assoc_id;  // key stored in AMF_APP for AMF_APP forward
                                  // NAS response to NGAP
  uint32_t gnb_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id;
  amf_ue_ngap_id_t amf_ue_ngap_id;
  bstring nas;
  tai_t tai; /* Indicating the Tracking Area from which the UE has sent the NAS
                message. */
  ecgi_t ecgi; /* Indicating the cell from which the UE has sent the NAS
                  message. */
  m5g_rrc_establishment_cause_t
      m5g_rrc_establishment_cause; /* Establishment cause */
  bool is_s_tmsi_valid;
  bool is_csg_id_valid;
  bool is_guamfi_valid;
  // s_tmsi_t opt_s_tmsi;
  s_tmsi_m5_t opt_s_tmsi;
  csg_id_t opt_csg_id;
  guamfi_t opt_guamfi;
  // void                opt_cell_access_mode;
  // void                opt_cell_gw_transport_address;
  // void                opt_relay_node_indicator;
  /* Transparent message from ngap to be forwarded to AMF_APP or
   * to NGAP if connection establishment is rejected by NAS.
   */
  ngap_initial_ue_message_t transparent;
} itti_ngap_initial_ue_message_t;

#define NGAP_ITTI_UE_PER_DEREGISTER_MESSAGE 128
typedef struct itti_ngap_gNB_deregistered_ind_s {
  uint16_t nb_ue_to_deregister;
  gnb_ue_ngap_id_t gnb_ue_ngap_id[NGAP_ITTI_UE_PER_DEREGISTER_MESSAGE];
  amf_ue_ngap_id_t amf_ue_ngap_id[NGAP_ITTI_UE_PER_DEREGISTER_MESSAGE];
  uint32_t gnb_id;
} itti_ngap_gNB_deregistered_ind_t;

typedef struct itti_ngap_path_switch_request_s {
  uint32_t sctp_assoc_id;
  uint32_t gnb_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  //  e_rab_to_be_switched_in_downlink_list_t e_rab_to_be_switched_dl_list;
  amf_ue_ngap_id_t amf_ue_ngap_id;
  tai_t tai;
  ecgi_t ecgi;
  uint16_t encryption_algorithm_capabilities;
  uint16_t integrity_algorithm_capabilities;
} itti_ngap_path_switch_request_t;

typedef struct itti_ngap_path_switch_request_ack_s {
  uint32_t sctp_assoc_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  amf_ue_ngap_id_t amf_ue_ngap_id;
  // Security key
  // uint8_t nh[AUTH_NEXT_HOP_SIZE]; /* Next Hop security key*/
  uint8_t ncc; /* next hop chaining count */
} itti_ngap_path_switch_request_ack_t;

typedef struct itti_ngap_path_switch_request_failure_s {
  uint32_t sctp_assoc_id;
  gnb_ue_ngap_id_t gnb_ue_ngap_id : 24;
  amf_ue_ngap_id_t amf_ue_ngap_id;
} itti_ngap_path_switch_request_failure_t;
#endif /* FILE_NGAP_MESSAGES_TYPES_SEEN */
