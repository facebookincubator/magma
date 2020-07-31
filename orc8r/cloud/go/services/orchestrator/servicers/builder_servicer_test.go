/*
 Copyright 2020 The Magma Authors.

 This source code is licensed under the BSD-style license found in the
 LICENSE file in the root directory of this source tree.

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package servicers_test

import (
	"testing"

	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/pluginimpl"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/configurator/mconfig"
	"magma/orc8r/cloud/go/services/orchestrator"
	"magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	orchestrator_test_init "magma/orc8r/cloud/go/services/orchestrator/test_init"
	"magma/orc8r/lib/go/protos"
	mconfig_protos "magma/orc8r/lib/go/protos/mconfig"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBaseOrchestratorMconfigBuilder_Build(t *testing.T) {
	assert.NoError(t, plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{}))
	orchestrator_test_init.StartTestService(t)

	t.Run("no tier", func(t *testing.T) {
		nw := configurator.Network{ID: "n1"}
		gw := configurator.NetworkEntity{
			Type: orc8r.MagmadGatewayType,
			Key:  "gw1",
			Config: &models.MagmadGatewayConfigs{
				AutoupgradeEnabled:      swag.Bool(true),
				AutoupgradePollInterval: 300,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				DynamicServices:         []string{},
				FeatureFlags:            map[string]bool{},
			},
		}
		graph := configurator.EntityGraph{
			Entities: []configurator.NetworkEntity{gw},
		}

		expected := map[string]proto.Message{
			"control_proxy": &mconfig_protos.ControlProxy{LogLevel: protos.LogLevel_INFO},
			"magmad": &mconfig_protos.MagmaD{
				LogLevel:                protos.LogLevel_INFO,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				AutoupgradeEnabled:      true,
				AutoupgradePollInterval: 300,
				PackageVersion:          "0.0.0-0",
				Images:                  nil,
				DynamicServices:         nil,
				FeatureFlags:            nil,
			},
			"metricsd": &mconfig_protos.MetricsD{LogLevel: protos.LogLevel_INFO},
			"td-agent-bit": &mconfig_protos.FluentBit{
				ExtraTags:        map[string]string{"network_id": "n1", "gateway_id": "gw1"},
				ThrottleRate:     1000,
				ThrottleWindow:   5,
				ThrottleInterval: "1m",
			},
			"eventd": &mconfig_protos.EventD{
				LogLevel:       protos.LogLevel_INFO,
				EventVerbosity: -1,
			},
		}

		actual, err := BuildBaseOrchestrator(&nw, &graph, "gw1")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	// Put a tier in the graph
	t.Run("tiers work correctly", func(t *testing.T) {
		nw := configurator.Network{ID: "n1"}
		gw := configurator.NetworkEntity{
			Type: orc8r.MagmadGatewayType,
			Key:  "gw1",
			Config: &models.MagmadGatewayConfigs{
				AutoupgradeEnabled:      swag.Bool(true),
				AutoupgradePollInterval: 300,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				DynamicServices:         []string{},
				FeatureFlags:            map[string]bool{},
			},
		}
		tier := configurator.NetworkEntity{
			Type: orc8r.UpgradeTierEntityType,
			Key:  "default",
			Config: &models.Tier{
				Name:    "default",
				Version: "1.0.0-0",
				Images: []*models.TierImage{
					{Name: swag.String("Image1"), Order: swag.Int64(42)},
					{Name: swag.String("Image2"), Order: swag.Int64(1)},
				},
			},
		}
		graph := configurator.EntityGraph{
			Entities: []configurator.NetworkEntity{gw, tier},
			Edges: []configurator.GraphEdge{
				{From: tier.GetTypeAndKey(), To: gw.GetTypeAndKey()},
			},
		}

		expected := map[string]proto.Message{
			"control_proxy": &mconfig_protos.ControlProxy{LogLevel: protos.LogLevel_INFO},
			"magmad": &mconfig_protos.MagmaD{
				LogLevel:                protos.LogLevel_INFO,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				AutoupgradeEnabled:      true,
				AutoupgradePollInterval: 300,
				PackageVersion:          "1.0.0-0",
				Images: []*mconfig_protos.ImageSpec{
					{Name: "Image1", Order: 42},
					{Name: "Image2", Order: 1},
				},
				DynamicServices: nil,
				FeatureFlags:    nil,
			},
			"metricsd": &mconfig_protos.MetricsD{LogLevel: protos.LogLevel_INFO},
			"td-agent-bit": &mconfig_protos.FluentBit{
				ExtraTags:        map[string]string{"network_id": "n1", "gateway_id": "gw1"},
				ThrottleRate:     1000,
				ThrottleWindow:   5,
				ThrottleInterval: "1m",
			},
			"eventd": &mconfig_protos.EventD{
				LogLevel:       protos.LogLevel_INFO,
				EventVerbosity: -1,
			},
		}

		actual, err := BuildBaseOrchestrator(&nw, &graph, "gw1")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("set list of files for log aggregation", func(t *testing.T) {
		testThrottleInterval := "30h"
		testThrottleWindow := uint32(808)
		testThrottleRate := uint32(305)

		nw := configurator.Network{ID: "n1"}
		gw := configurator.NetworkEntity{
			Type: orc8r.MagmadGatewayType,
			Key:  "gw1",
			Config: &models.MagmadGatewayConfigs{
				AutoupgradeEnabled:      swag.Bool(true),
				AutoupgradePollInterval: 300,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				DynamicServices:         nil,
				FeatureFlags:            nil,
				Logging: &models.GatewayLoggingConfigs{
					Aggregation: &models.AggregationLoggingConfigs{
						TargetFilesByTag: map[string]string{
							"thing": "/var/log/thing.log",
							"blah":  "/some/directory/blah.log",
						},
						ThrottleRate:     &testThrottleRate,
						ThrottleWindow:   &testThrottleWindow,
						ThrottleInterval: &testThrottleInterval,
					},
					EventVerbosity: swag.Int32(0),
				},
			},
		}
		tier := configurator.NetworkEntity{
			Type: orc8r.UpgradeTierEntityType,
			Key:  "default",
			Config: &models.Tier{
				Name:    "default",
				Version: "1.0.0-0",
				Images: []*models.TierImage{
					{Name: swag.String("Image1"), Order: swag.Int64(42)},
					{Name: swag.String("Image2"), Order: swag.Int64(1)},
				},
			},
		}
		graph := configurator.EntityGraph{
			Entities: []configurator.NetworkEntity{gw, tier},
			Edges: []configurator.GraphEdge{
				{From: tier.GetTypeAndKey(), To: gw.GetTypeAndKey()},
			},
		}

		expected := map[string]proto.Message{
			"control_proxy": &mconfig_protos.ControlProxy{LogLevel: protos.LogLevel_INFO},
			"magmad": &mconfig_protos.MagmaD{
				LogLevel:                protos.LogLevel_INFO,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				AutoupgradeEnabled:      true,
				AutoupgradePollInterval: 300,
				PackageVersion:          "1.0.0-0",
				Images: []*mconfig_protos.ImageSpec{
					{Name: "Image1", Order: 42},
					{Name: "Image2", Order: 1},
				},
				DynamicServices: nil,
				FeatureFlags:    nil,
			},
			"metricsd": &mconfig_protos.MetricsD{LogLevel: protos.LogLevel_INFO},
			"td-agent-bit": &mconfig_protos.FluentBit{
				ExtraTags:        map[string]string{"network_id": "n1", "gateway_id": "gw1"},
				ThrottleRate:     305,
				ThrottleWindow:   808,
				ThrottleInterval: "30h",
				FilesByTag: map[string]string{
					"thing": "/var/log/thing.log",
					"blah":  "/some/directory/blah.log",
				},
			},
			"eventd": &mconfig_protos.EventD{
				LogLevel:       protos.LogLevel_INFO,
				EventVerbosity: 0,
			},
		}

		actual, err := BuildBaseOrchestrator(&nw, &graph, "gw1")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("check default values for log throttling", func(t *testing.T) {
		nw := configurator.Network{ID: "n1"}
		gw := configurator.NetworkEntity{
			Type: orc8r.MagmadGatewayType,
			Key:  "gw1",
			Config: &models.MagmadGatewayConfigs{
				AutoupgradeEnabled:      swag.Bool(true),
				AutoupgradePollInterval: 300,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				DynamicServices:         nil,
				FeatureFlags:            nil,
				Logging: &models.GatewayLoggingConfigs{
					Aggregation: &models.AggregationLoggingConfigs{
						TargetFilesByTag: map[string]string{
							"thing": "/var/log/thing.log",
							"blah":  "/some/directory/blah.log",
						},
						// No throttle values
					},
				},
			},
		}
		tier := configurator.NetworkEntity{
			Type: orc8r.UpgradeTierEntityType,
			Key:  "default",
			Config: &models.Tier{
				Name:    "default",
				Version: "1.0.0-0",
				Images: []*models.TierImage{
					{Name: swag.String("Image1"), Order: swag.Int64(42)},
					{Name: swag.String("Image2"), Order: swag.Int64(1)},
				},
			},
		}
		graph := configurator.EntityGraph{
			Entities: []configurator.NetworkEntity{gw, tier},
			Edges: []configurator.GraphEdge{
				{From: tier.GetTypeAndKey(), To: gw.GetTypeAndKey()},
			},
		}

		expected := map[string]proto.Message{
			"control_proxy": &mconfig_protos.ControlProxy{LogLevel: protos.LogLevel_INFO},
			"magmad": &mconfig_protos.MagmaD{
				LogLevel:                protos.LogLevel_INFO,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				AutoupgradeEnabled:      true,
				AutoupgradePollInterval: 300,
				PackageVersion:          "1.0.0-0",
				Images: []*mconfig_protos.ImageSpec{
					{Name: "Image1", Order: 42},
					{Name: "Image2", Order: 1},
				},
				DynamicServices: nil,
				FeatureFlags:    nil,
			},
			"metricsd": &mconfig_protos.MetricsD{LogLevel: protos.LogLevel_INFO},
			"td-agent-bit": &mconfig_protos.FluentBit{
				ExtraTags:        map[string]string{"network_id": "n1", "gateway_id": "gw1"},
				ThrottleRate:     1000,
				ThrottleWindow:   5,
				ThrottleInterval: "1m",
				FilesByTag: map[string]string{
					"thing": "/var/log/thing.log",
					"blah":  "/some/directory/blah.log",
				},
			},
			"eventd": &mconfig_protos.EventD{
				LogLevel:       protos.LogLevel_INFO,
				EventVerbosity: -1,
			},
		}

		actual, err := BuildBaseOrchestrator(&nw, &graph, "gw1")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestDnsdMconfigBuilder_Build(t *testing.T) {
	assert.NoError(t, plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{}))
	orchestrator_test_init.StartTestService(t)

	t.Run("empty dnsd network config", func(t *testing.T) {
		nw := configurator.Network{ID: "n1"}
		gw := configurator.NetworkEntity{
			Type: orc8r.MagmadGatewayType,
			Key:  "gw1",
			Config: &models.MagmadGatewayConfigs{
				AutoupgradeEnabled:      swag.Bool(true),
				AutoupgradePollInterval: 300,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				DynamicServices:         []string{},
				FeatureFlags:            map[string]bool{},
			},
		}
		graph := configurator.EntityGraph{
			Entities: []configurator.NetworkEntity{gw},
		}

		expected := map[string]proto.Message{
			"dnsd": &mconfig_protos.DnsD{},
		}

		actual, err := BuildDnsd(&nw, &graph, "gw1")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("non-empty dnsd network config", func(t *testing.T) {
		nw := configurator.Network{
			ID: "n1",
			Configs: map[string]interface{}{
				"dnsd_network": &models.NetworkDNSConfig{
					EnableCaching: swag.Bool(true),
					LocalTTL:      swag.Uint32(100),
					Records: []*models.DNSConfigRecord{
						{
							ARecord:     []strfmt.IPv4{"127.0.0.1", "127.0.0.2"},
							AaaaRecord:  []strfmt.IPv6{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "1234:0db8:85a3:0000:0000:8a2e:0370:1234"},
							CnameRecord: []string{"baz"},
							Domain:      "facebook.com",
						},
						{
							ARecord: []strfmt.IPv4{"quz"},
						},
					},
				},
			},
		}
		gw := configurator.NetworkEntity{
			Type: orc8r.MagmadGatewayType,
			Key:  "gw1",
			Config: &models.MagmadGatewayConfigs{
				AutoupgradeEnabled:      swag.Bool(true),
				AutoupgradePollInterval: 300,
				CheckinInterval:         60,
				CheckinTimeout:          10,
				DynamicServices:         []string{},
				FeatureFlags:            map[string]bool{},
			},
		}
		graph := configurator.EntityGraph{
			Entities: []configurator.NetworkEntity{gw},
		}

		expected := map[string]proto.Message{
			"dnsd": &mconfig_protos.DnsD{
				LogLevel:      protos.LogLevel_INFO,
				EnableCaching: true,
				LocalTTL:      100,
				Records: []*mconfig_protos.NetworkDNSConfigRecordsItems{
					{
						ARecord:     []string{"127.0.0.1", "127.0.0.2"},
						AaaaRecord:  []string{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "1234:0db8:85a3:0000:0000:8a2e:0370:1234"},
						CnameRecord: []string{"baz"},
						Domain:      "facebook.com",
					},
					{
						ARecord: []string{"quz"},
					},
				},
			},
		}

		actual, err := BuildDnsd(&nw, &graph, "gw1")
		assert.NoError(t, err)
		assert.Equal(t, expected["dnsd"].String(), actual["dnsd"].String())
	})
}

func BuildBaseOrchestrator(network *configurator.Network, graph *configurator.EntityGraph, gatewayID string) (map[string]proto.Message, error) {
	networkProto, err := network.ToStorageProto()
	if err != nil {
		return nil, err
	}
	graphProto, err := graph.ToStorageProto()
	if err != nil {
		return nil, err
	}
	builder := mconfig.NewRemoteBuilder(orchestrator.ServiceName)
	res, err := builder.Build(networkProto, graphProto, gatewayID)
	if err != nil {
		return nil, err
	}

	configs := map[string]proto.Message{}

	magmadIConfig, ok := res["magmad"]
	if ok {
		magmadConfig := &mconfig_protos.MagmaD{}
		err = ptypes.UnmarshalAny(magmadIConfig, magmadConfig)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		configs["magmad"] = magmadConfig

		fluentBitConfig := &mconfig_protos.FluentBit{}
		err = ptypes.UnmarshalAny(res["td-agent-bit"], fluentBitConfig)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		configs["td-agent-bit"] = fluentBitConfig

		eventdConfig := &mconfig_protos.EventD{}
		err = ptypes.UnmarshalAny(res["eventd"], eventdConfig)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		configs["eventd"] = eventdConfig
	}

	controlProxyConfig := &mconfig_protos.ControlProxy{}
	err = ptypes.UnmarshalAny(res["control_proxy"], controlProxyConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	configs["control_proxy"] = controlProxyConfig

	metricsdConfig := &mconfig_protos.MetricsD{}
	err = ptypes.UnmarshalAny(res["metricsd"], metricsdConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	configs["metricsd"] = metricsdConfig

	return configs, nil
}

func BuildDnsd(network *configurator.Network, graph *configurator.EntityGraph, gatewayID string) (map[string]proto.Message, error) {
	networkProto, err := network.ToStorageProto()
	if err != nil {
		return nil, err
	}
	graphProto, err := graph.ToStorageProto()
	if err != nil {
		return nil, err
	}
	builder := mconfig.NewRemoteBuilder(orchestrator.ServiceName)
	res, err := builder.Build(networkProto, graphProto, gatewayID)
	if err != nil {
		return nil, err
	}

	configs := map[string]proto.Message{}

	dnsdConfig := &mconfig_protos.DnsD{}
	err = ptypes.UnmarshalAny(res["dnsd"], dnsdConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	configs["dnsd"] = dnsdConfig
	return configs, nil
}
