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
 *------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include "ngap_state_manager.h"
#include "bstrlib.h"
#include "log.h"
typedef unsigned int uint32_t;
namespace {
constexpr char NGAP_GNB_COLL[]             = "ngap_eNB_coll";
constexpr char NGAP_AMF_ID2ASSOC_ID_COLL[] = "ngap_amf_id2assoc_id_coll";
constexpr char NGAP_IMSI_MAP_TABLE_NAME[]  = "ngap_imsi_map";
}  // namespace

using magma::lte::oai::UeDescription;

namespace magma5g {

NgapStateManager::NgapStateManager() : max_gnbs_(0), max_ues_(0) {}

NgapStateManager::~NgapStateManager() {
  free_state();
}

NgapStateManager& NgapStateManager::getInstance() {
  static NgapStateManager instance;
  return instance;
}

void NgapStateManager::init(
    uint32_t max_ues, uint32_t max_gnbs, bool use_stateless) {
  log_task              = LOG_NGAP;
  table_key             = NGAP_STATE_TABLE;
  task_name             = NGAP_TASK_NAME;
  persist_state_enabled = use_stateless;
  max_ues_              = max_ues;
  max_gnbs_             = max_gnbs;
  create_state();
  if (read_state_from_db() != RETURNok) {
    OAILOG_ERROR(LOG_NGAP, "Failed to read state from redis");
  }
  read_ue_state_from_db();
  is_initialized = true;
}

void NgapStateManager::create_state() {
  bstring ht_name;

  state_cache_p = (ngap_state_t*) calloc(1, sizeof(ngap_state_t));

  ht_name = bfromcstr(NGAP_GNB_COLL);
  hashtable_ts_init(
      &state_cache_p->gnbs, max_gnbs_, nullptr, free_wrapper, ht_name);

  state_ue_ht = hashtable_ts_create(max_ues_, nullptr, free_wrapper, ht_name);
  bdestroy(ht_name);

  ht_name = bfromcstr(NGAP_AMF_ID2ASSOC_ID_COLL);
  hashtable_ts_init(
      &state_cache_p->amfid2associd, max_ues_, nullptr, hash_free_int_func,
      ht_name);
  bdestroy(ht_name);

  state_cache_p->num_gnbs = 0;

  create_ngap_imsi_map();
}

void NgapStateManager::free_state() {
  AssertFatal(
      is_initialized,
      "NgapStateManager init() function should be called to initialize state.");

  if (state_cache_p == nullptr) {
    return;
  }

  int i;
  hashtable_rc_t ht_rc;
  hashtable_key_array_t* keys;
  sctp_assoc_id_t assoc_id;
  gnb_description_t* gnb;

  keys = hashtable_ts_get_keys(&state_cache_p->gnbs);
  if (!keys) {
    OAILOG_DEBUG(LOG_NGAP, "No keys in the amf hashtable");
  } else {
    for (i = 0; i < keys->num_keys; i++) {
      assoc_id = (sctp_assoc_id_t) keys->keys[i];
      ht_rc    = hashtable_ts_get(
          &state_cache_p->gnbs, (hash_key_t) assoc_id, (void**) &gnb);
      AssertFatal(ht_rc == HASH_TABLE_OK, "eNB UE id not in assoc_id");
      AssertFatal(ht_rc == HASH_TABLE_OK, "eNB UE id not in assoc_id");
    }
    FREE_HASHTABLE_KEY_ARRAY(keys);
  }

  if (hashtable_ts_destroy(&state_cache_p->gnbs) != HASH_TABLE_OK) {
    OAI_FPRINTF_ERR("An error occurred while destroying s1 eNB hash table");
  }
  if (hashtable_ts_destroy(&state_cache_p->amfid2associd) != HASH_TABLE_OK) {
    OAI_FPRINTF_ERR("An error occurred while destroying assoc_id hash table");
  }
  if (hashtable_ts_destroy(state_ue_ht) != HASH_TABLE_OK) {
    OAI_FPRINTF_ERR("An error occurred while destroying assoc_id hash table");
  }
  free_wrapper((void**) &state_cache_p);

  clear_ngap_imsi_map();
}

int NgapStateManager::read_ue_state_from_db() {
  if (!persist_state_enabled) {
    return RETURNok;
  }
  auto keys = redis_client->get_keys("IMSI*" + task_name + "*");

  for (const auto& key : keys) {
    UeDescription ue_proto = UeDescription();
    ue_description_t* ue_context =
        (ue_description_t*) calloc(1, sizeof(ue_description_t));
    if (redis_client->read_proto(key.c_str(), ue_proto) != RETURNok) {
      return RETURNerror;
    }

    NgapStateConverter::proto_to_ue(ue_proto, ue_context);

    hashtable_ts_insert(
        state_ue_ht, ue_context->comp_ngap_id, (void*) ue_context);
    OAILOG_DEBUG(log_task, "Reading UE state from db for %s", key.c_str());
  }
  return RETURNok;
}

void NgapStateManager::create_ngap_imsi_map() {
  ngap_imsi_map_ = (ngap_imsi_map_t*) calloc(1, sizeof(ngap_imsi_map_t));

  ngap_imsi_map_->amf_ue_id_imsi_htbl =
      hashtable_uint64_ts_create(max_ues_, nullptr, nullptr);

  oai::NgapImsiMap imsi_proto = oai::NgapImsiMap();
  redis_client->read_proto(NGAP_IMSI_MAP_TABLE_NAME, imsi_proto);

  NgapStateConverter::proto_to_ngap_imsi_map(imsi_proto, ngap_imsi_map_);
}

void NgapStateManager::clear_ngap_imsi_map() {
  if (!ngap_imsi_map_) {
    return;
  }
  hashtable_uint64_ts_destroy(ngap_imsi_map_->amf_ue_id_imsi_htbl);

  free_wrapper((void**) &ngap_imsi_map_);
}

ngap_imsi_map_t* NgapStateManager::get_ngap_imsi_map() {
  return ngap_imsi_map_;
}

void NgapStateManager::put_ngap_imsi_map() {
  oai::NgapImsiMap imsi_proto = oai::NgapImsiMap();
  NgapStateConverter::ngap_imsi_map_to_proto(ngap_imsi_map_, &imsi_proto);
  redis_client->write_proto(NGAP_IMSI_MAP_TABLE_NAME, imsi_proto);
}

}  // namespace magma5g
