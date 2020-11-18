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
#define HA
#define HA_TASK_C

#include <stdio.h>

#include "ha_defs.h"
#include "ha_messages_types.h"
#include "log.h"
#include "common_defs.h"
#include "intertask_interface_types.h"
#include "itti_free_defined_msg.h"

static void ha_exit(void);

task_zmq_ctx_t ha_task_zmq_ctx;

static int handle_message(zloop_t* loop, zsock_t* reader, void* arg) {
  zframe_t* msg_frame = zframe_recv(reader);
  assert(msg_frame);
  MessageDef* received_message_p = (MessageDef*) zframe_data(msg_frame);

  switch (ITTI_MSG_ID(received_message_p)) {
    case AGW_OFFLOAD_REQ: {
      OAILOG_DEBUG(
          LOG_UTIL, "Received AGW_OFFLOAD_REQ message for eNB ID %d ",
          AGW_OFFLOAD_REQ(received_message_p).eNB_id);
      handle_agw_offload_req(&received_message_p->ittiMsg.ha_agw_offload_req);
    } break;

    case TERMINATE_MESSAGE: {
      itti_free_msg_content(received_message_p);
      zframe_destroy(&msg_frame);
      ha_exit();
    } break;

    default: {
      OAILOG_DEBUG(
          LOG_UTIL, "Unknown message ID %d:%s\n",
          ITTI_MSG_ID(received_message_p), ITTI_MSG_NAME(received_message_p));
    } break;
  }
  itti_free_msg_content(received_message_p);
  zframe_destroy(&msg_frame);
  return 0;
}

//------------------------------------------------------------------------------
static void* ha_thread(__attribute__((unused)) void* args_p) {
  task_zmq_ctx_t* task_zmq_ctx_p = &ha_task_zmq_ctx;

  itti_mark_task_ready(TASK_HA);
  init_task_context(
      TASK_HA, (task_id_t[]){TASK_MME_APP}, 1, handle_message, task_zmq_ctx_p);

  zloop_start(task_zmq_ctx_p->event_loop);
  ha_exit();
  return NULL;
}

//------------------------------------------------------------------------------
int ha_init(const mme_config_t* mme_config_p) {
  OAILOG_DEBUG(LOG_UTIL, "Initializing HA task interface\n");

  if (itti_create_task(TASK_HA, &ha_thread, NULL) < 0) {
    OAILOG_ERROR(LOG_UTIL, "ha create task\n");
    return RETURNerror;
  }
  OAILOG_DEBUG(LOG_UTIL, "Initializing HA task interface: DONE\n");
  return RETURNok;
}

//------------------------------------------------------------------------------
static void ha_exit(void) {
  destroy_task_context(&ha_task_zmq_ctx);
  OAI_FPRINTF_INFO("TASK_HA terminated\n");
  pthread_exit(NULL);
}
