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
#include <future>
#include <memory>
#include <utility>

#include <glog/logging.h>
#include <gtest/gtest.h>

#include "SessiondMocks.h"

using ::testing::Test;

namespace magma {

MATCHER_P(CheckCount, count, "") {
  int arg_count = arg.size();
  return arg_count == count;
}

MATCHER_P(CheckCreateSession, imsi, "") {
  auto req = static_cast<const CreateSessionRequest*>(arg);
  return req->common_context().sid().id() == imsi;
}

MATCHER_P(CheckStaticRuleSize, size, "") {
  return arg.static_rules_size() == size;
}

MATCHER_P(CheckUpdateSessionRequestNumber, request_number, "") {
  auto request = static_cast<const UpdateSessionRequest&>(arg);
  for (const auto& credit_usage_update : request.updates()) {
    int req_number = credit_usage_update.request_number();
    return req_number == request_number;
  }
  return false;
}

MATCHER_P(CheckSingleUpdateSession, expected_update, "") {
  auto request = static_cast<const UpdateSessionRequest*>(arg);
  if (request->updates_size() != 1) {
    return false;
  }

  auto& update = request->updates(0);
  bool val =
      update.usage().type() == expected_update.usage().type() &&
      update.usage().bytes_tx() == expected_update.usage().bytes_tx() &&
      update.usage().bytes_rx() == expected_update.usage().bytes_rx() &&
      update.sid() == expected_update.sid() &&
      update.usage().charging_key() == expected_update.usage().charging_key();
  return val;
}

MATCHER_P(CheckTerminateImsi, imsi, "") {
  auto request = static_cast<const SessionTerminateRequest*>(arg);
  return request->sid() == imsi;
}

MATCHER_P(CheckDeactivateFlows, imsi, "") {
  auto request = static_cast<const DeactivateFlowsRequest*>(arg);
  return request->sid().id() == imsi;
}

MATCHER_P2(CheckUpdateRequestCount, monitorCount, chargingCount, "") {
  auto req = static_cast<const UpdateSessionRequest>(arg);
  return req.updates().size() == chargingCount &&
         req.usage_monitors().size() == monitorCount;
}

MATCHER_P3(CheckTerminateRequestCount, imsi, monitorCount, chargingCount, "") {
  auto req = static_cast<const SessionTerminateRequest>(arg);
  return req.sid() == imsi && req.credit_usages().size() == chargingCount &&
         req.monitor_usages().size() == monitorCount;
}

MATCHER_P2(CheckActivateFlows, imsi, rule_count, "") {
  auto request = static_cast<const ActivateFlowsRequest*>(arg);
  return request->sid().id() == imsi && request->rule_ids_size() == rule_count;
}

MATCHER_P2(CheckQuotaUpdateState, size, expected_states, "") {
  auto updates     = static_cast<const std::vector<SubscriberQuotaUpdate>>(arg);
  int updates_size = updates.size();
  if (updates_size != size) {
    return false;
  }
  for (int i = 0; i < updates_size; i++) {
    if (updates[i].update_type() != expected_states[i]) {
      return false;
    }
  }
  return true;
}

MATCHER_P5(
    CheckSessionInfos, imsi_list, ip_address_list, cfg, static_rule_lists,
    dynamic_rule_ids_lists, "") {
  auto infos = static_cast<const std::vector<SessionState::SessionInfo>>(arg);

  if (infos.size() != imsi_list.size()) return false;

  for (size_t i = 0; i < infos.size(); i++) {
    if (infos[i].imsi != imsi_list[i]) return false;
    if (infos[i].ip_addr != ip_address_list[i]) return false;
    if (infos[i].static_rules.size() != static_rule_lists[i].size())
      return false;
    if (infos[i].dynamic_rules.size() != dynamic_rule_ids_lists[i].size())
      return false;
    for (size_t r_index = 0; i < infos[i].static_rules.size(); i++) {
      if (infos[i].static_rules[r_index] != static_rule_lists[i][r_index])
        return false;
    }
    for (size_t r_index = 0; i < infos[i].dynamic_rules.size(); i++) {
      if (infos[i].dynamic_rules[r_index].id() !=
          dynamic_rule_ids_lists[i][r_index])
        return false;
    }
    // check ambr field if config has qos_info
    if (cfg.rat_specific_context.has_lte_context() &&
        cfg.rat_specific_context.lte_context().has_qos_info()) {
      const auto& qos_info = cfg.rat_specific_context.lte_context().qos_info();
      if (!infos[i].ambr) {
        return false;
      } else if (infos[i].ambr->max_bandwidth_ul() != qos_info.apn_ambr_ul()) {
        return false;
      } else if (infos[i].ambr->max_bandwidth_dl() != qos_info.apn_ambr_dl()) {
        return false;
      }
    }
  }
  return true;
}

MATCHER_P(CheckEventType, expectedEventType, "") {
  return (arg.event_type() == expectedEventType);
}

MATCHER_P2(CheckCreateBearerReq, imsi, rule_count, "") {
  auto request = static_cast<const CreateBearerRequest>(arg);
  return request.sid().id() == imsi &&
         request.policy_rules().size() == rule_count;
}

MATCHER_P3(CheckDeleteOneBearerReq, imsi, link_bearer_id, eps_bearer_id, "") {
  auto request = static_cast<const DeleteBearerRequest>(arg);

  return request.sid().id() == imsi &&
         request.link_bearer_id() == uint32_t(link_bearer_id) &&
         request.eps_bearer_ids_size() == 1 &&
         request.eps_bearer_ids(0) == uint32_t(eps_bearer_id);
}

MATCHER_P(CheckSubset, ids, "") {
  auto request = static_cast<const std::vector<std::string>>(arg);
  for (size_t i = 0; i < request.size(); i++) {
    if (ids.find(request[i]) != ids.end()) {
      return true;
    }
  }
  return false;
}

};  // namespace magma