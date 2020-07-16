/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// NOTE: to run these tests outside the testing environment, e.g. from IntelliJ,
// ensure postgres_test container is running, and use the following environment
// variables to point to the relevant DB endpoints:
//	- DATABASE_SOURCE=host=localhost port=5433 dbname=magma_test user=magma_test password=magma_test sslmode=disable

package statemachines_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	ltemodels "magma/lte/cloud/go/services/lte/obsidian/models"
	"magma/orc8r/cloud/go/clock"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/pluginimpl"
	"magma/orc8r/cloud/go/services/configurator"
	cfgTestInit "magma/orc8r/cloud/go/services/configurator/test_init"
	"magma/orc8r/cloud/go/services/device"
	deviceTestInit "magma/orc8r/cloud/go/services/device/test_init"
	models2 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	stateTestInit "magma/orc8r/cloud/go/services/state/test_init"
	"magma/orc8r/cloud/go/services/state/test_utils"
	"magma/orc8r/cloud/go/storage"
	"magma/orc8r/lib/go/protos"
	plugin2 "orc8r/fbinternal/cloud/go/plugin"
	"orc8r/fbinternal/cloud/go/services/testcontroller/obsidian/models"
	"orc8r/fbinternal/cloud/go/services/testcontroller/statemachines"
	storage2 "orc8r/fbinternal/cloud/go/services/testcontroller/storage"
	tcTestInit "orc8r/fbinternal/cloud/go/services/testcontroller/test_init"

	"github.com/go-openapi/swag"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// don't test intermediate failure conditions (e.g. unexpected config types,
