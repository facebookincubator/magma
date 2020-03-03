/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package config

import (
	"log"
	"time"

	"magma/orc8r/lib/go/definitions"
	"magma/orc8r/lib/go/service/config"
)

const (
	CONTROL_PROXY_YML_FRESHNESS_CHECK_INTERVAL = time.Minute * 10

	// Defaults
	DefaultNghttpxConfigLocation = "/var/tmp/nghttpx.conf"
	DefaultRootCaFile            = "/var/opt/magma/certs/rootCA.pem"
	DefaultGwCertFile            = "/var/opt/magma/certs/gateway.crt"
	DefaultGwCertKeyFile         = "/var/opt/magma/certs/gateway.key"
	DefaultCloudPort             = 443
	DefaultBootstrapPort         = 443
	DefaultLocalPort             = 8443
	DefaultProxyCloudConnection  = true
)

// ControlProxyCfg represents control_proxy.yml configuration struct
type ControlProxyCfg struct {
	NghttpxConfigLocation string `yaml:"nghttpx_config_location"`

	// GW Certificate params
	RootCaFile    string `yaml:"rootca_cert"`
	GwCertFile    string `yaml:"gateway_cert"`
	GwCertKeyFile string `yaml:"gateway_key"`

	LocalPort            int    `yaml:"local_port"`
	CloudAddr            string `yaml:"cloud_address"`
	CloudPort            int    `yaml:"cloud_port"`
	BootstrapAddr        string `yaml:"bootstrap_address"`
	BootstrapPort        int    `yaml:"bootstrap_port"`
	ProxyCloudConnection bool   `yaml:"proxy_cloud_connections"`
}

// NewDefaultControlProxyCfg returns new Bootstrapper struct with default configuration
func NewDefaultControlProxyCfg() *ControlProxyCfg {
	return &ControlProxyCfg{
		NghttpxConfigLocation: DefaultNghttpxConfigLocation,
		RootCaFile:            DefaultRootCaFile,
		GwCertFile:            DefaultGwCertFile,
		GwCertKeyFile:         DefaultGwCertKeyFile,
		LocalPort:             DefaultLocalPort,
		CloudAddr:             "",
		CloudPort:             DefaultCloudPort,
		BootstrapAddr:         "",
		BootstrapPort:         DefaultBootstrapPort,
		ProxyCloudConnection:  DefaultProxyCloudConnection,
	}
}

// UpdateFromYml of StructuredConfign interface - updates given control proxy config struct from corresponding YML file
// returns updated ControlProxyCfg, main YML CFG file path & owerwrite YML CFG file path (if any)
func (cpc *ControlProxyCfg) UpdateFromYml() (StructuredConfig, string, string) {
	var newCfg *ControlProxyCfg
	if cpc != nil {
		newCfg = &ControlProxyCfg{}
		*newCfg = *cpc // copy current configs
	} else {
		newCfg = NewDefaultControlProxyCfg()
		cpc = newCfg
	}
	ymlFile, ymlOWFile, err := config.GetStructuredServiceConfig("", definitions.ControlProxyServiceName, newCfg)
	if err != nil {
		log.Printf("Error Getting Control Proxy Configs: %v,\n\tcontinue using old configs: %+v", err, cpc)
	} else {
		if cpc != newCfg { // success, copy if needed
			*cpc = *newCfg
		}
	}
	return cpc, ymlFile, ymlOWFile
}

// FreshnessCheckInterval of StructuredConfig interface
func (_ *ControlProxyCfg) FreshnessCheckInterval() time.Duration {
	return CONTROL_PROXY_YML_FRESHNESS_CHECK_INTERVAL
}

var controlProxyConfigs AtomicStore

func controlProxyCfgFactory() StructuredConfig {
	return NewDefaultControlProxyCfg()
}

// GetControlProxyConfigs returns current magmad configuration
func GetControlProxyConfigs() *ControlProxyCfg {
	return controlProxyConfigs.GetCurrent(controlProxyCfgFactory).(*ControlProxyCfg)
}

// OverwriteControlProxyConfigs overwrites current control proxy configs
func OverwriteControlProxyConfigs(cfg *ControlProxyCfg) {
	controlProxyConfigs.Overwrite(cfg)
}
