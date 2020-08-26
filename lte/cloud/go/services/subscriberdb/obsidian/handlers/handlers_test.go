/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers_test

import (
	"testing"
	"time"

	"magma/lte/cloud/go/lte"
	ltePlugin "magma/lte/cloud/go/plugin"
	lteHandlers "magma/lte/cloud/go/services/lte/obsidian/handlers"
	lteModels "magma/lte/cloud/go/services/lte/obsidian/models"
	policydbHandlers "magma/lte/cloud/go/services/policydb/obsidian/handlers"
	policydbModels "magma/lte/cloud/go/services/policydb/obsidian/models"
	"magma/lte/cloud/go/services/subscriberdb/obsidian/handlers"
	subscriberModels "magma/lte/cloud/go/services/subscriberdb/obsidian/models"
	"magma/orc8r/cloud/go/clock"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/obsidian/tests"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/pluginimpl"
	"magma/orc8r/cloud/go/services/configurator"
	configuratorTestInit "magma/orc8r/cloud/go/services/configurator/test_init"
	deviceTestInit "magma/orc8r/cloud/go/services/device/test_init"
	"magma/orc8r/cloud/go/services/directoryd"
	"magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"magma/orc8r/cloud/go/services/state"
	stateTestInit "magma/orc8r/cloud/go/services/state/test_init"
	"magma/orc8r/cloud/go/services/state/test_utils"
	"magma/orc8r/cloud/go/storage"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubscriber(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})

	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers"
	handlers := handlers.GetHandlers()
	createSubscriber := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot, obsidian.POST).HandlerFunc

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	// default sub profile should always succeed
	payload := &subscriberModels.MutableSubscriber{
		ID:   "IMSI1234567890",
		Name: "Jane Doe",
		Lte: &subscriberModels.LteSubscription{
			AuthAlgo:   "MILENAGE",
			AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			State:      "ACTIVE",
			SubProfile: "default",
		},
		StaticIps: subscriberModels.SubscriberStaticIps{
			apn1: "192.168.100.1",
		},
		ActiveApns: subscriberModels.ApnList{apn2, apn1},
	}
	tc := tests.Test{
		Method:         "POST",
		URL:            testURLRoot,
		Payload:        payload,
		Handler:        createSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 201,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	expected := configurator.NetworkEntity{
		NetworkID: "n1",
		Type:      lte.SubscriberEntityType,
		Key:       "IMSI1234567890",
		Name:      "Jane Doe",
		Config: &subscriberModels.SubscriberConfig{
			Lte:       payload.Lte,
			StaticIps: payload.StaticIps,
		},
		GraphID:      "2",
		Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
	}
	assert.Equal(t, expected, actual)

	// no cellular config on network and a non-default sub profile should be 500
	payload = &subscriberModels.MutableSubscriber{
		ID: "IMSI0987654321",
		Lte: &subscriberModels.LteSubscription{
			AuthAlgo:   "MILENAGE",
			AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			State:      "ACTIVE",
			SubProfile: "foo",
		},
		ActiveApns: subscriberModels.ApnList{apn2, apn1},
	}
	tc = tests.Test{
		Method:         "POST",
		URL:            testURLRoot,
		Payload:        payload,
		Handler:        createSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 500,
		ExpectedError:  "no cellular config found for network",
	}
	tests.RunUnitTest(t, e, tc)

	_, err = configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI0987654321", configurator.FullEntityLoadCriteria())
	assert.EqualError(t, err, "Not found")

	// nonexistent sub profile should be 400
	err = configurator.UpdateNetworkConfig(
		"n1", lte.CellularNetworkConfigType,
		&lteModels.NetworkCellularConfigs{
			Epc: &lteModels.NetworkEpcConfigs{
				SubProfiles: map[string]lteModels.NetworkEpcConfigsSubProfilesAnon{
					"blah": {
						MaxDlBitRate: 100,
						MaxUlBitRate: 100,
					},
				},
			},
		},
	)
	assert.NoError(t, err)
	payload = &subscriberModels.MutableSubscriber{
		ID: "IMSI0987654321",
		Lte: &subscriberModels.LteSubscription{
			AuthAlgo:   "MILENAGE",
			AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			State:      "ACTIVE",
			SubProfile: "foo",
		},
		ActiveApns: subscriberModels.ApnList{apn2, apn1},
	}
	tc = tests.Test{
		Method:         "POST",
		URL:            testURLRoot,
		Payload:        payload,
		Handler:        createSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 400,
		ExpectedError:  "subscriber profile foo does not exist for the network",
	}
	tests.RunUnitTest(t, e, tc)

	// other validation failure
	tc = tests.Test{
		Method: "POST",
		URL:    testURLRoot,
		Payload: &subscriberModels.Subscriber{
			ID: "IMSI1234567898",
			Lte: &subscriberModels.LteSubscription{
				AuthAlgo:   "MILENAGE",
				AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:      "ACTIVE",
				SubProfile: "default",
			},
			ActiveApns: subscriberModels.ApnList{apn2, apn1},
		},
		Handler:        createSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 400,
		ExpectedError:  "expected lte auth key to be 16 bytes but got 15 bytes",
	}
	tests.RunUnitTest(t, e, tc)

	// Can't assign static IP for inactive APN
	tc = tests.Test{
		Method: "POST",
		URL:    testURLRoot,
		Payload: &subscriberModels.MutableSubscriber{
			ID: "IMSI1234567898",
			Lte: &subscriberModels.LteSubscription{
				AuthAlgo:   "MILENAGE",
				AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:      "ACTIVE",
				SubProfile: "default",
			},
			StaticIps: subscriberModels.SubscriberStaticIps{
				"asdf": "192.168.100.1",
			},
			ActiveApns: subscriberModels.ApnList{apn2, apn1},
		},
		Handler:        createSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 400,
		ExpectedError:  "static IP assigned to APN asdf which is not active for the subscriber",
	}
	tests.RunUnitTest(t, e, tc)
}

