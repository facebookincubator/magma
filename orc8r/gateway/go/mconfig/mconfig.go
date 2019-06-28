/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Package mconfig provides gateway Go support for cloud managed configuration (mconfig)
package mconfig

import (
	"fmt"

	"magma/orc8r/cloud/go/protos"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func GetServiceConfigs(service string, result proto.Message) error {
	current := GetGatewayConfigs()
	anyCfg, found := current.ConfigsByKey[service]
	if !found {
		cfgMu.Lock()
		defer cfgMu.Unlock()
		return fmt.Errorf("No configs found for service: '%s' in %s", service, lastFilePath)
	}
	return ptypes.UnmarshalAny(anyCfg, result)
}

func GetGatewayConfigs() *protos.GatewayConfigs {
	current := localConfig.Load().(*protos.GatewayConfigs)
	if current == nil {
		// initial refresh, only do it once
		err := RefreshConfigs()
		if err != nil || localConfig.Load() == nil {
			localConfig.Store(new(protos.GatewayConfigs))
		}
		current = localConfig.Load().(*protos.GatewayConfigs)
	}
	return current
}
