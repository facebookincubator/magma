/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import (
	"magma/orc8r/cloud/go/models"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/orc8r"
	models2 "magma/orc8r/cloud/go/pluginimpl/models"
)

const (
	Networks                           = "networks"
	ListNetworksPath                   = obsidian.V1Root + Networks
	RegisterNetworkPath                = obsidian.V1Root + Networks
	ManageNetworkPath                  = obsidian.V1Root + Networks + obsidian.UrlSep + ":network_id"
	ManageNetworkNamePath              = ManageNetworkPath + obsidian.UrlSep + "name"
	ManageNetworkTypePath              = ManageNetworkPath + obsidian.UrlSep + "type"
	ManageNetworkDescriptionPath       = ManageNetworkPath + obsidian.UrlSep + "description"
	ManageNetworkFeaturesPath          = ManageNetworkPath + obsidian.UrlSep + "features"
	ManageNetworkDNSPath               = ManageNetworkPath + obsidian.UrlSep + "dns"
	ManageNetworkDNSRecordsPath        = ManageNetworkDNSPath + obsidian.UrlSep + "records"
	ManageNetworkDNSRecordByDomainPath = ManageNetworkDNSRecordsPath + obsidian.UrlSep + ":domain"

	Gateways                     = "gateways"
	ListGatewaysPath             = ManageNetworkPath + obsidian.UrlSep + Gateways
	ManageGatewayPath            = ListGatewaysPath + obsidian.UrlSep + ":gateway_id"
	ManageGatewayNamePath        = ManageGatewayPath + obsidian.UrlSep + "name"
	ManageGatewayDescriptionPath = ManageGatewayPath + obsidian.UrlSep + "description"
	ManageGatewayConfigPath      = ManageGatewayPath + obsidian.UrlSep + "magmad"
	ManageGatewayDevicePath      = ManageGatewayPath + obsidian.UrlSep + "device"
	ManageGatewayStatePath       = ManageGatewayPath + obsidian.UrlSep + "state"
	ManageGatewayTierPath        = ManageGatewayPath + obsidian.UrlSep + "tier"

	Channels          = "channels"
	ListChannelsPath  = obsidian.V1Root + Channels
	ManageChannelPath = obsidian.V1Root + Channels + obsidian.UrlSep + ":channel_id"
	Tiers             = "tiers"
	ListTiersPath     = ManageNetworkPath + obsidian.UrlSep + Tiers
	ManageTiersPath   = ListTiersPath + obsidian.UrlSep + ":tier_id"
)

// GetObsidianHandlers returns all plugin-level obsidian handlers for orc8r
func GetObsidianHandlers() []obsidian.Handler {
	ret := []obsidian.Handler{
		// Magma V1 Network
		{Path: ListNetworksPath, Methods: obsidian.GET, HandlerFunc: listNetworks},
		{Path: RegisterNetworkPath, Methods: obsidian.POST, HandlerFunc: registerNetwork},
		{Path: ManageNetworkPath, Methods: obsidian.GET, HandlerFunc: getNetwork},
		{Path: ManageNetworkPath, Methods: obsidian.PUT, HandlerFunc: updateNetwork},
		{Path: ManageNetworkPath, Methods: obsidian.DELETE, HandlerFunc: deleteNetwork},

		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.POST, HandlerFunc: CreateDNSRecord},
		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.GET, HandlerFunc: ReadDNSRecord},
		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.PUT, HandlerFunc: UpdateDNSRecord},
		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.DELETE, HandlerFunc: DeleteDNSRecord},

		// Magma V1 Gateways
		{Path: ListGatewaysPath, Methods: obsidian.GET, HandlerFunc: ListGatewaysHandler},
		{Path: ListGatewaysPath, Methods: obsidian.POST, HandlerFunc: CreateGatewayHandler},
		{Path: ManageGatewayPath, Methods: obsidian.GET, HandlerFunc: GetGatewayHandler},
		{Path: ManageGatewayPath, Methods: obsidian.PUT, HandlerFunc: UpdateGatewayHandler},
		{Path: ManageGatewayPath, Methods: obsidian.DELETE, HandlerFunc: DeleteGatewayHandler},
		{Path: ManageGatewayStatePath, Methods: obsidian.GET, HandlerFunc: GetStateHandler},

		// Upgrades
		{Path: ListChannelsPath, Methods: obsidian.GET, HandlerFunc: listChannelsHandler},
		{Path: ListChannelsPath, Methods: obsidian.POST, HandlerFunc: createChannelHandler},
		{Path: ManageChannelPath, Methods: obsidian.GET, HandlerFunc: readChannelHandler},
		{Path: ManageChannelPath, Methods: obsidian.PUT, HandlerFunc: updateChannelHandler},
		{Path: ManageChannelPath, Methods: obsidian.DELETE, HandlerFunc: deleteChannelHandler},
		{Path: ListTiersPath, Methods: obsidian.GET, HandlerFunc: listTiersHandler},
		{Path: ListTiersPath, Methods: obsidian.POST, HandlerFunc: createTierHandler},
		{Path: ManageTiersPath, Methods: obsidian.GET, HandlerFunc: readTierHandler},
		{Path: ManageTiersPath, Methods: obsidian.PUT, HandlerFunc: updateTierHandler},
		{Path: ManageTiersPath, Methods: obsidian.DELETE, HandlerFunc: deleteTierHandler},
	}
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkNamePath, new(models.NetworkName), "")...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkTypePath, new(models.NetworkType), "")...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkDescriptionPath, new(models.NetworkDescription), "")...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkFeaturesPath, &models2.NetworkFeatures{}, orc8r.NetworkFeaturesConfig)...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkDNSPath, &models2.NetworkDNSConfig{}, orc8r.DnsdNetworkType)...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkDNSRecordsPath, new(models2.NetworkDNSRecords), "")...)

	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayNamePath, new(models.GatewayName))...)
	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayDescriptionPath, new(models.GatewayDescription))...)
	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayConfigPath, &models2.MagmadGatewayConfigs{})...)
	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayTierPath, new(models2.TierID))...)
	ret = append(ret, GetGatewayDeviceHandlers(ManageGatewayDevicePath)...)
	return ret
}