func TestListSubscribers(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})

	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	stateTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers"
	handlers := handlers.GetHandlers()
	listSubscribers := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot, obsidian.GET).HandlerFunc

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	tc := tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        listSubscribers,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler(map[string]*subscriberModels.Subscriber{}),
	}
	tests.RunUnitTest(t, e, tc)

	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{
				Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
				Config: &subscriberModels.SubscriberConfig{
					Lte: &subscriberModels.LteSubscription{
						AuthAlgo: "MILENAGE",
						AuthKey:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
						AuthOpc:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
						State:    "ACTIVE",
					},
					StaticIps: subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1", apn2: "10.10.10.5"},
				},
				Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
			},
			{
				Type: lte.SubscriberEntityType, Key: "IMSI0987654321",
				Config: &subscriberModels.SubscriberConfig{
					Lte: &subscriberModels.LteSubscription{
						AuthAlgo:   "MILENAGE",
						AuthKey:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
						AuthOpc:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
						State:      "ACTIVE",
						SubProfile: "foo",
					},
				},
				Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn1}},
			},
		},
	)
	assert.NoError(t, err)

	tc = tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        listSubscribers,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler(map[string]*subscriberModels.Subscriber{
			"IMSI1234567890": {
				ID: "IMSI1234567890",
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo:   "MILENAGE",
					AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:      "ACTIVE",
					SubProfile: "default",
				},
				Config: &subscriberModels.SubscriberConfig{
					Lte: &subscriberModels.LteSubscription{
						AuthAlgo:   "MILENAGE",
						AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
						AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
						State:      "ACTIVE",
						SubProfile: "default",
					},
					StaticIps: subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1", apn2: "10.10.10.5"},
				},
				ActiveApns: subscriberModels.ApnList{apn2, apn1},
			},
			"IMSI0987654321": {
				ID: "IMSI0987654321",
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo:   "MILENAGE",
					AuthKey:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
					AuthOpc:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
					State:      "ACTIVE",
					SubProfile: "foo",
				},
				Config: &subscriberModels.SubscriberConfig{
					Lte: &subscriberModels.LteSubscription{
						AuthAlgo:   "MILENAGE",
						AuthKey:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
						AuthOpc:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
						State:      "ACTIVE",
						SubProfile: "foo",
					},
				},
				ActiveApns: subscriberModels.ApnList{apn1},
			},
		}),
	}
	tests.RunUnitTest(t, e, tc)

	// Now create some AGW-reported state for 1234567890
	// First we need to register a gateway which can report state
	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{Type: orc8r.MagmadGatewayType, Key: "g1", Config: &models.MagmadGatewayConfigs{}, PhysicalID: "hw1"},
	)
	assert.NoError(t, err)
	frozenClock := int64(1000000)
	clock.SetAndFreezeClock(t, time.Unix(frozenClock, 0))
	defer clock.UnfreezeClock(t)

	icmpStatus := &subscriberModels.IcmpStatus{LatencyMs: f32Ptr(12.34)}
	ctx := test_utils.GetContextWithCertificate(t, "hw1")
	test_utils.ReportState(t, ctx, lte.ICMPStateType, "IMSI1234567890", icmpStatus)
	mmeState := state.ArbitraryJSON{"mme": "foo"}
	test_utils.ReportState(t, ctx, lte.MMEStateType, "IMSI1234567890", &mmeState)
	spgwState := state.ArbitraryJSON{"spgw": "foo"}
	test_utils.ReportState(t, ctx, lte.SPGWStateType, "IMSI1234567890", &spgwState)
	s1apState := state.ArbitraryJSON{"s1ap": "foo"}
	test_utils.ReportState(t, ctx, lte.S1APStateType, "IMSI1234567890", &s1apState)
	// Report 2 allocated IP addresses for the subscriber
	mobilitydState1 := state.ArbitraryJSON{
		"ip": map[string]interface{}{
			"address": "wKiArg==",
		},
	}
	mobilitydState2 := state.ArbitraryJSON{
		"ip": map[string]interface{}{
			"address": "wKiAhg==",
		},
	}
	test_utils.ReportState(t, ctx, lte.MobilitydStateType, "IMSI1234567890.oai.ipv4", &mobilitydState1)
	test_utils.ReportState(t, ctx, lte.MobilitydStateType, "IMSI1234567890.magma.apn", &mobilitydState2)
	directoryState := directoryd.DirectoryRecord{LocationHistory: []string{"foo", "bar"}}
	test_utils.ReportState(t, ctx, orc8r.DirectoryRecordType, "IMSI1234567890", &directoryState)

	tc = tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        listSubscribers,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n1"},
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler(map[string]*subscriberModels.Subscriber{
			"IMSI1234567890": {
				ID: "IMSI1234567890",
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo:   "MILENAGE",
					AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:      "ACTIVE",
					SubProfile: "default",
				},
				Config: &subscriberModels.SubscriberConfig{
					Lte: &subscriberModels.LteSubscription{
						AuthAlgo:   "MILENAGE",
						AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
						AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
						State:      "ACTIVE",
						SubProfile: "default",
					},
					StaticIps: subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1", apn2: "10.10.10.5"},
				},
				ActiveApns: subscriberModels.ApnList{apn2, apn1},
				Monitoring: &subscriberModels.SubscriberStatus{
					Icmp: &subscriberModels.IcmpStatus{
						LastReportedTime: frozenClock,
						LatencyMs:        f32Ptr(12.34),
					},
				},
				State: &subscriberModels.SubscriberState{
					Mme:  mmeState,
					S1ap: s1apState,
					Spgw: spgwState,
					Mobility: []*subscriberModels.SubscriberIPAllocation{
						{
							Apn: "magma.apn",
							IP:  "192.168.128.134",
						},
						{
							Apn: "oai.ipv4",
							IP:  "192.168.128.174",
						},
					},
					Directory: &subscriberModels.SubscriberDirectoryRecord{
						LocationHistory: []string{"foo", "bar"},
					},
				},
			},
			"IMSI0987654321": {
				ID: "IMSI0987654321",
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo:   "MILENAGE",
					AuthKey:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
					AuthOpc:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
					State:      "ACTIVE",
					SubProfile: "foo",
				},
				Config: &subscriberModels.SubscriberConfig{
					Lte: &subscriberModels.LteSubscription{
						AuthAlgo:   "MILENAGE",
						AuthKey:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
						AuthOpc:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
						State:      "ACTIVE",
						SubProfile: "foo",
					},
				},
				ActiveApns: subscriberModels.ApnList{apn1},
			},
		}),
	}
	tests.RunUnitTest(t, e, tc)
}

