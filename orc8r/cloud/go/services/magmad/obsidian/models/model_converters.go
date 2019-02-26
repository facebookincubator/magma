/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package models

import (
	"magma/orc8r/cloud/go/protos"
	checkind_models "magma/orc8r/cloud/go/services/checkind/obsidian/models"
	"magma/orc8r/cloud/go/services/magmad/obsidian/handlers/view_factory"
	magmadprotos "magma/orc8r/cloud/go/services/magmad/protos"
)

// GatewayStateType is the manually defined model type for Gateway State
type GatewayStateType struct {
	Config    map[string]interface{}         `json:"config"`
	GatewayID string                         `json:"gateway_id"`
	Record    *AccessGatewayRecord           `json:"record"`
	Status    *checkind_models.GatewayStatus `json:"status"`
}

// GatewayStateToModel converts a storage.GatewayState object to the equivalent
// model.GatewayStateType
func GatewayStateToModel(state *view_factory.GatewayState) (*GatewayStateType, error) {
	modelState := &GatewayStateType{
		GatewayID: state.GatewayID,
		Config:    state.Config,
	}
	modelStatus, err := gatewayStatusToModel(state.Status)
	if err != nil {
		return nil, err
	}
	modelRecord, err := gatewayRecordToModel(state.Record)
	if err != nil {
		return nil, err
	}
	modelState.Status = modelStatus
	modelState.Record = modelRecord
	return modelState, nil
}

// GatewayStateMapToModelList converts a map of storage.GatewayState objects
// to an equivalent list of model.GatewayStateType objects
func GatewayStateMapToModelList(states map[string]*view_factory.GatewayState) ([]*GatewayStateType, error) {
	models := make([]*GatewayStateType, 0, len(states))
	for _, state := range states {
		gatewayState, err := GatewayStateToModel(state)
		if err != nil {
			return nil, err
		}
		models = append(models, gatewayState)
	}
	return models, nil
}

func gatewayStatusToModel(status *protos.GatewayStatus) (*checkind_models.GatewayStatus, error) {
	if status == nil {
		return nil, nil
	}
	modelStatus := &checkind_models.GatewayStatus{}
	err := modelStatus.FromMconfig(status)
	return modelStatus, err
}

func gatewayRecordToModel(record *magmadprotos.AccessGatewayRecord) (*AccessGatewayRecord, error) {
	if record == nil {
		return nil, nil
	}
	modelRecord := &AccessGatewayRecord{}
	err := modelRecord.FromMconfig(record)
	return modelRecord, err
}
