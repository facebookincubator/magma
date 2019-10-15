// Copyright (c) 2016-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#include <arpa/inet.h>
#include <fcntl.h>
#include <netdb.h>
#include <netinet/in.h>
#include <netinet/ip_icmp.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <time.h>
#include <unistd.h>

#include <algorithm>

#include <folly/GLog.h>

#include <devmand/channels/ping/Engine.h>
#include <devmand/utils/Time.h>

namespace devmand {
namespace channels {
namespace ping {

Engine::Engine(
    folly::EventBase& _eventBase,
    const std::chrono::milliseconds& pingTimeout_,
    const std::chrono::milliseconds& timeoutFrequency_)
    : folly::EventHandler(&_eventBase),
      eventBase(_eventBase),
      pingTimeout(pingTimeout_),
      timeoutFrequency(timeoutFrequency_) {
  icmpSocket = ::socket(AF_INET, SOCK_DGRAM, IPPROTO_ICMP);
  if (icmpSocket < 0) {
    throw std::system_error(errno, std::generic_category());
  }

  if (fcntl(icmpSocket, F_SETFL, O_NONBLOCK) < 0) {
    throw std::system_error(errno, std::generic_category());
  }

  folly::EventHandler::changeHandlerFD(
      folly::NetworkSocket::fromFd(icmpSocket));

  registerHandler(folly::EventHandler::READ | folly::EventHandler::PERSIST);
}

Engine::~Engine() {
  unregisterHandler();
}

void Engine::start() {
  // TODO I implement a very simple type of timeout here where ever n
  // milliseconds we walk the pending requests and timeout ones that have
  // exceeded their time. This is neither the most efficient (we walk the entire
  // list) or the most accurate (we only guarentee and eventual timeout not a
  // precise timeout). Given this usecase I dont think either of these are that
  // important but I'm making this note so we know in the future we could
  // improve this by using one of the nice timeout queues that exist.
  EventBaseUtils::scheduleEvery(
      eventBase, [this]() { timeout(); }, timeoutFrequency);
}

void Engine::timeout() {
  sharedOutstandingRequests.withULockPtr([this](auto uOutstandingRequests) {
    auto outstandingRequests = uOutstandingRequests.moveFromUpgradeToWrite();
    LOG(INFO) << "Processing ping timeouts";
    for (auto it = outstandingRequests->begin();
         it != outstandingRequests->end();) {
      utils::TimePoint end = utils::Time::now();
      if ((end - it->second.start) > pingTimeout) {
        LOG(ERROR) << "Ping request timed out";
        it->second.promise.setValue(0);
        it = outstandingRequests->erase(it);
      } else {
        ++it;
      }
    }
  });
}

folly::Future<Rtt> Engine::ping(
    const icmphdr& hdr,
    const folly::IPAddress& destination) {
  return sharedOutstandingRequests.withULockPtr([this, &hdr, &destination](
                                                    auto uOutstandingRequests) {
    auto outstandingRequests = uOutstandingRequests.moveFromUpgradeToWrite();
    auto request = outstandingRequests->emplace(
        std::piecewise_construct,
        std::forward_as_tuple(
            std::make_pair(destination, hdr.un.echo.sequence)),
        std::forward_as_tuple(Request{}));
    if (request.second) {
      sockaddr_storage dst;
      destination.toSockaddrStorage(&dst);
      request.first->second.start = utils::Time::now();
      auto result = sendto(
          icmpSocket,
          &hdr,
          sizeof(hdr),
          0,
          reinterpret_cast<const sockaddr*>(&dst),
          sizeof(dst));
      if (result <= 0) {
        switch (result) {
          case EAGAIN: // case EWOULDBLOCK:
            // TODO if the ping fail because of a kernel buffer I'm not going to
            // implement retry logic as something is filling up the buffers. We
            // should probably alarm if this is the case.
            LOG(ERROR) << "Buffer full so ping failed";
            break;
          default:
            // TODO BOOTCAMP get errno string from syserror
            LOG(ERROR) << "Failed to send packet with errno " << errno;
            break;
        }
        outstandingRequests->erase(request.first);
        return folly::makeFuture<Rtt>(0);
      } else {
        auto ret = request.first->second.promise.getFuture();
        return ret;
      }
    } else {
      LOG(ERROR) << "ICMP Echo Id rollover with outstanding requests";
      return folly::makeFuture<Rtt>(0);
    }
  });
}

IcmpPacket Engine::read() {
  IcmpPacket pkt;
  pkt.success = recvfrom(
                    icmpSocket,
                    &pkt.hdr,
                    sizeof(pkt.hdr),
                    0,
                    reinterpret_cast<sockaddr*>(&pkt.src),
                    &pkt.srcLen) > 0;
  return pkt;
}

void Engine::handlerReady(uint16_t) noexcept {
  // TODO end time isn't really precise here as we don't have a kernel time
  // need to implement kernel timestamping
  utils::TimePoint end = utils::Time::now();
  for (IcmpPacket pkt = read(); pkt.success; pkt = read()) {
    bool processed{false};
    if (pkt.hdr.type == 0 and pkt.hdr.code == 0) {
      sharedOutstandingRequests.withULockPtr([this, &end, &pkt, &processed](
                                                 auto uOutstandingRequests) {
        auto outstandingRequests =
            uOutstandingRequests.moveFromUpgradeToWrite();
        auto request = outstandingRequests->find(std::make_pair(
            folly::IPAddress(reinterpret_cast<sockaddr*>(&pkt.src)),
            pkt.hdr.un.echo.sequence));
        if (request != outstandingRequests->end()) {
          auto duration = std::chrono::duration_cast<std::chrono::microseconds>(
              end - request->second.start);
          LOG(INFO) << "Received ICMP response after " << duration.count()
                    << " microseconds";
          request->second.promise.setValue(duration.count());
          outstandingRequests->erase(request);
          processed = true;
        }
      });
    }

    if (not processed) {
      LOG(INFO) << "Packet received with ICMP type "
                << static_cast<int>(pkt.hdr.type) << " code "
                << static_cast<int>(pkt.hdr.code);
    }
  }
}

} // namespace ping
} // namespace channels
} // namespace devmand
