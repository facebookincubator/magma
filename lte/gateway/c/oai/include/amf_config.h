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

  Source      amf_message.h

  Version     0.1

  Date        2020/07/28

  Product     NAS stack

  Subsystem   Access and Mobility Management Function

  Author      Sandeep Kumar Mall

  Description Defines Access and Mobility Management Messages

*****************************************************************************/
#ifndef FILE_AMF_CONFIG_SEEN
#define FILE_AMF_CONFIG_SEEN

#include <pthread.h>
#include <stdint.h>
#include <arpa/inet.h>
#include <stdlib.h>
//#include "amf_default_values.h"
#include "common_types.h"
#include "3gpp_23.003.h"
#include "3gpp_24.008.h"
#include "log.h"
#include "service303.h"

#define MIN_GUMMEI 1
#define MAX_GUMMEI 5
#define MAX_APN_CORRECTION_MAP_LIST 10

typedef struct m5g_apn_map_s {
  bstring imsi_prefix;
  bstring apn_override;
} m5g_apn_map_t;

typedef struct m5g_apn_map_config_s {
  int nb;
  m5g_apn_map_t apn_map[MAX_APN_CORRECTION_MAP_LIST];
} m5g_apn_map_config_t;

typedef struct m5g_nas_config_s {
  uint8_t prefered_integrity_algorithm[8];
  uint8_t prefered_ciphering_algorithm[8];
  uint32_t t3402_min;
  uint32_t t3412_min;
  uint32_t t3422_sec;
  uint32_t t3450_sec;
  uint32_t t3460_sec;
  uint32_t t3470_sec;
  uint32_t t3485_sec;
  uint32_t t3486_sec;
  uint32_t t3489_sec;
  uint32_t t3495_sec;
  // non standard features
  bool force_reject_tau;
  bool force_reject_sr;
  bool disable_esm_information;
  // apn correction
  bool enable_apn_correction;
  m5g_apn_map_config_t m5g_apn_map_config;
} m5g_nas_config_t;
typedef uint64_t imsi64_t;
typedef uint32_t amf_ue_ngap_id_t;

typedef struct m5g_served_tai_s {
  uint8_t list_type;
  uint8_t nb_tai;
  uint16_t* plmn_mcc;
  uint16_t* plmn_mnc;
  uint16_t* plmn_mnc_len;
  uint16_t* tac;
} m5g_served_tai_t;

typedef struct ngap_config_s {
  uint16_t port_number;
  uint8_t outcome_drop_timer_sec;
} ngap_config_t;
#if 0
typedef struct sctp_config_s {
  uint16_t in_streams;
  uint16_t out_streams;
} sctp_config_t;

typedef struct itti_config_s {
  uint32_t queue_size;
  bstring log_file;
} itti_config_t;
#endif
typedef struct guamfi_config_s {
  int nb;
  guamfi_t guamfi[MAX_GUMMEI];
} guamfi_config_t;
// typedef enum { RUN_MODE_TEST = 0, RUN_MODE_OTHER } run_mode_t;

typedef struct amf_config_s {
  /* Reader/writer lock for this configuration */
  pthread_rwlock_t rw_lock;

  bstring config_file;
  bstring pid_dir;
  bstring realm;
  bstring full_network_name;
  bstring short_network_name;
  uint8_t daylight_saving_time;
  //  run_mode_t run_mode;
  uint32_t max_gnbs;
  uint32_t max_ues;
  uint8_t relative_capacity;
  bstring ip_capability;
  uint8_t unauthenticated_imsi_supported;
  guamfi_config_t guamfi;
  m5g_served_tai_t served_tai;
  service303_data_t service303_config;
  // sctp_config_t sctp_config;
  ngap_config_t ngap_config;
  // itti_config_t itti_config;
  m5g_nas_config_t m5g_nas_config;
  // sgs_config_t sgs_config;
  log_config_t log_config;
  // e_dns_config_t e_dns_emulation;

  // ip_t ip;

  // lai_t lai;

  bool use_stateless;
} amf_config_t;

extern amf_config_t amf_config;

int amf_config_find_mnc_length(
    const char mcc_digit1P, const char mcc_digit2P, const char mcc_digit3P,
    const char mnc_digit1P, const char mnc_digit2P, const char mnc_digit3P);

void amf_config_init(amf_config_t*);
int amf_config_parse_opt_line(int argc, char* argv[], amf_config_t* amf_config);
int amf_config_parse_file(amf_config_t*);
void amf_config_display(amf_config_t*);

void mme_config_exit(void);
#define amf_config_read_lock(aMFcONFIG)                                        \
  pthread_rwlock_rdlock(&(aMFcONFIG)->rw_lock)
#define amf_config_write_lock(aMFcONFIG)                                       \
  pthread_rwlock_wrlock(&(aMFcONFIG)->rw_lock)
#define amf_config_unlock(aMFcONFIG)                                           \
  pthread_rwlock_unlock(&(aMFcONFIG)->rw_lock)

#endif
