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
#include <thread>
#ifdef __cplusplus
extern "C" {
#endif
#include "log.h"
#include "conversions.h"
#include "3gpp_24.008.h"
#ifdef __cplusplus
}
#endif
#include "common_defs.h"
#include "M5gNasMessage.h"
#include "amf_app_defs.h"
#include "amf_app_ue_context_and_proc.h"
#include "amf_authentication.h"
#include "amf_as.h"
#include "amf_fsm.h"
#include "amf_recv.h"
#include "dynamic_memory_check.h"
#include "M5GDLNASTransport.h"
#include "S6aClient.h"
#include "proto_msg_to_itti_msg.h"

using namespace magma;
typedef uint32_t amf_ue_ngap_id_t;
#define QUADLET 4
#define AMF_GET_BYTE_ALIGNED_LENGTH(LENGTH)                                    \
  LENGTH += QUADLET - (LENGTH % QUADLET)
#define AMF_CAUSE_SUCCESS (1)
namespace magma5g {
/*forward declaration*/
static int amf_as_establish_req(amf_as_establish_t* msg, int* amf_cause);
static int amf_as_security_req(
    const amf_as_security_t* msg, m5g_dl_info_transfer_req_t* as_msg);
// Setup the security header of the given NAS message
static AMFMsg* amf_as_set_header(
    amf_nas_message_t* msg, const amf_as_security_data_t* security);
/**************************************************************************
**                                                                       **
** Name        : amf_as_send()                                           **
**                                                                       **
** Description : Processes the AMF-AS Service Access Point primitive.    **
**                                                                       **
** Inputs      : msg    :  The AMF-AS-SAP primitive to process           **
**               Others :  None                                          **
**                                                                       **
** Outputs     : None                                                    **
**      Return : RETURNok, RETURNerror                                   **
**      Others : None                                                    **
**                                                                       **
**************************************************************************/
int amf_as_send(amf_as_t* msg) {
  int rc                       = RETURNok;
  int amf_cause                = AMF_CAUSE_SUCCESS;
  amf_as_primitive_t primitive = msg->primitive;

  switch (primitive) {
    case _AMFAS_DATA_IND:
    case _AMFAS_ESTABLISH_REQ:
      // Process UE's establishment request
      rc = amf_as_establish_req(&msg->u.establish, &amf_cause);
      break;
    case _AMFAS_RELEASE_IND:
    default:
      // Other primitives are forwarded to NGAP
      rc = amf_as_send_ng(msg);
      break;
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, rc);
}

/***************************************************************************
  uint8_t* lenPtr;
**                                                                        **
** Name:    amf_as_establish_req()                                        **
**                                                                        **
** Description: Processes the _AMFAS-SAP connection establish request      **
**      primitive                                                         **
**                                                                        **
** _AMFAS-SAP - AS->AMF : ESTABLISH_REQ - NAS signalling connection        **
**     The AS notifies the NAS that establishment of the signal-          **
**     ling connection has been requested to tranfer initial NAS          **
**     message from the UE.                                               **
**                                                                        **
**      Inputs:  msg:       The _AMFAS-SAP primitive to process            **
**      Others:    None                                                   **
**                                                                        **
**      Outputs:   amf_cause: AMF cause code                              **
**      Return:    RETURNok, RETURNerror                                  **
**      Others:    None                                                   **
**                                                                        **
***************************************************************************/
static int amf_as_establish_req(amf_as_establish_t* msg, int* amf_cause) {
  amf_security_context_t* amf_security_context = NULL;
  amf_nas_message_decode_status_t decode_status;
  int decoder_rc        = 1;
  int rc                = RETURNerror;
  tai_t originating_tai = {0};
  amf_nas_message_t nas_msg;
  ue_m5gmm_context_s* ue_m5gmm_context = NULL;
  ue_m5gmm_context = amf_ue_context_exists_amf_ue_ngap_id(msg->ue_id);
  if (ue_m5gmm_context == NULL) {
    OAILOG_ERROR(
        LOG_AMF_APP, "ue context not found for the ue_id=%u\n", msg->ue_id);
    OAILOG_FUNC_RETURN(LOG_AMF_APP, rc);
  }

  ue_m5gmm_context->mm_state = UNREGISTERED;

  // Decode initial NAS message
  decoder_rc = nas5g_message_decode(
      msg->nas_msg->data, &nas_msg, blength(msg->nas_msg), amf_security_context,
      &decode_status);
  bdestroy_wrapper(&msg->nas_msg);

  // conditional IE error
  if (decoder_rc < 0) {
    if (decoder_rc < TLV_FATAL_ERROR) {
      *amf_cause = AMF_CAUSE_PROTOCOL_ERROR;
    } else if (decoder_rc == TLV_MANDATORY_FIELD_NOT_PRESENT) {
      *amf_cause = AMF_CAUSE_INVALID_MANDATORY_INFO;
    } else if (decoder_rc == TLV_UNEXPECTED_IEI) {
      *amf_cause = AMF_CAUSE_IE_NOT_IMPLEMENTED;
    } else {
      *amf_cause = AMF_CAUSE_PROTOCOL_ERROR;
    }
  }

  // Process initial NAS message
  AMFMsg* amf_msg = &nas_msg.plain.amf;
  switch (amf_msg->header.message_type) {
    case REG_REQUEST:
      memcpy(&originating_tai, &msg->tai, sizeof(originating_tai));
      rc = amf_handle_registration_request(
          msg->ue_id, &originating_tai, &msg->ecgi,
          &amf_msg->msg.registrationrequestmsg, msg->is_initial,
          msg->is_amf_ctx_new, *amf_cause, decode_status);
      break;
    case M5G_SERVICE_REQUEST:  // SERVICE_REQUEST:
      rc = amf_handle_service_request(
          msg->ue_id, &amf_msg->msg.service_request, decode_status);
      break;
    case M5G_IDENTITY_RESPONSE:
      rc = amf_handle_identity_response(
          msg->ue_id, &amf_msg->msg.identityresponsemsg.m5gs_mobile_identity,
          *amf_cause, decode_status);
      break;
    case AUTH_RESPONSE:
      rc = amf_handle_authentication_response(
          msg->ue_id, &amf_msg->msg.authenticationresponsemsg, *amf_cause,
          decode_status);
      break;
    case SEC_MODE_COMPLETE:
      rc = amf_handle_security_complete_response(msg->ue_id, decode_status);
      break;
    case REG_COMPLETE:
      rc = amf_handle_registration_complete_response(
          msg->ue_id, &amf_msg->msg.registrationcompletemsg, *amf_cause,
          decode_status);
      break;
    case DE_REG_REQUEST_UE_ORIGIN:
      rc = amf_handle_deregistration_ue_origin_req(
          msg->ue_id, &amf_msg->msg.deregistrationequesmsg, *amf_cause,
          decode_status);
      break;
    case ULNASTRANSPORT:
      rc = amf_smf_send(
          msg->ue_id, &amf_msg->msg.uplinknas5gtransport, *amf_cause);
      break;
    default:
      OAILOG_INFO(
          LOG_NAS_AMF, "unknown message type: %d, in %s ",
          amf_msg->header.message_type, __FUNCTION__);
  }
  return rc;
}

/**************************************************************************
 **                                                                      **
 ** Name       : amf_as_send_ng()                                        **
 **                                                                      **
 ** Description: Builds NAS message according to the given _AMFAS Service **
 **      Access Point primitive and sends it to the Access Stratum       **
 **      sublayer                                                        **
 **                                                                      **
 ** Inputs     : msg: The _AMFAS-SAP primitive to be sent                 **
 **      Others: None                                                    **
 **                                                                      **
 ** Outputs:     None                                                    **
 **      Return: RETURNok, RETURNerror                                   **
 **      Others: None                                                    **
 **                                                                      **
 *************************************************************************/
int amf_as_send_ng(const amf_as_t* msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  amf_as_message_t as_msg = {0};

  switch (msg->primitive) {
    case _AMFAS_DATA_REQ:
      as_msg.msg_id =
          amf_as_data_req(&msg->u.data, &as_msg.msg.dl_info_transfer_req);
      break;
    case _AMFAS_ESTABLISH_CNF:
      as_msg.msg_id = amf_as_establish_cnf(
          &msg->u.establish, &as_msg.msg.nas_establish_rsp);
      break;
    case _AMFAS_SECURITY_REQ:
      as_msg.msg_id = amf_as_security_req(
          &msg->u.security, &as_msg.msg.dl_info_transfer_req);
      break;
    default:
      as_msg.msg_id = 0;
      break;
  }

  /*
   * Send the message to the Access Stratum or NGAP in case of AMF
   */
  if (as_msg.msg_id > 0) {
    switch (as_msg.msg_id) {
      case AS_DL_INFO_TRANSFER_REQ_: {
        amf_app_handle_nas_dl_req(
            as_msg.msg.dl_info_transfer_req.ue_id,
            as_msg.msg.dl_info_transfer_req.nas_msg,
            as_msg.msg.dl_info_transfer_req.err_code);
        OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNok);
      } break;
      case AS_NAS_ESTABLISH_RSP_:
      case AS_NAS_ESTABLISH_CNF_: {
        if (as_msg.msg.nas_establish_rsp.err_code == M5G_AS_SUCCESS) {
          // This flow is to release the UE context after sending the NAS
          // message.
          amf_app_handle_nas_dl_req(
              as_msg.msg.nas_establish_rsp.ue_id,
              as_msg.msg.nas_establish_rsp.nas_msg,
              as_msg.msg.nas_establish_rsp.err_code);
          as_msg.msg.nas_establish_rsp.nas_msg = NULL;
          OAILOG_FUNC_RETURN(LOG_NAS_EMM, RETURNok);
        } else {
          OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNok);
        }
      } break;
      case AS_NAS_RELEASE_REQ_:
        OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNok);
        break;
      default:
        break;
    }
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNerror);
}

