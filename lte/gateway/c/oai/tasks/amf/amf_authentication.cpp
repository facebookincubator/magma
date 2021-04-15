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

#include <sstream>
#ifdef __cplusplus
extern "C" {
#endif
#include "intertask_interface.h"
#include "conversions.h"
#include "log.h"
#ifdef __cplusplus
}
#endif
#include "common_defs.h"
#include "amf_app_ue_context_and_proc.h"
#include "amf_authentication.h"
#include "amf_recv.h"
#include "amf_identity.h"
#include "dynamic_memory_check.h"
#include "amf_sap.h"
#define AMF_CAUSE_SUCCESS (1)
#define MAX_5G_AUTH_VECTORS 1

namespace magma5g {
extern task_zmq_ctx_t amf_app_task_zmq_ctx;

amf_as_data_t amf_data_sec_auth;

/****************************************************************************
 **                                                                        **
 ** Name:        nas_itti_auth_info_req()                                  **
 **                                                                        **
 ** Description: Sends Authentication Request to UDM via S6a Task          **
 **                                                                        **
 ** Inputs: ue_idP:         UE context Identifier                          **
 **         imsiP:          IMSI of UE                                     **
 **         is_initial_reqP:Flag to indicate, whether Auth Req is sent     **
 **                      for first time or initited as part of             **
 **                      re-synchronisation                                **
 **      visited_plmnP : Visited PLMN                                      **
 **      num_vectorsP : Number of Auth vectors in case of                  **
 **                    re-synchronisation                                  **
 **      auts_pP : sent in case of re-synchronisation                      **
 ** Outputs:                                                               **
 **     Return: None                                                       **
 **                                                                        **
 ***************************************************************************/
static void nas_itti_auth_info_req(
    const amf_ue_ngap_id_t ue_id, const imsi_t* const imsiP,
    const bool is_initial_reqP, plmn_t* const visited_plmnP,
    const uint8_t num_vectorsP, const_bstring const auts_pP) {
  OAILOG_FUNC_IN(LOG_NAS);
  MessageDef* message_p             = NULL;
  n6_auth_info_req_t* auth_info_req = NULL;
  OAILOG_DEBUG(
      LOG_NAS_AMF,
      "Sending authentication information request message to S6A for ue_id "
      "= " AMF_UE_NGAP_ID_FMT "\n",
      ue_id);
  message_p = itti_alloc_new_message(TASK_AMF_APP, S6A_AUTH_INFO_REQ);
  if (!message_p) {
    OAILOG_CRITICAL(
        LOG_NAS_AMF,
        "itti_alloc_new_message failed for authentication "
        "information request message to S6A for ue-id = " AMF_UE_NGAP_ID_FMT
        "\n",
        ue_id);
    OAILOG_FUNC_OUT(LOG_NAS);
  }
  memset(auth_info_req, 0, sizeof(n6_auth_info_req_t));
  IMSI_TO_STRING(imsiP, auth_info_req->imsi, IMSI_BCD_DIGITS_MAX + 1);
  auth_info_req->imsi_length = (uint8_t) strlen(auth_info_req->imsi);

  if (!(auth_info_req->imsi_length > MSIN_MAX_LENGTH) &&
      (auth_info_req->imsi_length < (IMSI_BCD_DIGITS_MAX + 1))) {
    OAILOG_WARNING(
        LOG_NAS_AMF, "Invalid IMSI length %d", auth_info_req->imsi_length);
    OAILOG_FUNC_OUT(LOG_NAS);
  }
  auth_info_req->nb_of_vectors = num_vectorsP;
  if (is_initial_reqP) {
    auth_info_req->re_synchronization = 0;
    memset(auth_info_req->resync_param, 0, sizeof auth_info_req->resync_param);
  } else {
    if (!auts_pP) {
      OAILOG_WARNING(LOG_NAS_AMF, "Auts is Null during resynchronization \n");
      OAILOG_FUNC_OUT(LOG_NAS);
    }
    auth_info_req->re_synchronization = 1;
    memcpy(
        auth_info_req->resync_param, auts_pP->data,
        sizeof auth_info_req->resync_param);
  }
  send_msg_to_task(&amf_app_task_zmq_ctx, TASK_S6A, message_p);
  OAILOG_FUNC_OUT(LOG_NAS);
}

nas_amf_smc_proc_t* get_nas5g_common_procedure_smc(const amf_context_t* ctxt) {
  return (nas_amf_smc_proc_t*) get_nas5g_common_procedure(
      ctxt, AMF_COMM_PROC_SMC);
}

nas5g_cn_proc_t* get_nas5g_cn_procedure(
    const amf_context_t* ctxt, cn5g_proc_type_t proc_type) {
  if (ctxt) {
    if (ctxt->amf_procedures) {
      nas5g_cn_procedure_t* p1 = LIST_FIRST(&ctxt->amf_procedures->cn_procs);
      nas5g_cn_procedure_t* p2 = NULL;
      while (p1) {
        p2 = LIST_NEXT(p1, entries);
        if (p1->proc->type == proc_type) {
          return p1->proc;
        }
        p1 = p2;
      }
    }
  }
  return NULL;
}

/***************************************************************************
**                                                                        **
** Name:    get_nas5g_cn_procedure_auth_info()                            **
**                                                                        **
** Description: Invokes get_nas5g_cn_procedure                            **
**              to fetch new security context                             **
**                                                                        **
**                                                                        **
***************************************************************************/
nas5g_auth_info_proc_t* get_nas5g_cn_procedure_auth_info(
    const amf_context_t* ctxt) {
  return (nas5g_auth_info_proc_t*) get_nas5g_cn_procedure(
      ctxt, CN5G_PROC_AUTH_INFO);
}

/***************************************************************************
**                                                                        **
** Name:    start_authentication_information_procedure()                  **
**                                                                        **
** Description: Invokes get_nas5g_cn_proceduree_auth_info                 **
**              to fetch new security context                             **
**                                                                        **
**                                                                        **
***************************************************************************/
static int start_authentication_information_procedure(
    amf_context_t* amf_context, nas5g_amf_auth_proc_t* const auth_proc,
    const_bstring auts) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  amf_ue_ngap_id_t ue_id =
      PARENT_STRUCT(amf_context, ue_m5gmm_context_s, amf_context)
          ->amf_ue_ngap_id;
  // Upper layer to fetch new security context
  nas5g_auth_info_proc_t* auth_info_proc =
      get_nas5g_cn_procedure_auth_info(amf_context);
  if (!auth_info_proc) {
    auth_info_proc               = nas5g_cn_auth_info_procedure(amf_context);
    auth_info_proc->request_sent = false;
  }
  auth_info_proc->ue_id        = ue_id;
  auth_info_proc->resync       = auth_info_proc->request_sent;
  plmn_t visited_plmn          = {0};
  bool is_initial_req          = !(auth_info_proc->request_sent);
  auth_info_proc->request_sent = true;
  nas_itti_auth_info_req(
      ue_id, &amf_context->imsi, is_initial_req, &visited_plmn,
      MAX_EPS_AUTH_VECTORS, auts);
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNok);
}