// service errors)
func Test_EnodebdE2ETestStateMachine_HappyPath(t *testing.T) {
	SetupTests(t)
	RegisterAGW(t)
	cli := &mockClient{}
	testConfig := GetEnodebTestConfig()
	mockMagmad, mockGenericCommandResp := GetMockObjects()

	mockMagmad.On("RebootEnodeb", "n1", "g1", "1202000038269KP0037").Return(mockGenericCommandResp, nil)
	mockMagmad.On("GenerateTraffic", "n1", "g2", "magmawifi", "magmamagma").Return(mockGenericCommandResp, nil)
	mockString := ""
	mockEnodebState := &ltemodels.EnodebState{
		EnodebConfigured:   new(bool),
		EnodebConnected:    new(bool),
		FsmState:           &mockString,
		GpsConnected:       new(bool),
		GpsLatitude:        &mockString,
		GpsLongitude:       &mockString,
		MmeConnected:       new(bool),
		OpstateEnabled:     new(bool),
		PtpConnected:       new(bool),
		ReportingGatewayID: "g1",
		RfTxDesired:        new(bool),
		RfTxOn:             new(bool),
		TimeReported:       1000,
	}
	mockMagmad.On("GetEnodebStatus", "n1", "1202000038269KP0037").Return(mockEnodebState, nil)

	// New test
	sm := statemachines.NewEnodebdE2ETestStateMachine(tcTestInit.GetTestTestcontrollerStorage(t), cli, mockMagmad)
	actualState, actualDuration, err := sm.Run(storage2.CommonStartState, testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, time.Minute, actualDuration)

	// ---
	// Check for upgrade, find version equal to what tier is configured to; expect epsilon transition, 20 minute delay
	// ---
	testdata, err := ioutil.ReadFile("../testdata/testdata")
	assert.NoError(t, err)
	mockResp := &http.Response{Status: "200", Body: ioutil.NopCloser(bytes.NewBuffer(testdata))}
	cli.On("Get", mock.AnythingOfType("string")).Return(mockResp, nil).Times(1)

	actualState, actualDuration, err = sm.Run("check_for_upgrade", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 20*time.Minute, actualDuration)

	// ---
	// Check for upgrade find version ahead of what tier is configured to
	// ---
	err = configurator.CreateOrUpdateEntityConfig("n1", orc8r.UpgradeTierEntityType, "t1", &models2.Tier{Version: "0.0.0-0-abcdefg"})
	assert.NoError(t, err)
	mockResp = &http.Response{Status: "200", Body: ioutil.NopCloser(bytes.NewBuffer(testdata))}
	cli.On("Get", mock.AnythingOfType("string")).Return(mockResp, nil).Times(1)

	actualState, actualDuration, err = sm.Run("check_for_upgrade", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "verify_upgrade_1", actualState)
	assert.Equal(t, 10*time.Minute, actualDuration)

	// Tier should get updated
	actualTierCfg, err := configurator.LoadEntityConfig("n1", orc8r.UpgradeTierEntityType, "t1")
	assert.NoError(t, err)
	assert.Equal(t, &models2.Tier{Version: "0.3.74-1560824953-b50f1bab"}, actualTierCfg)

	// ---
	// Check upgrade status, gateway hasn't upgraded yet
	// ---
	gatewayRecord := &models2.GatewayDevice{HardwareID: "hw1", Key: &models2.ChallengeKey{KeyType: "ECHO"}}
	err = device.RegisterDevice("n1", orc8r.AccessGatewayRecordType, "hw1", gatewayRecord)
	assert.NoError(t, err)
	ctx := test_utils.GetContextWithCertificate(t, "hw1")
	test_utils.ReportGatewayStatus(t, ctx, &models2.GatewayStatus{
		HardwareID: "hw1",
		PlatformInfo: &models2.PlatformInfo{
			Packages: []*models2.Package{
				{Name: "magma", Version: "0.0.0-0-abcdefg"},
			},
		},
	})

	actualState, actualDuration, err = sm.Run("verify_upgrade_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "verify_upgrade_2", actualState)
	assert.Equal(t, 10*time.Minute, actualDuration)

	// ---
	// Upgrade successful
	// ---
	mockResp = &http.Response{Status: "200", StatusCode: 200}
	// Should test for the payload eventually
	cli.On("Post", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(mockResp, nil).Times(4)
	test_utils.ReportGatewayStatus(t, ctx, &models2.GatewayStatus{
		HardwareID: "hw1",
		PlatformInfo: &models2.PlatformInfo{
			Packages: []*models2.Package{
				{Name: "magma", Version: "0.3.74-1560824953-b50f1bab"},
			},
		},
	})

	actualState, actualDuration, err = sm.Run("verify_upgrade_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "traffic_test1_1", actualState)
	assert.Equal(t, 20*time.Minute, actualDuration)

	// ---
	// Traffic test 1
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test1_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "reboot_enodeb_1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	// ---
	// Reboot enodeb
	// ---
	actualState, actualDuration, err = sm.Run("reboot_enodeb_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "verify_conn", actualState)
	assert.Equal(t, 15*time.Minute, actualDuration)

	// ---
	// Verify enodeb connectivity
	// ---
	actualState, actualDuration, err = sm.Run("verify_conn", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "traffic_test2_1", actualState)
	assert.Equal(t, 15*time.Minute, actualDuration)

	// ---
	// Traffic test 2
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test2_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "reconfig_enodeb1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	cli.On("Post", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(mockResp, nil)
	// ---
	// Upgrade unsuccessful
	// ---
	test_utils.ReportGatewayStatus(t, ctx, &models2.GatewayStatus{
		HardwareID: "hw1",
		PlatformInfo: &models2.PlatformInfo{
			Packages: []*models2.Package{
				{Name: "magma", Version: "0.0.0-0-abcdefg"},
			},
		},
	})
	actualState, actualDuration, err = sm.Run("verify_upgrade_3", testConfig, nil)
	assert.EqualError(t, err, "gateway g1 did not upgrade within 3 tries")
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 20*time.Minute, actualDuration)

	cli.AssertExpectations(t)
	mockMagmad.AssertExpectations(t)
}

func Test_EnodebdE2ETestStateMachine_VerifyConnection(t *testing.T) {
	SetupTests(t)
	RegisterAGW(t)
	cli := &mockClient{}
	testConfig := GetEnodebTestConfig()
	mockMagmad, mockGenericCommandResp := GetMockObjects()
	mockString := ""
	mockEnodebState := &ltemodels.EnodebState{
		EnodebConfigured:   new(bool),
		EnodebConnected:    new(bool),
		FsmState:           &mockString,
		GpsConnected:       new(bool),
		GpsLatitude:        &mockString,
		GpsLongitude:       &mockString,
		MmeConnected:       new(bool),
		OpstateEnabled:     new(bool),
		PtpConnected:       new(bool),
		ReportingGatewayID: "g1",
		RfTxDesired:        new(bool),
		RfTxOn:             new(bool),
		TimeReported:       1000,
	}

	mockResp := &http.Response{Status: "200", StatusCode: 200}
	// Should test for the payload eventually
	cli.On("Post", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(mockResp, nil)

	// New test
	sm := statemachines.NewEnodebdE2ETestStateMachine(tcTestInit.GetTestTestcontrollerStorage(t), cli, mockMagmad)

	mockMagmad.On("RebootEnodeb", "n1", "g1", "1202000038269KP0037").Return(mockGenericCommandResp, errors.New("")).Twice()
	// ---
	// reboot_enodeb_1 transition to reboot_enodeb_2
	// --
	actualState, actualDuration, err := sm.Run("reboot_enodeb_1", testConfig, nil)
	assert.EqualError(t, err, "")
	assert.Equal(t, "reboot_enodeb_2", actualState)
	assert.Equal(t, 5*time.Minute, actualDuration)

	// ---
	// Reboot unsuccessful
	// --
	actualState, actualDuration, err = sm.Run("reboot_enodeb_3", testConfig, nil)
	assert.EqualError(t, err, "enodeb 1202000038269KP0037 did not reboot within 3 tries")
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 15*time.Minute, actualDuration)

	mockMagmad.On("GetEnodebStatus", "n1", "1202000038269KP0037").Return(mockEnodebState, errors.New("")).Once()
	// ---
	// Unable to get enodeb status
	// --
	actualState, actualDuration, err = sm.Run("verify_conn", testConfig, nil)
	assert.EqualError(t, err, "error getting enodeb status: ")
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 5*time.Minute, actualDuration)

	mockMagmad.On("RebootEnodeb", "n1", "g1", "1202000038269KP0037").Return(mockGenericCommandResp, nil)
	mockMagmad.On("GetEnodebStatus", "n1", "1202000038269KP0037").Return(mockEnodebState, nil)
	// ---
	// Reboot enodeb
	// ---
	actualState, actualDuration, err = sm.Run("reboot_enodeb_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "verify_conn", actualState)
	assert.Equal(t, 15*time.Minute, actualDuration)

	// ---
	// Verify enodeb connectivity
	// ---
	actualState, actualDuration, err = sm.Run("verify_conn", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "traffic_test2_1", actualState)
	assert.Equal(t, 15*time.Minute, actualDuration)

	cli.AssertExpectations(t)
	mockMagmad.AssertExpectations(t)
}

func Test_EnodebdE2ETestStateMachine_TrafficScript(t *testing.T) {
	SetupTests(t)
	RegisterAGW(t)
	cli := &mockClient{}
	testConfig := GetEnodebTestConfig()
	mockMagmad, mockGenericCommandResp := GetMockObjects()

	// New test
	sm := statemachines.NewEnodebdE2ETestStateMachine(tcTestInit.GetTestTestcontrollerStorage(t), cli, mockMagmad)

	mockResp := &http.Response{Status: "200", StatusCode: 200}
	// Should test for the payload eventually
	cli.On("Post", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(mockResp, nil)

	mockMagmad.On("GenerateTraffic", "n1", "g2", "magmawifi", "magmamagma").Return(mockGenericCommandResp, errors.New("")).Times(4)
	// ---
	// Unsuccessful traffic test 1
	// ---
	actualState, actualDuration, err := sm.Run("traffic_test1_1", testConfig, nil)
	assert.EqualError(t, err, "")
	assert.Equal(t, "traffic_test1_2", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	actualState, actualDuration, err = sm.Run("traffic_test1_3", testConfig, nil)
	assert.EqualError(t, err, "Traffic test number 1 failed on gwID g2 after 3 tries")
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	// ---
	// Unsuccessful traffic test 2
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test2_1", testConfig, nil)
	assert.EqualError(t, err, "")
	assert.Equal(t, "traffic_test2_2", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	actualState, actualDuration, err = sm.Run("traffic_test2_3", testConfig, nil)
	assert.EqualError(t, err, "Traffic test number 2 failed on gwID g2 after 3 tries")
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	mockMagmad.On("GenerateTraffic", "n1", "g2", "magmawifi", "magmamagma").Return(mockGenericCommandResp, nil)
	// ---
	// Traffic Test 1
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test1_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "reboot_enodeb_1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	// ---
	// Traffic Test 2
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test2_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "reconfig_enodeb1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	// ---
	// Successful traffic test in state 3
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test1_3", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "reboot_enodeb_1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	actualState, actualDuration, err = sm.Run("traffic_test2_3", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "reconfig_enodeb1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	cli.AssertExpectations(t)
	mockMagmad.AssertExpectations(t)
}

func Test_EnodebdE2ETestStateMachine_ReconfigEnb(t *testing.T) {
	SetupTests(t)
	RegisterAGW(t)
	cli := &mockClient{}
	testConfig := GetEnodebTestConfig()
	mockMagmad, mockGenericCommandResp := GetMockObjects()
	mockResp := &http.Response{Status: "200", StatusCode: 200}
	// Should test for the payload eventually
	cli.On("Post", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(mockResp, nil)

	// New test
	sm := statemachines.NewEnodebdE2ETestStateMachine(tcTestInit.GetTestTestcontrollerStorage(t), cli, mockMagmad)

	// ---
	// Reconfig Enodeb
	// ---
	actualState, actualDuration, err := sm.Run("reconfig_enodeb1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "verify_config1", actualState)
	assert.Equal(t, 10*time.Minute, actualDuration)

	mockMagmad.On("GenerateTraffic", "n1", "g2", "magmawifi", "magmamagma").Return(mockGenericCommandResp, nil)
	// ---
	// Traffic Test 3
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test3_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "restore_enodeb1", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	// ---
	// Restore Enodeb config
	// ---
	actualState, actualDuration, err = sm.Run("restore_enodeb1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "verify_config2", actualState)
	assert.Equal(t, 10*time.Minute, actualDuration)

	// ---
	// Traffic Test 4
	// ---
	actualState, actualDuration, err = sm.Run("traffic_test4_1", testConfig, nil)
	assert.NoError(t, err)
	assert.Equal(t, "check_for_upgrade", actualState)
	assert.Equal(t, 1*time.Minute, actualDuration)

	cli.AssertExpectations(t)
	mockMagmad.AssertExpectations(t)
}

func SetupTests(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &pluginimpl.BaseOrchestratorPlugin{})
	_ = plugin.RegisterPluginForTests(t, &plugin2.FbinternalOrchestratorPlugin{})
	tcTestInit.StartTestService(t)
	cfgTestInit.StartTestService(t)
	stateTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)

	frozenClock := 1000 * time.Hour
	clock.SetAndFreezeClock(t, time.Unix(0, 0).Add(frozenClock))
	defer clock.UnfreezeClock(t)
}

func RegisterAGW(t *testing.T) {
	// Register an AGW
	err := configurator.CreateNetwork(configurator.Network{ID: "n1"})
	assert.NoError(t, err)
	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{Type: orc8r.UpgradeTierEntityType, Key: "t1", Config: &models2.Tier{Name: "t1", Version: "0.3.74-1560824953-b50f1bab"}},
	)
	assert.NoError(t, err)
	_, err = configurator.CreateEntity(
		"n1",
		configurator.NetworkEntity{
			Type:         orc8r.MagmadGatewayType,
			Key:          "g1",
			Config:       &models2.MagmadGatewayConfigs{},
			PhysicalID:   "hw1",
			Associations: []storage.TypeAndKey{{Type: orc8r.UpgradeTierEntityType, Key: "t1"}},
		},
	)
	assert.NoError(t, err)
}

func GetEnodebTestConfig() *models.EnodebdTestConfig {
	testConfig := &models.EnodebdTestConfig{
		AgwConfig: &models.AgwTestConfig{
			PackageRepo:     swag.String("https://packages.magma.etagecom.io"),
			ReleaseChannel:  swag.String("stretch-beta"),
			SLACKWebhook:    swag.String("https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"),
			TargetGatewayID: swag.String("g1"),
			TargetTier:      swag.String("t1"),
		},
		EnodebSN:    swag.String("1202000038269KP0037"),
		NetworkID:   swag.String("n1"),
		Ssid:        ("magmawifi"),
		SsidPw:      ("magmamagma"),
		TrafficGwID: swag.String("g2"),
		EnodebConfig: &ltemodels.EnodebConfiguration{
			BandwidthMhz:           20,
			CellID:                 swag.Uint32(138777000),
			DeviceClass:            "Baicells ID TDD/FDD",
			Earfcndl:               44590,
			Pci:                    260,
			SpecialSubframePattern: 7,
			SubframeAssignment:     2,
			Tac:                    1,
			TransmitEnabled:        swag.Bool(true),
		},
	}
	return testConfig
}

func GetMockObjects() (*mockMagmadClient, *protos.GenericCommandResponse) {
	mockMagmad := &mockMagmadClient{}
	mockResponse := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"response": &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "200"}},
		},
	}
	mockGenericCommandResp := &protos.GenericCommandResponse{
		Response: mockResponse,
	}
	return mockMagmad, mockGenericCommandResp
}

type mockMagmadClient struct {
	mock.Mock
}

func (m *mockMagmadClient) GenerateTraffic(networkId string, trafficGatewayId string, ssid string, pw string) (*protos.GenericCommandResponse, error) {
	args := m.Called(networkId, trafficGatewayId, ssid, pw)
	return args.Get(0).(*protos.GenericCommandResponse), args.Error(1)
}

func (m *mockMagmadClient) RebootEnodeb(networkId string, gatewayId string, enodebSerial string) (*protos.GenericCommandResponse, error) {
	args := m.Called(networkId, gatewayId, enodebSerial)
	return args.Get(0).(*protos.GenericCommandResponse), args.Error(1)
}

func (m *mockMagmadClient) GetEnodebStatus(networkId string, hwId string) (*ltemodels.EnodebState, error) {
	args := m.Called(networkId, hwId)
	return args.Get(0).(*ltemodels.EnodebState), args.Error(1)
}

type mockClient struct {
	mock.Mock
}

func (client *mockClient) Get(url string) (resp *http.Response, err error) {
	args := client.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (client *mockClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	args := client.Called(url, contentType, body)
	return args.Get(0).(*http.Response), args.Error(1)
}
