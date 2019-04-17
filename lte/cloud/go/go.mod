// Copyright (c) Facebook, Inc. and its affiliates.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
module magma/lte/cloud/go

replace (
	magma/feg/cloud/go/protos => ../../../feg/cloud/go/protos
	magma/orc8r/cloud/go => ../../../orc8r/cloud/go
)

require (
	github.com/aws/aws-sdk-go v1.16.19
	github.com/fiorix/go-diameter v3.0.2+incompatible
	github.com/go-openapi/errors v0.18.0
	github.com/go-openapi/strfmt v0.18.0
	github.com/go-openapi/swag v0.18.0
	github.com/go-openapi/validate v0.18.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/labstack/echo v0.0.0-20181123063414-c54d9e8eed6c
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc v1.19.1

	magma/feg/cloud/go/protos v0.0.0
	magma/orc8r/cloud/go v0.0.0
)
