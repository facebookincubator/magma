// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Run make gen at FeG to re-generate

package mocks

import (
	context "context"
	protos "magma/lte/cloud/go/protos"

	mock "github.com/stretchr/testify/mock"
)

// SessionProxyResponderServer is an autogenerated mock type for the SessionProxyResponderServer type
type SessionProxyResponderServer struct {
	mock.Mock
}

// ChargingReAuth provides a mock function with given fields: _a0, _a1
func (_m *SessionProxyResponderServer) ChargingReAuth(_a0 context.Context, _a1 *protos.ChargingReAuthRequest) (*protos.ChargingReAuthAnswer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.ChargingReAuthAnswer
	if rf, ok := ret.Get(0).(func(context.Context, *protos.ChargingReAuthRequest) *protos.ChargingReAuthAnswer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.ChargingReAuthAnswer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.ChargingReAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PolicyReAuth provides a mock function with given fields: _a0, _a1
func (_m *SessionProxyResponderServer) PolicyReAuth(_a0 context.Context, _a1 *protos.PolicyReAuthRequest) (*protos.PolicyReAuthAnswer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.PolicyReAuthAnswer
	if rf, ok := ret.Get(0).(func(context.Context, *protos.PolicyReAuthRequest) *protos.PolicyReAuthAnswer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.PolicyReAuthAnswer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.PolicyReAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
