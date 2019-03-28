/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package config_test

import (
	"testing"

	lteplugin "magma/lte/cloud/go/plugin"
	"magma/lte/cloud/go/protos/mconfig"
	cellular_config "magma/lte/cloud/go/services/cellular/config"
	"magma/lte/cloud/go/services/cellular/test_utils"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/pluginimpl"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/config"
	config_test_init "magma/orc8r/cloud/go/services/config/test_init"
	dnsd_config "magma/orc8r/cloud/go/services/dnsd/config"
	dnsd_protos "magma/orc8r/cloud/go/services/dnsd/protos"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestCellularBuilder_Build(t *testing.T) {
	plugin.RegisterPluginForTests(t, &lteplugin.LteOrchestratorPlugin{})
	plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	config_test_init.StartTestService(t)
	builder := &cellular_config.CellularBuilder{}
	actual, err := builder.Build("network", "gw1")
	assert.NoError(t, err)
	assert.Equal(t, map[string]proto.Message{}, actual)

	err = config.CreateConfig("network", cellular_config.CellularNetworkType, "network", test_utils.NewDefaultTDDNetworkConfig())
	assert.NoError(t, err)
	err = config.CreateConfig("network", dnsd_config.DnsdNetworkType, "network", &dnsd_protos.NetworkDNSConfig{EnableCaching: false, LocalTTL: 0})
	assert.NoError(t, err)
	err = config.CreateConfig("network", cellular_config.CellularGatewayType, "gw1", test_utils.NewDefaultGatewayConfig())
	assert.NoError(t, err)

	actual, err = builder.Build("network", "gw1")
	assert.NoError(t, err)

	expected := map[string]proto.Message{
		"enodebd": &mconfig.EnodebD{
			LogLevel:               protos.LogLevel_INFO,
			Earfcndl:               44590,
			SubframeAssignment:     2,
			SpecialSubframePattern: 7,
			Pci:                    260,
			TddConfig: &mconfig.EnodebD_TDDConfig{
				Earfcndl:               44590,
				SubframeAssignment:     2,
				SpecialSubframePattern: 7,
			},
			BandwidthMhz:        20,
			AllowEnodebTransmit: true,
			Tac:                 1,
			PlmnidList:          "00101",
			CsfbRat:             mconfig.EnodebD_CSFBRAT_2G,
			Arfcn_2G:            []int32{},
		},
		"mobilityd": &mconfig.MobilityD{
			LogLevel: protos.LogLevel_INFO,
			IpBlock:  "192.168.128.0/24",
		},
		"mme": &mconfig.MME{
			LogLevel:                 protos.LogLevel_INFO,
			Mcc:                      "001",
			Mnc:                      "01",
			Tac:                      1,
			MmeCode:                  1,
			MmeGid:                   1,
			NonEpsServiceControl:     mconfig.MME_NON_EPS_SERVICE_CONTROL_OFF,
			CsfbMcc:                  "",
			CsfbMnc:                  "",
			Lac:                      1,
			RelayEnabled:             false,
			CloudSubscriberdbEnabled: false,
		},
		"pipelined": &mconfig.PipelineD{
			LogLevel:      protos.LogLevel_INFO,
			UeIpBlock:     "192.168.128.0/24",
			NatEnabled:    true,
			DefaultRuleId: "",
			Services: []mconfig.PipelineD_NetworkServices{
				mconfig.PipelineD_METERING,
				mconfig.PipelineD_DPI,
				mconfig.PipelineD_ENFORCEMENT,
			},
		},
		"subscriberdb": &mconfig.SubscriberDB{
			LogLevel:     protos.LogLevel_INFO,
			LteAuthOp:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			LteAuthAmf:   []byte("\x80\x00"),
			SubProfiles:  map[string]*mconfig.SubscriberDB_SubscriptionProfile{},
			RelayEnabled: false,
		},
		"policydb": &mconfig.PolicyDB{
			LogLevel: protos.LogLevel_INFO,
		},
		"sessiond": &mconfig.SessionD{
			LogLevel:     protos.LogLevel_INFO,
			RelayEnabled: false,
		},
	}
	assert.Equal(t, expected, actual)
}

// Should still stream even if no dnsd config exists
func TestCellularBuilder_Build_NullDnsdConfig(t *testing.T) {
	plugin.RegisterPluginForTests(t, &lteplugin.LteOrchestratorPlugin{})
	plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	config_test_init.StartTestService(t)
	builder := &cellular_config.CellularBuilder{}
	actual, err := builder.Build("network", "gw1")
	assert.NoError(t, err)
	assert.Equal(t, map[string]proto.Message{}, actual)

	err = config.CreateConfig("network", cellular_config.CellularNetworkType, "network", test_utils.NewDefaultTDDNetworkConfig())
	assert.NoError(t, err)
	err = config.CreateConfig("network", cellular_config.CellularGatewayType, "gw1", test_utils.NewDefaultGatewayConfig())
	assert.NoError(t, err)

	actual, err = builder.Build("network", "gw1")
	assert.NoError(t, err)

	expected := map[string]proto.Message{
		"enodebd": &mconfig.EnodebD{
			LogLevel:               protos.LogLevel_INFO,
			Earfcndl:               44590,
			SubframeAssignment:     2,
			SpecialSubframePattern: 7,
			Pci:                    260,
			TddConfig: &mconfig.EnodebD_TDDConfig{
				Earfcndl:               44590,
				SubframeAssignment:     2,
				SpecialSubframePattern: 7,
			},
			BandwidthMhz:        20,
			AllowEnodebTransmit: true,
			Tac:                 1,
			PlmnidList:          "00101",
			CsfbRat:             mconfig.EnodebD_CSFBRAT_2G,
			Arfcn_2G:            []int32{},
		},
		"mobilityd": &mconfig.MobilityD{
			LogLevel: protos.LogLevel_INFO,
			IpBlock:  "192.168.128.0/24",
		},
		"mme": &mconfig.MME{
			LogLevel:                 protos.LogLevel_INFO,
			Mcc:                      "001",
			Mnc:                      "01",
			Tac:                      1,
			MmeCode:                  1,
			MmeGid:                   1,
			NonEpsServiceControl:     mconfig.MME_NON_EPS_SERVICE_CONTROL_OFF,
			CsfbMcc:                  "",
			CsfbMnc:                  "",
			Lac:                      1,
			RelayEnabled:             false,
			CloudSubscriberdbEnabled: false,
		},
		"pipelined": &mconfig.PipelineD{
			LogLevel:      protos.LogLevel_INFO,
			UeIpBlock:     "192.168.128.0/24",
			NatEnabled:    true,
			DefaultRuleId: "",
			Services: []mconfig.PipelineD_NetworkServices{
				mconfig.PipelineD_METERING,
				mconfig.PipelineD_DPI,
				mconfig.PipelineD_ENFORCEMENT,
			},
		},
		"subscriberdb": &mconfig.SubscriberDB{
			LogLevel:     protos.LogLevel_INFO,
			LteAuthOp:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			LteAuthAmf:   []byte("\x80\x00"),
			SubProfiles:  map[string]*mconfig.SubscriberDB_SubscriptionProfile{},
			RelayEnabled: false,
		},
		"policydb": &mconfig.PolicyDB{
			LogLevel: protos.LogLevel_INFO,
		},
		"sessiond": &mconfig.SessionD{
			LogLevel:     protos.LogLevel_INFO,
			RelayEnabled: false,
		},
	}
	assert.Equal(t, expected, actual)

}
