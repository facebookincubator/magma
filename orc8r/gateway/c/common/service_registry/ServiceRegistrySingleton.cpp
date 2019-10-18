/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */
#include <string>
#include <iostream>
#include <fstream>
#include <sstream>
#include "ServiceRegistrySingleton.h"

using magma::ServiceRegistrySingleton;
using grpc::Channel;
using grpc::CreateCustomChannel;
using grpc::InsecureChannelCredentials;
using grpc::SslCredentials;
using grpc::SslCredentialsOptions;

namespace magma {

ServiceRegistrySingleton* ServiceRegistrySingleton::instance_ = nullptr;

ServiceRegistrySingleton* ServiceRegistrySingleton::Instance() {
  if (instance_ == nullptr) {
    instance_ = new ServiceRegistrySingleton();
  }
  return instance_;
}

YAML::Node ServiceRegistrySingleton::GetProxyConfig() {
  YAML::Node node = service_config_loader_.load_service_config("control_proxy");
  return node;
}

YAML::Node ServiceRegistrySingleton::GetRegistry() {
  YAML::Node node = service_config_loader_.load_service_config(
    "service_registry");
  return node;
}

void ServiceRegistrySingleton::flush() {
  delete instance_;
  instance_ = new ServiceRegistrySingleton();
}

ServiceRegistrySingleton::ServiceRegistrySingleton()
{
  proxy_config_ = std::unique_ptr<YAML::Node>(
    new YAML::Node(ServiceRegistrySingleton::GetProxyConfig()));
  registry_ = std::unique_ptr<YAML::Node>(
    new YAML::Node(ServiceRegistrySingleton::GetRegistry()));
}

ip_port_pair_t ServiceRegistrySingleton::GetServiceAddr(
  const std::string& service
) {
  YAML::Node registry = *(this->registry_);
  YAML::Node node = registry["services"];
  assert(node.IsMap());
  if (node[service]) {
    YAML::Node serviceNode = node[service];
    assert(serviceNode.IsMap());
    ip_port_pair_t ip_port_pair;
    ip_port_pair.ip = serviceNode["ip_address"].as<std::string>();
    ip_port_pair.port = serviceNode["port"].as<std::string>();
    return ip_port_pair;
  } else {
    throw std::invalid_argument("Invalid service name: " + service);
  }

}

std::string ServiceRegistrySingleton::GetServiceAddrString(
    const std::string& service) {
  auto ip_pair = GetServiceAddr(service);
  return ip_pair.ip + ":" + ip_pair.port;
}

std::string
ServiceRegistrySingleton::LoadCertFile(const std::string& file) {
  std::ifstream certFile(file);
  std::stringstream certBuffer;
  certBuffer << certFile.rdbuf();
  return certBuffer.str();
}

std::shared_ptr<ChannelCredentials>
ServiceRegistrySingleton::GetSslCredentials() {
  YAML::Node proxyConfig = *(this->proxy_config_);
  SslCredentialsOptions options;

  options.pem_root_certs = LoadCertFile(
      proxyConfig["rootca_cert"].as<std::string>());
  options.pem_cert_chain = LoadCertFile(
      proxyConfig["gateway_cert"].as<std::string>());
  options.pem_private_key = LoadCertFile(
      proxyConfig["gateway_key"].as<std::string>());

  return SslCredentials(options);
}

const std::shared_ptr<Channel> ServiceRegistrySingleton::GetGrpcChannel(
  const std::string& service,
  const std::string& destination){
    create_grpc_channel_args_t args
      = GetCreateGrpcChannelArgs(service, destination);
    return ServiceRegistrySingleton::CreateGrpcChannel(
      args.ip, args.port, args.authority, args.creds);
}
const create_grpc_channel_args_t
ServiceRegistrySingleton::GetCreateGrpcChannelArgs(
  const std::string& service,
  const std::string& destination) {
    create_grpc_channel_args_t args;
    YAML::Node proxyConfig = *(this->proxy_config_);

    std::string cloud_address = proxyConfig["cloud_address"].as<std::string>();
    // control proxy uses the :authority: HTTP header to route to services
    if (destination.compare(ServiceRegistrySingleton::LOCAL) == 0) {
      args.authority = service + ".local";
    } else {
      args.authority = service + "-" + cloud_address;
    }

    if (destination.compare(ServiceRegistrySingleton::LOCAL) == 0) {
      // connect to local service
      ip_port_pair_t pair = ServiceRegistrySingleton::GetServiceAddr(service);
      args.ip = pair.ip;
      args.port = pair.port;
    } else if (proxyConfig["proxy_cloud_connections"].as<bool>()) {
      // connect to the cloud via local control proxy
      args.ip = "127.0.0.1";
      args.port = proxyConfig["local_port"].as<std::string>();
    } else {
      // connect to the cloud directly
      args.ip = cloud_address;
      args.port = proxyConfig["cloud_port"].as<std::string>();
      args.creds = GetSslCredentials();
    }

    return args;
  }
const std::shared_ptr<Channel> ServiceRegistrySingleton::CreateGrpcChannel(
  const std::string& ip,
  const std::string& port,
  const std::string& authority,
  std::shared_ptr<ChannelCredentials> creds){

  if (creds == nullptr) {
    creds = InsecureChannelCredentials();
  }
  grpc::ChannelArguments arg;

  arg.SetString("grpc.default_authority", authority);
  std::ostringstream ss;
  ss << ip << ":" << port;

  return CreateCustomChannel(ss.str(), creds, arg);
}
}
