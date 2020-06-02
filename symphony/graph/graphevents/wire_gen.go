// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package graphevents

import (
	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/pubsub"
	"github.com/facebookincubator/symphony/pkg/viewer"
)

// Injectors from wire.go:

func NewServer(cfg Config) (*Server, func(), error) {
	tenancy := cfg.Tenancy
	logger := cfg.Logger
	subscriber := cfg.Subscriber
	grapheventsServerConfig, err := newServerConfig(tenancy, logger, subscriber)
	if err != nil {
		return nil, nil, err
	}
	server, cleanup, err := newServer(grapheventsServerConfig)
	if err != nil {
		return nil, nil, err
	}
	return server, func() {
		cleanup()
	}, nil
}

// wire.go:

// Config defines the events server config.
type Config struct {
	Tenancy    viewer.Tenancy
	Subscriber pubsub.Subscriber
	Logger     log.Logger
}

func newServerConfig(tenancy viewer.Tenancy, logger log.Logger, subscriber pubsub.Subscriber) (cfg serverConfig, err error) {
	cfg = serverConfig{
		tenancy:    tenancy,
		logger:     logger,
		subscriber: subscriber,
	}
	return cfg, nil
}
