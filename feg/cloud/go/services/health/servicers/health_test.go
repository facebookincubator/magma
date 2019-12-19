/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers_test

import (
	"context"
	"testing"
	"time"

	fegprotos "magma/feg/cloud/go/protos"
	"magma/feg/cloud/go/services/health"
	"magma/feg/cloud/go/services/health/servicers"
	"magma/feg/cloud/go/services/health/test_utils"
	"magma/feg/cloud/go/services/health/util"
	"magma/orc8r/cloud/go/blobstore"
	"magma/orc8r/cloud/go/blobstore/mocks"
	"magma/orc8r/cloud/go/clock"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/pluginimpl/models"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/serde"
	configurator_test_init "magma/orc8r/cloud/go/services/configurator/test_init"
	"magma/orc8r/cloud/go/services/device"
	device_test_init "magma/orc8r/cloud/go/services/device/test_init"
	"magma/orc8r/cloud/go/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHealthServer_GetHealth(t *testing.T) {
	configurator_test_init.StartTestService(t)
	store := &mocks.TransactionalBlobStorage{}
	factory := &mocks.BlobStorageFactory{}
	clock.SetAndFreezeClock(t, time.Unix(1551916956, 0))
	service := servicers.NewTestHealthServer(factory)

	gwStatusReq := &fegprotos.GatewayStatusRequest{
		NetworkId: test_utils.TestFegNetwork,
		LogicalId: test_utils.TestFegLogicalId1,
	}
	healthyReq := test_utils.GetHealthyRequest()
	id1 := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  test_utils.TestFegLogicalId1,
	}
	marshaledHealthStats, err := protos.Marshal(healthyReq.HealthStats)
	assert.NoError(t, err)
	healthBlob := blobstore.Blob{
		Value: marshaledHealthStats,
	}
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Once()
	store.On("Get", test_utils.TestFegNetwork, id1).Return(healthBlob, nil).Once()
	store.On("Commit").Return(nil).Once()

	stats, err := service.GetHealth(context.Background(), gwStatusReq)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthStatus_HEALTHY, stats.Health.Health)
	assert.Equal(t, healthyReq.HealthStats.SystemStatus, stats.SystemStatus)
	assert.Equal(t, healthyReq.HealthStats.ServiceStatus, stats.ServiceStatus)

	store.AssertExpectations(t)

	unhealthyReq := test_utils.GetUnhealthyRequest()

	marshaledUnhealthyStats, err := protos.Marshal(unhealthyReq.HealthStats)
	unhealthyBlob := blobstore.Blob{
		Value: marshaledUnhealthyStats,
	}
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Once()
	store.On("Get", test_utils.TestFegNetwork, id1).Return(unhealthyBlob, nil).Once()
	store.On("Commit").Return(nil).Once()

	stats, err = service.GetHealth(context.Background(), gwStatusReq)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthStatus_UNHEALTHY, stats.Health.Health)
	assert.Equal(t, unhealthyReq.HealthStats.SystemStatus, stats.SystemStatus)
	assert.Equal(t, unhealthyReq.HealthStats.ServiceStatus, stats.ServiceStatus)

	store.AssertExpectations(t)
}

