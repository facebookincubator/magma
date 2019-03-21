/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package store_test

import (
	"reflect"
	"testing"
	"time"

	"magma/orc8r/cloud/go/protos"
	checkin_store "magma/orc8r/cloud/go/services/checkind/store"
	logger_test_init "magma/orc8r/cloud/go/services/logger/test_init"
	"magma/orc8r/cloud/go/services/magmad"
	magmad_protos "magma/orc8r/cloud/go/services/magmad/protos"
	magmad_test_init "magma/orc8r/cloud/go/services/magmad/test_init"
	"magma/orc8r/cloud/go/test_utils"

	"github.com/stretchr/testify/assert"
)

const testAgHwId = "test_ag_HW_id"

func TestCheckinStore(t *testing.T) {
	magmad_test_init.StartTestService(t)
	logger_test_init.StartTestService(t)
	testNetworkName := "Gateway Checkin Test Network"
	store, err := checkin_store.NewCheckinStore(test_utils.NewMockDatastore())
	assert.NoError(t, err)

	testNetworkId, err := magmad.RegisterNetwork(
		&magmad_protos.MagmadNetworkRecord{Name: testNetworkName},
		"checkind_store_test_network")
	assert.NoError(t, err)

	logicalId, err :=
		magmad.RegisterGateway(testNetworkId, &magmad_protos.AccessGatewayRecord{HwId: &protos.AccessGatewayID{Id: testAgHwId}})
	assert.NoError(t, err)
	assert.NotEqual(t, logicalId, "")

	status := protos.GatewayStatus{
		Time: uint64(time.Now().UnixNano() / int64(time.Millisecond)),
		Checkin: &protos.CheckinRequest{
			GatewayId: testAgHwId,
			Status: &protos.ServiceStatus{
				Meta: map[string]string{
					"hello": "world",
				},
			},
			SystemStatus: &protos.SystemStatus{
				Time:       1495484735606,
				CpuUser:    31498,
				CpuSystem:  8361,
				CpuIdle:    1869111,
				MemTotal:   1016084,
				MemUsed:    54416,
				MemFree:    412772,
				UptimeSecs: 1234,
				SwapTotal:  1016081,
				SwapUsed:   54415,
				SwapFree:   412771,
				DiskPartitions: []*protos.DiskPartition{
					{
						Device:     "/dev/sda1",
						MountPoint: "/",
						Total:      1,
						Used:       2,
						Free:       3,
					},
				},
			},
			PlatformInfo: &protos.PlatformInfo{
				VpnIp: "facebook.com",
				Packages: []*protos.Package{
					{
						Name:    "magma",
						Version: "0.0.0.0",
					},
				},
				KernelVersion:           "42",
				KernelVersionsInstalled: []string{"42", "43"},
			},
			MachineInfo: &protos.MachineInfo{
				CpuInfo: &protos.CPUInfo{
					CoreCount:      4,
					ThreadsPerCore: 1,
					Architecture:   "x86_64",
					ModelName:      "Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz",
				},
				NetworkInfo: &protos.NetworkInfo{
					NetworkInterfaces: []*protos.NetworkInterface{
						{
							NetworkInterfaceId: "gtp_br0",
							Status:             protos.NetworkInterface_UP,
							MacAddress:         "08:00:27:1e:8a:32",
							IpAddresses:        []string{"10.10.10.1"},
							Ipv6Addresses:      []string{"fe80::a00:27ff:fe1e:8332"},
						},
					},
					RoutingTable: []*protos.Route{
						{
							DestinationIp:      "0.0.0.0",
							GatewayIp:          "10.10.10.1",
							Genmask:            "255.255.255.0",
							NetworkInterfaceId: "eth0",
						},
					},
				},
			},
			MagmaPkgVersion:         "",
			VpnIp:                   "",
			KernelVersion:           "",
			KernelVersionsInstalled: []string{},
		},
	}

	err = store.UpdateGatewayStatus(&status)
	assert.NoError(t, err)

	status_req := protos.GatewayStatusRequest{NetworkId: testNetworkId, LogicalId: logicalId}
	read_status, err := store.GetGatewayStatus(&status_req)
	assert.NoError(t, err)

	if !reflect.DeepEqual(status, *read_status) {
		t.Errorf("GW Status Mismatch: %#v != %#v", status, *read_status)
	}

	err = store.DeleteGatewayStatus(&status_req)
	assert.NoError(t, err)

	_, err = store.GetGatewayStatus(&status_req)
	// Error since the gateway is deleted
	assert.Error(t, err)

	status.Time = uint64(time.Now().UnixNano() / int64(time.Millisecond))

	err = store.UpdateGatewayStatus(&status)
	assert.NoError(t, err)

	status_req = protos.GatewayStatusRequest{NetworkId: testNetworkId, LogicalId: logicalId}
	read_status, err = store.GetGatewayStatus(&status_req)
	assert.NoError(t, err)

	if !reflect.DeepEqual(status, *read_status) {
		t.Errorf("GW Status Mismatch: %#v != %#v", status, *read_status)
	}

	err = store.DeleteNetworkTable(testNetworkId)
	// Error since there is still the gateway table
	assert.Error(t, err)

	err = store.DeleteGatewayStatus(&status_req)
	assert.NoError(t, err)

	err = store.DeleteNetworkTable(testNetworkId)
	assert.NoError(t, err)

	_, err = store.GetGatewayStatus(&status_req)
	// Error since the network is deleted
	assert.Error(t, err)
}