func TestGetSubscriber(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})

	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	stateTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers/:subscriber_id"
	handlers := handlers.GetHandlers()
	getSubscriber := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot, obsidian.GET).HandlerFunc

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	tc := tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        getSubscriber,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 404,
		ExpectedError:  "Not Found",
	}
	tests.RunUnitTest(t, e, tc)

	// No sub profile configured, we should return "default"
	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{
			Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
			Name: "Jane Doe",
			Config: &subscriberModels.SubscriberConfig{
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo: "MILENAGE",
					AuthKey:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:    "ACTIVE",
				},
				StaticIps: subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1"},
			},
			Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
		},
	)
	assert.NoError(t, err)

	tc = tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        getSubscriber,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 200,
		ExpectedResult: &subscriberModels.Subscriber{
			ID:   "IMSI1234567890",
			Name: "Jane Doe",
			Lte: &subscriberModels.LteSubscription{
				AuthAlgo:   "MILENAGE",
				AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:      "ACTIVE",
				SubProfile: "default",
			},
			Config: &subscriberModels.SubscriberConfig{
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo:   "MILENAGE",
					AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:      "ACTIVE",
					SubProfile: "default",
				},
				StaticIps: subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1"},
			},
			ActiveApns: subscriberModels.ApnList{apn2, apn1},
		},
	}
	tests.RunUnitTest(t, e, tc)

	// Now create AGW
	// First we need to register a gateway which can report state
	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{Type: orc8r.MagmadGatewayType, Key: "g1", Config: &models.MagmadGatewayConfigs{}, PhysicalID: "hw1"},
	)
	assert.NoError(t, err)
	frozenClock := int64(1000000)
	clock.SetAndFreezeClock(t, time.Unix(frozenClock, 0))
	defer clock.UnfreezeClock(t)
	icmpStatus := &subscriberModels.IcmpStatus{LatencyMs: f32Ptr(12.34)}
	ctx := test_utils.GetContextWithCertificate(t, "hw1")
	test_utils.ReportState(t, ctx, lte.ICMPStateType, "IMSI1234567890", icmpStatus)
	mmeState := state.ArbitraryJSON{"mme": "foo"}
	test_utils.ReportState(t, ctx, lte.MMEStateType, "IMSI1234567890", &mmeState)
	spgwState := state.ArbitraryJSON{"spgw": "foo"}
	test_utils.ReportState(t, ctx, lte.SPGWStateType, "IMSI1234567890", &spgwState)
	s1apState := state.ArbitraryJSON{"s1ap": "foo"}
	test_utils.ReportState(t, ctx, lte.S1APStateType, "IMSI1234567890", &s1apState)
	// Report 2 allocated IP addresses for the subscriber
	mobilitydState1 := state.ArbitraryJSON{
		"ip": map[string]interface{}{
			"address": "wKiArg==",
		},
	}
	mobilitydState2 := state.ArbitraryJSON{
		"ip": map[string]interface{}{
			"address": "wKiAhg==",
		},
	}
	test_utils.ReportState(t, ctx, lte.MobilitydStateType, "IMSI1234567890.oai.ipv4", &mobilitydState1)
	test_utils.ReportState(t, ctx, lte.MobilitydStateType, "IMSI1234567890.magma.apn", &mobilitydState2)
	directoryState := directoryd.DirectoryRecord{LocationHistory: []string{"foo", "bar"}}
	test_utils.ReportState(t, ctx, orc8r.DirectoryRecordType, "IMSI1234567890", &directoryState)

	tc = tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        getSubscriber,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 200,
		ExpectedResult: &subscriberModels.Subscriber{
			ID:   "IMSI1234567890",
			Name: "Jane Doe",
			Lte: &subscriberModels.LteSubscription{
				AuthAlgo:   "MILENAGE",
				AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:      "ACTIVE",
				SubProfile: "default",
			},
			Config: &subscriberModels.SubscriberConfig{
				Lte: &subscriberModels.LteSubscription{
					AuthAlgo:   "MILENAGE",
					AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:      "ACTIVE",
					SubProfile: "default",
				},
				StaticIps: subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1"},
			},
			ActiveApns: subscriberModels.ApnList{apn2, apn1},
			Monitoring: &subscriberModels.SubscriberStatus{
				Icmp: &subscriberModels.IcmpStatus{
					LastReportedTime: frozenClock,
					LatencyMs:        f32Ptr(12.34),
				},
			},
			State: &subscriberModels.SubscriberState{
				Mme:  mmeState,
				S1ap: s1apState,
				Spgw: spgwState,
				Mobility: []*subscriberModels.SubscriberIPAllocation{
					{
						Apn: "magma.apn",
						IP:  "192.168.128.134",
					},
					{
						Apn: "oai.ipv4",
						IP:  "192.168.128.174",
					},
				},
				Directory: &subscriberModels.SubscriberDirectoryRecord{
					LocationHistory: []string{"foo", "bar"},
				},
			},
		},
	}
	tests.RunUnitTest(t, e, tc)
}