// Test that a single feg will always remain ACTIVE
func TestHealthServer_UpdateHealth_SingleGateway(t *testing.T) {
	configurator_test_init.StartTestService(t)
	device_test_init.StartTestService(t)
	store := &mocks.TransactionalBlobStorage{}
	factory := &mocks.BlobStorageFactory{}
	clock.SetAndFreezeClock(t, time.Unix(1551916956, 0))
	service := servicers.NewTestHealthServer(factory)

	test_utils.RegisterNetwork(t, test_utils.TestFegNetwork)
	_ = serde.RegisterSerdes(serde.NewBinarySerde(device.SerdeDomain, orc8r.AccessGatewayRecordType, &models.GatewayDevice{}))
	test_utils.RegisterGateway(t, test_utils.TestFegNetwork, test_utils.TestFegHwId1, test_utils.TestFegLogicalId1)

	// Use Healthy Request metrics
	healthyRequest := test_utils.GetHealthyRequest()
	clusterBlob, err := util.ClusterToBlob(test_utils.TestFegNetwork, test_utils.TestFegLogicalId1)
	assert.NoError(t, err)
	healthBlob, err := util.HealthToBlob(test_utils.TestFegLogicalId1, healthyRequest.GetHealthStats())
	assert.NoError(t, err)
	clusterTK := storage.TypeAndKey{
		Type: health.ClusterStatusType,
		Key:  test_utils.TestFegNetwork,
	}
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(2)
	store.On("CreateOrUpdate", test_utils.TestFegNetwork, []blobstore.Blob{healthBlob}).Return(nil).Once()
	store.On("GetExistingKeys", []string{test_utils.TestFegNetwork}, mock.AnythingOfType("SearchFilter")).Return([]string{}, nil)
	store.On("CreateOrUpdate", test_utils.TestFegNetwork, []blobstore.Blob{clusterBlob}).Return(nil).Once()
	store.On("Get", test_utils.TestFegNetwork, clusterTK).Return(clusterBlob, nil).Once()
	store.On("Commit").Return(nil).Times(2)

	res, err := service.UpdateHealth(context.Background(), healthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_UP, res.Action)
	store.AssertExpectations(t)

	// Ensure we stay active with only one feg, even if it is unhealthy
	unhealthyRequest := test_utils.GetUnhealthyRequest()
	unhealthyBlob, err := util.HealthToBlob(test_utils.TestFegLogicalId1, unhealthyRequest.GetHealthStats())
	assert.NoError(t, err)
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(2)
	store.On("CreateOrUpdate", test_utils.TestFegNetwork, []blobstore.Blob{unhealthyBlob}).Return(nil)
	store.On("CreateOrUpdate", test_utils.TestFegNetwork, []blobstore.Blob{clusterBlob}).Return(nil)
	store.On("GetExistingKeys", []string{test_utils.TestFegNetwork}, mock.Anything).Return([]string{test_utils.TestFegNetwork}, nil)
	store.On("Get", test_utils.TestFegNetwork, clusterTK).Return(clusterBlob, nil).Once()
	store.On("Commit").Return(nil).Times(2)

	res, err = service.UpdateHealth(context.Background(), unhealthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_UP, res.Action)
	store.AssertExpectations(t)
}

func TestHealthServer_UpdateHealth_DualFeg_HealthyActive(t *testing.T) {
	configurator_test_init.StartTestService(t)
	device_test_init.StartTestService(t)
	store := &mocks.TransactionalBlobStorage{}
	factory := &mocks.BlobStorageFactory{}
	clock.SetAndFreezeClock(t, time.Unix(1551916956, 0))
	service := servicers.NewTestHealthServer(factory)

	testNetworkID, gwId, gwId2 := registerTwoFegs(t)

	healthyRequest := test_utils.GetHealthyRequest()
	healthBlob, err := util.HealthToBlob(gwId, healthyRequest.GetHealthStats())
	assert.NoError(t, err)
	healthBlob2, err := util.HealthToBlob(gwId2, healthyRequest.GetHealthStats())
	assert.NoError(t, err)
	clusterBlob, err := util.ClusterToBlob(test_utils.TestFegNetwork, gwId)
	assert.NoError(t, err)
	clusterTK := storage.TypeAndKey{
		Type: health.ClusterStatusType,
		Key:  test_utils.TestFegNetwork,
	}
	healthTK := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  gwId,
	}
	healthTK2 := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  gwId2,
	}
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(3)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{healthBlob}).Return(nil)
	store.On("GetExistingKeys", []string{testNetworkID}, mock.AnythingOfType("SearchFilter")).Return([]string{testNetworkID}, nil)
	store.On("Get", testNetworkID, clusterTK).Return(clusterBlob, nil)
	store.On("Get", testNetworkID, healthTK2).Return(healthBlob2, nil)
	store.On("Commit").Return(nil).Times(3)

	res, err := service.UpdateHealth(context.Background(), healthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_UP, res.Action)
	store.AssertExpectations(t)

	// Update test servicer to simulate like this request is coming from standby feg
	service.Feg1 = false
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(3)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{healthBlob2}).Return(nil)
	store.On("GetExistingKeys", []string{testNetworkID}, mock.AnythingOfType("SearchFilter")).Return([]string{testNetworkID}, nil)
	store.On("Get", testNetworkID, clusterTK).Return(clusterBlob, nil)
	store.On("Get", testNetworkID, healthTK).Return(healthBlob, nil)
	store.On("Commit").Return(nil).Times(3)

	// Standby FeG receives SYSTEM_DOWN, since active is still healthy
	res, err = service.UpdateHealth(context.Background(), healthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_DOWN, res.Action)
	store.AssertExpectations(t)
}

