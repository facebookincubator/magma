// Copyright (c) 2016-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#pragma once

#include <string>

#include <devmand/channels/Engine.h>

namespace devmand {
namespace channels {
namespace snmp {

class Engine final : public channels::Engine {
 public:
  Engine(const std::string& appName);

  Engine() = delete;
  ~Engine() override = default;
  Engine(const Engine&) = delete;
  Engine& operator=(const Engine&) = delete;
  Engine(Engine&&) = delete;
  Engine& operator=(Engine&&) = delete;

 public:
  void run();
  void stopEventually();

 private:
  bool stopping{false};
};

} // namespace snmp
} // namespace channels
} // namespace devmand
