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

/*! \file s6a_fd_iface.c
  \brief
  \author Lionel Gauthier
  \company Eurecom
*/

#ifdef __cplusplus
extern "C" {
#endif

#include "common_defs.h"
#include "s6a_messages.h"
#include "s6a_messages_types.h"
#include "intertask_interface.h"
#include "mme_config.h"
#include "timer.h"
#include "dynamic_memory_check.h"

#ifdef __cplusplus
}
#endif

#include "s6a_fd_iface.h"

#include <iostream>
#include <exception>

using namespace std;

static int gnutls_log_level = 9;
static long timer_id = 0;

static void fd_gnutls_debug(int level, const char *str);
static void oai_fd_logger(int loglevel, const char *format, va_list args);

#define S6A_PEER_CONNECT_TIMEOUT_MICRO_SEC (0)
#define S6A_PEER_CONNECT_TIMEOUT_SEC (1)

// LG-EURECOM should be member of S6aFdIface, but requires modifications that will take
// 1 or 2 more hours, a bit late for today
s6a_fd_cnf_t s6a_fd_cnf;

//------------------------------------------------------------------------------
static void fd_gnutls_debug(int loglevel, const char *str)
{
  OAILOG_EXTERNAL(loglevel, LOG_S6A, "[GTLS] %s", str);
}

//------------------------------------------------------------------------------
// callback for freeDiameter logs
static void oai_fd_logger(int loglevel, const char *format, va_list args)
{
#define FD_LOG_MAX_MESSAGE_LENGTH 8192
  char buffer[FD_LOG_MAX_MESSAGE_LENGTH];
  int rv = 0;

  rv = vsnprintf(buffer, sizeof(buffer), format, args);
  if ((0 > rv) || ((FD_LOG_MAX_MESSAGE_LENGTH) < rv)) {
    return;
  }
  OAILOG_EXTERNAL(OAILOG_LEVEL_TRACE - loglevel, LOG_S6A, "%s\n", buffer);
}

//------------------------------------------------------------------------------
S6aFdIface::S6aFdIface(const s6a_config_t * const config)
{
  int ret = RETURNok;
  memset(&s6a_fd_cnf, 0, sizeof(s6a_fd_cnf_t));

  /*
   * if (strcmp(fd_core_version(), free_wrapper_DIAMETER_MINIMUM_VERSION) ) {
   * S6A_ERROR("Freediameter version %s found, expecting %s\n", fd_core_version(),
   * free_wrapper_DIAMETER_MINIMUM_VERSION);
   * return RETURNerror;
   * } else {
   * S6A_DEBUG("Freediameter version %s\n", fd_core_version());
   * }
   */

  /*
   * Initializing freeDiameter logger
   */
  ret = fd_log_handler_register(oai_fd_logger);
  if (ret) {
    OAILOG_ERROR(
      LOG_S6A,
      "An error occurred during freeDiameter log handler registration: %d\n",
      ret);
    std::runtime_error("An error occurred during freeDiameter log handler registration"); ;
  } else {
    OAILOG_DEBUG(LOG_S6A, "Initializing freeDiameter log handler done\n");
  }

  /*
   * Initializing freeDiameter core
   */
  OAILOG_DEBUG(LOG_S6A, "Initializing freeDiameter core...\n");
  ret = fd_core_initialize();
  if (ret) {
    OAILOG_ERROR(
      LOG_S6A,
      "An error occurred during freeDiameter core library initialization: %d\n",
      ret);
    std::runtime_error("An error occurred during freeDiameter core library initialization"); ;
  } else {
    OAILOG_DEBUG(LOG_S6A, "Initializing freeDiameter core done\n");
  }

  OAILOG_DEBUG(LOG_S6A, "Default ext path: %s\n", DEFAULT_EXTENSIONS_PATH);

  ret = fd_core_parseconf(bdata(config->conf_file));
  if (ret) {
    OAILOG_ERROR(
      LOG_S6A,
      "An error occurred during fd_core_parseconf file : %s.\n",
      bdata(config->conf_file));
    std::runtime_error("An error occurred during fd_core_parseconf file"); ;
  } else {
    OAILOG_DEBUG(LOG_S6A, "fd_core_parseconf done\n");
  }

  /*
   * Set gnutls debug level ?
   */
  if (gnutls_log_level) {
    gnutls_global_set_log_function((gnutls_log_func) fd_gnutls_debug);
    gnutls_global_set_log_level(gnutls_log_level);
    OAILOG_DEBUG(
      LOG_S6A, "Enabled GNUTLS debug at level %d\n", gnutls_log_level);
  }

  /*
   * Starting freeDiameter core
   */
  ret = fd_core_start();
  if (ret) {
    OAILOG_ERROR(
      LOG_S6A, "An error occurred during freeDiameter core library start\n");
    std::runtime_error("An error occurred during freeDiameter core library start"); ;
  } else {
    OAILOG_DEBUG(LOG_S6A, "fd_core_start done\n");
  }

  ret = fd_core_waitstartcomplete();
  if (ret) {
    OAILOG_ERROR(
      LOG_S6A, "An error occurred during freeDiameter core library start\n");
    std::runtime_error("An error occurred during freeDiameter core library start\n");
  } else {
    OAILOG_DEBUG(LOG_S6A, "fd_core_waitstartcomplete done\n");
  }

  ret = s6a_fd_init_dict_objs();
  if (ret) {
    OAILOG_ERROR(LOG_S6A, "An error occurred during s6a_fd_init_dict_objs.\n");
    std::runtime_error("An error occurred during s6a_fd_init_dict_obj\n");
  } else {
    OAILOG_DEBUG(LOG_S6A, "s6a_fd_init_dict_objs done\n");
  }

  OAILOG_DEBUG(
    LOG_S6A,
    "Initializing S6a interface over free-diameter:"
    "DONE\n");

  /* Add timer here to send message to connect to peer */
  timer_setup(
    S6A_PEER_CONNECT_TIMEOUT_SEC,
    S6A_PEER_CONNECT_TIMEOUT_MICRO_SEC,
    TASK_S6A,
    INSTANCE_DEFAULT,
    TIMER_ONE_SHOT,
    NULL,
    0,
    &timer_id);
}

