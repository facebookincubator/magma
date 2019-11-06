// Copyright (c) Facebook, Inc. and its affiliates.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
module magma/lte/cloud/go

replace magma/orc8r/cloud/go => ../../../orc8r/cloud/go

require (
	github.com/Masterminds/squirrel v1.1.1-0.20190513200039-d13326f0be73
	github.com/aws/aws-sdk-go v1.19.6
	github.com/go-openapi/errors v0.18.0
	github.com/go-openapi/strfmt v0.18.0
	github.com/go-openapi/swag v0.18.0
	github.com/go-openapi/validate v0.18.0
	github.com/go-sql-driver/mysql v1.4.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.2.0
	github.com/google/uuid v1.1.0
	github.com/labstack/echo v0.0.0-20181123063414-c54d9e8eed6c
	github.com/lib/pq v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc v1.19.0

	magma/orc8r/cloud/go v0.0.0
)

go 1.13
