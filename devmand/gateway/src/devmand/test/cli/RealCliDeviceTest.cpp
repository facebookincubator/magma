// Copyright (c) 2019-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#define LOG_WITH_GLOG
#include <magma_logging.h>

#include <boost/algorithm/string/trim.hpp>
#include <devmand/Application.h>
#include <devmand/cartography/DeviceConfig.h>
#include <devmand/channels/cli/CliFlavour.h>
#include <devmand/devices/Datastore.h>
#include <devmand/devices/Device.h>
#include <devmand/devices/cli/PlaintextCliDevice.h>
#include <devmand/test/cli/utils/Log.h>
#include <devmand/test/cli/utils/MockCli.h>
#include <devmand/test/cli/utils/Ssh.h>
#include <folly/executors/ThreadedExecutor.h>
#include <gtest/gtest.h>
#include <chrono>

namespace devmand {
namespace test {
namespace cli {

using namespace devmand::channels::cli;
using devmand::Application;
using devmand::cartography::ChannelConfig;
using devmand::cartography::DeviceConfig;
using devmand::channels::cli::UBIQUITI;
using devmand::devices::Datastore;
using devmand::devices::Device;
using devmand::devices::cli::PlaintextCliDevice;

class RealCliDeviceTest : public ::testing::Test {
 protected:
  unique_ptr<channels::cli::Engine> cliEngine;

  void SetUp() override {
    devmand::test::utils::log::initLog();
    cliEngine = make_unique<channels::cli::Engine>();
  }
};

TEST_F(RealCliDeviceTest, DISABLED_ubiquiti) {
  int i = 0;
  string output = "";
  Application app;
  cartography::DeviceConfig deviceConfig;
  devmand::cartography::ChannelConfig chnlCfg;
  std::map<std::string, std::string> kvPairs;
  kvPairs.insert(std::make_pair("stateCommand", "show mac access-lists"));
  kvPairs.insert(std::make_pair("port", "22"));
  kvPairs.insert(std::make_pair("username", "ubnt"));
  kvPairs.insert(std::make_pair("password", "ubnt"));
  kvPairs.insert(std::make_pair("flavour", UBIQUITI));
  chnlCfg.kvPairs = kvPairs;
  deviceConfig.channelConfigs.insert(std::make_pair("cli", chnlCfg));
  deviceConfig.ip = "10.19.0.245";
  deviceConfig.id = "ubiquiti-test-device";

  std::unique_ptr<devices::Device> dev =
      PlaintextCliDevice::createDeviceWithEngine(app, deviceConfig, *cliEngine);
  do {
    if (i > 0) {
      std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    i++;

    std::shared_ptr<Datastore> state = dev->getOperationalDatastore();
    const folly::dynamic& stateResult = state->collect().get();

    output = stateResult.getDefault(kvPairs.at("stateCommand"), "").asString();
    if (i > 20) {
      FAIL() << "Unable to execute command, probably not connected";
    }
  } while (output.empty());
  EXPECT_EQ("No ACLs are configured", boost::algorithm::trim_copy(output));
}

} // namespace cli
} // namespace test
} // namespace devmand
