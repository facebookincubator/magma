/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 *  LICENSE file in the root directory of this source tree.
 */

package tests

import (
	"context"
	"encoding/json"
	"testing"

	"magma/orc8r/cloud/go/errors"
	"magma/orc8r/cloud/go/identity"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/registry"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/service/middleware/unary/interceptors/tests"
	"magma/orc8r/cloud/go/services/magmad"
	magmad_protos "magma/orc8r/cloud/go/services/magmad/protos"
	magmad_test_init "magma/orc8r/cloud/go/services/magmad/test_init"
	"magma/orc8r/cloud/go/services/state"
	test_service "magma/orc8r/cloud/go/services/state/test_init"

	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	typeName   = "typeName"
	testAgHwId = "Test-AGW-Hw-Id"
)

type stateBundle struct {
	value interface{}
	state *protos.State
	ID    *protos.StateID
}

func makeStateBundle(typeVal string, key string, value interface{}) stateBundle {
	marshaledValue, _ := json.Marshal(value)
	ID := protos.StateID{Type: typeVal, DeviceID: key}
	state := protos.State{Type: typeVal, DeviceID: key, Value: marshaledValue}
	return stateBundle{state: &state, ID: &ID, value: value}
}

func TestStateService(t *testing.T) {
	// Set up test networkID, hwID, and encode into context
	magmad_test_init.StartTestService(t)
	networkID, err := magmad.RegisterNetwork(
		&magmad_protos.MagmadNetworkRecord{Name: "State Service Test"},
		"state_service_test_network")
	hwId := protos.AccessGatewayID{Id: testAgHwId}
	magmad.RegisterGateway(
		networkID,
		&magmad_protos.AccessGatewayRecord{HwId: &hwId, Name: "Test GW Name"})
	csn := tests.StartMockGwAccessControl(t, []string{testAgHwId})
	ctx := metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs(identity.CLIENT_CERT_SN_KEY, csn[0]))

	// Create States, IDs, values
	value0 := Name{Name: "name0"}
	value1 := Name{Name: "name1"}
	value2 := NameAndAge{Name: "name2", Age: 20}
	bundle0 := makeStateBundle(typeName, "key0", value0)
	bundle1 := makeStateBundle(typeName, "key1", value1)
	bundle2 := makeStateBundle(typeName, "key2", value2)

	test_service.StartTestService(t)
	err = serde.RegisterSerdes(&Serde{})
	assert.NoError(t, err)
	client, conn, err := getClient()
	defer conn.Close()

	// Check contract for empty network
	response, err := client.GetStates(ctx, makeGetStatesRequest(networkID, bundle0))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(response.States))

	// Report and read back
	_, err = client.ReportStates(ctx, makeReportStatesRequest(bundle0, bundle1))
	assert.NoError(t, err)
	response, err = client.GetStates(ctx, makeGetStatesRequest(networkID, bundle0, bundle1))
	assert.NoError(t, err)
	testGetStatesResponse(t, response, bundle0, bundle1)

	// Report a state with fields the corresponding serde does not expect
	_, err = client.ReportStates(ctx, makeReportStatesRequest(bundle2))
	assert.NoError(t, err)
	response, err = client.GetStates(ctx, makeGetStatesRequest(networkID, bundle2))
	assert.NoError(t, err)
	testGetStatesResponse(t, response, bundle2)

	// Delete and read back
	_, err = client.DeleteStates(ctx, makeDeleteStatesRequest(networkID, bundle0, bundle2))
	assert.NoError(t, err)
	response, err = client.GetStates(ctx, makeGetStatesRequest(networkID, bundle0, bundle1, bundle2))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response.States))
	testGetStatesResponse(t, response, bundle1)
}

type NameAndAge struct {
	// name
	Name string `json:"name"`
	// age
	Age int `json:"age"`
}

type Name struct {
	// name
	Name string `json:"name"`
}

type Serde struct {
}

func (*Serde) GetDomain() string {
	return state.SerdeDomain
}

func (*Serde) GetType() string {
	return typeName
}

func (*Serde) Serialize(in interface{}) ([]byte, error) {
	return json.Marshal(in)

}

func (*Serde) Deserialize(message []byte) (interface{}, error) {
	res := Name{}
	err := json.Unmarshal(message, &res)
	return res, err
}

func getClient() (protos.StateServiceClient, *grpc.ClientConn, error) {
	conn, err := registry.GetConnection(state.ServiceName)
	if err != nil {
		initErr := errors.NewInitError(err, state.ServiceName)
		glog.Error(initErr)
		return nil, nil, initErr
	}
	return protos.NewStateServiceClient(conn), conn, err
}

func testGetStatesResponse(t *testing.T, response *protos.GetStatesResponse, bundles ...stateBundle) {
	keyToValue := map[string][]byte{}
	states := response.States
	for _, state := range states {
		keyToValue[state.DeviceID] = state.Value
	}
	for _, bundle := range bundles {
		value := keyToValue[bundle.ID.DeviceID]
		assert.Equal(t, bundle.state.Value, value)
	}
}

func makeGetStatesRequest(networkID string, bundles ...stateBundle) *protos.GetStatesRequest {
	res := protos.GetStatesRequest{}
	res.NetworkID = networkID
	res.Ids = []*protos.StateID{}
	for _, bundle := range bundles {
		res.Ids = append(res.Ids, bundle.ID)
	}
	return &res
}

func makeReportStatesRequest(bundles ...stateBundle) *protos.ReportStatesRequest {
	res := protos.ReportStatesRequest{}
	res.States = makeStates(bundles)
	return &res
}

func makeDeleteStatesRequest(networkID string, bundles ...stateBundle) *protos.DeleteStatesRequest {
	res := protos.DeleteStatesRequest{}
	res.NetworkID = networkID
	res.Ids = makeIDs(bundles)
	return &res
}

func makeStates(bundles []stateBundle) []*protos.State {
	states := []*protos.State{}
	for _, bundle := range bundles {
		states = append(states, bundle.state)
	}
	return states
}

func makeIDs(bundles []stateBundle) []*protos.StateID {
	IDs := []*protos.StateID{}
	for _, bundle := range bundles {
		IDs = append(IDs, bundle.ID)
	}
	return IDs
}
