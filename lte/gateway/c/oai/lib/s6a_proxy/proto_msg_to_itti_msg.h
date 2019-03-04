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

#include "lte/protos/s6a_proxy.pb.h"
#include "lte/protos/s6a_proxy.grpc.pb.h"
#include <gmp.h>

extern "C" {
#include "intertask_interface.h"
}

namespace magma {
using namespace lte;

void convert_proto_msg_to_itti_s6a_auth_info_ans(
  AuthenticationInformationAnswer msg,
  s6a_auth_info_ans_t *itti_msg);

void convert_proto_msg_to_itti_s6a_update_location_ans(
  UpdateLocationAnswer msg,
  s6a_update_location_ans_t *itti_msg);
} // namespace magma