func TestUpdateSubscriber(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})

	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers/:subscriber_id"
	handlers := handlers.GetHandlers()
	updateSubscriber := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot, obsidian.PUT).HandlerFunc

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	// 404
	payload := &subscriberModels.MutableSubscriber{
		ID: "IMSI1234567890",
		Lte: &subscriberModels.LteSubscription{
			AuthAlgo:   "MILENAGE",
			AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			State:      "ACTIVE",
			SubProfile: "default",
		},
		ActiveApns: subscriberModels.ApnList{apn2, apn1},
	}
	tc := tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateSubscriber,
		Payload:        payload,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 404,
		ExpectedError:  "Not Found",
	}
	tests.RunUnitTest(t, e, tc)

	// Happy path
	err = configurator.UpdateNetworkConfig(
		"n1", lte.CellularNetworkConfigType,
		&lteModels.NetworkCellularConfigs{
			Epc: &lteModels.NetworkEpcConfigs{
				SubProfiles: map[string]lteModels.NetworkEpcConfigsSubProfilesAnon{
					"foo": {
						MaxUlBitRate: 100,
						MaxDlBitRate: 100,
					},
				},
			},
		},
	)
	assert.NoError(t, err)
	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{
			Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
			Config: &subscriberModels.SubscriberConfig{
				Lte: &subscriberModels.LteSubscription{
					AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:      "ACTIVE",
					SubProfile: "default",
				},
			},
			Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}},
		},
	)
	assert.NoError(t, err)

	payload = &subscriberModels.MutableSubscriber{
		ID:   "IMSI1234567890",
		Name: "Jane Doe",
		Lte: &subscriberModels.LteSubscription{
			AuthAlgo:   "MILENAGE",
			AuthKey:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
			AuthOpc:    []byte("\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22\x22"),
			State:      "INACTIVE",
			SubProfile: "foo",
		},
		StaticIps:  subscriberModels.SubscriberStaticIps{apn1: "192.168.100.1"},
		ActiveApns: subscriberModels.ApnList{apn2, apn1},
	}
	tc = tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateSubscriber,
		Payload:        payload,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	expected := configurator.NetworkEntity{
		NetworkID:    "n1",
		Type:         lte.SubscriberEntityType,
		Key:          "IMSI1234567890",
		Name:         "Jane Doe",
		Config:       &subscriberModels.SubscriberConfig{Lte: payload.Lte, StaticIps: payload.StaticIps},
		GraphID:      "2",
		Version:      1,
		Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
	}
	assert.Equal(t, expected, actual)

	// No profile matching
	payload.Lte.SubProfile = "bar"
	tc = tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateSubscriber,
		Payload:        payload,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 400,
		ExpectedError:  "subscriber profile bar does not exist for the network",
	}
	tests.RunUnitTest(t, e, tc)
}

