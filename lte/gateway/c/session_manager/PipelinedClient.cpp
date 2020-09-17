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

#include "PipelinedClient.h"

#include "ServiceRegistrySingleton.h"
#include "magma_logging.h"
#include "GrpcMagmaUtils.h"
#include <google/protobuf/util/time_util.h>

using grpc::Status;

namespace {  // anonymous
using std::experimental::optional;

magma::DeactivateFlowsRequest create_deactivate_req(
    const std::string& imsi, const std::string& ip_addr,
    const std::vector<std::string>& rule_ids,
    const std::vector<magma::PolicyRule>& dynamic_rules,
    const magma::RequestOriginType_OriginType origin_type) {
  magma::DeactivateFlowsRequest req;
  req.mutable_sid()->set_id(imsi);
  req.mutable_ip_addr()->set_address(ip_addr);
  req.mutable_request_origin()->set_type(origin_type);
  auto ids = req.mutable_rule_ids();
  for (const auto& id : rule_ids) {
    ids->Add()->assign(id);
  }
  for (const auto& rule : dynamic_rules) {
    ids->Add()->assign(rule.id());
  }
  return req;
}

magma::ActivateFlowsRequest create_activate_req(
    const std::string& imsi, const std::string& ip_addr,
    const optional<magma::AggregatedMaximumBitrate>& ambr,
    const std::vector<std::string>& static_rules,
    const std::vector<magma::PolicyRule>& dynamic_rules,
    const magma::RequestOriginType_OriginType origin_type) {
  magma::ActivateFlowsRequest req;
  req.mutable_sid()->set_id(imsi);
  req.mutable_ip_addr()->set_address(ip_addr);
  req.mutable_request_origin()->set_type(origin_type);
  if (ambr) {
    // TODO remove log once feature is stable
    MLOG(MINFO) << "Sending AMBR info for " << imsi << ", ip addr=" << ip_addr
                << ", dl=" << ambr->max_bandwidth_dl()
                << ", ul=" << ambr->max_bandwidth_ul();
    req.mutable_apn_ambr()->CopyFrom(*ambr);
  }
  auto ids = req.mutable_rule_ids();
  for (const auto& id : static_rules) {
    ids->Add()->assign(id);
  }
  auto mut_dyn_rules = req.mutable_dynamic_rules();
  for (const auto& dyn_rule : dynamic_rules) {
    mut_dyn_rules->Add()->CopyFrom(dyn_rule);
  }
  return req;
}

magma::UEMacFlowRequest create_add_ue_mac_flow_req(
    const magma::SubscriberID& sid, const std::string& ue_mac_addr,
    const std::string& msisdn, const std::string& ap_mac_addr,
    const std::string& ap_name, const std::uint64_t& pdp_start_time) {
  magma::UEMacFlowRequest req;
  req.mutable_sid()->CopyFrom(sid);
  req.set_mac_addr(ue_mac_addr);
  req.set_msisdn(msisdn);
  req.set_ap_mac_addr(ap_mac_addr);
  req.set_ap_name(ap_name);
  req.set_pdp_start_time(pdp_start_time);
  return req;
}

magma::UEMacFlowRequest create_delete_ue_mac_flow_req(
    const magma::SubscriberID& sid, const std::string& ue_mac_addr) {
  magma::UEMacFlowRequest req;
  req.mutable_sid()->CopyFrom(sid);
  req.set_mac_addr(ue_mac_addr);
  return req;
}

magma::SetupPolicyRequest create_setup_policy_req(
    const std::vector<magma::SessionState::SessionInfo>& infos,
    const std::uint64_t& epoch) {
  magma::SetupPolicyRequest req;
  std::vector<magma::ActivateFlowsRequest> activation_reqs;
  for (auto it = infos.begin(); it != infos.end(); it++) {
    auto gx_activate_req = create_activate_req(
        it->imsi, it->ip_addr, it->ambr, it->static_rules, it->dynamic_rules,
        magma::RequestOriginType::GX);
    activation_reqs.push_back(gx_activate_req);
    if (!it->gy_dynamic_rules.empty()) {
      std::vector<std::string> static_rules;
      auto gy_activate_req = create_activate_req(
          it->imsi, it->ip_addr, {}, static_rules, it->gy_dynamic_rules,
          magma::RequestOriginType::GY);
      activation_reqs.push_back(gy_activate_req);
    }
  }
  auto mut_requests = req.mutable_requests();
  for (const auto& act_req : activation_reqs) {
    mut_requests->Add()->CopyFrom(act_req);
  }
  req.set_epoch(epoch);
  return req;
}

magma::SetupUEMacRequest create_setup_ue_mac_req(
    const std::vector<magma::SessionState::SessionInfo>& infos,
    const std::vector<std::string> ue_mac_addrs,
    const std::vector<std::string> msisdns,
    const std::vector<std::string> apn_mac_addrs,
    const std::vector<std::string> apn_names,
    const std::vector<std::uint64_t> pdp_start_times,
    const std::uint64_t& epoch) {
  magma::SetupUEMacRequest req;
  std::vector<magma::UEMacFlowRequest> activation_reqs;

  for (unsigned i = 0; i < infos.size(); i++) {
    magma::SubscriberID sid;
    sid.set_id(infos[i].imsi);
    auto activate_req = create_add_ue_mac_flow_req(
        sid, ue_mac_addrs[i], msisdns[i], apn_mac_addrs[i], apn_names[i],
        pdp_start_times[i]);
    activation_reqs.push_back(activate_req);
  }
  auto mut_requests = req.mutable_requests();
  for (const auto& act_req : activation_reqs) {
    mut_requests->Add()->CopyFrom(act_req);
  }
  req.set_epoch(epoch);
  return req;
}

magma::UpdateSubscriberQuotaStateRequest create_subscriber_quota_state_req(
    const std::vector<magma::SubscriberQuotaUpdate>& updates) {
  magma::UpdateSubscriberQuotaStateRequest req;
  auto p_updates = req.mutable_updates();
  for (const auto& update : updates) {
    p_updates->Add()->CopyFrom(update);
  }
  return req;
}

}  // namespace