/************************************************************************
 **                                                                    **
 ** Name       : amf_as_encode()                                       **
 **                                                                    **
 ** Description: Encodes NAS message into NAS information container    **
 **                                                                    **
 ** Inputs     : msg : The NAS message to encode                       **
 **      length: The maximum length of the NAS message                 **
 **      Others: None                                                  **
 **                                                                    **
 ** Outputs    : info : The NAS information container                  **
 **      msg   : The NAS message to encode                             **
 **      Return: The number of bytes successfully encoded              **
 **      Others: None                                                  **
 **                                                                    **
 ***********************************************************************/
static int amf_as_encode(
    bstring* info, amf_nas_message_t* msg, size_t length,
    amf_security_context_t* amf_security_context) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int bytes = 1;

  /* Ciphering algorithms, EA1 and EA2 expects length to be mode of 4,
   * so length is modified such that it will be mode of 4 */
  AMF_GET_BYTE_ALIGNED_LENGTH(length);
  if (msg->header.security_header_type != SECURITY_HEADER_TYPE_NOT_PROTECTED) {
    amf_msg_header* header = &msg->security_protected.plain.amf.header;
    // Expand size of protected NAS message
    length += NAS_MESSAGE_SECURITY_HEADER_SIZE;
    // Set header of plain NAS message
    header->extended_protocol_discriminator = M5GS_MOBILITY_MANAGEMENT_MESSAGE;
    header->security_header_type = SECURITY_HEADER_TYPE_NOT_PROTECTED;
  }

  // Allocate memory to the NAS information container
  *info = bfromcstralloc(length, "\0");

  if (*info) {
    // Encode the NAS message
    AmfMsg amf_msg;
    bytes = amf_msg.M5gNasMessageEncodeMsg(
        (AmfMsg*) &msg->security_protected.plain.amf, (uint8_t*) (*info)->data,
        (uint32_t) length);

    if (bytes > 0) {
      (*info)->slen = bytes;
    } else {
      bdestroy_wrapper(info);
    }
  }

  OAILOG_FUNC_RETURN(LOG_NAS_AMF, bytes);
}

