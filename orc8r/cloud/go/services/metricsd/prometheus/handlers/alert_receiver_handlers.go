/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"magma/orc8r/cloud/go/obsidian/handlers"
	"magma/orc8r/cloud/go/services/metricsd/prometheus/alerting/receivers"

	"github.com/labstack/echo"
	"github.com/prometheus/alertmanager/config"
)

func GetConfigureAlertReceiverHandler(webServerURL string) func(c echo.Context) error {
	return func(c echo.Context) error {
		networkID, nerr := handlers.GetNetworkId(c)
		if nerr != nil {
			return nerr
		}
		url := makeNetworkReceiverPath(webServerURL, networkID)
		return configureAlertReceiver(c, url)
	}
}

func GetRetrieveAlertReceiverHandler(webServerURL string) func(c echo.Context) error {
	return func(c echo.Context) error {
		networkID, nerr := handlers.GetNetworkId(c)
		if nerr != nil {
			return nerr
		}
		url := makeNetworkReceiverPath(webServerURL, networkID)
		return retrieveAlertReceivers(c, url)
	}
}

func GetRetrieveAlertRouteHandler(webServerURL string) func(c echo.Context) error {
	return func(c echo.Context) error {
		networkID, nerr := handlers.GetNetworkId(c)
		if nerr != nil {
			return nerr
		}
		url := makeNetworkRoutePath(webServerURL, networkID)
		return retrieveAlertRoute(c, url)
	}
}

func GetUpdateAlertRouteHandler(webServerURL string) func(c echo.Context) error {
	return func(c echo.Context) error {
		networkID, nerr := handlers.GetNetworkId(c)
		if nerr != nil {
			return nerr
		}
		url := makeNetworkRoutePath(webServerURL, networkID)
		return updateAlertRoute(c, url)
	}
}

func configureAlertReceiver(c echo.Context, url string) error {
	receiver, err := buildReceiverFromContext(c)
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}

	err = sendConfig(receiver, url)
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func retrieveAlertReceivers(c echo.Context, url string) error {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body echo.HTTPError
		_ = json.NewDecoder(resp.Body).Decode(&body)
		return handlers.HttpError(fmt.Errorf("error reading receivers: %v", body.Message), resp.StatusCode)
	}
	var recs []receivers.Receiver
	err = json.NewDecoder(resp.Body).Decode(&recs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("error decoding server response %v", err))
	}
	return c.JSON(http.StatusOK, recs)
}

func retrieveAlertRoute(c echo.Context, url string) error {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body echo.HTTPError
		_ = json.NewDecoder(resp.Body).Decode(&body)
		return handlers.HttpError(fmt.Errorf("error reading alerting route: %v", body.Message), resp.StatusCode)
	}
	var route config.Route
	err = json.NewDecoder(resp.Body).Decode(&route)
	if err != nil {
		return handlers.HttpError(fmt.Errorf("error decoding server response %v", err), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, route)
}

func updateAlertRoute(c echo.Context, url string) error {
	route, err := buildRouteFromContext(c)
	if err != nil {
		return handlers.HttpError(fmt.Errorf("invalid route specification: %v\n", err), http.StatusBadRequest)
	}

	err = sendConfig(route, url)
	if err != nil {
		return handlers.HttpError(fmt.Errorf("error updating alert route: %v", err), http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func buildReceiverFromContext(c echo.Context) (receivers.Receiver, error) {
	wrapper := receivers.Receiver{}
	err := json.NewDecoder(c.Request().Body).Decode(&wrapper)
	if err != nil {
		return receivers.Receiver{}, err
	}
	return wrapper, nil
}

func buildRouteFromContext(c echo.Context) (config.Route, error) {
	jsonRoute := receivers.RouteJSONWrapper{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonRoute)
	if err != nil {
		return config.Route{}, err
	}
	return jsonRoute.ToPrometheusConfig()
}

func makeNetworkReceiverPath(webServerURL, networkID string) string {
	return webServerURL + "/" + networkID + "/receiver"
}

func makeNetworkRoutePath(webSeverURL, networkID string) string {
	return webSeverURL + "/" + networkID + "/receiver/route"
}
