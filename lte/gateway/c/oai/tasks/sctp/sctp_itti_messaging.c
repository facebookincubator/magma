/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
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

/*! \file sctp_itti_messaging.c
  \brief
  \author Sebastien ROUX, Lionel Gauthier
  \company Eurecom
  \email: lionel.gauthier@eurecom.fr
*/

#include <string.h>
#include <stdbool.h>

#include "intertask_interface.h"
#include "sctp_itti_messaging.h"
#include "itti_types.h"
#include "log.h"
#include "sctp_messages_types.h"

//------------------------------------------------------------------------------
int sctp_itti_send_lower_layer_conf(
    task_id_t origin_task_id,sctp_ppid_t ppid, sctp_assoc_id_t assoc_id, sctp_stream_id_t stream,
    uint32_t ap_id, bool is_success) {
  MessageDef* msg = itti_alloc_new_message(TASK_SCTP, SCTP_DATA_CNF);

  SCTP_DATA_CNF(msg).ppid	    = ppid;
  SCTP_DATA_CNF(msg).assoc_id       = assoc_id;
  SCTP_DATA_CNF(msg).stream         = stream;

  if (ppid == S1AP){
          OAILOG_INFO(LOG_SCTP, "ppid S1AP in sctp_itti_send_new_message_ind ");
	  SCTP_DATA_CNF(msg).mme_ue_s1ap_id = ap_id;
  }else if(ppid == NGAP){
	  OAILOG_INFO(LOG_SCTP, " ppid NGAP in sctp_itti_send_lower_layer_conf ");
	  SCTP_DATA_CNF(msg).amf_ue_ngap_id = ap_id;
  }else{
	OAILOG_ERROR(LOG_SCTP, "ppid not matching in sctp_itti_send_lower_layer_conf ");
  }
  
  SCTP_DATA_CNF(msg).is_success     = is_success;


  return send_msg_to_task(&sctp_task_zmq_ctx, origin_task_id, msg);
}

//------------------------------------------------------------------------------
int sctp_itti_send_new_association(
    sctp_ppid_t ppid, sctp_assoc_id_t assoc_id, sctp_stream_id_t instreams,
    sctp_stream_id_t outstreams) {
  MessageDef* msg = itti_alloc_new_message(TASK_SCTP, SCTP_NEW_ASSOCIATION);

  if (ppid == S1AP) {
	  SCTP_NEW_ASSOCIATION(msg).assoc_id = assoc_id;
	  SCTP_NEW_ASSOCIATION(msg).instreams = instreams;
	  SCTP_NEW_ASSOCIATION(msg).outstreams = outstreams;
	  OAILOG_INFO(LOG_SCTP, "ppid S1AP in sctp_itti_send_new_message_ind ");
	  return send_msg_to_task(&sctp_task_zmq_ctx, TASK_S1AP, msg);
  }else if(ppid==NGAP)  {
	  SCTP_NEW_ASSOCIATION(msg).assoc_id = assoc_id;
	  SCTP_NEW_ASSOCIATION(msg).instreams = instreams;
	  SCTP_NEW_ASSOCIATION(msg).outstreams = outstreams;
	  OAILOG_INFO(LOG_SCTP, " ppid NGAP in sctp_itti_send_new_association ");
	  return send_msg_to_task(&sctp_task_zmq_ctx, TASK_NGAP, msg);
  }else {
	 OAILOG_ERROR(LOG_SCTP, "ppid not matching in sctp_itti_send_new_association ");
	 return 0;
  }
}

//------------------------------------------------------------------------------
int sctp_itti_send_new_message_ind(
    STOLEN_REF bstring* payload,sctp_ppid_t ppid,sctp_assoc_id_t assoc_id,
    sctp_stream_id_t stream) {
    MessageDef* msg = itti_alloc_new_message(TASK_SCTP, SCTP_DATA_IND);

    SCTP_DATA_IND(msg).payload   = *payload;
    SCTP_DATA_IND(msg).stream    = stream;
    SCTP_DATA_IND(msg).assoc_id  = assoc_id;

  STOLEN_REF* payload = NULL;
  if (ppid == S1AP){
	  OAILOG_INFO(LOG_SCTP, "ppid S1AP in sctp_itti_send_new_message_ind ");
	  return send_msg_to_task(&sctp_task_zmq_ctx, TASK_S1AP, msg);
  }else if(ppid == NGAP)  {
	  OAILOG_INFO(LOG_SCTP, "ppid NGAP in sctp_itti_send_new_message_ind ");
	  return send_msg_to_task(&sctp_task_zmq_ctx, TASK_NGAP, msg); 
  }else {
	 OAILOG_ERROR(LOG_SCTP, "ppid not matching in sctp_itti_send_new_message_ind ");
	 return 0;
  }	
}

//------------------------------------------------------------------------------
int sctp_itti_send_com_down_ind(sctp_ppid_t ppid,sctp_assoc_id_t assoc_id, bool reset) {
  MessageDef* msg = itti_alloc_new_message(TASK_SCTP, SCTP_CLOSE_ASSOCIATION);

  SCTP_CLOSE_ASSOCIATION(msg).assoc_id = assoc_id;
  SCTP_CLOSE_ASSOCIATION(msg).reset    = reset;

  if (ppid == S1AP){
         OAILOG_INFO(LOG_SCTP, "ppid S1AP in sctp_itti_send_new_message_ind ");
	 return send_msg_to_task(&sctp_task_zmq_ctx, TASK_S1AP, msg);
  }else if(ppid == NGAP)  {
	 OAILOG_INFO(LOG_SCTP, "ppid NGAP in sctp_itti_send_com_down_ind ");
	 return send_msg_to_task(&sctp_task_zmq_ctx, TASK_NGAP, msg);
  }else {
	 OAILOG_ERROR(LOG_SCTP, " ppid not matching in sctp_itti_send_com_down_ind ");
	 return 0;
  }
 
}
