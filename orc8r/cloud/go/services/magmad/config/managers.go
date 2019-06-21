/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package config

import (
	"fmt"
	"reflect"

	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/config"
	magmad_protos "magma/orc8r/cloud/go/services/magmad/protos"
)

// To be deprecated! Config service DB tables have been seeded with magmad
// network configs that were migrated from legacy magmad network configs,
// so this will stick around for a bit. We can delete this after deleting
// all magmad network config types from the config service (to come in a
// future migration)
type MagmadNetworkConfigManager struct{}

func (*MagmadNetworkConfigManager) GetDomain() string {
	return config.SerdeDomain
}

func (*MagmadNetworkConfigManager) GetType() string {
	return orc8r.MagmadNetworkType
}

func (*MagmadNetworkConfigManager) Serialize(config interface{}) ([]byte, error) {
	castedConfig, ok := config.(*magmad_protos.MagmadNetworkRecord)
	if !ok {
		return nil, fmt.Errorf(
			"Invalid magmad network config type. Expected *MagmadNetworkRecord, received %s",
			reflect.TypeOf(config),
		)
	}
	if err := magmad_protos.ValidateNetworkConfig(castedConfig); err != nil {
		return nil, fmt.Errorf("Invalid network config: %s", err)
	}
	return protos.MarshalIntern(castedConfig)
}

func (*MagmadNetworkConfigManager) Deserialize(message []byte) (interface{}, error) {
	cfg := &magmad_protos.MagmadNetworkRecord{}
	err := protos.Unmarshal(message, cfg)
	return cfg, err
}

type MagmadGatewayConfigManager struct{}

func (*MagmadGatewayConfigManager) GetDomain() string {
	return config.SerdeDomain
}

func (*MagmadGatewayConfigManager) GetType() string {
	return orc8r.MagmadGatewayType
}

func (*MagmadGatewayConfigManager) Serialize(config interface{}) ([]byte, error) {
	castedConfig, ok := config.(*magmad_protos.MagmadGatewayConfig)
	if !ok {
		return nil, fmt.Errorf(
			"Invalid magmad gateway config type. Expected *MagmadGatewayConfig, received %s",
			reflect.TypeOf(config),
		)
	}
	if err := magmad_protos.ValidateGatewayConfig(castedConfig); err != nil {
		return nil, fmt.Errorf("Invalid gateway config: %s", err)
	}
	return protos.MarshalIntern(castedConfig)
}

func (*MagmadGatewayConfigManager) Deserialize(message []byte) (interface{}, error) {
	cfg := &magmad_protos.MagmadGatewayConfig{}
	err := protos.Unmarshal(message, cfg)
	return cfg, err
}
