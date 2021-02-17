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
/*****************************************************************************

  Source      amf_app_handler.cpp

  Version     0.1

  Date        2020/07/28

  Product     NAS stack

  Subsystem   Access and Mobility Management Function

  Author

  Description Defines Access and Mobility Management Messages

*****************************************************************************/
#include <sstream>

#ifdef __cplusplus
extern "C" {
#endif
#include "timer.h"
#include "log.h"
#include "3gpp_23.003.h"
#include "directoryd.h"
#include "amf_config.h"
#ifdef __cplusplus
}
#endif
#include "conversions.h"
#include "amf_data.h"
#include "amf_fsm.h"
#include "amf_asDefs.h"
#include "amf_sap.h"
#include "amf_app_ue_context_and_proc.h"
#include "amf_app_defs.h"
#include "ngap_messages_types.h"
#include "amf_app_state_manager.h"
#include "ngap_messages_types.h"
#include "M5gNasMessage.h"
#include "nas5g_network.h"

using namespace std;
#define QUADLET 4
#define AMF_GET_BYTE_ALIGNED_LENGTH(LENGTH)                                    \
  LENGTH += QUADLET - (LENGTH % QUADLET)

namespace magma5g {
amf_config_t amf_config_handler;
amf_sap_c amf_sap_handler;
//----------------------------------------------------------------------------
static void amf_directoryd_report_location(uint64_t imsi, uint8_t imsi_len) {
  char imsi_str[IMSI_BCD_DIGITS_MAX + 1];
  IMSI64_TO_STRING(imsi, imsi_str, imsi_len);
  directoryd_report_location(imsi_str);
  OAILOG_INFO_UE(LOG_AMF_APP, imsi, "Reported UE location to directoryd\n");
}
//------------------------------------------------------------------------------
void amf_ue_context_update_coll_keys(
    amf_ue_context_t* const amf_ue_context_p,
    ue_m5gmm_context_s* const ue_context_p,
    const gnb_ngap_id_key_t gnb_ngap_id_key,
    const amf_ue_ngap_id_t amf_ue_ngap_id, const imsi64_t imsi,
    const teid_t amf_teid_n11,
    const guti_m5_t* const
        guti_p)  //  never NULL, if none put &ue_context_p->guti
{
  hashtable_rc_t h_rc = HASH_TABLE_OK;
  // NEED-RECHECK
  hash_table_ts_t* amf_state_ue_id_ht =
      get_amf_ue_state();  // TODO -
                           // NEED-RECHECK as it is used in function
  OAILOG_FUNC_IN(LOG_AMF_APP);
  OAILOG_TRACE(
      LOG_AMF_APP,
      "Update ue context.old_gnb_ue_ngap_id_key %ld ue "
      "context.old_amf_ue_ngap_id " AMF_UE_NGAP_ID_FMT
      " ue context.old_IMSI " IMSI_64_FMT " ue context.old_GUTI " GUTI_FMT "\n",
      ue_context_p->gnb_ngap_id_key, ue_context_p->amf_ue_ngap_id,
      ue_context_p->amf_context._imsi64,
      GUTI_ARG_M5G(&ue_context_p->amf_context._guti));

  if ((INVALID_GNB_UE_NGAP_ID_KEY != gnb_ngap_id_key) &&
      (ue_context_p->gnb_ngap_id_key != gnb_ngap_id_key)) {
    h_rc = hashtable_uint64_ts_remove(
        amf_ue_context_p->gnb_ue_ngap_id_ue_context_htbl,
        (const hash_key_t) ue_context_p->gnb_ngap_id_key);
    h_rc = hashtable_uint64_ts_insert(
        amf_ue_context_p->gnb_ue_ngap_id_ue_context_htbl,
        (const hash_key_t) gnb_ngap_id_key, amf_ue_ngap_id);

    if (HASH_TABLE_OK != h_rc) {
      OAILOG_ERROR_UE(
          LOG_AMF_APP, imsi,
          "Error could not update this ue context %p "
          "gnb_ue_ngap_ue_id " GNB_UE_NGAP_ID_FMT
          " amf_ue_ngap_id " AMF_UE_NGAP_ID_FMT " %s\n",
          ue_context_p, ue_context_p->gnb_ue_ngap_id,
          ue_context_p->amf_ue_ngap_id, hashtable_rc_code2string(h_rc));
    }
    ue_context_p->gnb_ngap_id_key = gnb_ngap_id_key;
  }

  if (INVALID_AMF_UE_NGAP_ID != amf_ue_ngap_id) {
    if (ue_context_p->amf_ue_ngap_id != amf_ue_ngap_id) {
      h_rc = hashtable_ts_remove(
          amf_state_ue_id_ht, (const hash_key_t) ue_context_p->amf_ue_ngap_id,
          (void**) &ue_context_p);
      h_rc = hashtable_ts_insert(
          amf_state_ue_id_ht, (const hash_key_t) amf_ue_ngap_id,
          (void*) ue_context_p);

      if (HASH_TABLE_OK != h_rc) {
        // TODO
      }
      ue_context_p->amf_ue_ngap_id = amf_ue_ngap_id;
    }
  } else {
    // TODO
  }

  h_rc = hashtable_uint64_ts_remove(
      amf_ue_context_p->imsi_amf_ue_id_htbl,
      (const hash_key_t) ue_context_p->amf_context._imsi64);

  if (INVALID_AMF_UE_NGAP_ID != amf_ue_ngap_id) {
    h_rc = hashtable_uint64_ts_insert(
        amf_ue_context_p->imsi_amf_ue_id_htbl, (const hash_key_t) imsi,
        amf_ue_ngap_id);
  } else {
    h_rc = HASH_TABLE_KEY_NOT_EXISTS;
  }

  if (HASH_TABLE_OK != h_rc) {
    //    TODO
  }

  amf_directoryd_report_location(
      ue_context_p->amf_context._imsi64,
      ue_context_p->amf_context._imsi.length);
  h_rc = hashtable_uint64_ts_remove(
      amf_ue_context_p->tun11_ue_context_htbl,
      (const hash_key_t) ue_context_p->amf_teid_n11);

  if (INVALID_AMF_UE_NGAP_ID != amf_ue_ngap_id) {
    h_rc = hashtable_uint64_ts_insert(
        amf_ue_context_p->tun11_ue_context_htbl,
        (const hash_key_t) amf_teid_n11, (uint64_t) amf_ue_ngap_id);
  } else {
    h_rc = HASH_TABLE_KEY_NOT_EXISTS;
  }

  if (HASH_TABLE_OK != h_rc) {
    // TODO
  }

  ue_context_p->amf_teid_n11 = amf_teid_n11;

  if (guti_p) {
    if ((guti_p->guamfi.amf_code !=
         ue_context_p->amf_context._m5_guti.guamfi.amf_code) ||
        (guti_p->guamfi.amf_gid !=
         ue_context_p->amf_context._m5_guti.guamfi.amf_gid) ||
        (guti_p->m_tmsi != ue_context_p->amf_context._m5_guti.m_tmsi) ||
        (guti_p->guamfi.plmn.mcc_digit1 !=
         ue_context_p->amf_context._m5_guti.guamfi.plmn.mcc_digit1) ||
        (guti_p->guamfi.plmn.mcc_digit2 !=
         ue_context_p->amf_context._m5_guti.guamfi.plmn.mcc_digit2) ||
        (guti_p->guamfi.plmn.mcc_digit3 !=
         ue_context_p->amf_context._m5_guti.guamfi.plmn.mcc_digit3) ||
        (ue_context_p->amf_ue_ngap_id != amf_ue_ngap_id)) {
      h_rc = obj_hashtable_uint64_ts_remove(
          amf_ue_context_p->guti_ue_context_htbl,
          &ue_context_p->amf_context._m5_guti, sizeof(*guti_p));
      if (INVALID_AMF_UE_NGAP_ID != amf_ue_ngap_id) {
        h_rc = obj_hashtable_uint64_ts_insert(
            amf_ue_context_p->guti_ue_context_htbl, (const void* const) guti_p,
            sizeof(*guti_p), (uint64_t) amf_ue_ngap_id);
      } else {
        h_rc = HASH_TABLE_KEY_NOT_EXISTS;
      }
      if (HASH_TABLE_OK != h_rc) {
        // TODO
      }
      ue_context_p->amf_context._m5_guti = *guti_p;
    }
  }
  OAILOG_FUNC_OUT(LOG_AMF_APP);
}
//----------------------------------------------------------------------------------------------
static bool amf_app_construct_guti(
    const plmn_t* const plmn_p, const s_tmsi_m5_t* const s_tmsi_p,
    guti_m5_t* const guti_p) {
  /*
   * This is a helper function to construct GUTI from S-TMSI. It uses PLMN id
   * and AMF Group Id of the serving AMF for this purpose.
   *
   */
  bool is_guti_valid =
      false;  // Set to true if serving AMF is found and GUTI is constructed
  uint8_t num_amf         = 0;  // Number of configured AMF in the AMF pool
  guti_p->m_tmsi          = s_tmsi_p->m_tmsi;
  guti_p->guamfi.amf_code = s_tmsi_p->amf_code;
  // Create GUTI by using PLMN Id and AMF-Group Id of serving AMF
  OAILOG_DEBUG(
      LOG_AMF_APP,
      "Construct GUTI using S-TMSI received form UE and AMG Group Id and PLMN "
      "id from AMF Conf: %u, %u \n",
      s_tmsi_p->m_tmsi, s_tmsi_p->amf_code);
  amf_config_read_lock(&amf_config_handler);
  /*
   * Check number of MMEs in the pool.
   * At present it is assumed that one AMF is supported in AMF pool but in case
   * there are more than one AMF configured then search the serving AMF using
   * AMF code. Assumption is that within one PLMN only one pool of AMF will be
   * configured
   */
  if (amf_config_handler.guamfi.nb > 1) {
    OAILOG_DEBUG(LOG_AMF_APP, "More than one AMFs are configured.");
  }
  for (num_amf = 0; num_amf < amf_config_handler.guamfi.nb; num_amf++) {
    /*Verify that the AMF code within S-TMSI is same as what is configured in
     * AMF conf*/
    if ((plmn_p->mcc_digit2 ==
         amf_config_handler.guamfi.guamfi[num_amf].plmn.mcc_digit2) &&
        (plmn_p->mcc_digit1 ==
         amf_config_handler.guamfi.guamfi[num_amf].plmn.mcc_digit1) &&
        (plmn_p->mnc_digit3 ==
         amf_config_handler.guamfi.guamfi[num_amf].plmn.mnc_digit3) &&
        (plmn_p->mcc_digit3 ==
         amf_config_handler.guamfi.guamfi[num_amf].plmn.mcc_digit3) &&
        (plmn_p->mnc_digit2 ==
         amf_config_handler.guamfi.guamfi[num_amf].plmn.mnc_digit2) &&
        (plmn_p->mnc_digit1 ==
         amf_config_handler.guamfi.guamfi[num_amf].plmn.mnc_digit1) &&
        (guti_p->guamfi.amf_code ==
         amf_config_handler.guamfi.guamfi[num_amf].amf_code)) {
      break;
    }
  }
  if (num_amf >= amf_config_handler.guamfi.nb) {
    OAILOG_DEBUG(LOG_AMF_APP, "No AMF serves this UE");
  } else {
    guti_p->guamfi.plmn    = amf_config_handler.guamfi.guamfi[num_amf].plmn;
    guti_p->guamfi.amf_gid = amf_config_handler.guamfi.guamfi[num_amf].amf_gid;
    is_guti_valid          = true;
  }
  amf_config_unlock(&amf_config_handler);
  return is_guti_valid;
}
//------------------------------------------------------------------------------
ue_m5gmm_context_s* amf_ue_context_exists_guti(
    amf_ue_context_t* const amf_ue_context_p, const guti_m5_t* const guti_p) {
  hashtable_rc_t h_rc       = HASH_TABLE_OK;
  uint64_t amf_ue_ngap_id64 = 0;
  h_rc                      = obj_hashtable_uint64_ts_get(
      amf_ue_context_p->guti_ue_context_htbl, (const void*) guti_p,
      sizeof(*guti_p), &amf_ue_ngap_id64);

  if (HASH_TABLE_OK == h_rc) {
    // TODO
  } else {
    OAILOG_WARNING(LOG_AMF_APP, " No GUTI hashtable for GUTI ");
  }

  return NULL;
}
//-----------------------------------------------------------------------------------------
imsi64_t amf_app_defs::amf_app_handle_initial_ue_message(
    amf_app_desc_t* amf_app_desc_p,
    itti_ngap_initial_ue_message_t* const initial_pP) {
  OAILOG_FUNC_IN(LOG_AMF_APP);
  class ue_m5gmm_context_s* ue_context_p = NULL;
  bool is_guti_valid                     = false;
  bool is_mm_ctx_new                     = false;
  gnb_ngap_id_key_t gnb_ngap_id_key      = INVALID_GNB_UE_NGAP_ID_KEY;
  imsi64_t imsi64                        = INVALID_IMSI64;
  amf_app_msg amf_app_message;
  guti_m5_t guti;
  plmn_t plmn;
  nas_proc nas_procedure;

  if (initial_pP->amf_ue_ngap_id != INVALID_AMF_UE_NGAP_ID) {
    OAILOG_ERROR(
        LOG_AMF_APP,
        "AMF UE NGAP Id (" AMF_UE_NGAP_ID_FMT ") is already assigned\n",
        initial_pP->amf_ue_ngap_id);
  }

  // Check if there is any existing UE context using S-TMSI/GUTI
  if (initial_pP->is_s_tmsi_valid) {  // TODO -  NEED-RECHECK
    OAILOG_INFO(
        LOG_AMF_APP,
        "INITIAL UE Message: Valid amf_code and S-TMSI received from \n");
    guti.guamfi.plmn        = {0};
    guti.guamfi.amf_gid     = 0;
    guti.guamfi.amf_code    = 0;
    guti.guamfi.amf_Pointer = 0;
    guti.m_tmsi             = INVALID_M_TMSI;
    plmn.mcc_digit1         = initial_pP->tai.plmn.mcc_digit1;
    plmn.mcc_digit2         = initial_pP->tai.plmn.mcc_digit2;
    plmn.mcc_digit3         = initial_pP->tai.plmn.mcc_digit3;
    plmn.mnc_digit1         = initial_pP->tai.plmn.mnc_digit1;
    plmn.mnc_digit2         = initial_pP->tai.plmn.mnc_digit2;
    plmn.mnc_digit3         = initial_pP->tai.plmn.mnc_digit3;
    is_guti_valid =
        amf_app_construct_guti(&plmn, &(initial_pP->opt_s_tmsi), &guti);
    // create a new ue context if nothing is found
    if (is_guti_valid) {
      ue_context_p =
          amf_ue_context_exists_guti(&amf_app_desc_p->amf_ue_contexts, &guti);
      if (ue_context_p) {
        initial_pP->amf_ue_ngap_id = ue_context_p->amf_ue_ngap_id;
        if (ue_context_p->gnb_ngap_id_key != INVALID_GNB_UE_NGAP_ID_KEY) {
          /*
           * Ideally this should never happen. When UE moves to IDLE,
           * this key is set to INVALID.
           * Note - This can happen if eNB detects RLF late and by that time
           * UE sends Initial NAS message via new RRC connection.
           * However if this key is valid, remove the key from the hashtable.
           */
          OAILOG_ERROR(
              LOG_AMF_APP,
              "AMF_APP_INITAIL_UE_MESSAGE: gnb_ngap_id_key %ld has "
              "valid value \n",
              ue_context_p->gnb_ngap_id_key);
          amf_app_message.amf_app_ue_context_release(
              ue_context_p, ue_context_p->ue_context_rel_cause);
          hashtable_uint64_ts_remove(
              amf_app_desc_p->amf_ue_contexts.gnb_ue_ngap_id_ue_context_htbl,
              (const hash_key_t) ue_context_p->gnb_ngap_id_key);
          ue_context_p->gnb_ngap_id_key = INVALID_GNB_UE_NGAP_ID_KEY;
        }
        // Update AMF UE context with new gnb_ue_ngap_id
        ue_context_p->gnb_ue_ngap_id = initial_pP->gnb_ue_ngap_id;
        amf_ue_context_update_coll_keys(
            &amf_app_desc_p->amf_ue_contexts, ue_context_p, gnb_ngap_id_key,
            ue_context_p->amf_ue_ngap_id, ue_context_p->amf_context._imsi64,
            ue_context_p->amf_teid_n11, &guti);
        imsi64 = ue_context_p->amf_context._imsi64;
      }
    } else {
      //   TODO
    }
  } else {
    OAILOG_INFO(
        LOG_AMF_APP,
        " AMF_TEST: AMF_APP_INITIAL_UE_MESSAGE from NGAP,without S-TMSI. \n");
  }
  // create a new ue context if nothing is found
  if (ue_context_p == NULL) {
    OAILOG_INFO(
        LOG_AMF_APP, "AMF_TEST: UE context doesn't exist -> create one\n");
    if (!(ue_context_p = amf_create_new_ue_context())) {
      OAILOG_INFO(LOG_AMF_APP, "Failed to create context \n");
    }
    // Allocate new amf_ue_ngap_id
    ue_context_p->amf_ue_ngap_id =
        amf_app_ue_context::amf_app_ctx_get_new_ue_id(
            &amf_app_desc_p->amf_app_ue_ngap_id_generator);
    if (ue_context_p->amf_ue_ngap_id == INVALID_AMF_UE_NGAP_ID) {
      OAILOG_CRITICAL(
          LOG_AMF_APP,
          "AMF_APP_INITIAL_UE_MESSAGE. AMF_UE_NGAP_ID allocation Failed.\n");
      amf_app_ue_context::amf_remove_ue_context(
          &amf_app_desc_p->amf_ue_contexts, ue_context_p);
      OAILOG_FUNC_RETURN(LOG_AMF_APP, imsi64);
    }
    AMF_APP_GNB_NGAP_ID_KEY(
        ue_context_p->gnb_ngap_id_key, initial_pP->gnb_id,
        initial_pP->gnb_ue_ngap_id);
    amf_app_ue_context::amf_insert_ue_context(
        &amf_app_desc_p->amf_ue_contexts, ue_context_p);

    amf_app_ue_context::notify_ngap_new_ue_amf_ngap_id_association(
        ue_context_p);
    s_tmsi_m5_t s_tmsi = {0};
    if (initial_pP->is_s_tmsi_valid) {
      s_tmsi = initial_pP->opt_s_tmsi;
    } else {
      s_tmsi.amf_code = 0;
      s_tmsi.m_tmsi   = INVALID_M_TMSI;
    }

    OAILOG_INFO(
        LOG_AMF_APP,
        "AMF_TEST: Sending NAS Establishment Indication to NAS for ue_id = "
        "(%d)\n",
        ue_context_p->amf_ue_ngap_id);
    amf_ue_ngap_id_t ue_id = ue_context_p->amf_ue_ngap_id;
    nas_procedure.nas_proc_establish_ind(
        ue_context_p->amf_ue_ngap_id, is_mm_ctx_new, initial_pP->tai,
        initial_pP->ecgi, initial_pP->m5g_rrc_establishment_cause, s_tmsi,
        initial_pP->nas);
  }
}

int amf_app_defs::amf_app_handle_uplink_nas_message(
    amf_app_desc_t* amf_app_desc_p, bstring msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int rc = RETURNerror;
  OAILOG_INFO(
      LOG_AMF_APP,
      "AMF_TEST: Received NAS UPLINK DATA IND from NGAP\n");  // for ue_id =
                                                              // (%u)\n",
  if (msg) {
    amf_sap_t amf_sap;
    /*
     * Notify the AMF procedure call manager that data transfer
     * indication has been received from the Access-Stratum sublayer
     */
    amf_sap.primitive                  = AMFAS_ESTABLISH_REQ;
    amf_sap.u.amf_as.u.establish.ue_id = 1;  // TODO AMF_TEST, generate the
                                             // ue_id
    amf_sap.u.amf_as.u.establish.nas_msg = msg;
    msg                                  = NULL;
    rc = amf_sap_handler.amf_sap_send(&amf_sap);
  } else {
    OAILOG_WARNING(
        LOG_NAS, "Received NAS message in uplink is NULL for ue_id = (%u)\n",
        amf_app_desc_p->amf_app_ue_ngap_id_generator);
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

void amf_app_defs::amf_app_handle_pdu_session_response(
    itti_n11_create_pdu_session_response_t* pdu_session_resp) {
  extern ue_m5gmm_context_s
      ue_m5gmm_global_context;  // TODO AMF_TEST global var to temporarily store
                                // context inserted to ht
  DLNASTransportMsg* encode_msg;
  // amf_app_defs amf_app_def_as;
  nas_network nas_networks;
  SmfMsg* smf_msg;
  bstring buffer;
  uint32_t len;
  nas5g_error_code_t rc = M5G_AS_SUCCESS;
  int amf_rc            = RETURNerror;
  ue_m5gmm_context_s* ue_context;
  smf_context_t* smf_ctx;
  uint32_t bytes         = 0;
  uint32_t ue_id         = 1;  // TODO AMF_TEST get the ue_id from imsi from ht
  uint32_t container_len = 0;
  uint16_t ambr_len      = 0;

  // Handle smf_context
  ue_context = amf_ue_context_exists_amf_ue_ngap_id(ue_id);
  if (ue_context) {
    smf_ctx = &(ue_context->amf_context.smf_context);
  } else {
    ue_context = &ue_m5gmm_global_context;
    smf_ctx    = &ue_m5gmm_global_context.amf_context
                   .smf_context;  // TODO AMF_TEST global var to temporarily
                                  // store context inserted to ht
  }
  smf_ctx->dl_session_ambr = pdu_session_resp->session_ambr.downlink_units;
  smf_ctx->dl_ambr_unit    = pdu_session_resp->session_ambr.downlink_unit_type;
  smf_ctx->ul_session_ambr = pdu_session_resp->session_ambr.uplink_units;
  smf_ctx->ul_ambr_unit    = pdu_session_resp->session_ambr.uplink_unit_type;
  /*required for PDUSessionResourceSetupRequest to gNB with UPF teid*/
  memcpy(
      &(smf_ctx->pdu_resource_setup_req
            .pdu_session_resource_setup_request_transfer
            .qos_flow_setup_request_list),
      &(pdu_session_resp->qos_list), sizeof(qos_flow_request_list));
  memcpy(
      smf_ctx->gtp_tunnel_id.upf_gtp_teid_ip_addr,
      pdu_session_resp->upf_endpoint.end_ipv4_addr,
      sizeof(smf_ctx->gtp_tunnel_id.upf_gtp_teid_ip_addr));
  memcpy(
      smf_ctx->gtp_tunnel_id.upf_gtp_teid, pdu_session_resp->upf_endpoint.teid,
      sizeof(smf_ctx->gtp_tunnel_id.upf_gtp_teid));

  // Sent message to gNB, for PDUSessionResourceSetupRequest
  OAILOG_INFO(
      LOG_AMF_APP,
      "Sending message to gNB for PDUSessionResourceSetupRequest\n");
  OAILOG_INFO(
      LOG_AMF_APP, "#######TIED: %02x %02x %02x %02x \n",
      smf_ctx->gtp_tunnel_id.upf_gtp_teid[0],
      smf_ctx->gtp_tunnel_id.upf_gtp_teid[1],
      smf_ctx->gtp_tunnel_id.upf_gtp_teid[2],
      smf_ctx->gtp_tunnel_id.upf_gtp_teid[3]);
  OAILOG_INFO(
      LOG_AMF_APP, "#######TIED: %s \n", smf_ctx->gtp_tunnel_id.upf_gtp_teid);
  OAILOG_INFO(
      LOG_AMF_APP, "#######IP: %02x %02x %02x %02x \n",
      smf_ctx->gtp_tunnel_id.upf_gtp_teid_ip_addr[0],
      smf_ctx->gtp_tunnel_id.upf_gtp_teid_ip_addr[1],
      smf_ctx->gtp_tunnel_id.upf_gtp_teid_ip_addr[2],
      smf_ctx->gtp_tunnel_id.upf_gtp_teid_ip_addr[3]);

  amf_rc = pdu_session_resource_setup_request(ue_context, ue_id);
  if (amf_rc != RETURNok) {
    OAILOG_INFO(
        LOG_AMF_APP,
        "Failure in sending message to gNB for "
        "PDUSessionResourceSetupRequest\n");
    // in this negative case handling, send pdu reject command to UE and release
    // message to SMF
  }
  // smf_msg = &encode_msg.payload_container.smf_msg.pdu_session_estab_accept;

  amf_nas_message_t msg;
  msg.security_protected.plain.amf.header.extended_protocol_discriminator =
      M5G_MOBILITY_MANAGEMENT_MESSAGES;
  msg.security_protected.plain.amf.header.message_type = DLNASTRANSPORT;
  // amf_as::amf_as_set_header(&msg, ue_context->amf_context._security);
  msg.header.security_header_type =
      SECURITY_HEADER_TYPE_INTEGRITY_PROTECTED_NEW;
  msg.header.extended_protocol_discriminator = M5G_MOBILITY_MANAGEMENT_MESSAGES;
  msg.header.sequence_number =
      ue_context->amf_context._security.dl_count.seq_num;

  encode_msg = &msg.security_protected.plain.amf.downlinknas5gtransport;
  smf_msg    = &encode_msg->payload_container.smf_msg;

  // NAS AmfHeader
  encode_msg->extended_protocol_discriminator.extended_proto_discriminator =
      M5G_MOBILITY_MANAGEMENT_MESSAGES;
  encode_msg->spare_half_octet.spare     = 0x00;
  encode_msg->sec_header_type.sec_hdr    = 0x00;
  encode_msg->message_type.msg_type      = DLNASTRANSPORT;
  encode_msg->payload_container_type.iei = PAYLOAD_CONTAINER_TYPE;
  encode_msg->pdu_session_identity.iei   = 0x12;
  encode_msg->pdu_session_identity.pdu_session_id =
#define N1_SM_INFO 0x1  // TODO define in "M5gNasMessage.h" //pdu_change
      pdu_session_resp->pdu_session_id;

  // NAS SmfMsg
  encode_msg->payload_container_type.type_val = N1_SM_INFO;
  encode_msg->payload_container.iei           = PAYLOAD_CONTAINER;

  smf_msg->header.extended_protocol_discriminator =
      M5G_SESSION_MANAGEMENT_MESSAGES;
  container_len++;
  smf_msg->header.pdu_session_id = pdu_session_resp->pdu_session_id;
  container_len++;
  smf_msg->header.message_type = PDU_SESSION_ESTABLISHMENT_ACCEPT;
  container_len++;
  smf_msg->header.procedure_transaction_id =
      smf_ctx->smf_proc_data.pti.pti;  // TODO get it from SMF reply
  container_len++;

  smf_msg->pdu_session_estab_accept.extended_protocol_discriminator
      .extended_proto_discriminator = M5G_SESSION_MANAGEMENT_MESSAGES;
  container_len++;
  smf_msg->pdu_session_estab_accept.pdu_session_identity.pdu_session_id =
      pdu_session_resp->pdu_session_id;
  container_len++;
  smf_msg->pdu_session_estab_accept.pti.pti =
      smf_ctx->smf_proc_data.pti.pti;  // TODO get it from SMF reply

  OAILOG_INFO(LOG_AMF_APP, "AMF_TEST: pti: %d", smf_ctx->smf_proc_data.pti.pti);
  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: pti: %d",
      pdu_session_resp->procedure_trans_identity[0]);

  smf_msg->pdu_session_estab_accept.message_type.msg_type =
      PDU_SESSION_ESTABLISHMENT_ACCEPT;
  container_len++;
  smf_msg->pdu_session_estab_accept.pdu_session_type.type_val =
      pdu_session_resp->pdu_session_type;

  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: pdu_session_type: %d",
      pdu_session_resp->pdu_session_type);

  smf_msg->pdu_session_estab_accept.ssc_mode.mode_val =
      0x1;  // TODO fix mapping from NAS not covered in amf_smf_send

  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: selected_ssc_mode: %d",
      pdu_session_resp->selected_ssc_mode);

