/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

/*! \file mme_app_sgs_paging.c
   \brief Handles  SGSAP Paging Request message
   \author
   \version
   \company
   \email:
*/

#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stdint.h>

#include "assertions.h"
#include "conversions.h"
#include "log.h"
#include "service303.h"
#include "intertask_interface.h"
#include "mme_app_defs.h"
#include "mme_app_sgs_fsm.h"
#include "mme_app_ue_context.h"
#include "3gpp_23.003.h"
#include "3gpp_36.401.h"
#include "bstrlib.h"
#include "common_defs.h"
#include "common_types.h"
#include "emm_data.h"
#include "intertask_interface_types.h"
#include "itti_types.h"
#include "mme_app_state.h"
#include "nas_messages_types.h"
#include "s1ap_messages_types.h"
#include "sgs_messages_types.h"

static int _mme_app_send_sgsap_ue_unreachable(
  struct ue_mm_context_s *ue_context_p,
  SgsCause_t sgs_cause);

static int _sgsap_handle_paging_request_without_lai(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP);

static int _sgs_handle_paging_request_for_mt_call(const sgs_fsm_t *evt);

static int _sgs_handle_paging_request_for_mt_call_in_connected(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP);

static int _sgs_handle_paging_request_for_mt_call_in_idle(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP);

static int _sgs_handle_paging_request_for_mt_sms(const sgs_fsm_t *evt);

static int _sgs_handle_paging_request_for_mt_sms_in_connected(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP);

static int _sgs_handle_paging_request_for_mt_sms_in_idle(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP);
/**********************************************************************************
 **                                                                              **
 ** Name:    sgs_handle_associated_paging_request()                              **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 ** Inputs:  sgs_fsm_t: pointer for sgs_fsm_primitive structure                  **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/
int sgs_handle_associated_paging_request(const sgs_fsm_t *evt)
{
  int rc = RETURNerror;
  itti_sgsap_paging_request_t *sgsap_paging_req_pP = NULL;

  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(evt);

  OAILOG_DEBUG(
    LOG_MME_APP,
    "Handle paging request in Associated state for ue-id :%u \n",
    evt->ue_id);
  sgs_context_t *sgs_context = (sgs_context_t *) evt->ctx;
  sgsap_paging_req_pP = (itti_sgsap_paging_request_t *) sgs_context->sgsap_msg;

#define SGSAP_SMS_INDICATOR 0x02
  if (sgsap_paging_req_pP->service_indicator == SGSAP_SMS_INDICATOR) {
    rc = _sgs_handle_paging_request_for_mt_sms(evt);
  } else {
    rc = _sgs_handle_paging_request_for_mt_call(evt);
  }
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgs_handle_paging_request_for_mt_sms()                            **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 **               for Mobile terminating sms                                    **
 ** Inputs:  sgs_fsm_t: pointer for sgs_fsm_primitive structure                  **
 **          ue_context_p: UE context
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

static int _sgs_handle_paging_request_for_mt_sms(const sgs_fsm_t *evt)
{
  int rc = RETURNerror;
  ue_mm_context_t *ue_context_p = NULL;
  sgs_context_t *sgs_context = NULL;
  itti_sgsap_paging_request_t *sgsap_paging_req_pP = NULL;
  imsi64_t imsi64 = INVALID_IMSI64;
  OAILOG_FUNC_IN(LOG_MME_APP);

  mme_app_desc_t *mme_app_desc_p = get_mme_nas_state(false);
  ue_context_p = mme_ue_context_exists_mme_ue_s1ap_id(
    &mme_app_desc_p->mme_ue_contexts, evt->ue_id);
  if (!ue_context_p) {
    OAILOG_WARNING(
      LOG_MME_APP,
      "Received paging request- UE context not found for ue-id :%u \n",
      evt->ue_id);
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }
  sgs_context = (sgs_context_t *) evt->ctx;
  sgsap_paging_req_pP = (itti_sgsap_paging_request_t *) sgs_context->sgsap_msg;
  if (ue_context_p->granted_service == GRANTED_SERVICE_EPS_ONLY) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "Send SGSAP-Pagaing Reject due to Service mis-match"
      "requested service :%u granted service to UE :%u for imsi:" IMSI_64_FMT
      "\n",
      sgsap_paging_req_pP->service_indicator,
      (uint8_t) ue_context_p->granted_service,
      ue_context_p->emm_context._imsi64);
    IMSI_STRING_TO_IMSI64((char *) sgsap_paging_req_pP->imsi, &imsi64);
    mme_app_send_sgsap_paging_reject(
      ue_context_p,
      imsi64,
      sgsap_paging_req_pP->imsi_length,
      SGS_CAUSE_IMSI_IMPLICITLY_DETACHED_FOR_NONEPS_SERVICE);
    increment_counter(
      "sgsap_paging_reject", 1, 1, "cause", "ue_requested_only_eps");
    unlock_ue_contexts(ue_context_p);
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  /* Fetch LAI if present */
  OAILOG_DEBUG(
    LOG_MME_APP,
    " : LAI : %d\n",
    (sgsap_paging_req_pP->presencemask & PAGING_REQUEST_LAI_PARAMETER_PRESENT));
  if (!(sgsap_paging_req_pP->presencemask &
        PAGING_REQUEST_LAI_PARAMETER_PRESENT)) {
    rc = _sgsap_handle_paging_request_without_lai(
      ue_context_p, sgsap_paging_req_pP);
    unlock_ue_contexts(ue_context_p);
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }
  if (ue_context_p->ecm_state == ECM_CONNECTED) {
    rc = _sgs_handle_paging_request_for_mt_sms_in_connected(
      ue_context_p, sgsap_paging_req_pP);
  } else if (ue_context_p->ecm_state == ECM_IDLE) {
    rc = _sgs_handle_paging_request_for_mt_sms_in_idle(
      ue_context_p, sgsap_paging_req_pP);
  }
  unlock_ue_contexts(ue_context_p);
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgs_handle_paging_request_for_mt_call()                            **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 **               for Mobile terminating call                                    **
 ** Inputs:  sgs_fsm_t: pointer for sgs_fsm_primitive structure                  **
 **          ue_context_p: UE context
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