/****************************************************************************
 **                                                                        **
 ** Name:        amf_reg_acceptmsg()                                       **
 **                                                                        **
 ** Description: Builds Registration accept message                        **
 **                                                                        **
 **              The Registration Accept message is sent by the            **
 **              network to the UEi.                                       **
 **                                                                        **
 ** Inputs:      msg:           The AMFMsg    primitive to process         **
 **              Others:        None                                       **
 **                                                                        **
 ** Outputs:     amf_msg:       The AMF message to be sent                 **
 **              Return:        The size of the AMF message                **
 **              Others:        None                                       **
 **                                                                        **
 ***************************************************************************/
static int amf_reg_acceptmsg(
    const amf_as_establish_t* msg, amf_nas_message_t* nas_msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int size = REGISTRATION_ACCEPT_MINIMUM_LENGTH;
  nas_msg->security_protected.plain.amf.header.message_type = REG_ACCEPT;
  nas_msg->security_protected.plain.amf.header.extended_protocol_discriminator =
      M5G_MOBILITY_MANAGEMENT_MESSAGES;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg
      .extended_protocol_discriminator.extended_proto_discriminator =
      M5G_MOBILITY_MANAGEMENT_MESSAGES;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg
      .sec_header_type.sec_hdr = SECURITY_HEADER_TYPE_NOT_PROTECTED;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.message_type
      .msg_type = REG_ACCEPT;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg
      .m5gs_reg_result.sms_allowed = NOT_ALLOWED;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg
      .m5gs_reg_result.reg_result_val = M3GPP_ACCESS;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.odd_even = EVEN_IDENTITY_DIGITS;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .iei = M5GS_MOBILE_IDENTITY;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .len = M5GSMobileIdentityMsg_GUTI_LENGTH;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.type_of_identity = M5GSMobileIdentityMsg_GUTI;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.mcc_digit1 = msg->guti.guamfi.plmn.mcc_digit1;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.mcc_digit2 = msg->guti.guamfi.plmn.mcc_digit2;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.mcc_digit3 = msg->guti.guamfi.plmn.mcc_digit3;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.mnc_digit1 = msg->guti.guamfi.plmn.mnc_digit1;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.mnc_digit2 = msg->guti.guamfi.plmn.mnc_digit2;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.mnc_digit3 = msg->guti.guamfi.plmn.mnc_digit3;
  uint8_t* offset;
  offset = (uint8_t*) &msg->guti.m_tmsi;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.tmsi1 = *offset;
  offset++;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.tmsi2 = *offset;
  offset++;
  nas_msg->security_protected.plain.amf.msg.registrationacceptmsg.mobile_id
      .mobile_identity.guti.tmsi3 = *offset;
  size += MOBILE_IDENTITY_MAX_LENGTH;
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, size);
}

/****************************************************************************
 **                                                                        **
 ** Name:        amf_dl_nas_transport_msg()                            **
 **                                                                        **
 ** Description: Builds Downlink Nas Transport message                     **
 **                                                                        **
 **              The Downlink Nas Transport message is sent by the         **
 **              network to the UE to transfer the data in DL              **
 **              This function is used to send DL NAS Transport message    **
 **              via S1AP DL NAS Transport message.                        **
 **                                                                        **
 ** Inputs:      msg:           The AMFMsg    primitive to process         **
 **              Others:        None                                       **
 **                                                                        **
 ** Outputs:     amf_msg:       The AMF message to be sent                 **
 **              Return:        The size of the AMF message                **
 **              Others:        None                                       **
 **                                                                        **
 ***************************************************************************/