func TestDeleteSubscriber(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})

	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers/:subscriber_id"
	handlers := handlers.GetHandlers()
	deleteSubscriber := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot, obsidian.DELETE).HandlerFunc

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{
			Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
			// Intentionally populate with invalid config
			Config: &subscriberModels.LteSubscription{
				AuthAlgo: "MILENAGE",
				AuthKey:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:    "ACTIVE",
			},
			Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
		},
	)
	assert.NoError(t, err)

	tc := tests.Test{
		Method:         "DELETE",
		URL:            testURLRoot,
		Handler:        deleteSubscriber,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := configurator.LoadAllEntitiesInNetwork("n1", lte.SubscriberEntityType, configurator.EntityLoadCriteria{})
	assert.Equal(t, 0, len(actual))
}

func TestActivateDeactivateSubscriber(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})

	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers/:subscriber_id"
	handlers := handlers.GetHandlers()
	activateSubscriber := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot+"/activate", obsidian.POST).HandlerFunc
	deactivateSubscriber := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot+"/deactivate", obsidian.POST).HandlerFunc

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	expected := configurator.NetworkEntity{
		Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
		Config: &subscriberModels.SubscriberConfig{
			Lte: &subscriberModels.LteSubscription{
				AuthAlgo: "MILENAGE",
				AuthKey:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:  []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:    "ACTIVE",
			},
		},
		Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
	}
	_, err = configurator.CreateEntity("n1", expected)
	assert.NoError(t, err)
	expected.NetworkID = "n1"
	expected.GraphID = "2"
	expected.Version = 1

	// activate already activated subscriber
	tc := tests.Test{
		Method:         "POST",
		URL:            testURLRoot + "/activate",
		Handler:        activateSubscriber,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 200,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	// deactivate
	tc.URL = testURLRoot + "/deactivate"
	tc.Handler = deactivateSubscriber
	tests.RunUnitTest(t, e, tc)

	actual, err = configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	expected.Config.(*subscriberModels.SubscriberConfig).Lte.State = "INACTIVE"
	expected.Version = 2
	assert.Equal(t, expected, actual)

	// deactivate deactivated sub
	tests.RunUnitTest(t, e, tc)
	actual, err = configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	expected.Config.(*subscriberModels.SubscriberConfig).Lte.State = "INACTIVE"
	expected.Version = 3
	assert.Equal(t, expected, actual)

	// activate
	tc.URL = testURLRoot + "/activate"
	tc.Handler = activateSubscriber
	tests.RunUnitTest(t, e, tc)
	actual, err = configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	expected.Config.(*subscriberModels.SubscriberConfig).Lte.State = "ACTIVE"
	expected.Version = 4
	assert.Equal(t, expected, actual)
}

