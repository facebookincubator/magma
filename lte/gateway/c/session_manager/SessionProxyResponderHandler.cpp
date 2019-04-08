/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */
#include <chrono>
#include <thread>

#include "SessionProxyResponderHandler.h"
#include "magma_logging.h"

using grpc::Status;

namespace magma {

SessionProxyResponderHandlerImpl::SessionProxyResponderHandlerImpl(
  LocalEnforcer *enforcer):
  enforcer_(enforcer)
{
}

void SessionProxyResponderHandlerImpl::ChargingReAuth(
  ServerContext *context,
  const ChargingReAuthRequest *request,
  std::function<void(Status, ChargingReAuthAnswer)> response_callback)
{
  auto &request_cpy = *request;
  enforcer_->get_event_base().runInEventBaseThread(
    [this, request_cpy, response_callback]() {
      auto result = enforcer_->init_charging_reauth(request_cpy);
      ChargingReAuthAnswer ans;
      ans.set_result(result);
      response_callback(Status::OK, ans);
    });
}

void SessionProxyResponderHandlerImpl::PolicyReAuth(
  ServerContext *context,
  const PolicyReAuthRequest *request,
  std::function<void(Status, PolicyReAuthAnswer)> response_callback)
{
  auto &request_cpy = *request;
  enforcer_->get_event_base().runInEventBaseThread(
    [this, request_cpy, response_callback]() {
      PolicyReAuthAnswer ans;
      enforcer_->init_policy_reauth(request_cpy, ans);
      response_callback(Status::OK, ans);
    });
}
} // namespace magma