static int amf_dl_nas_transport_msg(
    const amf_as_data_t* msg, DLNASTransportMsg* amf_msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int size = AMF_HEADER_LENGTH;
  // Mandatory - Message type
  amf_msg->message_type.msg_type = DOWNLINK_NAS_TRANSPORT;
  // Mandatory - Nas message container
  size += NAS5G_MESSAGE_CONTAINER_MAXIMUM_LENGTH;
  memcpy(
      amf_msg->payload_container.contents, &(msg->nas_msg),
      sizeof(msg->nas_msg));
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, size);
}

/****************************************************************************
 **                                                                        **
 ** Name:        amf_de_reg_acceptmsg()                                    **
 **                                                                        **
 ** Description: Builds De-Registration Accept message                     **
 **                                                                        **
 **              The De-Registration accept message is sent by the network **
 **                                                                        **
 ** Inputs:      msg:           The AMFMsg    primitive to process         **
 **              Others:        None                                       **
 **                                                                        **
 ** Outputs:     amf_msg:       The AMF message to be sent                 **
 **              Return:        The size of the AMF message                **
 **              Others:        None                                       **
 **                                                                        **
 ***************************************************************************/
static int amf_de_reg_acceptmsg(
    const amf_as_data_t* msg, DeRegistrationAcceptUEInitMsg* amf_msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int size = DEREGISTRATION_ACCEPT_UEINIT_MINIMUM_LENGTH;
  // Mandatory - Message type
  amf_msg->message_type.msg_type = DE_REG_ACCEPT_UE_ORIGIN;
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, size);
}

/****************************************************************************
 **                                                                        **
 ** Name:    amf_as_data_req()                                             **
 **                                                                        **
 ** Description: Processes the AMFAS-SAP data transfer request             **
 **      primitive                                                         **
 **                                                                        **
 ** AMFAS-SAP - AMF->AS : DATA_REQ - Data transfer procedure               **
 **                                                                        **
 ** Inputs:  msg:       The AMFAS-SAP primitive to process                 **
 **      Others:    None                                                   **
 **                                                                        **
 ** Outputs:     as_msg:    The message to send to the AS                  **
 **      Return:    The identifier of the AS message                       **
 **      Others:    None                                                   **
 **                                                                        **
 ***************************************************************************/
uint16_t amf_as_data_req(
    const amf_as_data_t* msg, m5g_dl_info_transfer_req_t* as_msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int size       = 0;
  int is_encoded = false;
  amf_nas_message_t nas_msg;
  nas_msg.security_protected.header           = {0};
  nas_msg.security_protected.plain.amf.header = {0};
  nas_msg.security_protected.plain.amf.header = {0};

  // Setup the AS message
  if (msg->guti) {
    as_msg->s_tmsi.amf_set_id = msg->guti->guamfi.amf_set_id;
    as_msg->s_tmsi.m_tmsi     = msg->guti->m_tmsi;
  } else {
    as_msg->ue_id = msg->ue_id;
  }

  // Setup the NAS security header
  AMFMsg* amf_msg = amf_as_set_header(&nas_msg, &msg->sctx);

  // Setup the NAS information message
  if (amf_msg) {
    switch (msg->nas_info) {
      case AMF_AS_NAS_DATA_REGISTRATION_ACCEPT:
        size = amf_send_registration_accept_dl_nas(
            msg, &amf_msg->msg.registrationacceptmsg);
        break;
      case AMF_AS_NAS_DL_NAS_TRANSPORT:
        // DL messages to NGAP on Identity/Authentication request
        size =
            amf_dl_nas_transport_msg(msg, &amf_msg->msg.downlinknas5gtransport);
        break;
      case AMF_AS_NAS_DATA_DEREGISTRATION_ACCEPT: {
        size = amf_de_reg_acceptmsg(msg, &amf_msg->msg.deregistrationacceptmsg);
      } break;
      default:
        // Send other NAS messages as already encoded SMF messages
        size = msg->nas_msg.length();
        break;
    }
  }

  if (size > 0) {
    int bytes                                    = 0;
    amf_security_context_t* amf_security_context = NULL;
    OAILOG_DEBUG(LOG_AMF_APP, "start NAS encoding\n");
    amf_context_t* amf_ctx = NULL;
    ue_m5gmm_context_s* ue_m5gmm_context =
        amf_ue_context_exists_amf_ue_ngap_id(msg->ue_id);

    if (ue_m5gmm_context) {
      amf_ctx = &ue_m5gmm_context->amf_context;
#if 1  // TODO-RECHECK for NW initiated derestration and security
      if (amf_ctx) {
        // if (amf_msg->nw_deregister_request.nw_deregistertype ==
        //    NW_DEREGISTER_TYPE_IMSI_DEREGISTER) {
        //  amf_ctx->is_imsi_only_deregister = true;
        //}
        if (IS_AMF_CTXT_PRESENT_SECURITY(amf_ctx)) {
          amf_security_context = &amf_ctx->_security;
          // is_encoded           = true;// TODO
        }
      }
#endif
    } else {
      OAILOG_ERROR(
          LOG_AMF_APP, "ue context not found for the ue_id=%u\n", msg->ue_id);
      OAILOG_FUNC_RETURN(LOG_AMF_APP, RETURNerror);
    }

    if (!is_encoded) {
      /*
       * Encode the NAS information message
       */
      bytes =
          amf_as_encode(&as_msg->nas_msg, &nas_msg, size, amf_security_context);
    }

    if (bytes > 0) {
      OAILOG_DEBUG(LOG_AMF_APP, "NAS encoding successful\n");
      as_msg->err_code = M5G_AS_SUCCESS;
    } else {
      OAILOG_ERROR(LOG_AMF_APP, "NAS encoding failed\n");
    }

    OAILOG_FUNC_RETURN(LOG_NAS_AMF, AS_DL_INFO_TRANSFER_REQ_);
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, 0);
}

