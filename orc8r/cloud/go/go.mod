// Copyright (c) Facebook, Inc. and its affiliates.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
module magma/orc8r/cloud/go

replace (
	magma/gateway => ../../gateway/go
	magma/orc8r/lib/go => ../../lib/go
	magma/orc8r/lib/go/protos => ../../lib/go/protos
)

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3
	github.com/Masterminds/squirrel v1.1.1-0.20190513200039-d13326f0be73
	github.com/emakeev/snowflake v0.0.0-20200206205012-767080b052fe
	github.com/facebookincubator/ent v0.0.0-20191128071424-29c7b0a0d805
	github.com/go-openapi/errors v0.19.4
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.8
	github.com/go-sql-driver/mysql v1.4.1-0.20190510102335-877a9775f068
	github.com/go-swagger/go-swagger v0.24.0
	github.com/go-swagger/scan-repo-boundary v0.0.0-20180623220736-973b3573c013 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.0 // indirect
	github.com/labstack/echo v0.0.0-20181123063414-c54d9e8eed6c
	github.com/labstack/gommon v0.2.8 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0
	github.com/lib/pq v1.2.0
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/olivere/elastic/v7 v7.0.6
	github.com/pkg/errors v0.9.1
	github.com/prometheus/alertmanager v0.21.0
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.10.0
	github.com/prometheus/procfs v0.0.11
	github.com/prometheus/prometheus v0.0.0-20190115164134-b639fe140c1f
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/thoas/go-funk v0.4.0
	github.com/toqueteos/webbrowser v1.1.0 // indirect
	github.com/vektra/mockery v0.0.0-20181123154057-e78b021dcbb5
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120
	golang.org/x/tools v0.0.0-20200513201620-d5fe73897c97
	google.golang.org/grpc v1.27.1
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
	gopkg.in/yaml.v2 v2.3.0
	magma/gateway v0.0.0
	magma/orc8r/lib/go v0.0.0-00010101000000-000000000000
	magma/orc8r/lib/go/protos v0.0.0
)

go 1.12