namespace magma {

AsyncPipelinedClient::AsyncPipelinedClient(
    std::shared_ptr<grpc::Channel> channel)
    : stub_(Pipelined::NewStub(channel)) {}

AsyncPipelinedClient::AsyncPipelinedClient()
    : AsyncPipelinedClient(ServiceRegistrySingleton::Instance()->GetGrpcChannel(
          "pipelined", ServiceRegistrySingleton::LOCAL)) {}

bool AsyncPipelinedClient::setup_cwf(
    const std::vector<SessionState::SessionInfo>& infos,
    const std::vector<SubscriberQuotaUpdate>& quota_updates,
    const std::vector<std::string> ue_mac_addrs,
    const std::vector<std::string> msisdns,
    const std::vector<std::string> apn_mac_addrs,
    const std::vector<std::string> apn_names,
    const std::vector<std::uint64_t> pdp_start_times,
    const std::uint64_t& epoch,
    std::function<void(Status status, SetupFlowsResult)> callback) {
  SetupPolicyRequest setup_policy_req = create_setup_policy_req(infos, epoch);
  setup_policy_rpc(setup_policy_req, callback);

  SetupUEMacRequest setup_ue_mac_req = create_setup_ue_mac_req(
      infos, ue_mac_addrs, msisdns, apn_mac_addrs, apn_names,
      pdp_start_times, epoch);
  setup_ue_mac_rpc(setup_ue_mac_req, callback);

  update_subscriber_quota_state(quota_updates);
  return true;
}

bool AsyncPipelinedClient::setup_lte(
    const std::vector<SessionState::SessionInfo>& infos,
    const std::uint64_t& epoch,
    std::function<void(Status status, SetupFlowsResult)> callback) {
  SetupPolicyRequest setup_policy_req = create_setup_policy_req(infos, epoch);
  setup_policy_rpc(setup_policy_req, callback);
  return true;
}

bool AsyncPipelinedClient::deactivate_all_flows(const std::string& imsi) {
  DeactivateFlowsRequest req;
  req.mutable_sid()->set_id(imsi);
  MLOG(MDEBUG) << "Deactivating all flows for subscriber " << imsi;
  deactivate_flows_rpc(req, [imsi](Status status, DeactivateFlowsResult resp) {
    if (!status.ok()) {
      MLOG(MERROR) << "Could not deactivate flows for subscriber " << imsi
                   << ": " << status.error_message();
    }
  });
  return true;
}

bool AsyncPipelinedClient::deactivate_flows_for_rules(
    const std::string& imsi, const std::string& ip_addr,
    const std::vector<std::string>& rule_ids,
    const std::vector<PolicyRule>& dynamic_rules,
    const RequestOriginType_OriginType origin_type) {
  auto req = create_deactivate_req(
      imsi, ip_addr, rule_ids, dynamic_rules, origin_type);
  MLOG(MDEBUG) << "Deactivating " << rule_ids.size() << " static rules and "
               << dynamic_rules.size() << " dynamic rules for subscriber "
               << imsi << " IP " << ip_addr;
  deactivate_flows_rpc(req, [imsi](Status status, DeactivateFlowsResult resp) {
    if (!status.ok()) {
      MLOG(MERROR) << "Could not deactivate flows for subscriber " << imsi
                   << ": " << status.error_message();
    }
  });
  return true;
}

bool AsyncPipelinedClient::activate_flows_for_rules(
    const std::string& imsi, const std::string& ip_addr,
    const optional<AggregatedMaximumBitrate>& ambr,
    const std::vector<std::string>& static_rules,
    const std::vector<PolicyRule>& dynamic_rules,
    std::function<void(Status status, ActivateFlowsResult)> callback) {
  MLOG(MDEBUG) << "Activating " << static_rules.size() << " static rules and "
               << dynamic_rules.size() << " dynamic rules for " << imsi
               << " and ip " << ip_addr;
  // Activate static rules and dynamic rules separately until bug is fixed in
  // pipelined which crashes if activated at the same time
  auto static_req = create_activate_req(
      imsi, ip_addr, ambr, static_rules, std::vector<PolicyRule>(),
      RequestOriginType::GX);
  activate_flows_rpc(static_req, callback);
  auto dynamic_req = create_activate_req(
      imsi, ip_addr, ambr, std::vector<std::string>(), dynamic_rules,
      RequestOriginType::GX);
  activate_flows_rpc(dynamic_req, callback);
  return true;
}

bool AsyncPipelinedClient::add_ue_mac_flow(
    const SubscriberID& sid, const std::string& ue_mac_addr,
    const std::string& msisdn, const std::string& ap_mac_addr,
    const std::string& ap_name,
    std::function<void(Status status, FlowResponse)> callback) {
  auto req = create_add_ue_mac_flow_req(
      sid, ue_mac_addr, msisdn, ap_mac_addr, ap_name, 0);
  add_ue_mac_flow_rpc(req, callback);
  return true;
}

bool AsyncPipelinedClient::update_ipfix_flow(
    const SubscriberID& sid, const std::string& ue_mac_addr,
    const std::string& msisdn, const std::string& ap_mac_addr,
    const std::string& ap_name, const uint64_t& pdp_start_time) {
  auto req = create_add_ue_mac_flow_req(
      sid, ue_mac_addr, msisdn, ap_mac_addr, ap_name, pdp_start_time);
  update_ipfix_flow_rpc(req, [ue_mac_addr](Status status, FlowResponse resp) {
    if (!status.ok()) {
      MLOG(MERROR) << "Could not update ipfix flow for subscriber with MAC"
                   << ue_mac_addr << ": " << status.error_message();
    }
  });
  return true;
}

bool AsyncPipelinedClient::delete_ue_mac_flow(
    const SubscriberID& sid, const std::string& ue_mac_addr) {
  auto req = create_delete_ue_mac_flow_req(sid, ue_mac_addr);
  delete_ue_mac_flow_rpc(req, [ue_mac_addr](Status status, FlowResponse resp) {
    if (!status.ok()) {
      MLOG(MERROR) << "Could not delete flow for subscriber with UE MAC"
                   << ue_mac_addr << ": " << status.error_message();
    }
  });
  return true;
}

bool AsyncPipelinedClient::update_subscriber_quota_state(
    const std::vector<SubscriberQuotaUpdate>& updates) {
  auto req = create_subscriber_quota_state_req(updates);
  update_subscriber_quota_state_rpc(req, [](Status status, FlowResponse resp) {
    if (!status.ok()) {
      MLOG(MERROR) << "Could send quota update " << status.error_message();
    }
  });
  return true;
}

bool AsyncPipelinedClient::add_gy_final_action_flow(
    const std::string& imsi, const std::string& ip_addr,
    const std::vector<std::string>& static_rules,
    const std::vector<PolicyRule>& dynamic_rules) {
  MLOG(MDEBUG) << "Activating GY final action for subscriber " << imsi;
  auto static_req = create_activate_req(
      imsi, ip_addr, {}, static_rules, std::vector<PolicyRule>(),
      RequestOriginType::GY);
  activate_flows_rpc(
      static_req, [imsi](Status status, ActivateFlowsResult resp) {
        if (!status.ok()) {
          MLOG(MERROR) << "Could not activate flows through pipelined for UE "
                       << imsi << ": " << status.error_message();
        }
      });
  auto dynamic_req = create_activate_req(
      imsi, ip_addr, {}, std::vector<std::string>(), dynamic_rules,
      RequestOriginType::GY);
  activate_flows_rpc(
      dynamic_req, [imsi](Status status, ActivateFlowsResult resp) {
        if (!status.ok()) {
          MLOG(MERROR) << "Could not activate flows through pipelined for UE "
                       << imsi << ": " << status.error_message();
        }
      });
  return true;
}

void AsyncPipelinedClient::setup_policy_rpc(
    const SetupPolicyRequest& request,
    std::function<void(Status, SetupFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<SetupFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncSetupPolicyFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::setup_ue_mac_rpc(
    const SetupUEMacRequest& request,
    std::function<void(Status, SetupFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<SetupFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncSetupUEMacFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::deactivate_flows_rpc(
    const DeactivateFlowsRequest& request,
    std::function<void(Status, DeactivateFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<DeactivateFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncDeactivateFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::activate_flows_rpc(
    const ActivateFlowsRequest& request,
    std::function<void(Status, ActivateFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<ActivateFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(
      stub_->AsyncActivateFlows(local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::add_ue_mac_flow_rpc(
    const UEMacFlowRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(
      stub_->AsyncAddUEMacFlow(local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::update_ipfix_flow_rpc(
    const UEMacFlowRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncUpdateIPFIXFlow(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::delete_ue_mac_flow_rpc(
    const UEMacFlowRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncDeleteUEMacFlow(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::update_subscriber_quota_state_rpc(
    const UpdateSubscriberQuotaStateRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT);
  PrintGrpcMessage(
      static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(
      std::move(stub_->AsyncUpdateSubscriberQuotaState(
          local_resp->get_context(), request, &queue_)));
}

}  // namespace magma