/***************************************************************************
 **                                                                       **
 ** Name:        amf_as_set_header()                                      **
 **                                                                       **
 ** Description: Setup the security header of the given NAS message       **
 **                                                                       **
 ** Inputs:      security: The NAS security data to use                   **
 **              Others:   None                                           **
 **                                                                       **
 ** Outputs:     msg:     The NAS message                                 **
 **              Return:  Pointer to the plain NAS message to be se-      **
 **                       curity protected if setting of the securi-      **
 **                       ty header succeed;                              **
 **                       NULL pointer otherwise                          **
 **              Others:  None                                            **
 **                                                                       **
 **************************************************************************/
AMFMsg* amf_as_set_header(
    amf_nas_message_t* msg, const amf_as_security_data_t* security) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  msg->header.extended_protocol_discriminator =
      M5GS_MOBILITY_MANAGEMENT_MESSAGE;

  if (security && (security->ksi != KSI_NO_KEY_AVAILABLE)) {
    /*
     * A valid 5G CN security context exists
     */
    if (security->is_new) {
      /*
       * New 5G CN security context is taken into use
       */
      if (security->is_knas_int_present) {
        if (security->is_knas_enc_present) {
          /*
           * NAS integrity and cyphering keys are available
           */
          msg->header.security_header_type =
              SECURITY_HEADER_TYPE_INTEGRITY_PROTECTED_CYPHERED_NEW;
        } else {
          /*
           * NAS integrity key only is available
           */
          msg->header.security_header_type =
              SECURITY_HEADER_TYPE_INTEGRITY_PROTECTED_NEW;
        }
        OAILOG_FUNC_RETURN(LOG_NAS_AMF, &msg->security_protected.plain.amf);
      }
    } else if (security->is_knas_int_present) {
      if (security->is_knas_enc_present) {
        /*
         * NAS integrity and cyphering keys are available
         */
        msg->header.security_header_type =
            SECURITY_HEADER_TYPE_INTEGRITY_PROTECTED_CYPHERED;
      } else {
        /*
         * NAS integrity key only is available
         */
        msg->header.security_header_type =
            SECURITY_HEADER_TYPE_INTEGRITY_PROTECTED;
      }
      OAILOG_FUNC_RETURN(LOG_NAS_AMF, &msg->security_protected.plain.amf);
    } else {
      /*
       * No valid 5G CN security context exists
       */
      msg->header.security_header_type = SECURITY_HEADER_TYPE_NOT_PROTECTED;
      OAILOG_FUNC_RETURN(LOG_NAS_AMF, &msg->plain.amf);
    }
  } else {
    /*
     * No valid 5G CN security context exists
     */
    msg->header.security_header_type = SECURITY_HEADER_TYPE_NOT_PROTECTED;
    OAILOG_FUNC_RETURN(LOG_NAS_AMF, &msg->plain.amf);
  }

  OAILOG_FUNC_RETURN(LOG_NAS_AMF, NULL);
}

/***************************************************************************
 **                                                                       **
 ** Name:        amf_identity_request()                              **
 **                                                                       **
 ** Description: Send Identity Request to UE                              **
 **                                                                       **
 ** Inputs:      msg: Security msg                                        **
 **              amf_msg :   amf msg                                      **
 **                                                                       **
 ** Return:   size                                                        **
 **                                                                       **
 **************************************************************************/
static int amf_identity_request(
    const amf_as_security_t* msg, IdentityRequestMsg* amf_msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int size = AMF_HEADER_LENGTH;
  /*
   * Mandatory - Message type
   */
  amf_msg->message_type.msg_type  = IDENTITY_REQUEST;
  amf_msg->m5gs_identity_type.toi = M5G_IDENTITY_SUCI;
  size += IDENTITY_TYPE_2_IE_MAX_LENGTH;
  if (msg->ident_type == IDENTITY_TYPE_2_IMSI) {
    amf_msg->m5gs_identity_type.toi = IDENTITY_TYPE_2_IMSI;
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, size);
}

static void _s6a_handle_authentication_info_ans(
    const std::string& imsi, uint8_t imsi_length, const grpc::Status& status,
    feg::AuthenticationInformationAnswer response, s6a_auth_info_ans_t* aia_p) {
  strncpy(aia_p->imsi, imsi.c_str(), imsi_length);
  aia_p->imsi_length = imsi_length;

  if (status.ok()) {
    if (response.error_code() < feg::ErrorCode::COMMAND_UNSUPORTED) {
      magma::convert_proto_msg_to_itti_s6a_auth_info_ans(response, aia_p);
      OAILOG_DEBUG(
          LOG_AMF_APP,
          "Received S6A-AUTHENTICATION_INFORMATION_ANSWER for IMSI: %s",
          imsi.c_str());
    }
  } else {
    OAILOG_INFO(
        LOG_AMF_APP,
        "S6A-AUTHENTICATION_INFORMATION_ANSWER failed with "
        "status:%s, StatusCode:%d\n",
        status.error_message().c_str(), response.error_code());
  }
}