static int _sgs_handle_paging_request_for_mt_call(const sgs_fsm_t *evt)
{
  int rc = RETURNerror;
  ue_mm_context_t *ue_context_p = NULL;
  sgs_context_t *sgs_context = NULL;
  itti_sgsap_paging_request_t *sgsap_paging_req_pP = NULL;
  imsi64_t imsi64 = INVALID_IMSI64;
  OAILOG_FUNC_IN(LOG_MME_APP);

  mme_app_desc_t *mme_app_desc_p = get_mme_nas_state(false);
  ue_context_p = mme_ue_context_exists_mme_ue_s1ap_id(
    &mme_app_desc_p->mme_ue_contexts, evt->ue_id);
  if (!ue_context_p) {
    OAILOG_WARNING(
      LOG_MME_APP,
      "Received paging request- UE context not found for ue-id :%u \n",
      evt->ue_id);
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }
  sgs_context = (sgs_context_t *) evt->ctx;
  /*If call_cancelled is set to TRUE when a new Paging message is received.Set call_cancelled to false*/
  if (sgs_context->call_cancelled == true) {
    sgs_context->call_cancelled = false;
  }
  sgsap_paging_req_pP = (itti_sgsap_paging_request_t *) sgs_context->sgsap_msg;
  IMSI_STRING_TO_IMSI64((char *) sgsap_paging_req_pP->imsi, &imsi64);
  if (ue_context_p->granted_service != GRANTED_SERVICE_CSFB_SMS) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "Send SGSAP-Pagaing Reject due to Service mis-match"
      "requested service :%u granted service to UE :%u for imsi:" IMSI_64_FMT
      "\n",
      sgsap_paging_req_pP->service_indicator,
      (uint8_t) ue_context_p->granted_service,
      ue_context_p->emm_context._imsi64);
    mme_app_send_sgsap_paging_reject(
      ue_context_p,
      imsi64,
      ue_context_p->emm_context._imsi.length,
      SGS_CAUSE_MT_CSFB_CALL_REJECTED_BY_USER);
    increment_counter(
      "sgsap_paging_reject", 1, 1, "cause", "ue_requested_only_sms");
    unlock_ue_contexts(ue_context_p);
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  /* Fetch LAI if present */
  if (!(sgsap_paging_req_pP->presencemask &
        PAGING_REQUEST_LAI_PARAMETER_PRESENT)) {
    rc = _sgsap_handle_paging_request_without_lai(
      ue_context_p, sgsap_paging_req_pP);
    unlock_ue_contexts(ue_context_p);
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }

  /* Check the vlr-reliable flag */
  if (sgs_context->vlr_reliable == false) {
    OAILOG_DEBUG(
      LOG_MME_APP,
      "Received Paging Request while vlr-rliable is :%d for imsi" IMSI_64_FMT
      "\n",
      sgs_context->vlr_reliable,
      ue_context_p->emm_context._imsi64);
    /* Handling for paging received without LAI and vlr-reliable flag set to false is same */
    rc = _sgsap_handle_paging_request_without_lai(
      ue_context_p, sgsap_paging_req_pP);
    unlock_ue_contexts(ue_context_p);
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }

  if (ue_context_p->ecm_state == ECM_CONNECTED) {
    rc = _sgs_handle_paging_request_for_mt_call_in_connected(
      ue_context_p, sgsap_paging_req_pP);
  } else if (ue_context_p->ecm_state == ECM_IDLE) {
    rc = _sgs_handle_paging_request_for_mt_call_in_idle(
      ue_context_p, sgsap_paging_req_pP);
  }
  unlock_ue_contexts(ue_context_p);
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    mme_app_send_nas_detach_request                                     **
 ** Description   Build and send nas detach request                              **
 ** Inputs:  ue_id       : ue identity                                           **
 **          detach_type : Network detach type                                   **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **
