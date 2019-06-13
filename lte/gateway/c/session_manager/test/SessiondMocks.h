/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */
#pragma once

#include <grpc++/grpc++.h>
#include <gtest/gtest.h>
#include <gmock/gmock.h>

#include <lte/protos/policydb.pb.h>
#include <lte/protos/pipelined.grpc.pb.h>
#include <lte/protos/session_manager.grpc.pb.h>

#include <folly/io/async/EventBase.h>

#include "CloudReporter.h"
#include "LocalSessionManagerHandler.h"
#include "PipelinedClient.h"
#include "RuleStore.h"

using grpc::Status;
using ::testing::_;
using ::testing::Return;

namespace magma {
/**
 * Mock handler to mock actual request handling and just test server
 */
class MockPipelined final : public Pipelined::Service {
 public:
  MockPipelined(): Pipelined::Service()
  {
    ON_CALL(*this, AddRule(_, _, _)).WillByDefault(Return(Status::OK));
    ON_CALL(*this, ActivateFlows(_, _, _)).WillByDefault(Return(Status::OK));
    ON_CALL(*this, DeactivateFlows(_, _, _)).WillByDefault(Return(Status::OK));
  }

  MOCK_METHOD3(
    AddRule,
    Status(grpc::ServerContext *, const PolicyRule *, Void *));
  MOCK_METHOD3(
    ActivateFlows,
    Status(
      grpc::ServerContext *,
      const ActivateFlowsRequest *,
      ActivateFlowsResult *));
  MOCK_METHOD3(
    DeactivateFlows,
    Status(
      grpc::ServerContext *,
      const DeactivateFlowsRequest *,
      DeactivateFlowsResult *));
};

class MockPipelinedClient : public PipelinedClient {
 public:
  MockPipelinedClient()
  {
    ON_CALL(*this, deactivate_all_flows(_)).WillByDefault(Return(true));
    ON_CALL(*this, deactivate_flows_for_rules(_, _, _))
      .WillByDefault(Return(true));
    ON_CALL(*this, activate_flows_for_rules(_, _, _, _))
      .WillByDefault(Return(true));
  }

  MOCK_METHOD1(deactivate_all_flows, bool(const std::string &imsi));
  MOCK_METHOD3(
    deactivate_flows_for_rules,
    bool(
      const std::string &imsi,
      const std::vector<std::string> &rule_ids,
      const std::vector<PolicyRule> &dynamic_rules));
  MOCK_METHOD4(
    activate_flows_for_rules,
    bool(
      const std::string &imsi,
      const std::string &ip_addr,
      const std::vector<std::string> &static_rules,
      const std::vector<PolicyRule> &dynamic_rules));
};

/**
 * Mock handler to mock actual request handling and just test server
 */
class MockCentralController final : public CentralSessionController::Service {
 public:
  MOCK_METHOD3(
    CreateSession,
    Status(
      grpc::ServerContext *,
      const CreateSessionRequest *,
      CreateSessionResponse *));

  MOCK_METHOD3(
    UpdateSession,
    Status(
      grpc::ServerContext *,
      const UpdateSessionRequest *,
      UpdateSessionResponse *));

  MOCK_METHOD3(
    TerminateSession,
    Status(
      grpc::ServerContext *,
      const SessionTerminateRequest *,
      SessionTerminateResponse *));
};

class MockCallback {
 public:
  MOCK_METHOD2(
    update_callback,
    void(Status status, const UpdateSessionResponse &));
  MOCK_METHOD2(
    create_callback,
    void(Status status, const CreateSessionResponse &));
};

/**
 * Mock handler to mock actual request handling and just test server
 */
class MockSessionHandler final : public LocalSessionManagerHandler {
 public:
  ~MockSessionHandler() {}

  MOCK_METHOD3(
    ReportRuleStats,
    void(
      grpc::ServerContext *,
      const RuleRecordTable *,
      std::function<void(Status, Void)>));

  MOCK_METHOD3(
    CreateSession,
    void(
      grpc::ServerContext *,
      const LocalCreateSessionRequest *,
      std::function<void(Status, LocalCreateSessionResponse)>));

  MOCK_METHOD3(
    EndSession,
    void(
      grpc::ServerContext *,
      const SubscriberID *,
      std::function<void(Status, LocalEndSessionResponse)>));
};

class MockSessionCloudReporter : public SessionCloudReporter {
  public:
    MOCK_METHOD2(
      report_updates,
      void(
        const UpdateSessionRequest &,
        std::function<void(grpc::Status, UpdateSessionResponse)>));

    MOCK_METHOD2(
      report_create_session,
      void(
        const CreateSessionRequest &,
        std::function<void(Status, CreateSessionResponse)>));

    MOCK_METHOD2(
      report_terminate_session,
      void(
        const SessionTerminateRequest &,
        std::function<void(Status, SessionTerminateResponse)>));

};

} // namespace magma