/***************************************************************************
 **                                                                       **
 ** Name:        amf_auth_request()                                  **
 **                                                                       **
 ** Description: Send authentication Request to UE                        **
 **                                                                       **
 ** Inputs:      msg: Security msg                                        **
 **              amf_msg :   amf msg                                      **
 **                                                                       **
 ** Return:   size                                                        **
 **                                                                       **
 **************************************************************************/
static int amf_auth_request(
    const amf_as_security_t* msg, AuthenticationRequestMsg* amf_msg) {
  s6a_auth_info_req_t air_t;
  memset(&air_t, 0, sizeof(s6a_auth_info_req_t));

  ue_m5gmm_context_s* ue_context =
      amf_ue_context_exists_amf_ue_ngap_id(msg->ue_id);
  if (ue_context) {
    IMSI64_TO_STRING(ue_context->amf_context.imsi64, air_t.imsi, IMSI_LENGTH);
  } else {
    OAILOG_ERROR(
        LOG_AMF_APP, "ue context not found for the ue_id=%u\n", msg->ue_id);
    OAILOG_FUNC_RETURN(LOG_AMF_APP, RETURNerror);
  }
  char temp_imsi[IMSI_BCD_DIGITS_MAX + 1] = "208950000000031";
  strcpy(air_t.imsi, temp_imsi);
  air_t.imsi_length             = IMSI_LENGTH;
  air_t.visited_plmn.mcc_digit1 = 0x2;
  air_t.visited_plmn.mcc_digit2 = 0x0;
  air_t.visited_plmn.mcc_digit3 = 0x8;
  air_t.visited_plmn.mnc_digit1 = 0x9;
  air_t.visited_plmn.mnc_digit2 = 0x5;
  air_t.visited_plmn.mnc_digit3 = 0x0;
  air_t.nb_of_vectors           = 1;
  air_t.re_synchronization      = 0;
  s6a_auth_info_ans_t aia_t;
  memset(&aia_t, 0, sizeof(s6a_auth_info_ans_t));
  auto imsi_len = air_t.imsi_length;
  OAILOG_DEBUG(LOG_AMF_APP, "Sending S6A-AUTHENTICATION_INFORMATION_REQUEST\n");
  magma::S6aClient::authentication_info_req(
      &air_t,
      [imsiStr = std::string(air_t.imsi), imsi_len, &aia_t](
          grpc::Status status, feg::AuthenticationInformationAnswer response) {
        _s6a_handle_authentication_info_ans(
            imsiStr, imsi_len, status, response, &aia_t);
      });
  std::this_thread::sleep_for(std::chrono::milliseconds(60));

  if (aia_t.auth_info.nb_of_vectors == 1) {
    amf_msg->auth_autn.AUTN.assign(
        (const char*) aia_t.auth_info.eutran_vector[0].autn,
        AUTN_LENGTH_OCTETS);
    amf_msg->auth_rand.rand_val.assign(
        (const char*) aia_t.auth_info.eutran_vector[0].rand,
        RAND_LENGTH_OCTETS);
  } else {
    OAILOG_DEBUG(LOG_AMF_APP, "s6a_air request failed\n");
  }
  OAILOG_DEBUG(LOG_AMF_APP, "Sending AUTHENTICATION_REQUEST to UE\n");
  int size = AUTHENTICATION_REQUEST_MINIMUM_LENGTH;
  amf_msg->extended_protocol_discriminator.extended_proto_discriminator =
      M5G_MOBILITY_MANAGEMENT_MESSAGES;
  amf_msg->message_type.msg_type      = AUTH_REQUEST;
  amf_msg->nas_key_set_identifier.tsc = NATIVE_SECURITY_CONTEXT;
  amf_msg->nas_key_set_identifier.nas_key_set_identifier = 0x1;
  uint8_t abba_buff[]                                    = {0x00, 0x00};
  amf_msg->abba.contents.assign((const char*) abba_buff, sizeof(abba_buff));
  size += RAND_MAX_LEN;
  amf_msg->auth_rand.iei = AUTH_PARAM_RAND;
  size += AUTN_MAX_LEN;
  amf_msg->auth_autn.iei = AUTH_PARAM_AUTN;
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, size);
}

/****************************************************************************
 **                                                                        **
 ** Name:        amf_security_mode_command()                          **
 **                                                                        **
 ** Description: Builds Security Mode Command message.                     **
 **      The Security Mode Command message is sent by the network          **
 **      to the UE to establish NAS signalling security.                   **
 **                                                                        **
 ** Inputs:      msg:     The AMFAS-SAP primitive to process               **
 **              Others:  None                                             **
 **                                                                        **
 ** Outputs:     amf_msg: The AMF message to be sent                       **
 **              Return:  The size of the AMF message                      **
 **              Others:  None                                             **
 **                                                                        **
 ***************************************************************************/