***********************************************************************************/
int mme_app_send_nas_detach_request(mme_ue_s1ap_id_t ue_id, uint8_t detach_type)
{
  int rc = RETURNerror;
  MessageDef *message_p = NULL;
  itti_nas_nw_initiated_detach_ue_req_t *sgsap_nas_detach_pP = NULL;
  OAILOG_FUNC_IN(LOG_MME_APP);

  message_p =
    itti_alloc_new_message(TASK_MME_APP, NAS_NW_INITIATED_DETACH_UE_REQ);
  AssertFatal(message_p, "itti_alloc_new_message Failed");
  sgsap_nas_detach_pP = &message_p->ittiMsg.nas_nw_initiated_detach_ue_req;
  memset(
    (void *) sgsap_nas_detach_pP,
    0,
    sizeof(itti_nas_nw_initiated_detach_ue_req_t));

  sgsap_nas_detach_pP->ue_id = ue_id;
  sgsap_nas_detach_pP->detach_type = detach_type;
  rc = itti_send_msg_to_task(TASK_NAS_MME, INSTANCE_DEFAULT, message_p);

  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgs_handle_paging_request_for_mt_call_in_connected()               **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 **               and UE connected state for Mobile terminating call             **
 ** Inputs:  ue_context_p: UE context                                            **
 **          itti_sgsap_paging_request_t : received sgs-paging request           **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

static int _sgs_handle_paging_request_for_mt_call_in_connected(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP)
{
  int rc = RETURNerror;
  uint8_t paging_id = NAS_PAGING_ID_IMSI;
  bstring cli = NULL;
  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(ue_context_p);
  DevAssert(sgsap_paging_req_pP);

  OAILOG_INFO(
    LOG_MME_APP,
    "Received SGSAP-Paging Request in UE Connected state for IMSI:" IMSI_64_FMT
    "\n",
    ue_context_p->emm_context._imsi64);

  /* Fetch TMSI if present */
  if (
    sgsap_paging_req_pP->presencemask & PAGING_REQUEST_TMSI_PARAMETER_PRESENT) {
    paging_id = NAS_PAGING_ID_TMSI;
  }
  /* Fetch CLI if present */
  if (
    sgsap_paging_req_pP->presencemask & PAGING_REQUEST_CLI_PARAMETER_PRESENT) {
    bassign(cli, sgsap_paging_req_pP->opt_cli);
  }
  if (
    RETURNok != (rc = mme_app_send_nas_cs_service_notification(
                   ue_context_p->mme_ue_s1ap_id, paging_id, cli))) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "Failed to send CS-Service Notification to NAS for ue-id :%u \n",
      ue_context_p->mme_ue_s1ap_id);
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  if (
    RETURNok != (rc = mme_app_send_sgsap_service_request(
                   sgsap_paging_req_pP->service_indicator, ue_context_p))) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "Failed to send CS-Service Request to SGS-Task for ue-id :%u \n",
      ue_context_p->mme_ue_s1ap_id);
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  ue_context_p->sgs_context->csfb_service_type = CSFB_SERVICE_MT_CALL;
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgs_handle_paging_request_for_mt_sms_in_connected()               **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 **               and UE connected state for Mobile terminating sms             **
 ** Inputs:  ue_context_p: UE context                                            **
 **          itti_sgsap_paging_request_t : received sgs-paging request           **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

static int _sgs_handle_paging_request_for_mt_sms_in_connected(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP)
{
  int rc = RETURNerror;

  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(ue_context_p);
  DevAssert(sgsap_paging_req_pP);

  OAILOG_INFO(
    LOG_MME_APP,
    "Received SGSAP-Paging Request in UE Connected state for IMSI:" IMSI_64_FMT
    "\n",
    ue_context_p->emm_context._imsi64);

  if (
    RETURNok != (rc = mme_app_send_sgsap_service_request(
                   sgsap_paging_req_pP->service_indicator, ue_context_p))) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "Failed to send CS-Service Request to SGS-Task for ue-id :%u \n",
      ue_context_p->mme_ue_s1ap_id);
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  ue_context_p->sgs_context->csfb_service_type = CSFB_SERVICE_MT_SMS;
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgs_handle_paging_request_for_mt_call_in_idle                      **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 **               and UE idle state for Mobile terminating call                  **
 ** Inputs:  ue_context_p: UE context                                            **
 **          itti_sgsap_paging_request_t : received sgs-paging request           **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

