/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package config

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	"magma/feg/cloud/go/protos/mconfig"
	config_protos "magma/feg/cloud/go/services/controller/protos"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/config"
)

type Builder struct{}

func (builder *Builder) Build(networkId string, gatewayId string) (map[string]proto.Message, error) {
	emptyRet := map[string]proto.Message{}
	gwConfig, err := GetGatewayConfig(networkId, gatewayId)
	if err != nil {
		return emptyRet, err
	}
	if gwConfig == nil {
		return emptyRet, nil
	}

	s6ac := gwConfig.GetS6A()
	gxc := gwConfig.GetGx()
	gyc := gwConfig.GetGy()
	hss := gwConfig.GetHss()
	swxc := gwConfig.GetSwx()
	eapAka := gwConfig.GetEapAka()

	hssSubProfile := map[string]*mconfig.HSSConfig_SubscriptionProfile{}
	for imsi, profile := range hss.GetSubProfiles() {
		hssSubProfile[imsi] = profile.ToMconfig()
	}
	healthc := gwConfig.GetHealth()

	return map[string]proto.Message{
		"s6a_proxy": &mconfig.S6AConfig{
			LogLevel:                protos.LogLevel_INFO,
			Server:                  s6ac.GetServer().ToMconfig(),
			RequestFailureThreshold: healthc.GetRequestFailureThreshold(),
			MinimumRequestThreshold: healthc.GetMinimumRequestThreshold(),
		},
		"session_proxy": &mconfig.SessionProxyConfig{
			LogLevel: protos.LogLevel_INFO,
			Gx: &mconfig.GxConfig{
				Server: gxc.GetServer().ToMconfig(),
			},
			Gy: &mconfig.GyConfig{
				Server:     gyc.GetServer().ToMconfig(),
				InitMethod: mconfig.GyInitMethod(gyc.GetInitMethod()),
			},
			RequestFailureThreshold: healthc.GetRequestFailureThreshold(),
			MinimumRequestThreshold: healthc.GetMinimumRequestThreshold(),
		},
		"health": &mconfig.GatewayHealthConfig{
			RequiredServices:          healthc.GetHealthServices(),
			UpdateIntervalSecs:        healthc.GetUpdateIntervalSecs(),
			UpdateFailureThreshold:    healthc.GetUpdateFailureThreshold(),
			CloudDisconnectPeriodSecs: healthc.GetCloudDisablePeriodSecs(),
			LocalDisconnectPeriodSecs: healthc.GetLocalDisablePeriodSecs(),
		},
		"hss": &mconfig.HSSConfig{
			Server:            hss.GetServer().ToMconfig(),
			LteAuthOp:         hss.GetLteAuthOp(),
			LteAuthAmf:        hss.GetLteAuthAmf(),
			DefaultSubProfile: hss.GetDefaultSubProfile().ToMconfig(),
			SubProfiles:       hssSubProfile,
		},
		"swx_proxy": &mconfig.SwxConfig{
			LogLevel:            protos.LogLevel_INFO,
			Server:              swxc.GetServer().ToMconfig(),
			VerifyAuthorization: swxc.GetVerifyAuthorization(),
			CacheTTLSeconds:     swxc.GetCacheTTLSeconds(),
		},
		"eap_aka": &mconfig.EapAkaConfig{
			LogLevel: protos.LogLevel_INFO,
			Timeout:  eapAka.GetTimeout().ToMconfig(),
			PlmnIds:  eapAka.GetPlmnIds(),
		},
	}, nil
}

// GetGatewayConfig returns the specified GW's configs. gatewayId is Logical GW ID
func GetGatewayConfig(networkId string, gatewayId string) (*config_protos.Config, error) {
	var (
		networkCfg *config_protos.Config
		ok         bool
	)
	// Get Network configs to fill in GW configs or return
	netCfg, nerr := config.GetConfig(networkId, FegNetworkType, networkId)
	if nerr == nil && netCfg != nil {
		networkCfg, ok = netCfg.(*config_protos.Config)
		if !ok {
			nerr = fmt.Errorf(
				"received unexpected type for network record. Expected *Config but got %T", netCfg)
			networkCfg = nil
		}
	}
	// Check if there are GW configs
	cfg, err := config.GetConfig(networkId, FegGatewayType, gatewayId)
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return networkCfg, nerr
	}
	gatewayConfigs, ok := cfg.(*config_protos.Config)
	if !ok {
		return nil, fmt.Errorf(
			"received unexpected type for gateway record. Expected *Config but got %T", cfg)
	}
	// If there are network & gateway configs, blend them together with GW values overwriting network values
	if networkCfg != nil {
		protos.FillIn(gatewayConfigs, networkCfg)
		return networkCfg, nerr
	}
	return gatewayConfigs, nil
}