static int amf_security_mode_command(
    const amf_as_security_t* msg, SecurityModeCommandMsg* amf_msg) {
  int size = SECURITY_MODE_COMMAND_MINIMUM_LENGTH;
  amf_msg->extended_protocol_discriminator.extended_proto_discriminator =
      M5G_MOBILITY_MANAGEMENT_MESSAGES;
  amf_msg->sec_header_type.sec_hdr    = SECURITY_HEADER_TYPE_NOT_PROTECTED;
  amf_msg->message_type.msg_type      = SEC_MODE_COMMAND;
  amf_msg->nas_sec_algorithms.tca     = M5G_NAS_SECURITY_ALGORITHMS_5G_EA0;
  amf_msg->nas_sec_algorithms.tia     = M5G_NAS_SECURITY_ALGORITHMS_5G_IA0;
  amf_msg->nas_key_set_identifier.tsc = NATIVE_SECURITY_CONTEXT;
  amf_msg->nas_key_set_identifier.nas_key_set_identifier = 1;
  amf_msg->ue_sec_capability.length      = UE_SECURITY_CAPABILITY_MIN_LENGTH;
  amf_msg->ue_sec_capability.ea0         = 1;
  amf_msg->ue_sec_capability.ea1         = 0;
  amf_msg->ue_sec_capability.ea2         = 0;
  amf_msg->ue_sec_capability.ea3         = 0;
  amf_msg->ue_sec_capability.ea4         = 0;
  amf_msg->ue_sec_capability.ea5         = 0;
  amf_msg->ue_sec_capability.ea6         = 0;
  amf_msg->ue_sec_capability.ea7         = 0;
  amf_msg->ue_sec_capability.ia0         = 1;
  amf_msg->ue_sec_capability.ia1         = 0;
  amf_msg->ue_sec_capability.ia2         = 0;
  amf_msg->ue_sec_capability.ia3         = 0;
  amf_msg->ue_sec_capability.ia4         = 0;
  amf_msg->ue_sec_capability.ia5         = 0;
  amf_msg->ue_sec_capability.ia6         = 0;
  amf_msg->ue_sec_capability.ia7         = 0;
  amf_msg->imeisv_request.imeisv_request = IMEISV_REQUESTED;
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, size);
}

/****************************************************************************
 **                                                                        **
 ** Name:              amf_as_security_req()                               **
 **                                                                        **
 ** Description:       Processes the AMFAS-SAP security request primitive  **
 **                                                                        **
 ** AMFAS-SAP-AMF->AS: SECURITY_REQ - Security mode control procedure      **
 **                                                                        **
 ** Inputs:  msg:      The AMFAS-SAP primitive to process                  **
 **          Others:   None                                                **
 **                                                                        **
 ** Outputs: as_msg:   The message to send to the AS                       **
 **          Return:   The identifier of the AS message                    **
 **          Others:   None                                                **
 **                                                                        **
 ***************************************************************************/
static int amf_as_security_req(
    const amf_as_security_t* msg, m5g_dl_info_transfer_req_t* as_msg) {
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  int size = 0;
  amf_nas_message_t nas_msg;
  nas_msg.security_protected.header           = {0};
  nas_msg.security_protected.plain.amf.header = {0};
  nas_msg.security_protected.plain.amf.header = {0};

  /*
   * Setup the AS message
   */
  if (msg) {
    as_msg->s_tmsi.amf_set_id = msg->guti.guamfi.amf_set_id;
    as_msg->s_tmsi.m_tmsi     = msg->guti.m_tmsi;
    as_msg->ue_id             = msg->ue_id;
  } else {
    as_msg->ue_id = msg->ue_id;
  }
  /*
   * Setup the NAS security header
   */
  AMFMsg* amf_msg = amf_as_set_header(&nas_msg, &msg->sctx);
  /*
   * Setup the NAS security message
   */
  if (amf_msg) switch (msg->msg_type) {
      case AMF_AS_MSG_TYPE_IDENT:
        size = amf_identity_request(msg, &amf_msg->msg.identityrequestmsg);
        break;
      case AMF_AS_MSG_TYPE_AUTH:
        size = amf_auth_request(msg, &amf_msg->msg.authenticationrequestmsg);
        break;
      case AMF_AS_MSG_TYPE_SMC:
        size = amf_security_mode_command(
            msg, &amf_msg->msg.securitymodecommandmsg);
        break;
      default:
        OAILOG_WARNING(
            LOG_NAS_AMF,
            "AMFAS-SAP - Type of NAS security "
            "message 0x%.2x is not valid\n",
            msg->msg_type);
    }

  if (size > 0) {
    amf_context_t* amf_ctx                       = NULL;
    amf_security_context_t* amf_security_context = NULL;
    ue_m5gmm_context_s* ue_mm_context =
        amf_ue_context_exists_amf_ue_ngap_id(msg->ue_id);

    if (ue_mm_context) {
      amf_ctx = &ue_mm_context->amf_context;

      if (amf_ctx) {
        if (IS_AMF_CTXT_PRESENT_SECURITY(amf_ctx)) {
          amf_security_context           = &amf_ctx->_security;
          nas_msg.header.sequence_number = amf_ctx->_security.dl_count.seq_num;
        }
      }
    } else {
      OAILOG_ERROR(
          LOG_AMF_APP, "ue context not found for the ue_id=%u\n", msg->ue_id);
      OAILOG_FUNC_RETURN(LOG_AMF_APP, RETURNerror);
    }

    /*
     * Encode the NAS security message
     */
    OAILOG_DEBUG(LOG_AMF_APP, "Start NAS encoding");
    int bytes =
        amf_as_encode(&as_msg->nas_msg, &nas_msg, size, amf_security_context);

    if (bytes > 0) {
      OAILOG_DEBUG(LOG_AMF_APP, "NAS Encoding Success");
      as_msg->err_code = M5G_AS_SUCCESS;
      OAILOG_FUNC_RETURN(LOG_NAS_AMF, AS_DL_INFO_TRANSFER_REQ_);
    } else {
      OAILOG_INFO(LOG_AMF_APP, "NAS Encoding Failed");
      OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNerror);
    }
  }

  OAILOG_FUNC_RETURN(LOG_NAS_AMF, 0);
}

