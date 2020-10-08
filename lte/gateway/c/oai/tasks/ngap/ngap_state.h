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

#pragma once

#ifdef __cplusplus
extern "C" {
#endif

#include "hashtable.h"
#include "amf_config.h"
#include "ngap_types.h"

int ngap_state_init(uint32_t max_ues, uint32_t max_gnbs, bool use_stateless);

void ngap_state_exit(void);

ngap_state_t* get_ngap_state(bool read_from_db);

void put_ngap_state(void);

gnb_description_t* ngap_state_get_gnb(
  ngap_state_t* state,
  sctp_assoc_id_t assoc_id);

ue_description_t* ngap_state_get_ue_gnbid(
  sctp_assoc_id_t sctp_assoc_id,
  gnb_ue_ngap_id_t gnb_ue_ngap_id);

ue_description_t* ngap_state_get_ue_amfid(
  amf_ue_ngap_id_t amf_ue_ngap_id);

ue_description_t* ngap_state_get_ue_imsi(imsi64_t imsi64);

/**
 * Return unique composite id for NGAP UE context
 * @param sctp_assoc_id unique SCTP assoc id
 * @param gnb_ue_ngap_id unique UE ngap ID on eNB
 * @return uint64_t of composite id
 */
uint64_t ngap_get_comp_ngap_id(
    sctp_assoc_id_t sctp_assoc_id,
    gnb_ue_ngap_id_t gnb_ue_ngap_id);

/**
 * Converts ngap_imsi_map to protobuf and saves it into data store
 */
void put_ngap_imsi_map(void);

/**
 * @return ngap_imsi_map_t pointer
 */
ngap_imsi_map_t * get_ngap_imsi_map(void);

hash_table_ts_t* get_ngap_ue_state(void);

int read_ngap_ue_state_db(void);

void put_ngap_ue_state(imsi64_t imsi64);

void delete_ngap_ue_state(imsi64_t imsi64);

bool ngap_ue_compare_by_amf_ue_id_cb(
  __attribute__((unused)) hash_key_t keyP,
  void* elementP,
  void* parameterP,
  void** resultP);

bool ngap_ue_compare_by_imsi(
    __attribute__((unused)) hash_key_t keyP,
    void* elementP,
    void* parameterP,
    void** resultP);

#ifdef __cplusplus
}
#endif