/***************************************************************************
**                                                                        **
** Name:    get_nas5g_common_procedure()                                  **
**                                                                        **
** Description:  Generic function to fetch security context and others    **
**                                                                        **
**                                                                        **
***************************************************************************/
nas_amf_common_proc_t* get_nas5g_common_procedure(
    const amf_context_t* const ctxt, amf_common_proc_type_t proc_type) {
  if (ctxt) {
    if (ctxt->amf_procedures) {
      nas_amf_common_procedure_t* p1 =
          LIST_FIRST(&ctxt->amf_procedures->amf_common_procs);
      nas_amf_common_procedure_t* p2 = NULL;
      while (p1) {
        p2 = LIST_NEXT(p1, entries);
        if (p1->proc->type == proc_type) {
          return p1->proc;
        }
        p1 = p2;
      }
    }
  }
  return NULL;
}

/***************************************************************************
**                                                                        **
** Name:    get_nas5g_common_procedure_authentication() **
**                                                                        **
** Description:  Generic function to fetch security context and others    **
**                                                                        **
**                                                                        **
***************************************************************************/
nas5g_amf_auth_proc_t* get_nas5g_common_procedure_authentication(
    const amf_context_t* const ctxt) {
  return (nas5g_amf_auth_proc_t*) get_nas5g_common_procedure(
      ctxt, AMF_COMM_PROC_AUTH);
}

