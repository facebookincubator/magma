/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {symphony_agent} from '@fbcnms/magma-api';

export const RAW_AGENT: symphony_agent = {
  description: 'The agent running in the docker engine in the Boston lab',
  device: {
    hardware_id: 'faceb00c-face-b00c-face-000c2940b2bf',
    key: {
      key_type: 'ECHO',
    },
  },
  id: 'fbbosfbcdockerengine',
  magmad: {
    autoupgrade_enabled: false,
    autoupgrade_poll_interval: 300,
    checkin_interval: 15,
    checkin_timeout: 12,
    dynamic_services: [],
  },
  managed_devices: ['ping_fb_dns_from_lab', 'ping_google_ipv6'],
  name: 'FBC Boston Lab',
  status: {
    cert_expiration_time: 1574698700,
    checkin_time: 1574351656724,
    hardware_id: 'faceb00c-face-b00c-face-000c2940b2bf',
    machine_info: {
      cpu_info: {
        architecture: 'x86_64',
        core_count: 16,
        model_name: 'Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz',
        threads_per_core: 1,
      },
      network_info: {
        network_interfaces: [
          {
            ip_addresses: ['127.0.0.1/8'],
            mac_address: '00:00:00:00:00:00',
            network_interface_id: 'lo',
            status: 'UP',
          },
          {
            ip_addresses: ['10.1.128.62/16'],
            mac_address: '00:0c:29:40:b2:be',
            network_interface_id: 'ens160',
            status: 'UP',
          },
          {
            ip_addresses: ['172.18.0.1/16'],
            mac_address: '02:42:f8:2a:4a:3f',
            network_interface_id: 'br-1d31df74735e',
            status: 'UP',
          },
          {
            ip_addresses: ['172.17.0.1/16'],
            mac_address: '02:42:c3:0c:89:b0',
            network_interface_id: 'docker0',
            status: 'UP',
          },
        ],
        routing_table: [
          {
            destination_ip: '0.0.0.0',
            gateway_ip: '10.1.0.1',
            genmask: '0.0.0.0',
            network_interface_id: 'ens160',
          },
          {
            destination_ip: '10.1.0.0',
            gateway_ip: '0.0.0.0',
            genmask: '255.255.0.0',
            network_interface_id: 'ens160',
          },
          {
            destination_ip: '10.1.0.1',
            gateway_ip: '0.0.0.0',
            genmask: '255.255.255.255',
            network_interface_id: 'ens160',
          },
          {
            destination_ip: '172.17.0.0',
            gateway_ip: '0.0.0.0',
            genmask: '255.255.0.0',
            network_interface_id: 'docker0',
          },
          {
            destination_ip: '172.18.0.0',
            gateway_ip: '0.0.0.0',
            genmask: '255.255.0.0',
            network_interface_id: 'br-1d31df74735e',
          },
        ],
      },
    },
    meta: {
      devmand:
        '{"ping_google_ipv6":{"fbc-symphony-device:system":{"status":"UP","latencies":{"latency":[{"rtt":12296,"dst":"device","src":"agent","type":"ping"}]},"geo-location":{"height":0,"longitude":0,"latitude":0,"reference-frame":{"geodetic-system":{"geodetic-datum":"wgs-84"},"astronomical-body":"earth"}}}},"ping_fb_dns_from_lab":{"fbc-symphony-device:system":{"geo-location":{"reference-frame":{"astronomical-body":"earth","geodetic-system":{"geodetic-datum":"wgs-84"}},"latitude":0,"longitude":0,"height":0},"latencies":{"latency":[{"type":"ping","src":"agent","dst":"device","rtt":11797}]},"status":"UP"}}}',
    },
    platform_info: {
      config_info: {
        mconfig_created_at: 1574351615,
      },
      kernel_version: '5.0.0-32-generic',
      packages: [
        {
          name: 'magma',
          version: '0.0.0',
        },
      ],
      vpn_ip: 'N/A',
    },
    system_status: {
      cpu_idle: 23448782600,
      cpu_system: 6868530,
      cpu_user: 5526870,
      disk_partitions: [
        {
          device: '/dev/sda2',
          free: 89002266624,
          mount_point: '/etc/devman',
          total: 134742020096,
          used: 38851186688,
        },
        {
          device: '/dev/sda2',
          free: 89002266624,
          mount_point: '/etc/resolv.conf',
          total: 134742020096,
          used: 38851186688,
        },
        {
          device: '/dev/sda2',
          free: 89002266624,
          mount_point: '/etc/hostname',
          total: 134742020096,
          used: 38851186688,
        },
        {
          device: '/dev/sda2',
          free: 89002266624,
          mount_point: '/etc/hosts',
          total: 134742020096,
          used: 38851186688,
        },
        {
          device: '/dev/sda2',
          free: 89002266624,
          mount_point: '/var/opt/magma/certs/rootCA.pem',
          total: 134742020096,
          used: 38851186688,
        },
      ],
      mem_available: 9711292416,
      mem_free: 619540480,
      mem_total: 12586868736,
      mem_used: 1977106432,
      swap_free: 2147205120,
      swap_total: 2147479552,
      swap_used: 274432,
      time: 1574351656,
      uptime_secs: 1466771,
    },
  },
  tier: 'default',
};
