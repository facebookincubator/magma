// Copyright (c) 2016-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#include <arpa/inet.h>

#include <folly/GLog.h>

#include <devmand/channels/ping/Channel.h>

namespace devmand {
namespace channels {
namespace ping {

Channel::Channel(Engine& engine_, folly::IPAddress target_)
    : engine(engine_), target(target_) {}

folly::Future<Rtt> Channel::ping() {
  icmphdr hdr = makeIcmpPacket();

  LOG(INFO) << "Sending ping to " << target.str() << " with sequence "
            << hdr.un.echo.sequence;

  return engine.ping(hdr, target);
}

RequestId Channel::getSequence() {
  return ++sequence;
}

icmphdr Channel::makeIcmpPacket() {
  icmphdr hdr{};
  hdr.type = ICMP_ECHO;
  // hdr.un.echo.id = 0;
  hdr.un.echo.sequence = getSequence();
  // hdr.checksum = 0; // Let the kernel fill in the checksum
  return hdr;
}

} // namespace ping
} // namespace channels
} // namespace devmand