  container_len++;
  smf_msg->pdu_session_estab_accept.pti.pti = 0x01;
  container_len++;
  memset(smf_msg->pdu_session_estab_accept.pdu_address.address_info, '\0', 12);
  memcpy(
      smf_msg->pdu_session_estab_accept.pdu_address.address_info,
      pdu_session_resp->pdu_address.redirect_server_address, 4);
  smf_msg->pdu_session_estab_accept.pdu_address.type_val = 0x1;

  // QOSRulesMsg qos_rules;
  smf_msg->pdu_session_estab_accept.qos_rules.length = 0x9;
  QOSRule qos_rule;
  qos_rule.qos_rule_id         = 0x1;
  qos_rule.len                 = 0x6;
  qos_rule.rule_oper_code      = 0x1;
  qos_rule.dqr_bit             = 0x1;
  qos_rule.no_of_pkt_filters   = 0x1;
  qos_rule.qos_rule_precedence = 0xff;
  qos_rule.spare               = 0x0;
  qos_rule.segregation         = 0x0;
  qos_rule.qfi                 = 0x6;
  NewQOSRulePktFilter new_qos_rule_pkt_filter;
  new_qos_rule_pkt_filter.spare          = 0x0;
  new_qos_rule_pkt_filter.pkt_filter_dir = 0x3;
  new_qos_rule_pkt_filter.pkt_filter_id  = 0x1;
  new_qos_rule_pkt_filter.len            = 0x1;
  uint8_t contents                       = 0x1;
  memcpy(
      new_qos_rule_pkt_filter.contents, &contents, new_qos_rule_pkt_filter.len);
  memcpy(
      qos_rule.new_qos_rule_pkt_filter, &new_qos_rule_pkt_filter,
      1 * sizeof(NewQOSRulePktFilter));
  memcpy(
      smf_msg->pdu_session_estab_accept.qos_rules.qos_rule, &qos_rule,
      1 * sizeof(QOSRule));

  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: session_ambr.downlink_units: %d",
      pdu_session_resp->session_ambr.downlink_units);
  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: session_ambr.uplink_units: %d",
      pdu_session_resp->session_ambr.uplink_units);
  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: session_ambr.dl_unit: %d",
      pdu_session_resp->session_ambr.downlink_unit_type);
  OAILOG_INFO(
      LOG_AMF_APP, "AMF_TEST: session_ambr.ul_unit: %d",
      pdu_session_resp->session_ambr.uplink_unit_type);

  smf_msg->pdu_session_estab_accept.session_ambr.dl_unit = 0x4;
  ambr_len += 1;
  smf_msg->pdu_session_estab_accept.session_ambr.ul_unit = 0x4;
  ambr_len += 1;
  smf_msg->pdu_session_estab_accept.session_ambr.dl_session_ambr = 0x01;
  ambr_len += 2;
  smf_msg->pdu_session_estab_accept.session_ambr.ul_session_ambr = 0x01;
  ambr_len += 2;
  smf_msg->pdu_session_estab_accept.session_ambr.length = ambr_len;
  smf_msg->pdu_session_estab_accept.dnn.len             = 12;
  smf_msg->pdu_session_estab_accept.dnn.dnn             = "carrier.com";
  encode_msg->payload_container.len = 30;
  OAILOG_INFO(
      LOG_AMF_APP,
      "AMF_TEST: start NAS encoding for PDU Session Establishment Accept\n");

  len = 41;  

  /* Ciphering algorithms, EEA1 and EEA2 expects length to be mode of 4,
   * so length is modified such that it will be mode of 4
   */
  AMF_GET_BYTE_ALIGNED_LENGTH(len);
  if (msg.header.security_header_type != SECURITY_HEADER_TYPE_NOT_PROTECTED) {
    amf_msg_header* header = &msg.security_protected.plain.amf.header;
    /*
     * Expand size of protected NAS message
     */
    OAILOG_INFO(
        LOG_AMF_APP, "AMF_TEST:before adding sec header, length %d ", len);
    len += NAS_MESSAGE_SECURITY_HEADER_SIZE;
    OAILOG_INFO(
        LOG_AMF_APP, "AMF_TEST:after adding sec header, length %d ", len);
    /*
     * Set header of plain NAS message
     */
    header->extended_protocol_discriminator = M5GS_MOBILITY_MANAGEMENT_MESSAGE;
    header->security_header_type = SECURITY_HEADER_TYPE_NOT_PROTECTED;
  }

  buffer = bfromcstralloc(len, "\0");
  bytes = nas5g_message_encode(
      buffer->data, &msg, len, &ue_context->amf_context._security);
  OAILOG_INFO(LOG_AMF_APP, "bytes:%d \n", bytes);
  if (bytes > 0) {
    OAILOG_INFO(
        LOG_AMF_APP,
        "NAS encode success, sent PDU Establishment Accept to UE\n");
    buffer->slen = bytes;
    amf_app_handle_nas_dl_req(ue_id, buffer, rc);

  } else {
    nas_networks.bdestroy_wrapper(&buffer);
  }
}

