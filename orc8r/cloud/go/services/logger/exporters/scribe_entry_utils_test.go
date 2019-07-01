/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package exporters_test

import (
	"os"
	"testing"

	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/pluginimpl"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/serde"
	configuratorti "magma/orc8r/cloud/go/services/configurator/test_init"
	configuratortu "magma/orc8r/cloud/go/services/configurator/test_utils"
	deviceti "magma/orc8r/cloud/go/services/device/test_init"
	"magma/orc8r/cloud/go/services/logger/exporters"
	"magma/orc8r/cloud/go/services/magmad/obsidian/models"

	"github.com/stretchr/testify/assert"
)

func TestScribeEntryUtils(t *testing.T) {
	os.Setenv(orc8r.UseConfiguratorEnv, "1")
	logEntries := []*protos.LogEntry{
		{
			Category:  "test",
			NormalMap: map[string]string{"status": "ACTIVE"},
			IntMap:    map[string]int64{"port": 443},
			Time:      12345,
		},
	}
	scribeEntries, err := exporters.ConvertToScribeLogEntries(logEntries)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(scribeEntries))
	assert.Equal(t, logEntries[0].Category, scribeEntries[0].Category)
	expectedMsg := "{\"int\":{\"port\":443,\"time\":12345},\"normal\":{\"status\":\"ACTIVE\"}}"
	assert.Equal(t, expectedMsg, scribeEntries[0].Message)
}

func TestScribeEntryUtils_WithHWID(t *testing.T) {
	os.Setenv(orc8r.UseConfiguratorEnv, "1")
	configuratorti.StartTestService(t)
	deviceti.StartTestService(t)
	serde.RegisterSerdes(&pluginimpl.GatewayRecordSerde{})

	networkID := "test_network"
	gatewayID := "test_gateway"
	hwID := "test_hwID"
	configuratortu.RegisterNetwork(t, networkID, "")
	configuratortu.RegisterGateway(t, networkID, gatewayID, &models.AccessGatewayRecord{HwID: &models.HwGatewayID{ID: hwID}})

	logEntries := []*protos.LogEntry{
		{
			Category:  "test",
			NormalMap: map[string]string{"status": "ACTIVE"},
			IntMap:    map[string]int64{"port": 443},
			Time:      12345,
			HwId:      hwID,
		},
	}
	scribeEntries, err := exporters.ConvertToScribeLogEntries(logEntries)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(scribeEntries))
	assert.Equal(t, logEntries[0].Category, scribeEntries[0].Category)
	expectedMsg := "{\"int\":{\"port\":443,\"time\":12345},\"normal\":{\"gatewayId\":\"test_gateway\",\"networkId\":\"test_network\",\"status\":\"ACTIVE\"}}"
	assert.Equal(t, expectedMsg, scribeEntries[0].Message)
}
