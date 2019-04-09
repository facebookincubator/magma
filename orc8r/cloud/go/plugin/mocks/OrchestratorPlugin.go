// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"magma/orc8r/cloud/go/obsidian/handlers"
	goregistry "magma/orc8r/cloud/go/registry"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/service/config"
	"magma/orc8r/cloud/go/services/metricsd"
	"magma/orc8r/cloud/go/services/streamer/mconfig/factory"
	"magma/orc8r/cloud/go/services/streamer/providers"

	"github.com/stretchr/testify/mock"
)

// OrchestratorPlugin is an autogenerated mock type for the OrchestratorPlugin type
type OrchestratorPlugin struct {
	mock.Mock
}

// GetMconfigBuilders provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetMconfigBuilders() []factory.MconfigBuilder {
	ret := _m.Called()

	var r0 []factory.MconfigBuilder
	if rf, ok := ret.Get(0).(func() []factory.MconfigBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]factory.MconfigBuilder)
		}
	}

	return r0
}

// GetMetricsProfiles provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	ret := _m.Called()

	var r0 []metricsd.MetricsProfile
	if rf, ok := ret.Get(0).(func() []metricsd.MetricsProfile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]metricsd.MetricsProfile)
		}
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetObsidianHandlers provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []handlers.Handler {
	ret := _m.Called()

	var r0 []handlers.Handler
	if rf, ok := ret.Get(0).(func() []handlers.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]handlers.Handler)
		}
	}

	return r0
}

// GetSerdes provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetSerdes() []serde.Serde {
	ret := _m.Called()

	var r0 []serde.Serde
	if rf, ok := ret.Get(0).(func() []serde.Serde); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]serde.Serde)
		}
	}

	return r0
}

// GetServices provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetServices() []goregistry.ServiceLocation {
	ret := _m.Called()

	var r0 []goregistry.ServiceLocation
	if rf, ok := ret.Get(0).(func() []goregistry.ServiceLocation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]goregistry.ServiceLocation)
		}
	}

	return r0
}

// GetStreamerProviders provides a mock function with given fields:
func (_m *OrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	ret := _m.Called()

	var r0 []providers.StreamProvider
	if rf, ok := ret.Get(0).(func() []providers.StreamProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]providers.StreamProvider)
		}
	}

	return r0
}