/****************************************************************************
 **                                                                        **
 ** Name:             amf_as_establish_cnf()                               **
 **                                                                        **
 ** Description:      Processes the AMFAS-SAP connection establish confirm **
 **      primitive of PDU session                                          **
 **                                                                        **
 ** AMFAS-SAP-AMF->AS:ESTABLISH_CNF - NAS signalling connection            **
 **                                                                        **
 ** Inputs:   msg:    The AMFAS-SAP primitive to process                   **
 **           Others: None                                                 **
 **                                                                        **
 ** Outputs:  as_msg: The message to send to the AS                        **
 **           Return: The identifier of the AS message                     **
 **           Others: None                                                 **
 **                                                                        **
 ***************************************************************************/
uint16_t amf_as_establish_cnf(
    const amf_as_establish_t* msg, nas5g_establish_rsp_t* as_msg) {
  int size    = 0;
  int ret_val = 0;
  OAILOG_FUNC_IN(LOG_NAS_AMF);
  OAILOG_DEBUG(
      LOG_NAS_AMF,
      "Send AS connection establish confirmation for (ue_id = "
      "%d)\n",
      msg->ue_id);
  amf_nas_message_t nas_msg;
  // Setting-up the AS message
  as_msg->ue_id = msg->ue_id;

  if (msg->pds_id.guti == NULL) {
    OAILOG_WARNING(LOG_NAS_AMF, "AMFAS-SAP - GUTI is NULL...");
    OAILOG_FUNC_RETURN(LOG_NAS_AMF, ret_val);
  }

  as_msg->nas_msg          = msg->nas_msg;
  as_msg->presencemask     = msg->presencemask;
  as_msg->m5g_service_type = msg->service_type;
  amf_context_t* amf_ctx   = NULL;
  amf_security_context_t* amf_security_context = NULL;
  amf_ctx                                      = amf_context_get(msg->ue_id);
  if (amf_ctx) {
    if (IS_AMF_CTXT_PRESENT_SECURITY(amf_ctx)) {
      amf_security_context                  = &amf_ctx->_security;
      as_msg->selected_encryption_algorithm = (uint16_t) htons(
          0x10000 >> amf_security_context->selected_algorithms.encryption);
      as_msg->selected_integrity_algorithm = (uint16_t) htons(
          0x10000 >> amf_security_context->selected_algorithms.integrity);
      as_msg->nas_ul_count = 0x00000000 |
                             (amf_security_context->ul_count.overflow << 8) |
                             amf_security_context->ul_count.seq_num;
    }
  } else {
    OAILOG_WARNING(LOG_NAS_AMF, "AMFAS-SAP - AMF Context is NULL...!");
  }
  /*
   * Setup the NAS security header
   */
  amf_as_set_header(&nas_msg, &msg->sctx);
  switch (msg->nas_info) {
    case AMF_AS_NAS_INFO_REGISTERD:
      size = amf_reg_acceptmsg(msg, &nas_msg);
      break;
    case AMF_AS_NAS_INFO_TAU:
    case AMF_AS_NAS_INFO_NONE:  // Response to SR
      as_msg->err_code = M5G_AS_SUCCESS;
      ret_val          = AS_NAS_ESTABLISH_CNF_;
      OAILOG_FUNC_RETURN(LOG_NAS_AMF, ret_val);
    default:
      OAILOG_WARNING(
          LOG_NAS_AMF,
          "AMFAS-SAP - Type of initial NAS "
          "message 0x%.2x is not valid\n",
          msg->nas_info);
      break;
  }

  if (size > 0) {
    nas_msg.header.sequence_number = amf_security_context->dl_count.seq_num;
  } else {
    OAILOG_FUNC_RETURN(LOG_NAS_AMF, ret_val);
  }
  /*
   * Encode the initial NAS information message
   */
  OAILOG_DEBUG(LOG_AMF_APP, "start NAS encoding \n");
  int bytes =
      amf_as_encode(&as_msg->nas_msg, &nas_msg, size, amf_security_context);

  if (bytes > 0) {
    OAILOG_DEBUG(LOG_AMF_APP, "NAS encoding success\n");
    as_msg->err_code = M5G_AS_SUCCESS;
    ret_val          = AS_NAS_ESTABLISH_CNF_;
  } else {
    OAILOG_INFO(LOG_AMF_APP, "NAS Encoding Failed");
    OAILOG_FUNC_RETURN(LOG_NAS_AMF, RETURNerror);
  }
  OAILOG_FUNC_RETURN(LOG_NAS_AMF, ret_val);
}
}  // namespace magma5g