func TestNewHealthServer_UpdateHealth_FailoverFromActive(t *testing.T) {
	configurator_test_init.StartTestService(t)
	device_test_init.StartTestService(t)
	store := &mocks.TransactionalBlobStorage{}
	factory := &mocks.BlobStorageFactory{}
	clock.SetAndFreezeClock(t, time.Unix(1551916956, 0))
	service := servicers.NewTestHealthServer(factory)

	testNetworkID, gwId, gwId2 := registerTwoFegs(t)

	// Simulate an unhealthy active FeG, and thus a failover
	unhealthyRequest := test_utils.GetUnhealthyRequest()
	healthyRequest := test_utils.GetHealthyRequest()
	unhealthyBlob, err := util.HealthToBlob(gwId, unhealthyRequest.GetHealthStats())
	assert.NoError(t, err)
	healthyBlob, err := util.HealthToBlob(gwId2, healthyRequest.GetHealthStats())
	assert.NoError(t, err)
	clusterBlob, err := util.ClusterToBlob(testNetworkID, gwId)
	assert.NoError(t, err)

	clusterTK := storage.TypeAndKey{
		Type: health.ClusterStatusType,
		Key:  test_utils.TestFegNetwork,
	}
	healthTK2 := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  gwId2,
	}
	updatedClusterBlob, err := util.ClusterToBlob(testNetworkID, gwId2)
	assert.NoError(t, err)

	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(4)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{unhealthyBlob}).Return(nil)
	store.On("GetExistingKeys", []string{testNetworkID}, mock.AnythingOfType("SearchFilter")).Return([]string{testNetworkID}, nil)
	store.On("Get", testNetworkID, clusterTK).Return(clusterBlob, nil)
	store.On("Get", testNetworkID, healthTK2).Return(healthyBlob, nil)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{updatedClusterBlob}).Return(nil)
	store.On("Commit").Return(nil).Times(4)

	res, err := service.UpdateHealth(context.Background(), unhealthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_DOWN, res.Action)
	store.AssertExpectations(t)
}

func TestNewHealthServer_UpdateHealth_FailoverFromStandby(t *testing.T) {
	configurator_test_init.StartTestService(t)
	device_test_init.StartTestService(t)
	store := &mocks.TransactionalBlobStorage{}
	factory := &mocks.BlobStorageFactory{}
	clock.SetAndFreezeClock(t, time.Unix(1551916956, 0))
	service := servicers.NewTestHealthServer(factory)

	testNetworkID, gwId, gwId2 := registerTwoFegs(t)

	// Update test servicer to act as though this request is coming from the standby feg
	service.Feg1 = false

	// Simulate that the active's last update was too long ago
	healthyRequestTooLongAgo := test_utils.GetHealthyRequest()
	healthyRequestTooLongAgo.HealthStats.Time = 0
	healthyRequest := test_utils.GetHealthyRequest()
	healthyBlob, err := util.HealthToBlob(gwId2, healthyRequest.GetHealthStats())
	assert.NoError(t, err)
	unhealthyBlob, err := util.HealthToBlob(gwId, healthyRequestTooLongAgo.GetHealthStats())
	assert.NoError(t, err)
	clusterTK := storage.TypeAndKey{
		Type: health.ClusterStatusType,
		Key:  test_utils.TestFegNetwork,
	}
	healthTK := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  gwId,
	}
	clusterBlob, err := util.ClusterToBlob(testNetworkID, gwId)
	assert.NoError(t, err)
	updatedClusterBlob, err := util.ClusterToBlob(testNetworkID, gwId2)
	assert.NoError(t, err)

	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(4)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{healthyBlob}).Return(nil)
	store.On("GetExistingKeys", []string{testNetworkID}, mock.AnythingOfType("SearchFilter")).Return([]string{testNetworkID}, nil)
	store.On("Get", testNetworkID, clusterTK).Return(clusterBlob, nil)
	store.On("Get", testNetworkID, healthTK).Return(unhealthyBlob, nil)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{updatedClusterBlob}).Return(nil)
	store.On("Commit").Return(nil).Times(4)

	res, err := service.UpdateHealth(context.Background(), healthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_UP, res.Action)
	store.AssertExpectations(t)
}