//------------------------------------------------------------------------------
bool S6aFdIface::update_location_req(s6a_update_location_req_t *ulr_p)
{
 if (s6a_generate_update_location(ulr_p)) return false;
 else return true;
}
//------------------------------------------------------------------------------
bool S6aFdIface::authentication_info_req(s6a_auth_info_req_t *air_p)
{
  if (s6a_generate_authentication_info_req(air_p)) return false;
  else return true;
}
//------------------------------------------------------------------------------
bool  S6aFdIface::send_cancel_location_ans(s6a_cancel_location_ans_t *cla_pP)
{
  if (s6a_send_cancel_location_ans(cla_pP)) return false;
  else return true;
}
//------------------------------------------------------------------------------
bool S6aFdIface::purge_ue(const char *imsi)
{
  if (s6a_generate_purge_ue_req(imsi)) return false;
  else return true;
}
//------------------------------------------------------------------------------
void S6aFdIface::timer_expired (const long timer_idP)
{
  if (!timer_exists(timer_idP)) {
    return;
  }
  /*
   * Trying to connect to peers
   */
  timer_id = 0;
  if (s6a_fd_new_peer() != RETURNok) {
    /*
     * On failure, reschedule timer.
     * * Preferred over TIMER_PERIODIC because if s6a_fd_new_peer takes
     * * longer to return than the period, the timer will schedule while
     * * the previous one is active, causing a seg fault.
     */
    increment_counter(
      "s6a_subscriberdb_connection_failure", 1, NO_LABELS);
    OAILOG_ERROR(
      LOG_S6A,
      "s6a_fd_new_peer has failed (%s:%d)\n",
      __FILE__,
      __LINE__);
    timer_setup(
      S6A_PEER_CONNECT_TIMEOUT_SEC,
      S6A_PEER_CONNECT_TIMEOUT_MICRO_SEC,
      TASK_S6A,
      INSTANCE_DEFAULT,
      TIMER_ONE_SHOT,
      NULL,
      0,
      &timer_id);
  }
  timer_handle_expired(timer_idP);
}
//------------------------------------------------------------------------------
S6aFdIface::~S6aFdIface()
{
  if (timer_id) {
    timer_remove(timer_id, NULL);
  }
  // Release all resources
  free_wrapper((void **) &fd_g_config->cnf_diamid);
  fd_g_config->cnf_diamid_len = 0;
  int rv = RETURNok;
  /* Initialize shutdown of the framework */
  rv = fd_core_shutdown();
  if (rv) {
    OAI_FPRINTF_ERR("An error occurred during fd_core_shutdown().\n");
  }

  /* Wait for the shutdown to be complete -- this should always be called after fd_core_shutdown */
  rv = fd_core_wait_shutdown_complete();
  if (rv) {
    OAI_FPRINTF_ERR(
      "An error occurred during fd_core_wait_shutdown_complete().\n");
  }
}