static int _sgs_handle_paging_request_for_mt_call_in_idle(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP)

{
  int rc = RETURNerror;
  uint8_t paging_id = NAS_PAGING_ID_IMSI;
  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(ue_context_p);
  DevAssert(sgsap_paging_req_pP);

  OAILOG_INFO(
    LOG_MME_APP,
    "Received SGSAP-Paging Request in UE Idle state for IMSI:" IMSI_64_FMT "\n",
    ue_context_p->emm_context._imsi64);
  if (ue_context_p->ppf) {
    /* Paging timer shall not be started, if paging procedure initiated for CSFB
     * Reference: spec-24.301 section: 5.6.2.3
    */
    if (!IS_EMM_CTXT_PRESENT_GUTI(&(ue_context_p->emm_context))) {
      /* On reception of SGS-Paging request in idle and not able retrieve S-TMSI from IMSI,
       * page with IMSI and PS domain
       */
      OAILOG_INFO(
        LOG_MME_APP,
        "Received SGS-paging request Unable to retrieve S-TMSI from "
        "IMSI " IMSI_64_FMT "\n",
        ue_context_p->emm_context._imsi64);
      rc = mme_app_paging_request_helper(
        ue_context_p, false, NAS_PAGING_ID_IMSI, CN_DOMAIN_PS);
    } else {
      /* Fetch TMSI if present */
      if (
        sgsap_paging_req_pP->presencemask &
        PAGING_REQUEST_TMSI_PARAMETER_PRESENT) {
        paging_id = NAS_PAGING_ID_TMSI;
      }
      /* if TMSI is received, then page with S-TMSI otherwise page with IMSI */
      if (
        (rc = mme_app_paging_request_helper(
           ue_context_p, false, paging_id, CN_DOMAIN_CS)) != RETURNok) {
        OAILOG_ERROR(
          LOG_MME_APP,
          "Failed to send PAGING Message to UE for UE-id:%u \n",
          ue_context_p->mme_ue_s1ap_id);
      }
      ue_context_p->sgs_context->csfb_service_type = CSFB_SERVICE_MT_CALL;
      ue_context_p->sgs_context->service_indicator =
        sgsap_paging_req_pP->service_indicator;
    }
  } else {
    /* Send UE Unreachable to MSC/VLR */
    if (
      RETURNok != (rc = _mme_app_send_sgsap_ue_unreachable(
                     ue_context_p, SGS_CAUSE_UE_UNREACHABLE))) {
      OAILOG_ERROR(
        LOG_MME_APP,
        "Failed to send SGSAP-UE-UNREACHABLE for ue-id :%u \n",
        ue_context_p->mme_ue_s1ap_id);
    }
  }
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgs_handle_paging_request_for_mt_sms_in_idle                      **
 ** Description   Handle SGSAP-Paging request in SGS-Associated state            **
 **               and UE idle state for Mobile terminating sms                  **
 ** Inputs:  ue_context_p: UE context                                            **
 **          itti_sgsap_paging_request_t : received sgs-paging request           **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

static int _sgs_handle_paging_request_for_mt_sms_in_idle(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP)

{
  int rc = RETURNerror;
  uint8_t paging_id = NAS_PAGING_ID_IMSI;
  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(ue_context_p);
  DevAssert(sgsap_paging_req_pP);

  OAILOG_INFO(
    LOG_MME_APP,
    "Received SGSAP-Paging Request in UE Idle state for IMSI:" IMSI_64_FMT "\n",
    ue_context_p->emm_context._imsi64);
  if (ue_context_p->ppf) {
    /* Paging timer shall not be started, if paging procedure initiated for CSFB
     * Reference: spec-24.301 section: 5.6.2.3
    */
    if (!IS_EMM_CTXT_PRESENT_GUTI(&(ue_context_p->emm_context))) {
      /* On reception of SGS-Paging request in idle and not able retrieve S-TMSI from IMSI,
       * page with IMSI and PS domain
       */
      OAILOG_INFO(
        LOG_MME_APP,
        "Received SGS-paging request Unable to retrieve S-TMSI from "
        "IMSI " IMSI_64_FMT "\n",
        ue_context_p->emm_context._imsi64);
      rc = mme_app_paging_request_helper(
        ue_context_p, false, NAS_PAGING_ID_IMSI, CN_DOMAIN_PS);
    } else {
      /* Fetch TMSI if present */
      if (
        sgsap_paging_req_pP->presencemask &
        PAGING_REQUEST_TMSI_PARAMETER_PRESENT) {
        paging_id = NAS_PAGING_ID_TMSI;
      }
      /* if TMSI is received, then page with S-TMSI otherwise page with IMSI */
      if (
        (rc = mme_app_paging_request_helper(
           ue_context_p, false, paging_id, CN_DOMAIN_PS)) != RETURNok) {
        OAILOG_ERROR(
          LOG_MME_APP,
          "Failed to send PAGING Message to UE for UE-id:%u \n",
          ue_context_p->mme_ue_s1ap_id);
      }
      ue_context_p->sgs_context->csfb_service_type = CSFB_SERVICE_MT_SMS;
      ue_context_p->sgs_context->service_indicator =
        sgsap_paging_req_pP->service_indicator;
    }
  } else {
    /* Send UE Unreachable to MSC/VLR */
    if (
      RETURNok != (rc = _mme_app_send_sgsap_ue_unreachable(
                     ue_context_p, SGS_CAUSE_UE_UNREACHABLE))) {
      OAILOG_ERROR(
        LOG_MME_APP,
        "Failed to send SGSAP-UE-UNREACHABLE for ue-id :%u \n",
        ue_context_p->mme_ue_s1ap_id);
    }
  }
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    mme_app_send_sgsap_service_request()                                **
 ** Description  Build and Send Service Request to MSC/VLR                       **
 ** Inputs:                                                                      **
 **          service-indicator   Indicates type services: SMS or CS-CALL         **
 **          ue_context_p:  pointer to UE context                                **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/

int mme_app_send_sgsap_service_request(
  uint8_t service_indicator,
  struct ue_mm_context_s *ue_context_p)
{
  int rc = RETURNerror;
  MessageDef *message_p = NULL;
  itti_sgsap_service_request_t *sgsap_service_req_pP = NULL;

  OAILOG_FUNC_IN(LOG_MME_APP);
  message_p = itti_alloc_new_message(TASK_MME_APP, SGSAP_SERVICE_REQUEST);
  AssertFatal(message_p, "itti_alloc_new_message Failed");
  sgsap_service_req_pP = &message_p->ittiMsg.sgsap_service_request;
  memset(
    (void *) sgsap_service_req_pP, 0, sizeof(itti_sgsap_service_request_t));

  IMSI64_TO_STRING(
    ue_context_p->emm_context._imsi64, sgsap_service_req_pP->imsi,
    ue_context_p->emm_context._imsi.length);
  sgsap_service_req_pP->imsi_length = ue_context_p->emm_context._imsi.length;
  sgsap_service_req_pP->service_indicator = service_indicator;
  if (IS_EMM_CTXT_PRESENT_IMEISV(&(ue_context_p->emm_context))) {
    sgsap_service_req_pP->presencemask |=
      SERVICE_REQUEST_IMEISV_PARAMETER_PRESENT;
    hexa_to_ascii(
      (uint8_t *) ue_context_p->emm_context._imeisv.u.value,
      sgsap_service_req_pP->opt_imeisv,
      8);
    sgsap_service_req_pP->opt_imeisv[ue_context_p->emm_context._imeisv.length]
    = '\0';
    sgsap_service_req_pP->opt_imeisv_length = ue_context_p->emm_context._imeisv.length;
  }
  sgsap_service_req_pP->opt_ecgi = ue_context_p->e_utran_cgi;
  sgsap_service_req_pP->presencemask |= SERVICE_REQUEST_ECGI_PARAMETER_PRESENT;
  sgsap_service_req_pP->opt_ue_emm_mode = ue_context_p->ecm_state;
  sgsap_service_req_pP->presencemask |=
    SERVICE_REQUEST_UE_EMM_MODE_PARAMETER_PRESENT;
  /* TODO - Add other optional information like ue_time_zone, mobilestationclassmark2, tai in sgs  service request */
  OAILOG_INFO(
    LOG_MME_APP,
    "Send SGSAP-Service Request for IMSI " IMSI_64_FMT "\n",
    ue_context_p->emm_context._imsi64);
  rc = itti_send_msg_to_task(TASK_SGS, INSTANCE_DEFAULT, message_p);
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    mme_app_send_nas_cs_service_notification()                          **
 ** Description    If SGSAP-Paging request received in UE ECM_Connected state    **
 **                Send CS Service Notofication                                  **
 ** Inputs:  ue_id:      UE identifier                                           **
 **          paging_id   Indicates the identity used for paging non-eps services **
 **          cli         Calling Line Identification                             **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **                                                                              **
***********************************************************************************/
int mme_app_send_nas_cs_service_notification(
  mme_ue_s1ap_id_t ue_id,
  uint8_t paging_id,
  bstring cli)
{
  int rc = RETURNerror;
  MessageDef *message_p = NULL;
  itti_nas_cs_service_notification_t *cs_service_notification_p = NULL;

  OAILOG_FUNC_IN(LOG_MME_APP);
  message_p = itti_alloc_new_message(TASK_MME_APP, NAS_CS_SERVICE_NOTIFICATION);
  AssertFatal(message_p, "itti_alloc_new_message Failed");
  cs_service_notification_p = &message_p->ittiMsg.nas_cs_service_notification;
  memset(
    (void *) cs_service_notification_p,
    0,
    sizeof(itti_nas_cs_service_notification_t));

  cs_service_notification_p->ue_id = ue_id;
  cs_service_notification_p->paging_id = paging_id;
  bassign(cs_service_notification_p->cli, cli);

  OAILOG_INFO(
    LOG_MME_APP,
    "Send NAS CS Service Notification from MME app for ue_id:%u \n",
    ue_id);
  rc = itti_send_msg_to_task(TASK_NAS_MME, INSTANCE_DEFAULT, message_p);

  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    mme_app_send_sgsap_paging_reject()                                  **
 ** Description   Build and send Paging reject                                   **
 ** Inputs:  ue_context_p: pointer ue_context                                    **
 **          imsi        : imsi                                                  **
 **          sgs_cause   : paging reject cause                                   **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **
***********************************************************************************/
int mme_app_send_sgsap_paging_reject(
  struct ue_mm_context_s *ue_context_p,
  imsi64_t imsi,
  uint8_t imsi_len,
  SgsCause_t sgs_cause)
{
  int rc = RETURNerror;
  MessageDef *message_p = NULL;
  itti_sgsap_paging_reject_t *sgsap_paging_reject_pP = NULL;
  OAILOG_FUNC_IN(LOG_MME_APP);

  message_p = itti_alloc_new_message(TASK_MME_APP, SGSAP_PAGING_REJECT);
  AssertFatal(message_p, "itti_alloc_new_message Failed");
  sgsap_paging_reject_pP = &message_p->ittiMsg.sgsap_paging_reject;
  memset(
    (void *) sgsap_paging_reject_pP, 0, sizeof(itti_sgsap_paging_reject_t));

  //IMSI
  if (imsi) {
    IMSI64_TO_STRING(imsi, sgsap_paging_reject_pP->imsi, imsi_len);
  } else {
    OAILOG_ERROR(LOG_MME_APP, "Invalid imsi \n");
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }
  sgsap_paging_reject_pP->imsi_length = imsi_len;
  sgsap_paging_reject_pP->sgs_cause = sgs_cause;

  if (ue_context_p) {
    OAILOG_INFO(
      LOG_MME_APP,
      "Send SGSAP-Paging Reject for IMSI" IMSI_64_FMT " with sgs-cause :%d \n",
      ue_context_p->emm_context._imsi64,
      (int) sgs_cause);
  } else {
    OAILOG_INFO(
      LOG_MME_APP,
      "Send SGSAP-Paging Reject with sgs-cause :%d \n",
      (int) sgs_cause);
  }
  rc = itti_send_msg_to_task(TASK_SGS, INSTANCE_DEFAULT, message_p);
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    sgs_handle_null_paging_request()                                    **
 ** Description   Handle SGSAP-Paging request in SGS-NULL state                  **
 ** Inputs:  sgs_fsm_t: pointer for sgs_fsm_primitive structure                  **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **
***********************************************************************************/

int sgs_handle_null_paging_request(const sgs_fsm_t *evt)
{
  int rc = RETURNerror;
  struct ue_mm_context_s *ue_context_p = NULL;
  itti_sgsap_paging_request_t *sgsap_paging_req_pP = NULL;
  imsi64_t imsi64 = INVALID_IMSI64;

  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(evt);

  OAILOG_DEBUG(
    LOG_MME_APP,
    "Handle paging request in Null state for ue-id :%u \n",
    evt->ue_id);
  mme_app_desc_t *mme_app_desc_p = get_mme_nas_state(false);
  ue_context_p = mme_ue_context_exists_mme_ue_s1ap_id(
    &mme_app_desc_p->mme_ue_contexts, evt->ue_id);
  if (!ue_context_p) {
    OAILOG_WARNING(
      LOG_MME_APP,
      "Received paging request- UE context not found for ue-id :%u \n",
      evt->ue_id);
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }
  sgsap_paging_req_pP =
    (itti_sgsap_paging_request_t *) ue_context_p->sgs_context->sgsap_msg;
  IMSI_STRING_TO_IMSI64(sgsap_paging_req_pP->imsi, &imsi64);
  /* Send SGSAP-Paging reject, if SGSAP_paging request recived in NULL state */
  OAILOG_INFO(
    LOG_MME_APP,
    "Send SGSAP_Paging Reject for Paging Request received in"
    "SGS-NULL state for imsi: " IMSI_64_FMT "\n",
    ue_context_p->emm_context._imsi64);
  rc = mme_app_send_sgsap_paging_reject(
    ue_context_p,
    imsi64,
    sgsap_paging_req_pP->imsi_length,
    SGS_CAUSE_IMSI_DETACHED_FOR_NONEPS_SERVICE);
  increment_counter(
    "sgsap_paging_reject", 1, 1, "cause", "paging_request_rx in null_state");

  unlock_ue_contexts(ue_context_p);
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _mme_app_send_sgsap_ue_unreachable()                                 **
 ** Description   Build and send UE Unreachable                                  **
 ** Inputs:  ue_context_p: pointer to ue_context                                    **
 **          sgs_cause   : paging reject cause                                   **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **
***********************************************************************************/

static int _mme_app_send_sgsap_ue_unreachable(
  struct ue_mm_context_s *ue_context_p,
  SgsCause_t sgs_cause)
{
  int rc = RETURNerror;
  MessageDef *message_p = NULL;
  itti_sgsap_ue_unreachable_t *sgsap_ue_unreachable_pP = NULL;
  OAILOG_FUNC_IN(LOG_MME_APP);

  if (!ue_context_p) {
    OAILOG_WARNING(LOG_MME_APP, "Invalid Ue context \n");
    OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
  }

  message_p = itti_alloc_new_message(TASK_MME_APP, SGSAP_UE_UNREACHABLE);
  AssertFatal(message_p, "itti_alloc_new_message Failed");
  sgsap_ue_unreachable_pP = &message_p->ittiMsg.sgsap_ue_unreachable;
  memset(
    (void *) sgsap_ue_unreachable_pP, 0, sizeof(itti_sgsap_ue_unreachable_t));

  IMSI64_TO_STRING(
    ue_context_p->emm_context._imsi64, sgsap_ue_unreachable_pP->imsi,
    ue_context_p->emm_context._imsi.length);
  sgsap_ue_unreachable_pP->imsi_length =
    (uint8_t) strlen(sgsap_ue_unreachable_pP->imsi);
  sgsap_ue_unreachable_pP->sgs_cause = sgs_cause;

  OAILOG_INFO(
    LOG_MME_APP,
    "Send SGSAP-UE-unreachable for IMSI" IMSI_64_FMT " with sgs-cause :%d \n",
    ue_context_p->emm_context._imsi64,
    (int) sgs_cause);
  rc = itti_send_msg_to_task(TASK_SGS, INSTANCE_DEFAULT, message_p);

  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/**********************************************************************************
 **                                                                              **
 ** Name:    _sgsap_handle_paging_request_without_lai()                          **
 ** Description   Handles SGS-Paging request mesasage received without LAI       **
 ** Inputs:  ue_context_p: pointer ue_context                                    **
 **          itti_sgsap_paging_request_t : Received SGS-Paging request mesasage  **
 ** Outputs:                                                                     **
 **          Return:    RETURNok, RETURNerror                                    **
 **
***********************************************************************************/

static int _sgsap_handle_paging_request_without_lai(
  ue_mm_context_t *ue_context_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP)
{
  MessageDef *message_p = NULL;
  int rc = RETURNok;
  s1ap_cn_domain_t cn_domain = CN_DOMAIN_CS;
  uint8_t paging_id = NAS_PAGING_ID_IMSI;

  OAILOG_FUNC_IN(LOG_MME_APP);
  if (!ue_context_p) {
    OAILOG_ERROR(LOG_MME_APP, "Invalid ue_context_p \n");
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  if (!sgsap_paging_req_pP) {
    OAILOG_ERROR(LOG_MME_APP, "Null Pagaing Request Received \n");
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }

  OAILOG_INFO(
    LOG_MME_APP,
    "Handle sgsap-paging request received without LAI for IMSI " IMSI_64_FMT
    "\n",
    ue_context_p->emm_context._imsi64);
  if (ue_context_p->ecm_state == ECM_CONNECTED) {
    // Send N/W Initiated Detach Request to NAS
    message_p =
      itti_alloc_new_message(TASK_MME_APP, NAS_NW_INITIATED_DETACH_UE_REQ);
    AssertFatal(message_p, "itti_alloc_new_message Failed");
    message_p->ittiMsg.nas_nw_initiated_detach_ue_req.ue_id =
      ue_context_p->mme_ue_s1ap_id;
    message_p->ittiMsg.nas_nw_initiated_detach_ue_req.detach_type =
      SGS_INITIATED_IMSI_DETACH;
    rc = itti_send_msg_to_task(TASK_NAS_MME, INSTANCE_DEFAULT, message_p);
  } else if (ue_context_p->ecm_state == ECM_IDLE) {
    /* While UE is in ECM_IDLE and mobile reachability timer is still running
     * The value of ppf-paging proceeding flag will be "true"
     */
    if (ue_context_p->ppf) {
      /* if Paging request received without LAI for MT SMS, always page with S-TMSI */
      if (sgsap_paging_req_pP->service_indicator == SGSAP_SMS_INDICATOR) {
        paging_id = NAS_PAGING_ID_TMSI;
        cn_domain = CN_DOMAIN_PS;
      }
      /* if Paging request received without LAI for CS call, always page with IMSI */
      if (
        (rc = mme_app_paging_request_helper(
           ue_context_p, false, paging_id, cn_domain)) == RETURNok) {
        ue_context_p->sgs_context->csfb_service_type =
          CSFB_SERVICE_MT_CALL_OR_SMS_WITHOUT_LAI;
      }
    } else {
      /* Send UE Unreachable to MSC/VLR */
      if (
        RETURNok != (rc = _mme_app_send_sgsap_ue_unreachable(
                       ue_context_p, SGS_CAUSE_UE_UNREACHABLE))) {
        OAILOG_ERROR(
          LOG_MME_APP,
          "Failed to send SGSAP-UE-UNREACHABLE for ue-id :%u \n",
          ue_context_p->mme_ue_s1ap_id);
      }
    }
  }
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}

/****************************************************************************
 **                                                                        **
 ** Name:    mme_app_handle_sgsap_paging_request()                         **
 **                                                                        **
 ** Description: Processes the SGSAP Paging Request message re-            **
 **      ceived from the SGS task and invokes FSM handler based on state   **
 **                                                                        **
 ** Inputs:  itti_sgsap_paging_request_t: SGSAP Paging Request message     **
 **                                                                        **
 ** Outputs:                                                               **
 **      Return:    RETURNok, RETURNerror                                  **
 **                                                                        **
 ***************************************************************************/

int mme_app_handle_sgsap_paging_request(mme_app_desc_t *mme_app_desc_p,
  itti_sgsap_paging_request_t *const sgsap_paging_req_pP)
{
  struct ue_mm_context_s *ue_context_p = NULL;
  int rc = RETURNok;
  sgs_fsm_t sgs_fsm;
  imsi64_t imsi64 = INVALID_IMSI64;

  OAILOG_FUNC_IN(LOG_MME_APP);
  DevAssert(sgsap_paging_req_pP);

  IMSI_STRING_TO_IMSI64(sgsap_paging_req_pP->imsi, &imsi64);

  OAILOG_INFO(
    LOG_MME_APP,
    "Received SGS-PAGING REQUEST for IMSI " IMSI_64_FMT "\n",
    imsi64);
  if (
    (ue_context_p = mme_ue_context_exists_imsi(
       &mme_app_desc_p->mme_ue_contexts, imsi64)) == NULL) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "SGS-PAGING REQUEST: Failed to find UE context for IMSI " IMSI_64_FMT
      "\n",
      imsi64);
    mme_app_send_sgsap_paging_reject(
      NULL, imsi64, sgsap_paging_req_pP->imsi_length, SGS_CAUSE_IMSI_UNKNOWN);
    increment_counter("sgsap_paging_reject", 1, 1, "cause", "imsi_unknown");
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  if (ue_context_p->sgs_context == NULL) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "SGS context not created for IMSI " IMSI_64_FMT "\n",
      imsi64);
    mme_app_send_sgsap_paging_reject(
      NULL,
      imsi64,
      sgsap_paging_req_pP->imsi_length,
      SGS_CAUSE_IMSI_DETACHED_FOR_NONEPS_SERVICE);
    increment_counter(
      "sgsap_paging_reject", 1, 1, "cause", "SGS context not created");
    unlock_ue_contexts(ue_context_p);
    OAILOG_FUNC_RETURN(LOG_MME_APP, RETURNerror);
  }
  ue_context_p->sgs_context->sgsap_msg = (void *) sgsap_paging_req_pP;
  sgs_fsm.primitive = _SGS_PAGING_REQUEST;
  sgs_fsm.ue_id = ue_context_p->mme_ue_s1ap_id;
  sgs_fsm.ctx = (void *) ue_context_p->sgs_context;

  /* Invoke SGS FSM */
  if (RETURNok != (rc = sgs_fsm_process(&sgs_fsm))) {
    OAILOG_WARNING(
      LOG_MME_APP,
      "Failed  to execute SGS State machine for ue_id :%u \n",
      ue_context_p->mme_ue_s1ap_id);
  }
  ue_context_p->sgs_context->sgsap_msg = NULL;
  unlock_ue_contexts(ue_context_p);
  OAILOG_FUNC_RETURN(LOG_MME_APP, rc);
}