func TestNewHealtherServer_UpdateHealth_AllUnhealthy(t *testing.T) {
	configurator_test_init.StartTestService(t)
	device_test_init.StartTestService(t)
	store := &mocks.TransactionalBlobStorage{}
	factory := &mocks.BlobStorageFactory{}
	clock.SetAndFreezeClock(t, time.Unix(1551916956, 0))
	service := servicers.NewTestHealthServer(factory)

	testNetworkID, gwId, gwId2 := registerTwoFegs(t)

	// Simulate that both the active and standby are unhealthy
	unhealthyRequest := test_utils.GetUnhealthyRequest()
	unhealthyBlob, err := util.HealthToBlob(gwId, unhealthyRequest.HealthStats)
	unhealthyBlob2, err := util.HealthToBlob(gwId2, unhealthyRequest.HealthStats)
	clusterBlob, err := util.ClusterToBlob(testNetworkID, gwId)
	assert.NoError(t, err)
	clusterTK := storage.TypeAndKey{
		Type: health.ClusterStatusType,
		Key:  test_utils.TestFegNetwork,
	}
	healthTK := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  gwId,
	}
	healthTK2 := storage.TypeAndKey{
		Type: health.HealthStatusType,
		Key:  gwId2,
	}
	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(3)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{unhealthyBlob}).Return(nil)
	store.On("GetExistingKeys", []string{testNetworkID}, mock.AnythingOfType("SearchFilter")).Return([]string{testNetworkID}, nil)
	store.On("Get", testNetworkID, clusterTK).Return(clusterBlob, nil)
	store.On("Get", testNetworkID, healthTK2).Return(unhealthyBlob, nil)
	store.On("Commit").Return(nil).Times(3)

	res, err := service.UpdateHealth(context.Background(), unhealthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_UP, res.Action)
	store.AssertExpectations(t)

	// Update test servicer to simulate like this request is coming from standby feg
	service.Feg1 = false

	factory.On("StartTransaction", mock.Anything).Return(store, nil).Times(3)
	store.On("CreateOrUpdate", testNetworkID, []blobstore.Blob{unhealthyBlob2}).Return(nil)
	store.On("GetExistingKeys", []string{testNetworkID}, mock.AnythingOfType("SearchFilter")).Return([]string{testNetworkID}, nil)
	store.On("Get", testNetworkID, clusterTK).Return(clusterBlob, nil)
	store.On("Get", testNetworkID, healthTK).Return(unhealthyBlob, nil)
	store.On("Commit").Return(nil).Times(3)

	res, err = service.UpdateHealth(context.Background(), unhealthyRequest)
	assert.NoError(t, err)
	assert.Equal(t, fegprotos.HealthResponse_SYSTEM_DOWN, res.Action)
	store.AssertExpectations(t)
}

func registerTwoFegs(t *testing.T) (string, string, string) {
	test_utils.RegisterNetwork(t, test_utils.TestFegNetwork)
	_ = serde.RegisterSerdes(serde.NewBinarySerde(device.SerdeDomain, orc8r.AccessGatewayRecordType, &models.GatewayDevice{}))
	test_utils.RegisterGateway(
		t,
		test_utils.TestFegNetwork,
		test_utils.TestFegHwId1,
		test_utils.TestFegLogicalId1,
	)
	test_utils.RegisterGateway(
		t,
		test_utils.TestFegNetwork,
		test_utils.TestFegHwId2,
		test_utils.TestFegLogicalId2,
	)
	return test_utils.TestFegNetwork, test_utils.TestFegLogicalId1, test_utils.TestFegLogicalId2
}