// resource setup request and release UL procedure defination
void amf_app_handle_resource_setup_response(
    itti_ngap_pdusessionresource_setup_rsp_t session_seup_resp) {
  /* Check if failure message is not NULL and if NULL,
   * it is successful message from gNB.
   * Nothing to in this case. If failure message comes from gNB
   * AMF need to report this failed message to SMF
   *
   * NOTE: only handling success part not failure part
   * will be handled later
   */
  OAILOG_INFO(
      LOG_AMF_APP,
      "AMF_TEST: handling uplink PDU session setup response message\n");
  if (session_seup_resp.pduSessionResource_setup_list.no_of_items > 0) {
    /* This is success case and we need not to send message to SMF
     * and drop the message here
     */
    OAILOG_INFO(
        LOG_AMF_APP,
        "AMF_TEST: this is success case and no need to hadle anything and drop "
        "the message\n");
  } else {
    // TODO implement failure message from gNB. messagge to send to SMF
    OAILOG_INFO(
        LOG_AMF_APP,
        "AMF_TEST: Failure message not handled and dropping the message\n");
  }
}

void amf_app_handle_resource_release_response(
    itti_ngap_pdusessionresource_rel_rsp_t session_rel_resp) {
  /*
   * Release request always should be successful.
   * This response message will be dropped here as nothing to do.
   * as pdu_session_resource_release_response_transfer is
   * optional as per 38.413 - 9.3.4.2.1
   */
  OAILOG_INFO(
      LOG_AMF_APP,
      "AMF_TEST: handling uplink PDU session release response message\n");
  if (session_rel_resp.pduSessionResourceReleasedRspList.no_of_items > 0) {
    /* This is success case and we need not to send message to SMF
     * and drop the message here
     */
    OAILOG_INFO(
        LOG_AMF_APP,
        "AMF_TEST: this is success case of release response and no need to "
        "hadle anything and drop the message\n");
  } else {
    // TODO implement failure message from gNB. messagge to send to SMF
    OAILOG_INFO(
        LOG_AMF_APP,
        "AMF_TEST: Failure message not handled and dropping the message\n");
  }
}
}  // namespace magma5g
