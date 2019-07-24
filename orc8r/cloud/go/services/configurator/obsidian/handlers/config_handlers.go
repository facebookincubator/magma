/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package handlers

import (
	"net/http"

	"magma/orc8r/cloud/go/obsidian/handlers"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/services/configurator"

	"github.com/labstack/echo"
)

// GetCRUDNetworkConfigHandlers returns 4 Handlers which implement CRUD for
// a network config.
// Path should look like '.../networks/:network_id/{config_type}'
// ModelPtr is a pointer to the config structure to be stored
func GetCRUDNetworkConfigHandlers(path string, configType string, modelPtr serde.BinaryConvertible) []handlers.Handler {
	return []handlers.Handler{
		GetCreateNetworkConfigHandler(path, configType, modelPtr),
		GetReadNetworkConfigHandler(path, configType),
		GetUpdateNetworkConfigHandler(path, configType, modelPtr),
		GetDeleteNetworkConfigHandler(path, configType),
	}
}

func GetReadNetworkConfigHandler(path string, configType string) handlers.Handler {
	return handlers.Handler{
		Path:    path,
		Methods: handlers.GET,
		HandlerFunc: func(c echo.Context) error {
			networkID, nerr := handlers.GetNetworkId(c)
			if nerr != nil {
				return nerr
			}
			iConfig, err := configurator.GetNetworkConfigsByType(networkID, configType)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, iConfig)
		},
	}
}

func GetCreateNetworkConfigHandler(path string, configType string, modelPtr serde.BinaryConvertible) handlers.Handler {
	return handlers.Handler{
		Path:    path,
		Methods: handlers.POST,
		HandlerFunc: func(c echo.Context) error {
			return createOrUpdateNetworkConfig(c, configType, modelPtr)
		},
	}
}

func GetUpdateNetworkConfigHandler(path string, configType string, modelPtr serde.BinaryConvertible) handlers.Handler {
	return handlers.Handler{
		Path:    path,
		Methods: handlers.PUT,
		HandlerFunc: func(c echo.Context) error {
			return createOrUpdateNetworkConfig(c, configType, modelPtr)
		},
	}
}

func GetDeleteNetworkConfigHandler(path string, configType string) handlers.Handler {
	return handlers.Handler{
		Path:    path,
		Methods: handlers.DELETE,
		HandlerFunc: func(c echo.Context) error {
			networkID, nerr := handlers.GetNetworkId(c)
			if nerr != nil {
				return nerr
			}
			err := configurator.DeleteNetworkConfig(networkID, configType)
			if err != nil {
				return handlers.HttpError(err, http.StatusBadRequest)
			}
			return c.NoContent(http.StatusNoContent)
		},
	}
}

func createOrUpdateNetworkConfig(c echo.Context, configType string, modelPtr serde.BinaryConvertible) error {
	networkID, nerr := handlers.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	config, err := getConfigFromContext(c, modelPtr)
	if err != nil {
		return handlers.HttpError(err, http.StatusBadRequest)
	}

	updateCriteria := configurator.NetworkUpdateCriteria{
		ID:                   networkID,
		ConfigsToAddOrUpdate: map[string]interface{}{configType: config},
	}
	err = configurator.UpdateNetworks([]configurator.NetworkUpdateCriteria{updateCriteria})
	if err != nil {
		return handlers.HttpError(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, networkID)
}

func getConfigFromContext(c echo.Context, modelPtr serde.BinaryConvertible) (interface{}, error) {
	err := c.Bind(modelPtr)
	if err != nil {
		return nil, handlers.HttpError(err, http.StatusBadRequest)
	}
	return modelPtr, nil
}
