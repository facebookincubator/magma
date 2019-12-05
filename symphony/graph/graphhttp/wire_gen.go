// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package graphhttp

import (
	"github.com/facebookincubator/symphony/cloud/log"
	"github.com/facebookincubator/symphony/cloud/oc"
	"github.com/facebookincubator/symphony/cloud/orc8r"
	"github.com/facebookincubator/symphony/cloud/server"
	"github.com/facebookincubator/symphony/cloud/server/xserver"
	"github.com/facebookincubator/symphony/graph/viewer"
	"gocloud.dev/server/health"
	"net/http"
)

// Injectors from wire.go:

func NewServer(cfg Config) (*server.Server, func(), error) {
	mySQLTenancy := cfg.Tenancy
	logger := cfg.Logger
	config := cfg.Orc8r
	client := newOrc8rClient(config)
	router, err := newRouter(mySQLTenancy, logger, client)
	if err != nil {
		return nil, nil, err
	}
	zapLogger := xserver.NewRequestLogger(logger)
	v := newHealthChecker(mySQLTenancy)
	v2 := xserver.DefaultViews()
	exporter, err := xserver.NewPrometheusExporter(logger)
	if err != nil {
		return nil, nil, err
	}
	options := cfg.Census
	jaegerOptions := oc.JaegerOptions(options)
	traceExporter, cleanup, err := xserver.NewJaegerExporter(logger, jaegerOptions)
	if err != nil {
		return nil, nil, err
	}
	profilingEnabler := _wireProfilingEnablerValue
	sampler := oc.TraceSampler(options)
	handlerFunc := xserver.NewRecoveryHandler(logger)
	defaultDriver := _wireDefaultDriverValue
	serverOptions := &server.Options{
		RequestLogger:         zapLogger,
		HealthChecks:          v,
		Views:                 v2,
		ViewExporter:          exporter,
		TraceExporter:         traceExporter,
		EnableProfiling:       profilingEnabler,
		DefaultSamplingPolicy: sampler,
		RecoveryHandler:       handlerFunc,
		Driver:                defaultDriver,
	}
	serverServer := server.New(router, serverOptions)
	return serverServer, func() {
		cleanup()
	}, nil
}

var (
	_wireProfilingEnablerValue = server.ProfilingEnabler(true)
	_wireDefaultDriverValue    = &server.DefaultDriver{}
)

// wire.go:

// Config defines the http server config.
type Config struct {
	Tenancy *viewer.MySQLTenancy
	Logger  log.Logger
	Census  oc.Options
	Orc8r   orc8r.Config
}

func newHealthChecker(tenancy *viewer.MySQLTenancy) []health.Checker {
	return []health.Checker{tenancy}
}

func newOrc8rClient(config orc8r.Config) *http.Client {
	client, _ := orc8r.NewClient(config)
	return client
}