/****************************************************************************
 **                                                                        **
 ** Name:    amf_authentication_abort()                                    **
 **                                                                        **
 ** Description: Aborts the authentication procedure currently in progress **
 **                                                                        **
 ** Inputs:  args:      Authentication data to be released                 **
 **      Others:    None                                                   **
 **                                                                        **
 ** Outputs:     None                                                      **
 **     Return: None                                                       **
 **     Others: None                                                       **
 **                                                                        **
 ***************************************************************************/
static int amf_authentication_abort(
    amf_context_t* amf_ctx, struct nas5g_base_proc_t* base_proc) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int rc = RETURNerror;
  if ((base_proc) && (amf_ctx)) {
    ue_m5gmm_context_s* ue_mm_context =
        PARENT_STRUCT(amf_ctx, ue_m5gmm_context_s, amf_context);
    OAILOG_DEBUG(
        LOG_NAS_AMF,
        "AMF-PROC  - Abort authentication procedure invoked "
        "(ue_id= " AMF_UE_NGAP_ID_FMT ")\n",
        ue_mm_context->amf_ue_ngap_id);
    // TODO in future need to be implemented.
    rc = RETURNok;
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

/***************************************************************************
**                                                                        **
** Name:    nas5g_new_authentication_procedure()                          **
**                                                                        **
** Description:  Handler for nas5g Authenthication Procedure              **
**                                                                        **
**                                                                        **
***************************************************************************/
nas5g_amf_auth_proc_t* nas5g_new_authentication_procedure(
    amf_context_t* const amf_context) {
  if (!(amf_context->amf_procedures)) {
    amf_context->amf_procedures = nas_new_amf_procedures(amf_context);
  }
  nas5g_amf_auth_proc_t* auth_proc = new (nas5g_amf_auth_proc_t);
  auth_proc->amf_com_proc.amf_proc.base_proc.nas_puid =
      __sync_fetch_and_add(&nas_puid, 1);
  auth_proc->amf_com_proc.amf_proc.base_proc.type = NAS_PROC_TYPE_AMF;
  auth_proc->amf_com_proc.amf_proc.type           = NAS_AMF_PROC_TYPE_COMMON;
  auth_proc->amf_com_proc.type                    = AMF_COMM_PROC_AUTH;
  nas_amf_common_procedure_t* wrapper = new nas_amf_common_procedure_t;
  if (wrapper) {
    wrapper->proc = &auth_proc->amf_com_proc;
    LIST_INSERT_HEAD(
        &amf_context->amf_procedures->amf_common_procs, wrapper, entries);
    OAILOG_TRACE(LOG_NAS_AMF, "New AMF_COMM_PROC_AUTH\n");
    return auth_proc;
  } else {
    free_wrapper((void**) &auth_proc);
  }
  return NULL;
}

/***************************************************************************
**                                                                        **
** Name:    amf_proc_authentication                                       **
**                                                                        **
** Description:  Procedure to start Authentication procedure              **
**                                                                        **
**                                                                        **
***************************************************************************/
int amf_proc_authentication(
    amf_context_t* amf_context,
    nas_amf_specific_proc_t* const amf_specific_proc, success_cb_t success,
    failure_cb_t failure) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int rc                  = RETURNerror;
  bool run_auth_info_proc = false;
  ksi_t eksi              = 0;
  OAILOG_DEBUG(LOG_NGAP, "starting Authentication procedure");
  amf_ue_ngap_id_t ue_id =
      PARENT_STRUCT(amf_context, ue_m5gmm_context_s, amf_context)
          ->amf_ue_ngap_id;
  nas5g_amf_auth_proc_t* auth_proc =
      get_nas5g_common_procedure_authentication(amf_context);
  if (!auth_proc) {
    auth_proc = nas5g_new_authentication_procedure(amf_context);
  }
  if (auth_proc) {
    if (amf_specific_proc) {
      if (AMF_SPEC_PROC_TYPE_REGISTRATION == amf_specific_proc->type) {
        auth_proc->is_cause_is_registered = true;
      } else if (AMF_SPEC_PROC_TYPE_TAU == amf_specific_proc->type) {
        auth_proc->is_cause_is_registered = false;
      }
    }
    auth_proc->amf_cause            = AMF_CAUSE_SUCCESS;
    auth_proc->retransmission_count = 0;
    auth_proc->ue_id                = ue_id;
    ((nas5g_base_proc_t*) auth_proc)->parent =
        (nas5g_base_proc_t*) amf_specific_proc;
    auth_proc->amf_com_proc.amf_proc.delivered               = NULL;
    auth_proc->amf_com_proc.amf_proc.not_delivered           = NULL;
    auth_proc->amf_com_proc.amf_proc.not_delivered_ho        = NULL;
    auth_proc->amf_com_proc.amf_proc.base_proc.success_notif = success;
    auth_proc->amf_com_proc.amf_proc.base_proc.failure_notif = failure;
    auth_proc->amf_com_proc.amf_proc.base_proc.abort = amf_authentication_abort;
    auth_proc->amf_com_proc.amf_proc.base_proc.fail_in = NULL;  // only response
    // TODO Negative Scenarios to be taken in future.
    auth_proc->amf_com_proc.amf_proc.base_proc.time_out = NULL;
    if (!IS_AMF_CTXT_VALID_AUTH_VECTORS(amf_context)) {
      // Upper layer to fetch new security context
      nas5g_auth_info_proc_t* auth_info_proc =
          get_nas5g_cn_procedure_auth_info(amf_context);
      if (!auth_info_proc) {
        auth_info_proc = nas5g_cn_auth_info_procedure(amf_context);
      }
      if (!auth_info_proc->request_sent) {
        run_auth_info_proc = true;
      }
      rc = RETURNok;
    } else {
      if (amf_context->_security.eksi < KSI_NO_KEY_AVAILABLE) {
        eksi = (amf_context->_security.eksi + 1) % (EKSI_MAX_VALUE + 1);
      }
      for (; eksi < MAX_5G_AUTH_VECTORS; eksi++) {
        if (IS_AMF_CTXT_VALID_AUTH_VECTOR(
                amf_context, (eksi % MAX_5G_AUTH_VECTORS))) {
          break;
        }
      }
      // eksi should always be 0
      if (!IS_AMF_CTXT_VALID_AUTH_VECTOR(
              amf_context, (eksi % MAX_5G_AUTH_VECTORS))) {
        run_auth_info_proc = true;
      } else {
        rc = amf_proc_authentication_ksi(
            amf_context, amf_specific_proc, eksi,
            amf_context->_vector[eksi % MAX_5G_AUTH_VECTORS].rand,
            amf_context->_vector[eksi % MAX_5G_AUTH_VECTORS].autn, success,
            failure);
        OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
      }
    }
    if (run_auth_info_proc) {
      rc = start_authentication_information_procedure(
          amf_context, auth_proc, NULL);
    }
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

/****************************************************************************
 **                                                                        **
 ** Name:    amf_proc_authentication_ksi()                                 **
 **                                                                        **
 ** Description: Initiates authentication procedure to establish partial   **
 **      native 5G CN security context in the UE and the AMF.              **
 **                                                                        **
 **              3GPP TS 24.501, section 5.4.1.3                           **
 **      The network initiates the authentication procedure by             **
 **      sending an AUTHENTICATION REQUEST message to the UE and           **
 **      starting the timer T3560. The AUTHENTICATION REQUEST mes-         **
 **      sage contains the parameters necessary to calculate the           **
 **      authentication response.                                          **
 **                                                                        **
 ** Inputs:  ue_id:      UE lower layer identifier                         **
 **      ksi:       NAS key set identifier                                 **
 **      rand:      Random challenge number                                **
 **      autn:      Authentication token                                   **
 **      success:   Callback function executed when the authen-            **
 **             tication procedure successfully completes                  **
 **      reject:    Callback function executed when the authen-            **
 **             tication procedure fails or is rejected                    **
 **      failure:   Callback function executed whener a lower              **
 **             layer failure occured before the authenti-                 **
 **             cation procedure comnpletes                                **
 **      Others:    None                                                   **
 **                                                                        **
 ** Outputs:     None                                                      **
 **      Return:    RETURNok, RETURNerror                                  **
 **      Others:    None                                                   **
 **                                                                        **
 ***************************************************************************/
int amf_proc_authentication_ksi(
    amf_context_t* amf_context,
    nas_amf_specific_proc_t* const amf_specific_proc, ksi_t ksi,
    const uint8_t* const rand, const uint8_t* const autn, success_cb_t success,
    failure_cb_t failure) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int rc = RETURNerror;
  nas5g_amf_auth_proc_t* auth_proc;
  amf_ue_ngap_id_t ue_id;
  if ((amf_context) && ((AMF_DEREGISTERED == amf_context->amf_fsm_state) ||
                        (AMF_REGISTERED == amf_context->amf_fsm_state))) {
    ue_id = PARENT_STRUCT(amf_context, ue_m5gmm_context_s, amf_context)
                ->amf_ue_ngap_id;
    OAILOG_DEBUG(
        LOG_NAS_AMF,
        "ue_id= " AMF_UE_NGAP_ID_FMT
        " AMF-PROC  - Initiate Authentication KSI = %d\n",
        ue_id, ksi);
    auth_proc = get_nas5g_common_procedure_authentication(amf_context);
    if (!auth_proc) {
      auth_proc = nas5g_new_authentication_procedure(amf_context);
    }
    if (auth_proc) {
      if (AMF_SPEC_PROC_TYPE_REGISTRATION == amf_specific_proc->type)
        auth_proc->is_cause_is_registered = true;
    }
    // Set the RAND value
    auth_proc->ksi = ksi;
    if (rand) {
      memcpy(auth_proc->rand, rand, AUTH_RAND_SIZE);
    }
    // Set the authentication token
    if (autn) {
      memcpy(auth_proc->autn, autn, AUTH_AUTN_SIZE);
    }
    auth_proc->amf_cause            = AMF_CAUSE_SUCCESS;
    auth_proc->retransmission_count = 0;
    auth_proc->ue_id                = ue_id;
    ((nas5g_base_proc_t*) auth_proc)->parent =
        (nas5g_base_proc_t*) amf_specific_proc;
    auth_proc->amf_com_proc.amf_proc.delivered               = NULL;
    auth_proc->amf_com_proc.amf_proc.base_proc.success_notif = success;
    auth_proc->amf_com_proc.amf_proc.base_proc.failure_notif = failure;
    auth_proc->amf_com_proc.amf_proc.base_proc.abort = amf_authentication_abort;
    auth_proc->amf_com_proc.amf_proc.base_proc.fail_in = NULL;
  }

  /*
   * Send authentication request message to the UE
   */
  rc = amf_send_authentication_request(amf_context, auth_proc);

  if (rc != RETURNerror) {
    /*
     * Notify AMF that common procedure has been initiated
     */
    amf_sap_t amf_sap;
    amf_sap.primitive       = AMFREG_COMMON_PROC_REQ;
    amf_sap.u.amf_reg.ue_id = ue_id;
    amf_sap.u.amf_reg.ctx   = amf_context;
    rc                      = amf_sap_send(&amf_sap);
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

/****************************************************************************
 **                                                                        **
 ** Name:    amf_proc_authentication_complete()                            **
 **                                                                        **
 ** Description: Performs the authentication completion procedure executed **
 **      by the network.                                                   **
 **                                                                        **
 **              3GPP TS 24.501, section 5.4.1.3.4                         **
 **      Upon receiving the AUTHENTICATION RESPONSE message, the           **
 **      MME shall stop timer T3560 and check the correctness of           **
 **      the RES parameter.                                                **
 **                                                                        **
 ** Inputs:  ue_id:      UE lower layer identifier                         **
 **      emm_cause: Authentication failure AMF cause code                  **
 **      res:       Authentication response parameter. or auts             **
 **                 in case of sync failure                                **
 **      Others:    None                                                   **
 **                                                                        **
 ** Outputs:     None                                                      **
 **      Return:    RETURNok, RETURNerror                                  **
 **      Others:    amf_data, T3560                                        **
 **                                                                        **
 ***************************************************************************/
int amf_proc_authentication_complete(
    amf_ue_ngap_id_t ue_id, AuthenticationResponseMsg* msg, int amf_cause,
    const unsigned char* res) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int rc = RETURNerror;
  int idx;
  nas_amf_smc_proc_t nas_amf_smc_proc_autn;
  OAILOG_DEBUG(
      LOG_NAS_AMF,
      "Authentication  procedures complete for "
      "(ue_id=" AMF_UE_NGAP_ID_FMT ")\n",
      ue_id);
  ue_m5gmm_context_s* ue_mm_context = NULL;

  amf_context_t* amf_ctx = NULL;
  ue_mm_context          = amf_ue_context_exists_amf_ue_ngap_id(ue_id);

  if (!ue_mm_context) {
    OAILOG_WARNING(
        LOG_NAS_AMF,
        "AMF-PROC - Failed to authenticate the UE due to NULL"
        "ue_mm_context\n");
    OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
  }

  amf_ctx = &ue_mm_context->amf_context;
  nas5g_amf_auth_proc_t* auth_proc =
      get_nas5g_common_procedure_authentication(amf_ctx);

  if (auth_proc) {
    nas_amf_smc_proc_autn.amf_ctx_set_security_eksi(amf_ctx, auth_proc->ksi);

    for (idx = 0; idx < amf_ctx->_vector[auth_proc->ksi].xres_size; idx++) {
      if ((amf_ctx->_vector[auth_proc->ksi].xres[idx]) !=
          msg->autn_response_parameter.response_parameter[idx]) {
        break;
      }
    }

    OAILOG_DEBUG(LOG_NAS_AMF, "Authentication of the UE is Successful\n");

    /*
     * Notify AMF that the authentication procedure successfully completed
     */
    amf_sap_t amf_sap;
    amf_sap.primitive               = AMFREG_COMMON_PROC_CNF;
    amf_sap.u.amf_reg.ue_id         = ue_id;
    amf_sap.u.amf_reg.ctx           = amf_ctx;
    amf_sap.u.amf_reg.notify        = true;
    amf_sap.u.amf_reg.free_proc     = true;
    amf_sap.u.amf_reg.u.common_proc = &auth_proc->amf_com_proc;
    rc                              = amf_sap_send(&amf_sap);
  } else {
    OAILOG_ERROR(LOG_NAS_AMF, "Auth proc is null");
  }
  /* Completing Authentication response and invoking Security Request
   * Invoking success directly to handle security mode command
   * */
  rc = amf_registration_success_authentication_cb(amf_ctx);
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

/****************************************************************************
 **                                                                        **
 ** Name:    amf_send_authentication_request()                             **
 **                                                                        **
 ** Description: Sends AUTHENTICATION REQUEST message and start timer T3560**
 **                                                                        **
 ** Inputs:  args: pointer to amf context                                  **
 **                handler parameters                                      **
 **                                                                        **
 ** Outputs:     None                                                      **
 **      Return:    RETURNok, RETURNerror                                  **
 **                                                                        **
 ***************************************************************************/
int amf_send_authentication_request(
    amf_context_t* amf_ctx, nas5g_amf_auth_proc_t* auth_proc) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int rc = RETURNerror;

  if (auth_proc) {
    /*
     * Notify AMF-AS SAP that Authentication Request message has to be sent
     * to the UE
     */
    amf_sap_t amf_sap;
    amf_sap.primitive = AMFAS_SECURITY_REQ;
    amf_sap.u.amf_as.u.security.puid =
        auth_proc->amf_com_proc.amf_proc.base_proc.nas_puid;
    amf_sap.u.amf_as.u.security.guti     = {0};
    amf_sap.u.amf_as.u.security.ue_id    = auth_proc->ue_id;
    amf_sap.u.amf_as.u.security.msg_type = AMF_AS_MSG_TYPE_AUTH;
    amf_sap.u.amf_as.u.security.ksi      = auth_proc->ksi;
    memcpy(amf_sap.u.amf_as.u.security.rand, auth_proc->rand, AUTH_RAND_SIZE);
    memcpy(amf_sap.u.amf_as.u.security.autn, auth_proc->autn, AUTH_AUTN_SIZE);

    /*
     * Setup 5GCN NAS security data
     */
    amf_data_sec_auth.amf_as_set_security_data(
        &amf_sap.u.amf_as.u.security.sctx, &amf_ctx->_security, false, true);
    rc = amf_sap_send(&amf_sap);
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

}  // namespace magma5g
