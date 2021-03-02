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
// Preparation of Set Session request to UPF
magma::SessionSet create_session_set_req(
    magma::SessionState::SessionInfo info) {
  magma::SessionSet req;
  magma::lte::Fsm_state_FsmState state         = info.state;
  std::string subscriber_id                    = info.subscriber_id;
  uint32_t sess_ver_no                         = info.ver_no;
  magma::SessionState::SessionInfo::NodeId tmp = info.nodeId;
  std::string node_id                          = tmp.node_id;
  req.set_subscriber_id(subscriber_id);
  req.set_session_version(sess_ver_no);
  req.set_local_f_teid(info.local_f_teid);
  req.mutable_node_id()->set_node_id(node_id);
  req.mutable_node_id()->set_node_id_type(magma::NodeID::IPv4);
  req.mutable_state()->set_state(state);

  std::vector<magma::SetGroupPDR> pdr_reqs;
  std::vector<magma::SetGroupFAR> far_reqs;
  pdr_reqs = info.Pdr_rules_;

  auto mut_pdr_requests = req.mutable_set_gr_pdr();
  for (const auto& final_req : pdr_reqs) {
    mut_pdr_requests->Add()->CopyFrom(final_req);
  }
  return req;
}

magma::DeactivateFlowsRequest create_deactivate_req(
    const std::string& imsi, const std::string& ip_addr,
    const std::string& ipv6_addr, const magma::Teids teids,
    const std::vector<std::string>& rule_ids,
    const std::vector<magma::PolicyRule>& dynamic_rules,
    const magma::RequestOriginType_OriginType origin_type,
    const bool remove_default_drop_rules) {
  magma::DeactivateFlowsRequest req;
  req.mutable_sid()->set_id(imsi);
  req.set_ip_addr(ip_addr);
  req.set_ipv6_addr(ipv6_addr);
  req.set_downlink_tunnel(teids.enb_teid());
  req.set_uplink_tunnel(teids.agw_teid());
  req.set_remove_default_drop_flows(remove_default_drop_rules);
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
    const std::string& ipv6_addr, const magma::Teids teids,
    const std::string& msisdn,
    const optional<magma::AggregatedMaximumBitrate>& ambr,
    const std::vector<std::string>& static_rules,
    const std::vector<magma::PolicyRule>& dynamic_rules,
    const magma::RequestOriginType_OriginType origin_type) {
  magma::ActivateFlowsRequest req;
  req.mutable_sid()->set_id(imsi);
  req.set_ip_addr(ip_addr);
  req.set_ipv6_addr(ipv6_addr);
  req.set_downlink_tunnel(teids.enb_teid());
  req.set_uplink_tunnel(teids.agw_teid());
  req.set_msisdn(msisdn);
  req.mutable_request_origin()->set_type(origin_type);
  if (ambr) {
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

magma::SetupDefaultRequest create_setup_default_req(
    const std::uint64_t& epoch) {
  magma::SetupDefaultRequest req;
  req.set_epoch(epoch);
  return req;
}

magma::SetupPolicyRequest create_setup_policy_req(
    const std::vector<magma::SessionState::SessionInfo>& infos,
    const std::uint64_t& epoch) {
  magma::SetupPolicyRequest req;
  std::vector<magma::ActivateFlowsRequest> activation_reqs;
  for (auto it = infos.begin(); it != infos.end(); it++) {
    auto gx_activate_req = create_activate_req(
        it->imsi, it->ip_addr, it->ipv6_addr, it->teids, it->msisdn, it->ambr,
        it->static_rules, it->dynamic_rules, magma::RequestOriginType::GX);
    activation_reqs.push_back(gx_activate_req);
    if (!it->gy_dynamic_rules.empty()) {
      std::vector<std::string> static_rules;
      auto gy_activate_req = create_activate_req(
          it->imsi, it->ip_addr, it->ipv6_addr, it->teids, it->msisdn, {},
          static_rules, it->gy_dynamic_rules, magma::RequestOriginType::GY);
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
    : stub_(Pipelined::NewStub(channel)), ongoing_request_count_(0) {
  teid = M5G_MIN_TEID;
}

AsyncPipelinedClient::AsyncPipelinedClient()
    : AsyncPipelinedClient(ServiceRegistrySingleton::Instance()->GetGrpcChannel(
          "pipelined", ServiceRegistrySingleton::LOCAL)) {}

inline std::ostream& operator<<(
    std::ostream& s, const PendingRequest& pending_req) {
  std::string sid;
  std::string rule_ids = "{ ";
  switch (pending_req.request_type) {
    case ACTIVATE: {
      const ActivateFlowsRequest& req = pending_req.activate_req.request;
      sid                             = req.sid().id();
      std::string rule_ids            = "{ ";
      for (const auto& rule_id : req.rule_ids()) {
        rule_ids += rule_id + " ";
      }
      for (const auto& rule : req.dynamic_rules()) {
        rule_ids += rule.id() + " ";
      }
      break;
    }
    case DEACTIVATE:
      break;
  }
  rule_ids += "}";

  s << sid << " with rules " << rule_ids << " expiring at "
    << pending_req.expiry_time;
  return s;
}

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
  SetupDefaultRequest setup_default_req = create_setup_default_req(epoch);
  setup_default_controllers_rpc(setup_default_req, callback);
  SetupPolicyRequest setup_policy_req = create_setup_policy_req(infos, epoch);
  setup_policy_rpc(setup_policy_req, callback);

  SetupUEMacRequest setup_ue_mac_req = create_setup_ue_mac_req(
      infos, ue_mac_addrs, msisdns, apn_mac_addrs, apn_names, pdp_start_times,
      epoch);
  setup_ue_mac_rpc(setup_ue_mac_req, callback);

  update_subscriber_quota_state(quota_updates);
  return true;
}

bool AsyncPipelinedClient::setup_lte(
    const std::vector<SessionState::SessionInfo>& infos,
    const std::uint64_t& epoch,
    std::function<void(Status status, SetupFlowsResult)> callback) {
  SetupDefaultRequest setup_default_req = create_setup_default_req(epoch);
  setup_default_controllers_rpc(setup_default_req, callback);
  SetupPolicyRequest setup_policy_req = create_setup_policy_req(infos, epoch);
  setup_policy_rpc(setup_policy_req, callback);
  return true;
}

// Method to Setup UPF Session
bool AsyncPipelinedClient::set_upf_session(
    const SessionState::SessionInfo info,
    std::function<void(Status status, UPFSessionContextState)> callback) {
  SessionSet setup_session_req = create_session_set_req(info);
  set_upf_session_rpc(setup_session_req, callback);
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

bool AsyncPipelinedClient::deactivate_flows_for_rules_for_termination(
    const std::string& imsi, const std::string& ip_addr,
    const std::string& ipv6_addr, const Teids teids,
    const std::vector<std::string>& rule_ids,
    const std::vector<PolicyRule>& dynamic_rules,
    const RequestOriginType_OriginType origin_type) {
  MLOG(MDEBUG) << "Deactivating " << rule_ids.size() << " static rules and "
               << dynamic_rules.size()
               << " dynamic rules and default drop flows "
                  "for subscriber "
               << imsi << " IP " << ip_addr << " " << ipv6_addr;

  auto req = create_deactivate_req(
      imsi, ip_addr, ipv6_addr, teids, rule_ids, dynamic_rules, origin_type,
      true);
  return deactivate_flows(req);
}

bool AsyncPipelinedClient::deactivate_flows_for_rules(
    const std::string& imsi, const std::string& ip_addr,
    const std::string& ipv6_addr, const Teids teids,
    const std::vector<std::string>& rule_ids,
    const std::vector<PolicyRule>& dynamic_rules,
    const RequestOriginType_OriginType origin_type) {
  MLOG(MDEBUG) << "Deactivating " << rule_ids.size() << " static rules and "
               << dynamic_rules.size() << " dynamic rules for subscriber "
               << imsi << " IP " << ip_addr << " " << ipv6_addr;

  auto req = create_deactivate_req(
      imsi, ip_addr, ipv6_addr, teids, rule_ids, dynamic_rules, origin_type,
      false);
  return deactivate_flows(req);
}

bool AsyncPipelinedClient::deactivate_flows(DeactivateFlowsRequest& request) {
  auto imsi = request.sid().id();
  deactivate_flows_rpc(
      request, [imsi](Status status, DeactivateFlowsResult resp) {
        if (!status.ok()) {
          MLOG(MERROR) << "Could not deactivate flows for subscriber " << imsi
                       << ": " << status.error_message();
        }
      });
  return true;
}

bool AsyncPipelinedClient::activate_flows_for_rules(
    const std::string& imsi, const std::string& ip_addr,
    const std::string& ipv6_addr, const Teids teids, const std::string& msisdn,
    const optional<AggregatedMaximumBitrate>& ambr,
    const std::vector<std::string>& static_rules,
    const std::vector<PolicyRule>& dynamic_rules,
    std::function<void(Status status, ActivateFlowsResult)> callback) {
  MLOG(MDEBUG) << "Activating " << static_rules.size() << " static rules and "
               << dynamic_rules.size() << " dynamic rules for " << imsi
               << " msisdn " << msisdn << " and ip " << ip_addr << " "
               << ipv6_addr;
  // TODO: Activate static rules and dynamic rules separately until bug
  //  is fixed in pipelined which crashes if activated at the same time
  if (static_rules.size() > 0) {
    auto static_req = create_activate_req(
        imsi, ip_addr, ipv6_addr, teids, msisdn, ambr, static_rules,
        std::vector<PolicyRule>(), RequestOriginType::GX);
    activate_flows_rpc(static_req, callback);
  }
  if (dynamic_rules.size() > 0) {
    auto dynamic_req = create_activate_req(
        imsi, ip_addr, ipv6_addr, teids, msisdn, ambr,
        std::vector<std::string>(), dynamic_rules, RequestOriginType::GX);
    activate_flows_rpc(dynamic_req, callback);
  }
  // If they are both empty, that means this call is at the start of the
  // session. PipelineD requires at least one call to setup things.
  if (static_rules.size() == 0 && dynamic_rules.size() == 0) {
    auto req = create_activate_req(
        imsi, ip_addr, ipv6_addr, teids, msisdn, ambr, static_rules,
        dynamic_rules, RequestOriginType::GX);
    activate_flows_rpc(req, callback);
  }
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
    const std::string& ipv6_addr, const Teids teids, const std::string& msisdn,
    const std::vector<std::string>& static_rules,
    const std::vector<PolicyRule>& dynamic_rules) {
  MLOG(MDEBUG) << "Activating GY final action for subscriber " << imsi;
  auto static_req = create_activate_req(
      imsi, ip_addr, ipv6_addr, teids, msisdn, {}, static_rules,
      std::vector<PolicyRule>(), RequestOriginType::GY);
  activate_flows_rpc(
      static_req, [imsi](Status status, ActivateFlowsResult resp) {
        if (!status.ok()) {
          MLOG(MERROR) << "Could not activate flows through pipelined for UE "
                       << imsi << ": " << status.error_message();
        }
      });
  auto dynamic_req = create_activate_req(
      imsi, ip_addr, ipv6_addr, teids, msisdn, {}, std::vector<std::string>(),
      dynamic_rules, RequestOriginType::GY);
  activate_flows_rpc(
      dynamic_req, [imsi](Status status, ActivateFlowsResult resp) {
        if (!status.ok()) {
          MLOG(MERROR) << "Could not activate flows through pipelined for UE "
                       << imsi << ": " << status.error_message();
        }
      });
  return true;
}

// RPC definition to Send Set Session request to UPF
void AsyncPipelinedClient::set_upf_session_rpc(
    const SessionSet& request,
    std::function<void(Status, UPFSessionContextState)> callback) {
  auto local_resp = new AsyncLocalResponse<UPFSessionContextState>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(
      stub_->AsyncSetSMFSessions(local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::setup_default_controllers_rpc(
    const SetupDefaultRequest& request,
    std::function<void(Status, SetupFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<SetupFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncSetupDefaultControllers(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::setup_policy_rpc(
    const SetupPolicyRequest& request,
    std::function<void(Status, SetupFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<SetupFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncSetupPolicyFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::setup_ue_mac_rpc(
    const SetupUEMacRequest& request,
    std::function<void(Status, SetupFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<SetupFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncSetupUEMacFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::deactivate_flows_rpc(
    const DeactivateFlowsRequest& request,
    std::function<void(Status, DeactivateFlowsResult)> callback) {
  auto local_resp = new AsyncLocalResponse<DeactivateFlowsResult>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncDeactivateFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::set_rate_limiting_config(const YAML::Node& config) {
  // default values
  ENABLE_RATE_LIMITING = false;
  MAX_ACTIVE_REQS      = 10;
  MAX_QUEUE_LENGTH     = 20;

  if (config["enable_pipelined_client_rate_limiting"].IsDefined()) {
    ENABLE_RATE_LIMITING =
        config["enable_pipelined_client_rate_limiting"].as<bool>();
    if (config["rate_limiting_max_active_requests"].IsDefined()) {
      MAX_ACTIVE_REQS =
          config["rate_limiting_max_active_requests"].as<uint32_t>();
    }
    if (config["rate_limiting_max_queue_length"].IsDefined()) {
      MAX_QUEUE_LENGTH =
          config["rate_limiting_max_queue_length"].as<uint32_t>();
    }
  }
}

void AsyncPipelinedClient::pop_and_send_pending_request() {
  bool found_request = false;
  PendingRequest pending_req;
  {  // critical section
    std::unique_lock<std::mutex> lock(queue_lock_);
    if (!pending_reqs_.empty()) {
      pending_req = pending_reqs_.front();
      pending_reqs_.pop();
      found_request = true;
    }
  }  // end of critical section
  if (found_request) {
    MLOG(MINFO) << "Sending a previously pending request: " << pending_req;
    // the GRPC timeout here should be the remaining time since the request
    // was created
    auto timeout = pending_req.expiry_time - std::time(nullptr);
    switch (pending_req.request_type) {
      case ACTIVATE:
        send_activate_flows_rpc(
            pending_req.activate_req.request,
            pending_req.activate_req.callback_fn, timeout);
        return;
      case DEACTIVATE:
        return;
    }
  }
}

std::vector<PendingRequest>
AsyncPipelinedClient::get_pending_requests_to_cancel() {
  std::vector<PendingRequest> requests_to_cancel;
  std::time_t now = std::time(nullptr);
  // Pop items of the queue until
  // 1. The queue has only elements with expiry_time > now
  // 2. The queue has less than or equal to MAX_QUEUE_LENGTH elements
  while ((!pending_reqs_.empty() && pending_reqs_.front().expiry_time <= now) ||
         pending_reqs_.size() > MAX_QUEUE_LENGTH) {
    requests_to_cancel.push_back(pending_reqs_.front());
    pending_reqs_.pop();
  }
  return requests_to_cancel;
}

void AsyncPipelinedClient::cancel_requests(
    const std::vector<PendingRequest>& requests_to_cancel) {
  grpc::Status cancelled_status = Status::CANCELLED;
  ActivateFlowsResult empty_response;
  for (auto& pending : requests_to_cancel) {
    MLOG(MINFO) << "Cancelling: " << pending;
    pending.activate_req.callback_fn(cancelled_status, empty_response);
  }
}

void AsyncPipelinedClient::send_rpc_with_rate_limiting(
    const PendingRequest& pending_req,
    std::function<void(Status, ActivateFlowsResult)> callback) {
  bool send_request_now = false;
  std::vector<PendingRequest> requests_to_cancel;
  MLOG(MINFO) << "Ongoing FlowsRequests: " << ongoing_request_count_
              << ", Queueing FlowsRequests: " << pending_reqs_.size();
  {  // critical section
    std::unique_lock<std::mutex> lock(queue_lock_);
    requests_to_cancel = get_pending_requests_to_cancel();

    send_request_now = ongoing_request_count_ < MAX_ACTIVE_REQS;
    if (send_request_now) {
      ongoing_request_count_++;
    } else {
      MLOG(MINFO) << "Pushing a new pending request to queue: " << pending_req;
      pending_reqs_.push(pending_req);
    }
  }  // end of critical section

  if (send_request_now) {
    send_activate_flows_rpc(
        pending_req.activate_req.request,
        [this, callback](Status status, ActivateFlowsResult result) {
          // this cb function is called when receiving the response
          ongoing_request_count_--;
          callback(status, result);

          // Clean up the queue before sending a request
          std::vector<PendingRequest> requests_to_cancel;
          {  // critical section
            std::unique_lock<std::mutex> lock(queue_lock_);
            requests_to_cancel = get_pending_requests_to_cancel();
          }  // end of critical section
          cancel_requests(requests_to_cancel);

          pop_and_send_pending_request();
        },
        RESPONSE_TIMEOUT_SEC);
  }
  cancel_requests(requests_to_cancel);
}

void AsyncPipelinedClient::activate_flows_rpc(
    const ActivateFlowsRequest& request,
    std::function<void(Status, ActivateFlowsResult)> callback) {
  if (!ENABLE_RATE_LIMITING) {
    send_activate_flows_rpc(request, callback, RESPONSE_TIMEOUT_SEC);
    return;
  }
  send_rpc_with_rate_limiting(
      PendingRequest(
          request, callback, std::time(nullptr) + RESPONSE_TIMEOUT_SEC),
      callback);
}

void AsyncPipelinedClient::send_activate_flows_rpc(
    const ActivateFlowsRequest& request,
    std::function<void(Status, ActivateFlowsResult)> callback,
    const uint32_t timeout) {
  auto local_resp =
      new AsyncLocalResponse<ActivateFlowsResult>(callback, timeout);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(
      stub_->AsyncActivateFlows(local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::send_deactivate_flows_rpc(
    const DeactivateFlowsRequest& request,
    std::function<void(Status, DeactivateFlowsResult)> callback,
    const uint32_t timeout) {
  auto local_resp =
      new AsyncLocalResponse<DeactivateFlowsResult>(callback, timeout);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncDeactivateFlows(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::add_ue_mac_flow_rpc(
    const UEMacFlowRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(
      stub_->AsyncAddUEMacFlow(local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::update_ipfix_flow_rpc(
    const UEMacFlowRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncUpdateIPFIXFlow(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::delete_ue_mac_flow_rpc(
    const UEMacFlowRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(std::move(stub_->AsyncDeleteUEMacFlow(
      local_resp->get_context(), request, &queue_)));
}

void AsyncPipelinedClient::update_subscriber_quota_state_rpc(
    const UpdateSubscriberQuotaStateRequest& request,
    std::function<void(Status, FlowResponse)> callback) {
  auto local_resp = new AsyncLocalResponse<FlowResponse>(
      std::move(callback), RESPONSE_TIMEOUT_SEC);
  PrintGrpcMessage(static_cast<const google::protobuf::Message&>(request));
  local_resp->set_response_reader(
      std::move(stub_->AsyncUpdateSubscriberQuotaState(
          local_resp->get_context(), request, &queue_)));
}

uint32_t AsyncPipelinedClient::get_next_teid() {
  /* For now TEID we use current no, increment for next, later we plan to
     maintain  release/alloc table for reu sing */
  uint32_t allocated_teid = teid++;
  return allocated_teid;
}

uint32_t AsyncPipelinedClient::get_current_teid() {
  /* For now TEID we use current no, increment for next, later we plan to
     maintain  release/alloc table for reu sing */
  return teid;
}

}  // namespace magma
