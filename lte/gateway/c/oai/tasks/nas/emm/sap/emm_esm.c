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

#include "emm_esm.h"
#include "common_defs.h"
#include "log.h"
#include "LowerLayer.h"
#include "msc.h"

/****************************************************************************/
/****************  E X T E R N A L    D E F I N I T I O N S  ****************/
/****************************************************************************/

/****************************************************************************/
/*******************  L O C A L    D E F I N I T I O N S  *******************/
/****************************************************************************/

/*
   String representation of EMMESM-SAP primitives
*/
static const char *_emm_esm_primitive_str[] = {
  "EMMESM_RELEASE_IND",
  "EMMESM_UNITDATA_REQ",
  "EMMESM_ACTIVATE_BEARER_REQ",
  "EMMESM_UNITDATA_IND",
};

/****************************************************************************/
/******************  E X P O R T E D    F U N C T I O N S  ******************/
/****************************************************************************/

/****************************************************************************
 **                                                                        **
 ** Name:    emm_esm_initialize()                                      **
 **                                                                        **
 ** Description: Initializes the EMMESM Service Access Point               **
 **                                                                        **
 ** Inputs:  None                                                      **
 **      Others:    None                                       **
 **                                                                        **
 ** Outputs:     None                                                      **
 **      Return:    None                                       **
 **      Others:    NONE                                       **
 **                                                                        **
 ***************************************************************************/
void emm_esm_initialize(void)
{
  OAILOG_FUNC_IN(LOG_NAS_EMM);
  /*
   * TODO: Initialize the EMMESM-SAP
   */
  OAILOG_FUNC_OUT(LOG_NAS_EMM);
}

/****************************************************************************
 **                                                                        **
 ** Name:    emm_esm_send()                                            **
 **                                                                        **
 ** Description: Processes the EMMESM Service Access Point primitive       **
 **                                                                        **
 ** Inputs:  msg:       The EMMESM-SAP primitive to process        **
 **      Others:    None                                       **
 **                                                                        **
 ** Outputs:     None                                                      **
 **      Return:    RETURNok, RETURNerror                      **
 **      Others:    None                                       **
 **                                                                        **
 ***************************************************************************/
int emm_esm_send(const emm_esm_t *msg)
{
  OAILOG_FUNC_IN(LOG_NAS_EMM);
  int rc = RETURNerror;
  emm_esm_primitive_t primitive = msg->primitive;

  OAILOG_INFO(
    LOG_NAS_EMM,
    "EMMESM-SAP - Received primitive %s (%d)\n",
    _emm_esm_primitive_str[primitive - _EMMESM_START - 1],
    primitive);

  switch (primitive) {
    case _EMMESM_UNITDATA_REQ:
      MSC_LOG_RX_MESSAGE(
        MSC_NAS_EMM_MME,
        MSC_NAS_ESM_MME,
        NULL,
        0,
        "_EMMESM_UNITDATA_REQ ue id " MME_UE_S1AP_ID_FMT " ",
        msg->ue_id);
      /*
     * ESM requests EMM to transfer ESM data unit to lower layer
     */
      rc = lowerlayer_data_req(msg->ue_id, msg->u.data.msg);
      break;

    case _EMMESM_ACTIVATE_BEARER_REQ:
      MSC_LOG_RX_MESSAGE(
        MSC_NAS_EMM_MME,
        MSC_NAS_ESM_MME,
        NULL,
        0,
        "_EMMESM_ACTIVATE_BEARER_REQ ue id " MME_UE_S1AP_ID_FMT " ",
        msg->ue_id);
      rc = lowerlayer_activate_bearer_req(
        msg->ue_id,
        msg->u.activate_bearer.ebi,
        msg->u.activate_bearer.mbr_dl,
        msg->u.activate_bearer.mbr_ul,
        msg->u.activate_bearer.gbr_dl,
        msg->u.activate_bearer.gbr_ul,
        msg->u.activate_bearer.msg);
      break;

    case _EMMESM_DEACTIVATE_BEARER_REQ:
      MSC_LOG_RX_MESSAGE(
        MSC_NAS_EMM_MME,
        MSC_NAS_ESM_MME,
        NULL,
        0,
        "_EMMESM_ACTIVATE_BEARER_REQ ue id " MME_UE_S1AP_ID_FMT " ",
        msg->ue_id);
      rc = lowerlayer_deactivate_bearer_req(
        msg->ue_id,
        msg->u.activate_bearer.ebi,
        msg->u.activate_bearer.msg);
      break;


    default: break;
  }

  if (rc != RETURNok) {
    MSC_LOG_RX_DISCARDED_MESSAGE(
      MSC_NAS_EMM_MME,
      MSC_NAS_EMM_MME,
      NULL,
      0,
      "_EMMESM_UNKNOWN(primitive id %d) ue id " MME_UE_S1AP_ID_FMT " ",
      primitive,
      msg->ue_id);
    OAILOG_WARNING(
      LOG_NAS_EMM,
      "EMMESM-SAP - Failed to process primitive %s (%d)\n",
      _emm_esm_primitive_str[primitive - _EMMESM_START - 1],
      primitive);
  }

  OAILOG_FUNC_RETURN(LOG_NAS_EMM, rc);
}

/****************************************************************************/
/*********************  L O C A L    F U N C T I O N S  *********************/
/****************************************************************************/