func TestUpdateSubscriberProfile(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{})
	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)

	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)
	err = configurator.UpdateNetworkConfig(
		"n1", lte.CellularNetworkConfigType,
		&lteModels.NetworkCellularConfigs{
			Epc: &lteModels.NetworkEpcConfigs{
				SubProfiles: map[string]lteModels.NetworkEpcConfigsSubProfilesAnon{
					"foo": {
						MaxUlBitRate: 100,
						MaxDlBitRate: 100,
					},
				},
			},
		},
	)
	assert.NoError(t, err)

	//preseed 2 apns
	apn1, apn2 := "foo", "bar"
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{Type: lte.APNEntityType, Key: apn1},
			{Type: lte.APNEntityType, Key: apn2},
		},
	)
	assert.NoError(t, err)

	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{
			Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
			Config: &subscriberModels.SubscriberConfig{
				Lte: &subscriberModels.LteSubscription{
					AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
					State:      "ACTIVE",
					SubProfile: "default",
				},
			},
			Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
		},
	)
	assert.NoError(t, err)

	e := echo.New()
	testURLRoot := "/magma/v1/lte/:network_id/subscribers/:subscriber_id/lte/sub_profile"
	handlers := handlers.GetHandlers()
	updateProfile := tests.GetHandlerByPathAndMethod(t, handlers, testURLRoot, obsidian.PUT).HandlerFunc

	// 404
	payload := "foo"
	tc := tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateProfile,
		Payload:        tests.JSONMarshaler(payload),
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI0987654321"},
		ExpectedStatus: 404,
		ExpectedError:  "Not Found",
	}
	tests.RunUnitTest(t, e, tc)

	// bad profile
	payload = "bar"
	tc = tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateProfile,
		Payload:        tests.JSONMarshaler(payload),
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 400,
		ExpectedError:  "subscriber profile bar does not exist for the network",
	}
	tests.RunUnitTest(t, e, tc)

	// happy path
	payload = "foo"
	tc = tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateProfile,
		Payload:        tests.JSONMarshaler(payload),
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	assert.NoError(t, err)
	expected := configurator.NetworkEntity{
		NetworkID: "n1", Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
		Config: &subscriberModels.SubscriberConfig{
			Lte: &subscriberModels.LteSubscription{
				AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:      "ACTIVE",
				SubProfile: "foo",
			},
		},
		GraphID:      "2",
		Version:      1,
		Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
	}
	assert.Equal(t, expected, actual)

	// set to default
	payload = "default"
	tc = tests.Test{
		Method:         "PUT",
		URL:            testURLRoot,
		Handler:        updateProfile,
		Payload:        tests.JSONMarshaler(payload),
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n1", "IMSI1234567890"},
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err = configurator.LoadEntity("n1", lte.SubscriberEntityType, "IMSI1234567890", configurator.FullEntityLoadCriteria())
	expected = configurator.NetworkEntity{
		NetworkID: "n1", Type: lte.SubscriberEntityType, Key: "IMSI1234567890",
		Config: &subscriberModels.SubscriberConfig{
			Lte: &subscriberModels.LteSubscription{
				AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
				State:      "ACTIVE",
				SubProfile: "default",
			},
		},
		GraphID:      "2",
		Version:      2,
		Associations: []storage.TypeAndKey{{Type: lte.APNEntityType, Key: apn2}, {Type: lte.APNEntityType, Key: apn1}},
	}
	assert.Equal(t, expected, actual)
}

