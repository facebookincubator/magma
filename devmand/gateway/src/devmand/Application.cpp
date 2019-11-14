// Copyright (c) 2016-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#include <devmand/Application.h>

#include <chrono>
#include <future>
#include <iomanip>
#include <iostream>
#include <thread>

#include <folly/GLog.h>
#include <folly/executors/GlobalExecutor.h>
#include <folly/executors/IOExecutor.h>
#include <folly/json.h>

#include <devmand/Config.h>
#include <devmand/ErrorHandler.h>
#include <devmand/channels/ping/Engine.h>
#include <devmand/devices/Device.h>
#include <devmand/utils/LifetimeTracker.h>

using namespace std::chrono_literals;

namespace devmand {

using IPVersion = channels::ping::IPVersion;

Application::Application()
    : deviceFactory(*this),
      cartographer(
          [this](const cartography::DeviceConfig& deviceConfig) {
            add(deviceConfig);
          },
          [this](const cartography::DeviceConfig& deviceConfig) {
            del(deviceConfig);
          }) {}

void Application::init() {
  ErrorHandler::executeWithCatch(
      [this]() -> void {
        snmpEngine = addEngine<channels::snmp::Engine>(eventBase, name);
        pingEngine =
            addEngine<channels::ping::Engine>(eventBase, IPVersion::v4);
        pingEngineIpv6 =
            addEngine<channels::ping::Engine>(eventBase, IPVersion::v6);
        cliEngine = addEngine<channels::cli::Engine>();
      },
      [this]() { this->statusCode = EXIT_FAILURE; });
}

channels::snmp::Engine& Application::getSnmpEngine() {
  assert(snmpEngine != nullptr);
  return *snmpEngine;
}

channels::ping::Engine& Application::getPingEngine(IPVersion ipv) {
  if (ipv == IPVersion::v6) {
    assert(pingEngineIpv6 != nullptr);
    return *pingEngineIpv6;
  } else {
    assert(pingEngine != nullptr);
    return *pingEngine;
  }
}

// get the relevant ping engine for the given IP (ipv4 or ipv6)
channels::ping::Engine& Application::getPingEngine(folly::IPAddress ip) {
  if (ip.isV6()) {
    return getPingEngine(IPVersion::v6);
  } else {
    return getPingEngine(IPVersion::v4);
  }
}

std::string Application::getName() const {
  return name;
}

std::string Application::getVersion() const {
  return version;
}

// TODO move to device
void Application::pollDevices() {
  for (auto& device : devices) {
    device.second->updateSharedView(unifiedView);
  }
}

void Application::doDebug() {
  LOG(INFO) << "Debug Information";

  LOG(INFO) << "\tChannel Engines (" << channelEngines.size() << "):";
  for (auto& engine : channelEngines) {
    LOG(INFO) << "\t\t" << engine->getName()
              << ": iterations = " << engine->getNumIterations()
              << ", requests = " << engine->getNumRequests();
    setGauge(
        folly::sformat("channel.{}.engine.iterations", engine->getName()),
        engine->getNumIterations());
    setGauge(
        folly::sformat("channel.{}.engine.requests", engine->getName()),
        engine->getNumRequests());
  }

  LOG(INFO) << "\tDevices (" << devices.size() << "):";
  for (auto& device : devices) {
    LOG(INFO) << "\t\t" << device.second->getId();
  }
  setGauge("device.count", devices.size());

  LOG(INFO) << "\tLiving State Objects: "
            << utils::LifetimeTracker<devices::State>::getLivingCount();
  setGauge(
      "device.living_state_objects",
      utils::LifetimeTracker<devices::State>::getLivingCount());
}

UnifiedView Application::getUnifiedView() {
  UnifiedView cpy;
  unifiedView.withULock([&cpy](auto& map) { cpy = map; });
  return cpy;
}

void Application::scheduleEvery(
    std::function<void()> event,
    const std::chrono::seconds& seconds) {
  eventBase.runInEventBaseThread([this, event, seconds]() {
    ErrorHandler::executeWithCatch([this, &event]() { event(); });

    std::function<void()> recurse = [this, event, seconds]() {
      this->scheduleEvery(event, seconds);
    };

    eventBase.scheduleAt(recurse, eventBase.now() + seconds);
  });
}

void Application::scheduleIn(
    std::function<void()> event,
    const std::chrono::seconds& seconds) {
  eventBase.runInEventBaseThread([this, event, seconds]() {
    std::function<void()> recurse = [this, event]() {
      ErrorHandler::executeWithCatch([this, event]() { event(); });
    };

    eventBase.scheduleAt(recurse, eventBase.now() + seconds);
  });
}

void Application::run() {
  if (statusCode != EXIT_SUCCESS) {
    LOG(ERROR) << "Not running application " << name << ", error on setup.";
    return;
  }

  LOG(INFO) << "Starting " << name << ".";

  ErrorHandler::executeWithCatch([this]() {
    for (auto& service : services) {
      service->start();
    }

    // TODO move this to devices
    scheduleEvery(
        [this]() { pollDevices(); }, std::chrono::seconds(FLAGS_poll_interval));

    if (FLAGS_debug_print_interval != 0) {
      scheduleEvery(
          [this]() { doDebug(); },
          std::chrono::seconds(FLAGS_debug_print_interval));
    }

    setGauge("running", 1);

    eventBase.loopForever();

    setGauge("running", 0);

    for (auto& service : services) {
      service->stop();
    }

    for (auto& service : services) {
      service->wait();
    }
  });

  LOG(INFO) << "Stopping " << name << ".";
}

int Application::status() const {
  return statusCode;
}

void Application::add(const cartography::DeviceConfig& deviceConfig) {
  ErrorHandler::executeWithCatch([this, &deviceConfig]() {
    add(deviceFactory.createDevice(deviceConfig));
    devices[deviceConfig.id]->applyConfig(deviceConfig.yangConfig);
  });
}

void Application::del(const cartography::DeviceConfig& deviceConfig) {
  if (devices.erase(deviceConfig.id) != 1) {
    LOG(ERROR) << "Failed to delete device " << deviceConfig.id;
  }
}

void Application::add(std::unique_ptr<devices::Device>&& device) {
  ErrorHandler::executeWithCatch(
      [this, &device]() {
        devices.emplace(
            device->getId(),
            std::forward<std::unique_ptr<devices::Device>>(device));
      },
      [this]() { this->statusCode = EXIT_FAILURE; });
}

void Application::setGauge(const std::string& key, int value) {
  setGauge(key, static_cast<double>(value), "", "");
}

void Application::setGauge(const std::string& key, size_t value) {
  setGauge(key, static_cast<double>(value), "", "");
}

void Application::setGauge(const std::string& key, unsigned int value) {
  setGauge(key, static_cast<double>(value), "", "");
}

void Application::setGauge(
    const std::string& key,
    long long unsigned int value) {
  setGauge(key, static_cast<double>(value), "", "");
}

void Application::setGauge(const std::string& key, double value) {
  setGauge(key, value, "", "");
}

void Application::setGauge(
    const std::string& key,
    double value,
    const std::string& label_name,
    const std::string& label_value) {
  for (auto& service : services) {
    service->setGauge(key, value, label_name, label_value);
  }
}

void Application::add(std::unique_ptr<Service>&& service) {
  ErrorHandler::executeWithCatch(
      [this, &service]() {
        services.emplace_back(std::forward<std::unique_ptr<Service>>(service));
      },
      [this]() { this->statusCode = EXIT_FAILURE; });
}

void Application::addPlatform(
    const std::string& platform,
    devices::Factory::PlatformBuilder platformBuilder) {
  deviceFactory.addPlatform(platform, platformBuilder);
}

void Application::setDefaultPlatform(
    devices::Factory::PlatformBuilder platformBuilder) {
  deviceFactory.setDefaultPlatform(platformBuilder);
}

void Application::addDeviceDiscoveryMethod(
    const std::shared_ptr<cartography::Method>& method) {
  assert(method != nullptr);
  cartographer.addDeviceDiscoveryMethod(method);
}

folly::EventBase& Application::getEventBase() {
  return eventBase;
}

DhcpdConfig& Application::getDhcpdConfig() {
  return dhcpdConfig;
}

} // namespace devmand