func TestAPNPolicyProfile(t *testing.T) {
	assert.NoError(t, plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{}))
	assert.NoError(t, plugin.RegisterPluginForTests(t, &ltePlugin.LteOrchestratorPlugin{}))

	configuratorTestInit.StartTestService(t)
	stateTestInit.StartTestService(t)
	err := configurator.CreateNetwork(configurator.Network{ID: "n0"})
	assert.NoError(t, err)
	_, err = configurator.CreateEntities("n0", []configurator.NetworkEntity{
		{Type: lte.APNEntityType, Key: "apn0"},
		{Type: lte.APNEntityType, Key: "apn1"},
		{Type: lte.PolicyRuleEntityType, Key: "rule0"},
		{Type: lte.PolicyRuleEntityType, Key: "rule1"},
		{Type: lte.PolicyRuleEntityType, Key: "rule2"},
	})
	assert.NoError(t, err)

	e := echo.New()
	urlBase := "/magma/v1/lte/:network_id/subscribers"
	urlManage := urlBase + "/:subscriber_id"
	subscriberdbHandlers := handlers.GetHandlers()
	getAllSubscribers := tests.GetHandlerByPathAndMethod(t, subscriberdbHandlers, urlBase, obsidian.GET).HandlerFunc
	postSubscriber := tests.GetHandlerByPathAndMethod(t, subscriberdbHandlers, urlBase, obsidian.POST).HandlerFunc
	putSubscriber := tests.GetHandlerByPathAndMethod(t, subscriberdbHandlers, urlManage, obsidian.PUT).HandlerFunc
	getSubscriber := tests.GetHandlerByPathAndMethod(t, subscriberdbHandlers, urlManage, obsidian.GET).HandlerFunc
	deleteSubscriber := tests.GetHandlerByPathAndMethod(t, subscriberdbHandlers, urlManage, obsidian.DELETE).HandlerFunc

	deleteAPN := tests.GetHandlerByPathAndMethod(t, lteHandlers.GetHandlers(), "/magma/v1/lte/:network_id/apns/:apn_name", obsidian.DELETE).HandlerFunc
	deletePolicy := tests.GetHandlerByPathAndMethod(t, policydbHandlers.GetHandlers(), "/magma/v1/networks/:network_id/policies/rules/:rule_id", obsidian.DELETE).HandlerFunc

	imsi := "IMSI1234567890"
	imsi1 := "IMSI1234567800"
	mutableSub := newMutableSubscriber(imsi)
	sub := (&subscriberModels.Subscriber{}).FromMutable(mutableSub)

	// Get all, initially empty
	tc := tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers",
		Handler:        getAllSubscribers,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n0"},
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler(map[string]subscriberModels.Subscriber{}),
	}
	tests.RunUnitTest(t, e, tc)

	// Post err, APN doesn't exist
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{"apnXXX": policydbModels.PolicyIds{"rule0"}}
	tc = tests.Test{
		Method:                 "POST",
		URL:                    "/magma/v1/lte/n0/subscribers",
		Payload:                mutableSub,
		Handler:                postSubscriber,
		ParamNames:             []string{"network_id"},
		ParamValues:            []string{"n0"},
		ExpectedStatus:         500, // would make more sense as 400
		ExpectedErrorSubstring: `could not find entities matching [type:"apn" key:"apnXXX" ]`,
	}
	tests.RunUnitTest(t, e, tc)

	// Post err, rule doesn't exist
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{"apn0": policydbModels.PolicyIds{"ruleXXX"}}
	tc = tests.Test{
		Method:                 "POST",
		URL:                    "/magma/v1/lte/n0/subscribers",
		Payload:                mutableSub,
		Handler:                postSubscriber,
		ParamNames:             []string{"network_id"},
		ParamValues:            []string{"n0"},
		ExpectedStatus:         500, // would make more sense as 400
		ExpectedErrorSubstring: `could not find entities matching [type:"policy" key:"ruleXXX" ]`,
	}
	tests.RunUnitTest(t, e, tc)

	// Post, successful
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{
		"apn0": policydbModels.PolicyIds{"rule0"},
	}
	tc = tests.Test{
		Method:         "POST",
		URL:            "/magma/v1/lte/n0/subscribers",
		Payload:        mutableSub,
		Handler:        postSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n0"},
		ExpectedStatus: 201,
	}
	tests.RunUnitTest(t, e, tc)

	// Configurator confirms policy profile exists
	profiles, err := configurator.ListEntityKeys("n0", lte.APNPolicyProfileEntityType)
	assert.NoError(t, err)
	assert.Len(t, profiles, 1)

	// Get all, posted subscriber found
	sub = (&subscriberModels.Subscriber{}).FromMutable(mutableSub)
	tc = tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers",
		Handler:        getAllSubscribers,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n0"},
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler(map[string]*subscriberModels.Subscriber{imsi: sub}),
	}
	tests.RunUnitTest(t, e, tc)

	// Put err, APN doesn't exist
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{"apnXXX": policydbModels.PolicyIds{"rule0"}}
	tc = tests.Test{
		Method:                 "PUT",
		URL:                    "/magma/v1/lte/n0/subscribers/" + imsi,
		Payload:                mutableSub,
		ParamNames:             []string{"network_id", "subscriber_id"},
		ParamValues:            []string{"n0", imsi},
		Handler:                putSubscriber,
		ExpectedStatus:         500, // would make more sense as 400
		ExpectedErrorSubstring: `could not find entities matching [type:"apn" key:"apnXXX" ]`,
	}
	tests.RunUnitTest(t, e, tc)

	// Put err, rule doesn't exist
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{"apn0": policydbModels.PolicyIds{"ruleXXX"}}
	tc = tests.Test{
		Method:                 "PUT",
		URL:                    "/magma/v1/lte/n0/subscribers/" + imsi,
		Payload:                mutableSub,
		ParamNames:             []string{"network_id", "subscriber_id"},
		ParamValues:            []string{"n0", imsi},
		Handler:                putSubscriber,
		ExpectedStatus:         500, // would make more sense as 400
		ExpectedErrorSubstring: `could not find entities matching [type:"policy" key:"ruleXXX" ]`,
	}
	tests.RunUnitTest(t, e, tc)

	// Put, add new mappings
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{
		"apn0": policydbModels.PolicyIds{"rule0", "rule1"},
		"apn1": policydbModels.PolicyIds{"rule1", "rule2"},
	}
	tc = tests.Test{
		Method:         "PUT",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		Payload:        mutableSub,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        putSubscriber,
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	// Configurator confirms policy profiles exist
	profiles, err = configurator.ListEntityKeys("n0", lte.APNPolicyProfileEntityType)
	assert.NoError(t, err)
	assert.Len(t, profiles, 2)

	// Get, changes are reflected
	sub = (&subscriberModels.Subscriber{}).FromMutable(mutableSub)
	tc = tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        getSubscriber,
		ExpectedStatus: 200,
		ExpectedResult: sub,
	}
	tests.RunUnitTest(t, e, tc)

	// Delete
	tc = tests.Test{
		Method:         "DELETE",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        deleteSubscriber,
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	// Delete, subsequent delete still "succeeds"
	tc = tests.Test{
		Method:         "DELETE",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        deleteSubscriber,
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	// Get, delete confirmed
	tc = tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        getSubscriber,
		ExpectedStatus: 404,
		ExpectedError:  "Not Found",
	}
	tests.RunUnitTest(t, e, tc)

	// Configurator confirms deletion
	profiles, err = configurator.ListEntityKeys("n0", lte.APNPolicyProfileEntityType)
	assert.NoError(t, err)
	assert.Len(t, profiles, 0)

	// Post, add subscriber back
	mutableSub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{
		"apn0": policydbModels.PolicyIds{"rule0", "rule1"},
		"apn1": policydbModels.PolicyIds{"rule1", "rule2"},
	}
	tc = tests.Test{
		Method:         "POST",
		URL:            "/magma/v1/lte/n0/subscribers",
		Payload:        mutableSub,
		Handler:        postSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n0"},
		ExpectedStatus: 201,
	}
	tests.RunUnitTest(t, e, tc)

	// Get, successfully added back
	sub = (&subscriberModels.Subscriber{}).FromMutable(mutableSub)
	tc = tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        getSubscriber,
		ExpectedStatus: 200,
		ExpectedResult: sub,
	}
	tests.RunUnitTest(t, e, tc)

	// Delete linked policy rule
	tc = tests.Test{
		Method:         "DELETE",
		URL:            "/magma/v1/networks/n1/policies/rules/rule0",
		Payload:        nil,
		ParamNames:     []string{"network_id", "rule_id"},
		ParamValues:    []string{"n0", "rule0"},
		Handler:        deletePolicy,
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	// Get, policy rule changes reflected
	sub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{
		"apn0": policydbModels.PolicyIds{"rule1"}, // rule0 deleted
		"apn1": policydbModels.PolicyIds{"rule1", "rule2"},
	}
	tc = tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        getSubscriber,
		ExpectedStatus: 200,
		ExpectedResult: sub,
	}
	tests.RunUnitTest(t, e, tc)

	// Delete linked APN
	tc = tests.Test{
		Method:         "DELETE",
		URL:            "/magma/v1/lte/n0/apns/apn1",
		Handler:        deleteAPN,
		ParamNames:     []string{"network_id", "apn_name"},
		ParamValues:    []string{"n0", "apn0"},
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	// Get, APN change reflected
	sub.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{
		// DELETED: "apn0": policydbModels.PolicyIds{"rule1"},
		"apn1": policydbModels.PolicyIds{"rule1", "rule2"},
	}
	sub.ActiveApns = subscriberModels.ApnList{
		// DELETED: "apn0",
		"apn1",
	}
	tc = tests.Test{
		Method:         "GET",
		URL:            "/magma/v1/lte/n0/subscribers/" + imsi,
		ParamNames:     []string{"network_id", "subscriber_id"},
		ParamValues:    []string{"n0", imsi},
		Handler:        getSubscriber,
		ExpectedStatus: 200,
		ExpectedResult: sub,
	}
	tests.RunUnitTest(t, e, tc)

	// Configurator confirms deletion
	profiles, err = configurator.ListEntityKeys("n0", lte.APNPolicyProfileEntityType)
	assert.NoError(t, err)
	assert.Len(t, profiles, 1)

	// Post, add sub1, no namespacing issues
	mutableSub1 := newMutableSubscriber(imsi1)
	mutableSub1.ActivePoliciesByApn = policydbModels.PolicyIdsByApn{"apn1": policydbModels.PolicyIds{"rule1", "rule2"}}
	mutableSub1.ActiveApns = subscriberModels.ApnList{"apn1"}
	tc = tests.Test{
		Method:         "POST",
		URL:            "/magma/v1/lte/n0/subscribers",
		Payload:        mutableSub1,
		Handler:        postSubscriber,
		ParamNames:     []string{"network_id"},
		ParamValues:    []string{"n0"},
		ExpectedStatus: 201,
	}
	tests.RunUnitTest(t, e, tc)

	// Configurator non-shared apn_policy_profile
	profiles, err = configurator.ListEntityKeys("n0", lte.APNPolicyProfileEntityType)
	assert.NoError(t, err)
	assert.Len(t, profiles, 2)
}

func f32Ptr(f float32) *float32 {
	return &f
}

func newMutableSubscriber(id string) *subscriberModels.MutableSubscriber {
	sub := &subscriberModels.MutableSubscriber{
		ID:   policydbModels.SubscriberID(id),
		Name: "Jane Doe",
		Lte: &subscriberModels.LteSubscription{
			AuthAlgo:   "MILENAGE",
			AuthKey:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			AuthOpc:    []byte("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			State:      "ACTIVE",
			SubProfile: "default",
		},
		StaticIps: subscriberModels.SubscriberStaticIps{
			"apn1": "192.168.100.1",
		},
		ActiveApns: subscriberModels.ApnList{"apn0", "apn1"},
	}
	return sub
}
