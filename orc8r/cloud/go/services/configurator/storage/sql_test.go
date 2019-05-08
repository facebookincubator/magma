/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package storage_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"magma/orc8r/cloud/go/services/configurator/storage"
	storage2 "magma/orc8r/cloud/go/storage"

	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mockResult = sqlmock.NewResult(1, 1)

func TestSqlConfiguratorStorage_LoadNetworks(t *testing.T) {
	runFactory := func(ids []string, loadCriteria storage.NetworkLoadCriteria) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.LoadNetworks(ids, loadCriteria)
		}
	}

	idsOnly := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT cfg_networks.id, cfg_networks.version FROM cfg_networks").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "version"}).
						AddRow("hello", 1).
						AddRow("world", 2),
				)
		},
		run: runFactory([]string{"hello", "world"}, storage.NetworkLoadCriteria{}),

		expectedResult: storage.NetworkLoadResult{
			NetworkIDsNotFound: []string{},
			Networks: []storage.Network{
				{ID: "hello", Configs: map[string][]byte{}, Version: 1},
				{ID: "world", Configs: map[string][]byte{}, Version: 2},
			},
		},
	}

	metadataLoad := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT cfg_networks.id, cfg_networks.name, cfg_networks.description, cfg_networks.version FROM cfg_networks").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "description", "version"}).
						AddRow("hello", "Hello", "Hello network", 1).
						AddRow("world", "World", "World network", 2),
				)
		},
		run: runFactory([]string{"hello", "world"}, storage.NetworkLoadCriteria{LoadMetadata: true}),

		expectedResult: storage.NetworkLoadResult{
			NetworkIDsNotFound: []string{},
			Networks: []storage.Network{
				{
					ID:          "hello",
					Name:        "Hello",
					Description: "Hello network",
					Configs:     map[string][]byte{},
					Version:     1,
				},
				{
					ID:          "world",
					Name:        "World",
					Description: "World network",
					Configs:     map[string][]byte{},
					Version:     2,
				},
			},
		},
	}

	// 1 network has 2 configs, 1 has no configs
	idsAndConfigs := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT cfg_networks.id, cfg_network_configs.type, cfg_network_configs.value, cfg_networks.version FROM cfg_networks").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "type", "value", "version"}).
						AddRow("hello", "foo", []byte("foo"), 1).
						AddRow("hello", "bar", []byte("bar"), 1).
						AddRow("world", nil, nil, 3),
				)
		},
		run: runFactory([]string{"hello", "world", "foobar"}, storage.NetworkLoadCriteria{LoadConfigs: true}),

		expectedResult: storage.NetworkLoadResult{
			NetworkIDsNotFound: []string{"foobar"},
			Networks: []storage.Network{
				{
					ID: "hello",
					Configs: map[string][]byte{
						"foo": []byte("foo"),
						"bar": []byte("bar"),
					},
					Version: 1,
				},
				{
					ID:      "world",
					Configs: map[string][]byte{},
					Version: 3,
				},
			},
		},
	}

	// Same setup as above, plus 1 network not found
	fullLoad := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT cfg_networks.id, cfg_networks.name, cfg_networks.description, cfg_network_configs.type, cfg_network_configs.value, cfg_networks.version FROM cfg_networks").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "description", "type", "value", "version"}).
						AddRow("hello", "Hello", "Hello network", "foo", []byte("foo"), 1).
						AddRow("hello", "Hello", "Hello network", "bar", []byte("bar"), 1).
						AddRow("world", "World", "World network", nil, nil, 2),
				)
		},
		run: runFactory([]string{"hello", "world", "foobar"}, storage.NetworkLoadCriteria{LoadMetadata: true, LoadConfigs: true}),

		expectedResult: storage.NetworkLoadResult{
			NetworkIDsNotFound: []string{"foobar"},
			Networks: []storage.Network{
				{
					ID:          "hello",
					Name:        "Hello",
					Description: "Hello network",
					Configs: map[string][]byte{
						"foo": []byte("foo"),
						"bar": []byte("bar"),
					},
					Version: 1,
				},
				{
					ID:          "world",
					Name:        "World",
					Description: "World network",
					Configs:     map[string][]byte{},
					Version:     2,
				},
			},
		},
	}

	noneFound := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT cfg_networks.id, cfg_networks.version FROM cfg_networks").
				WillReturnRows(sqlmock.NewRows([]string{"id", "version"}))
		},
		run: runFactory([]string{"hello", "world"}, storage.NetworkLoadCriteria{}),

		expectedResult: storage.NetworkLoadResult{
			NetworkIDsNotFound: []string{"hello", "world"},
			Networks:           []storage.Network{},
		},
	}

	queryError := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT cfg_networks.id, cfg_networks.version FROM cfg_networks").
				WillReturnError(errors.New("mock query error"))
		},
		run: runFactory([]string{"hello", "world"}, storage.NetworkLoadCriteria{}),

		expectedError: errors.New("error querying for networks: mock query error"),
	}

	runCase(t, idsOnly)
	runCase(t, metadataLoad)
	runCase(t, idsAndConfigs)
	runCase(t, fullLoad)
	runCase(t, noneFound)
	runCase(t, queryError)
}

func TestSqlConfiguratorStorage_CreateNetwork(t *testing.T) {
	runFactory := func(network storage.Network) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.CreateNetwork(network)
		}
	}

	idOnly := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_networks`).
				WithArgs("n1").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			m.ExpectExec("INSERT INTO cfg_networks").
				WithArgs("n1", "", "").
				WillReturnResult(mockResult)
		},
		run: runFactory(storage.Network{ID: "n1"}),

		expectedResult: storage.Network{ID: "n1"},
	}

	allMetadata := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_networks`).
				WithArgs("n2").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			m.ExpectExec("INSERT INTO cfg_networks").
				WithArgs("n2", "hello", "world").
				WillReturnResult(mockResult)
		},
		run: runFactory(storage.Network{ID: "n2", Name: "hello", Description: "world"}),

		expectedResult: storage.Network{ID: "n2", Name: "hello", Description: "world"},
	}

	everythingNw := storage.Network{
		ID:          "n3",
		Name:        "hello",
		Description: "world",
		Configs: map[string][]byte{
			"foo": []byte("bar"),
			"baz": []byte("quz"),
		},
	}
	everything := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_networks`).
				WithArgs("n3").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			m.ExpectExec("INSERT INTO cfg_networks").
				WithArgs("n3", "hello", "world").
				WillReturnResult(mockResult)

			configStmt := m.ExpectPrepare("INSERT INTO cfg_network_configs")
			configStmt.ExpectExec().WithArgs("n3", "baz", []byte("quz")).WillReturnResult(mockResult)
			configStmt.ExpectExec().WithArgs("n3", "foo", []byte("bar")).WillReturnResult(mockResult)
			configStmt.WillBeClosed()
		},
		run: runFactory(everythingNw),

		expectedResult: everythingNw,
	}

	networkTableError := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_networks`).
				WithArgs("n4").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			m.ExpectExec("INSERT INTO cfg_networks").
				WithArgs("n4", "", "").
				WillReturnError(errors.New("mock exec error"))
		},
		run: runFactory(storage.Network{ID: "n4"}),

		expectedResult: storage.Network{ID: "n4"},
		expectedError:  errors.New("error inserting network: mock exec error"),
	}

	configTableErrNw := storage.Network{
		ID: "n5",
		Configs: map[string][]byte{
			"foo": []byte("bar"),
			"baz": []byte("quz"),
		},
	}
	configTableError := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_networks`).
				WithArgs("n5").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			m.ExpectExec("INSERT INTO cfg_networks").
				WithArgs("n5", "", "").
				WillReturnResult(mockResult)

			configStmt := m.ExpectPrepare("INSERT INTO cfg_network_configs")
			configStmt.ExpectExec().WithArgs("n5", "baz", []byte("quz")).
				WillReturnError(errors.New("mock exec error"))
			configStmt.WillBeClosed()
		},
		run: runFactory(configTableErrNw),

		expectedResult: configTableErrNw,
		expectedError:  errors.New("error inserting config baz: mock exec error"),
	}

	networkExists := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_networks`).
				WithArgs("n5").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
		},
		run: runFactory(storage.Network{ID: "n5"}),

		expectedResult: storage.Network{ID: "n5"},
		expectedError:  errors.New("a network with ID n5 already exists"),
	}

	runCase(t, idOnly)
	runCase(t, allMetadata)
	runCase(t, everything)
	runCase(t, networkTableError)
	runCase(t, configTableError)
	runCase(t, networkExists)
}

func TestSqlConfiguratorStorage_UpdateNetworks(t *testing.T) {
	runFactory := func(updates []storage.NetworkUpdateCriteria) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.UpdateNetworks(updates)
		}
	}

	// Delete 1 network (n1)
	// Update only metadata of another (n2)
	// Update and delete configs of another (n3)
	// Fill out all fields of the update criteria (n4)
	names := []string{"should be ignored", "name2", "name4"}
	descs := []string{"should be ignored", "desc2", ""}
	happyPath := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			deleteStmt := m.ExpectPrepare("DELETE FROM cfg_networks")
			upsertStmt := m.ExpectPrepare("INSERT INTO cfg_network_configs")
			deleteConfigStmt := m.ExpectPrepare("DELETE FROM cfg_network_configs")

			deleteStmt.ExpectExec().WithArgs("n1").WillReturnResult(mockResult)

			m.ExpectExec("UPDATE cfg_networks").WithArgs(names[1], descs[1], "n2").WillReturnResult(mockResult)

			m.ExpectExec("UPDATE cfg_networks").WithArgs("n3").WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n3", "baz", []byte("quz"), []byte("quz")).WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n3", "foo", []byte("bar"), []byte("bar")).WillReturnResult(mockResult)
			deleteConfigStmt.ExpectExec().WithArgs("n3", "hello").WillReturnResult(mockResult)
			deleteConfigStmt.ExpectExec().WithArgs("n3", "world").WillReturnResult(mockResult)

			m.ExpectExec("UPDATE cfg_networks").WithArgs(names[2], nil, "n4").WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n4", "baz", []byte("quz"), []byte("quz")).WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n4", "foo", []byte("bar"), []byte("bar")).WillReturnResult(mockResult)
			deleteConfigStmt.ExpectExec().WithArgs("n4", "hello").WillReturnResult(mockResult)
			deleteConfigStmt.ExpectExec().WithArgs("n4", "world").WillReturnResult(mockResult)

			deleteStmt.WillBeClosed()
			upsertStmt.WillBeClosed()
			deleteConfigStmt.WillBeClosed()
		},
		run: runFactory(
			[]storage.NetworkUpdateCriteria{
				{ID: "n1", DeleteNetwork: true, NewName: &names[0], NewDescription: &descs[0]},
				{ID: "n2", NewName: &names[1], NewDescription: &descs[1]},
				{
					ID:              "n3",
					ConfigsToDelete: []string{"hello", "world"},
					ConfigsToAddOrUpdate: map[string][]byte{
						"foo": []byte("bar"),
						"baz": []byte("quz"),
					},
				},
				{
					ID:              "n4",
					NewName:         &names[2],
					NewDescription:  &descs[2],
					ConfigsToDelete: []string{"hello", "world"},
					ConfigsToAddOrUpdate: map[string][]byte{
						"foo": []byte("bar"),
						"baz": []byte("quz"),
					},
				},
			},
		),

		expectedResult: storage.FailedOperations{},
	}

	// Error in 1 network should not block other networks (try with 3 networks/2 errors)
	errorCase := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			deleteStmt := m.ExpectPrepare("DELETE FROM cfg_networks")
			upsertStmt := m.ExpectPrepare("INSERT INTO cfg_network_configs")
			deleteConfigStmt := m.ExpectPrepare("DELETE FROM cfg_network_configs")

			deleteStmt.ExpectExec().WithArgs("n1").WillReturnError(errors.New("n1 delete error"))

			m.ExpectExec("UPDATE cfg_networks").WithArgs(names[1], descs[1], "n2").WillReturnError(errors.New("n2 update error"))

			m.ExpectExec("UPDATE cfg_networks").WithArgs("n3").WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n3", "baz", []byte("quz"), []byte("quz")).WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n3", "foo", []byte("bar"), []byte("bar")).WillReturnError(errors.New("n3foo update error"))

			m.ExpectExec("UPDATE cfg_networks").WithArgs(names[2], nil, "n4").WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n4", "baz", []byte("quz"), []byte("quz")).WillReturnResult(mockResult)
			upsertStmt.ExpectExec().WithArgs("n4", "foo", []byte("bar"), []byte("bar")).WillReturnResult(mockResult)
			deleteConfigStmt.ExpectExec().WithArgs("n4", "hello").WillReturnResult(mockResult)
			deleteConfigStmt.ExpectExec().WithArgs("n4", "world").WillReturnResult(mockResult)

			deleteStmt.WillBeClosed()
			upsertStmt.WillBeClosed()
			deleteConfigStmt.WillBeClosed()
		},
		run: runFactory(
			[]storage.NetworkUpdateCriteria{
				{ID: "n1", DeleteNetwork: true},
				{ID: "n2", NewName: &names[1], NewDescription: &descs[1]},
				{
					ID:              "n3",
					ConfigsToDelete: []string{"hello", "world"},
					ConfigsToAddOrUpdate: map[string][]byte{
						"foo": []byte("bar"),
						"baz": []byte("quz"),
					},
				},
				{
					ID:              "n4",
					NewName:         &names[2],
					NewDescription:  &descs[2],
					ConfigsToDelete: []string{"hello", "world"},
					ConfigsToAddOrUpdate: map[string][]byte{
						"foo": []byte("bar"),
						"baz": []byte("quz"),
					},
				},
			},
		),

		expectedResult: storage.FailedOperations{
			"n1": errors.New("error deleting network n1: n1 delete error"),
			"n2": errors.New("error updating network n2: n2 update error"),
			"n3": errors.New("error updating config foo on network n3: n3foo update error"),
		},
		expectedError: errors.New("some errors were encountered, see return value for details"),
	}

	validationFailure := &testCase{
		setup: func(m sqlmock.Sqlmock) {},

		run: runFactory(
			[]storage.NetworkUpdateCriteria{
				{ID: "n1", DeleteNetwork: true},
				{ID: "n1", NewName: &names[1]},
			},
		),

		expectedError: errors.New("multiple updates for a single network are not allowed"),
	}

	runCase(t, happyPath)
	runCase(t, errorCase)
	runCase(t, validationFailure)
}

func TestSqlConfiguratorStorage_LoadEntities(t *testing.T) {
	runFactory := func(networkID string, filter storage.EntityLoadFilter, loadCriteria storage.EntityLoadCriteria) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.LoadEntities(networkID, filter, loadCriteria)
		}
	}

	// Empty load criteria
	basicOnly := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			entStmt := m.ExpectPrepare("SELECT ent.pk, ent.key, ent.type, ent.physical_id, ent.version, ent.graph_id FROM cfg_entities")
			entStmt.ExpectQuery().
				WithArgs("network", "bar", "foo").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("abc", "bar", "foo", nil, 2, "42"),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "quz", "baz").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("def", "quz", "baz", nil, 1, "42"),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "world", "hello").
				WillReturnRows(sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}))
			entStmt.WillBeClosed()
		},
		run: runFactory(
			"network",
			storage.EntityLoadFilter{
				IDs: []storage2.TypeAndKey{
					{Type: "foo", Key: "bar"},
					{Type: "baz", Key: "quz"},
					{Type: "hello", Key: "world"},
				},
			},
			storage.EntityLoadCriteria{},
		),

		expectedResult: storage.EntityLoadResult{
			Entities: []storage.NetworkEntity{
				{Type: "baz", Key: "quz", GraphID: "42", Version: 1},
				{Type: "foo", Key: "bar", GraphID: "42", Version: 2},
			},
			EntitiesNotFound: []storage2.TypeAndKey{{Type: "hello", Key: "world"}},
		},
	}

	// Load everything, no assocs
	// foobar has 2 permissions, bazquz has 1 wildcarded permission
	loadEverything := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			entStmt := m.ExpectPrepare("SELECT .* FROM cfg_entities")
			entStmt.ExpectQuery().
				WithArgs("network", "bar", "foo").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id", "name", "description", "config", "id", "scope", "permission", "type", "id_filter", "acl.version"}).
						AddRow("abc", "bar", "foo", nil, 2, "42", "foobar", "foobar ent", []byte("foobar"), "foobar_acl_1", "n1,n2,n3", storage.OwnerPermission, "foo", nil, 1).
						AddRow("abc", "bar", "foo", nil, 2, "42", "foobar", "foobar ent", []byte("foobar"), "foobar_acl_2", "n4", storage.ReadPermission, "baz", nil, 2),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "quz", "baz").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id", "name", "description", "config", "id", "scope", "permission", "type", "id_filter", "acl.version"}).
						AddRow("def", "quz", "baz", nil, 1, "42", "bazquz", "bazquz ent", []byte("bazquz"), "bazquz_acl_1", "*", storage.WritePermission, "*", "1,2,3", 3),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "world", "hello").
				WillReturnRows(sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id", "name", "description", "config", "id", "scope", "permission", "type", "id_filter", "acl.version"}))
			entStmt.WillBeClosed()
		},
		run: runFactory(
			"network",
			storage.EntityLoadFilter{
				IDs: []storage2.TypeAndKey{
					{Type: "foo", Key: "bar"},
					{Type: "baz", Key: "quz"},
					{Type: "hello", Key: "world"},
				},
			},
			storage.EntityLoadCriteria{LoadMetadata: true, LoadConfig: true, LoadPermissions: true},
		),

		expectedResult: storage.EntityLoadResult{
			Entities: []storage.NetworkEntity{
				{
					Type: "baz", Key: "quz", GraphID: "42", Version: 1,
					Name:        "bazquz",
					Description: "bazquz ent",
					Config:      []byte("bazquz"),
					Permissions: []storage.ACL{
						{ID: "bazquz_acl_1", Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, Permission: storage.WritePermission, IDFilter: []string{"1", "2", "3"}, Version: 3},
					},
				},
				{
					Type: "foo", Key: "bar", GraphID: "42", Version: 2,
					Name:        "foobar",
					Description: "foobar ent",
					Config:      []byte("foobar"),
					Permissions: []storage.ACL{
						{ID: "foobar_acl_1", Type: storage.ACLType{EntityType: "foo"}, Scope: storage.ACLScope{NetworkIDs: []string{"n1", "n2", "n3"}}, Permission: storage.OwnerPermission, Version: 1},
						{ID: "foobar_acl_2", Type: storage.ACLType{EntityType: "baz"}, Scope: storage.ACLScope{NetworkIDs: []string{"n4"}}, Permission: storage.ReadPermission, Version: 2},
					},
				},
			},
			EntitiesNotFound: []storage2.TypeAndKey{{Type: "hello", Key: "world"}},
		},
	}

	// Load assocs to only
	assocsTo := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			entStmt := m.ExpectPrepare("SELECT ent.pk, ent.key, ent.type, ent.physical_id, ent.version, ent.graph_id FROM cfg_entities")
			entStmt.ExpectQuery().
				WithArgs("network", "bar", "foo").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("abc", "bar", "foo", nil, 2, "42"),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "quz", "baz").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("def", "quz", "baz", nil, 1, "42"),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "world", "hello").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("ghi", "world", "hello", nil, 3, "42"),
				)
			entStmt.WillBeClosed()

			expectAssocQuery(
				m, []driver.Value{"abc", "def", "ghi"},
				"abc", "def",
				"abc", "ghi",
				"ghi", "def",
			)
		},
		run: runFactory(
			"network",
			storage.EntityLoadFilter{
				IDs: []storage2.TypeAndKey{
					{Type: "foo", Key: "bar"},
					{Type: "baz", Key: "quz"},
					{Type: "hello", Key: "world"},
				},
			},
			storage.EntityLoadCriteria{LoadAssocsToThis: true},
		),

		expectedResult: storage.EntityLoadResult{
			Entities: []storage.NetworkEntity{
				{
					Type: "baz", Key: "quz", GraphID: "42", Version: 1,
					ParentAssociations: []storage2.TypeAndKey{
						{Type: "foo", Key: "bar"},
						{Type: "hello", Key: "world"},
					},
				},
				{Type: "foo", Key: "bar", GraphID: "42", Version: 2},
				{
					Type: "hello", Key: "world", GraphID: "42", Version: 3,
					ParentAssociations: []storage2.TypeAndKey{
						{Type: "foo", Key: "bar"},
					},
				},
			},
			EntitiesNotFound: []storage2.TypeAndKey{},
		},
	}

	// Load assocs from only
	assocsFrom := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			entStmt := m.ExpectPrepare("SELECT ent.pk, ent.key, ent.type, ent.physical_id, ent.version, ent.graph_id FROM cfg_entities")
			entStmt.ExpectQuery().
				WithArgs("network", "bar", "foo").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("abc", "bar", "foo", nil, 2, "42"),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "quz", "baz").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("def", "quz", "baz", nil, 1, "42"),
				)
			entStmt.ExpectQuery().
				WithArgs("network", "world", "hello").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("ghi", "world", "hello", nil, 3, "42"),
				)
			entStmt.WillBeClosed()

			expectAssocQuery(
				m,
				[]driver.Value{"abc", "def", "ghi"},
				"def", "abc",
				"ghi", "abc",
				"def", "ghi",
			)
		},
		run: runFactory(
			"network",
			storage.EntityLoadFilter{
				IDs: []storage2.TypeAndKey{
					{Type: "foo", Key: "bar"},
					{Type: "baz", Key: "quz"},
					{Type: "hello", Key: "world"},
				},
			},
			storage.EntityLoadCriteria{LoadAssocsFromThis: true},
		),

		expectedResult: storage.EntityLoadResult{
			Entities: []storage.NetworkEntity{
				{
					Type: "baz", Key: "quz", GraphID: "42", Version: 1,
					Associations: []storage2.TypeAndKey{
						{Type: "foo", Key: "bar"},
						{Type: "hello", Key: "world"},
					},
				},
				{Type: "foo", Key: "bar", GraphID: "42", Version: 2},
				{
					Type: "hello", Key: "world", GraphID: "42", Version: 3,
					Associations: []storage2.TypeAndKey{
						{Type: "foo", Key: "bar"},
					},
				},
			},
			EntitiesNotFound: []storage2.TypeAndKey{},
		},
	}

	// Load everything with type filter
	// (foo, bar) will have 2 permissions and 2 assocs - one to (foo, baz) and one to (bar, baz)
	// (foo, baz) will have 1 wildcard permission and 1 assoc to (baz, quz)
	// (hello, world) will have 1 assoc to (foo, bar)
	// We will only load entities of type foo
	fullLoadTypeFilter := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT .* FROM cfg_entities").
				WithArgs("network", "foo").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id", "name", "description", "config", "id", "scope", "permission", "type", "id_filter", "acl.version"}).
						// (foo, bar) comes from test case for full load above
						AddRow("foobar", "bar", "foo", nil, 1, "42", "foobar", "foobar ent", []byte("foobar"), "foobar_acl_1", "n1,n2,n3", storage.OwnerPermission, "foo", nil, 1).
						AddRow("foobar", "bar", "foo", nil, 1, "42", "foobar", "foobar ent", []byte("foobar"), "foobar_acl_2", "n4", storage.ReadPermission, "baz", nil, 2).
						AddRow("foobaz", "baz", "foo", nil, 2, "42", "foobaz", "foobaz ent", []byte("foobaz"), "foobaz_acl_1", "*", storage.WritePermission, "*", nil, 3),
				)

			expectAssocQuery(
				m,
				[]driver.Value{"foobar", "foobaz", "foobar", "foobaz"},
				"foobar", "foobaz",
				"foobar", "barbaz",
				"foobaz", "bazquz",
				"helloworld", "foobar",
			)

			// Since we don't query for (hello, world), we expect a query for its type and key given its PK
			m.ExpectQuery("SELECT pk, type, key FROM cfg_entities").
				WithArgs("barbaz", "bazquz", "helloworld").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "type", "key"}).
						AddRow("barbaz", "bar", "baz").
						AddRow("bazquz", "baz", "quz").
						AddRow("helloworld", "hello", "world"),
				)
		},
		run: runFactory(
			"network",
			storage.EntityLoadFilter{
				TypeFilter: stringPointer("foo"),
			},
			storage.FullEntityLoadCriteria,
		),

		expectedResult: storage.EntityLoadResult{
			Entities: []storage.NetworkEntity{
				{
					Type: "foo", Key: "bar", GraphID: "42", Version: 1,
					Name:        "foobar",
					Description: "foobar ent",
					Config:      []byte("foobar"),
					Permissions: []storage.ACL{
						{ID: "foobar_acl_1", Type: storage.ACLType{EntityType: "foo"}, Scope: storage.ACLScope{NetworkIDs: []string{"n1", "n2", "n3"}}, Permission: storage.OwnerPermission, Version: 1},
						{ID: "foobar_acl_2", Type: storage.ACLType{EntityType: "baz"}, Scope: storage.ACLScope{NetworkIDs: []string{"n4"}}, Permission: storage.ReadPermission, Version: 2},
					},
					Associations: []storage2.TypeAndKey{
						{Type: "bar", Key: "baz"},
						{Type: "foo", Key: "baz"},
					},
					ParentAssociations: []storage2.TypeAndKey{
						{Type: "hello", Key: "world"},
					},
				},
				{
					Type: "foo", Key: "baz", GraphID: "42", Version: 2,
					Name:        "foobaz",
					Description: "foobaz ent",
					Config:      []byte("foobaz"),
					Permissions: []storage.ACL{
						{ID: "foobaz_acl_1", Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, Permission: storage.WritePermission, Version: 3},
					},
					Associations: []storage2.TypeAndKey{
						{Type: "baz", Key: "quz"},
					},
					ParentAssociations: []storage2.TypeAndKey{
						{Type: "foo", Key: "bar"},
					},
				},
			},
			EntitiesNotFound: []storage2.TypeAndKey{},
		},
	}

	// Basic load with type and key filters
	typeAndKeyFilters := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery("SELECT ent.pk, ent.key, ent.type, ent.physical_id, ent.version, ent.graph_id FROM cfg_entities").
				WithArgs("network", "bar", "foo").
				WillReturnRows(
					sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"}).
						AddRow("abc", "bar", "foo", nil, 2, "42"),
				)
		},
		run: runFactory(
			"network",
			storage.EntityLoadFilter{
				TypeFilter: stringPointer("foo"),
				KeyFilter:  stringPointer("bar"),
			},
			storage.EntityLoadCriteria{},
		),

		expectedResult: storage.EntityLoadResult{
			Entities: []storage.NetworkEntity{
				{Type: "foo", Key: "bar", GraphID: "42", Version: 2},
			},
			EntitiesNotFound: []storage2.TypeAndKey{},
		},
	}

	runCase(t, basicOnly)
	runCase(t, loadEverything)
	runCase(t, assocsTo)
	runCase(t, assocsFrom)
	runCase(t, fullLoadTypeFilter)
	runCase(t, typeAndKeyFilters)
}

func TestSqlConfiguratorStorage_CreateEntity(t *testing.T) {
	runFactory := func(networkID string, entity storage.NetworkEntity) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.CreateEntity(networkID, entity)
		}
	}

	// basic fields
	basicCase := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_entities`).
				WithArgs("network", "foo", "bar").
				WillReturnRows(sqlmock.NewRows([]string{"count"}))

			insertStmt := m.ExpectPrepare("INSERT INTO cfg_entities").WillBeClosed()
			insertStmt.ExpectExec().
				WithArgs("1", "network", "foo", "bar", "2", "foobar", "foobar ent", nil, nil).
				WillReturnResult(mockResult)
		},
		run: runFactory(
			"network",
			storage.NetworkEntity{
				Type:        "foo",
				Key:         "bar",
				Name:        "foobar",
				Description: "foobar ent",
			},
		),

		expectedResult: storage.NetworkEntity{
			Type:        "foo",
			Key:         "bar",
			Name:        "foobar",
			Description: "foobar ent",
			GraphID:     "2",
		},
	}

	perms := []storage.ACL{
		{ID: "ignore this", Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, Permission: storage.WritePermission},
		{Type: storage.ACLTypeOf("foo"), Scope: storage.ACLScopeOf([]string{"n1", "n2"}), Permission: storage.ReadPermission, IDFilter: []string{"foo", "bar"}},
	}

	// with permissions
	withPerms := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_entities`).
				WithArgs("network", "foo", "bar").
				WillReturnRows(sqlmock.NewRows([]string{"count"}))

			insertStmt := m.ExpectPrepare("INSERT INTO cfg_entities").WillBeClosed()
			insertStmt.ExpectExec().
				WithArgs("1", "network", "foo", "bar", "2", "foobar", "foobar ent", nil, nil).
				WillReturnResult(mockResult)

			expectPermissionCreation(m, "1", 3, perms...)
		},
		run: runFactory(
			"network",
			storage.NetworkEntity{
				Type:        "foo",
				Key:         "bar",
				Name:        "foobar",
				Description: "foobar ent",
				Permissions: perms,
			},
		),

		expectedResult: storage.NetworkEntity{
			Type:        "foo",
			Key:         "bar",
			Name:        "foobar",
			Description: "foobar ent",
			GraphID:     "2",
			Permissions: []storage.ACL{
				{ID: "3", Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, Permission: storage.WritePermission},
				{ID: "4", Type: storage.ACLTypeOf("foo"), Scope: storage.ACLScopeOf([]string{"n1", "n2"}), Permission: storage.ReadPermission, IDFilter: []string{"foo", "bar"}},
			},
		},
	}

	// merge 2 graphs together
	mergeGraphs := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_entities`).
				WithArgs("network", "foo", "bar").
				WillReturnRows(sqlmock.NewRows([]string{"count"}))

			insertStmt := m.ExpectPrepare("INSERT INTO cfg_entities").WillBeClosed()
			insertStmt.ExpectExec().
				WithArgs("1", "network", "foo", "bar", "2", "foobar", "foobar ent", nil, nil).
				WillReturnResult(mockResult)

			assocs := []storage2.TypeAndKey{{Type: "bar", Key: "baz"}, {Type: "baz", Key: "quz"}}
			edgesByTk := map[storage2.TypeAndKey]expectedEntQueryResult{
				{Type: "bar", Key: "baz"}: {"bar", "baz", "42", "", "aaa", 1},
				{Type: "baz", Key: "quz"}: {"baz", "quz", "43", "", "zzz", 2},
			}
			expectEdgeQueries(m, assocs, edgesByTk)
			expectEdgeInsertions(m, assocsToEdges("1", assocs, edgesByTk))
			expectMergeGraphs(m, [][2]string{{"2", "aaa"}, {"zzz", "aaa"}})
		},
		run: runFactory(
			"network",
			storage.NetworkEntity{
				Type:        "foo",
				Key:         "bar",
				Name:        "foobar",
				Description: "foobar ent",
				Associations: []storage2.TypeAndKey{
					{Type: "bar", Key: "baz"},
					{Type: "baz", Key: "quz"},
					// Duplicate edge should only be loaded once
					{Type: "bar", Key: "baz"},
				},
			},
		),

		// We expect "aaa" as the returned graphID since we merged graphs
		expectedResult: storage.NetworkEntity{
			Type:        "foo",
			Key:         "bar",
			Name:        "foobar",
			Description: "foobar ent",
			GraphID:     "aaa",
			Associations: []storage2.TypeAndKey{
				{Type: "bar", Key: "baz"},
				{Type: "baz", Key: "quz"},
			},
		},
	}

	// already exists
	alreadyExists := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			m.ExpectQuery(`SELECT count\(1\) FROM cfg_entities`).
				WithArgs("network", "foo", "bar").
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
		},
		run: runFactory(
			"network",
			storage.NetworkEntity{
				Type:        "foo",
				Key:         "bar",
				Name:        "foobar",
				Description: "foobar ent",
			},
		),

		expectedResult: storage.NetworkEntity{},
		expectedError:  errors.New("an entity (foo-bar) already exists"),
	}

	runCase(t, basicCase)
	runCase(t, withPerms)
	runCase(t, mergeGraphs)
	runCase(t, alreadyExists)
}

func TestSqlConfiguratorStorage_UpdateEntity(t *testing.T) {
	runFactory := func(networkID string, update storage.EntityUpdateCriteria) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.UpdateEntity(networkID, update)
		}
	}

	// Delete entity, TODO: partition graph
	deleteCase := &testCase{
		setup: func(m sqlmock.Sqlmock) {
			expectBasicEntityQueries(m, expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0})
			m.ExpectExec("DELETE FROM cfg_entities").WithArgs("network", "foo", "bar").WillReturnResult(mockResult)
		},
		run: runFactory("network", storage.EntityUpdateCriteria{Type: "foo", Key: "bar", DeleteEntity: true}),

		expectedResult: storage.NetworkEntity{Type: "foo", Key: "bar"},
	}
	runCase(t, deleteCase)

	// Test some permutations of updating basic fields
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type:      "foo",
				Key:       "bar",
				NewName:   stringPointer("foobar"),
				NewConfig: bytesPointer([]byte("foobar config")),
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type:           "foo",
				Key:            "bar",
				NewDescription: stringPointer("foobar desc"),
				NewPhysicalID:  stringPointer("phys2"),
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type:           "foo",
				Key:            "bar",
				NewName:        stringPointer("foobar"),
				NewDescription: stringPointer("foobar desc"),
				NewPhysicalID:  stringPointer("phys2"),
				NewConfig:      bytesPointer([]byte("foobar config")),
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)

	// Test cases for permissions
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type: "foo",
				Key:  "bar",
				PermissionsToCreate: []storage.ACL{
					{Permission: storage.WritePermission, Type: storage.WildcardACLType, Scope: storage.ACLScopeOf([]string{"n3"})},
				},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)

	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type: "foo",
				Key:  "bar",
				PermissionsToUpdate: []storage.ACL{
					{ID: "42", Permission: storage.ReadPermission, Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, IDFilter: []string{"n1", "n2"}},
					{ID: "43", Permission: storage.WritePermission, Type: storage.ACLTypeOf("bar"), Scope: storage.ACLScopeOf([]string{"n3"})},
				},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)

	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type:                "foo",
				Key:                 "bar",
				PermissionsToDelete: []string{"100", "101"},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)

	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type: "foo",
				Key:  "bar",
				PermissionsToCreate: []storage.ACL{
					{ID: "ignore me", Permission: storage.WritePermission, Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, IDFilter: []string{"n1", "n2"}},
					{Permission: storage.ReadPermission, Type: storage.ACLTypeOf("foo"), Scope: storage.ACLScopeOf([]string{"n1", "n2"})},
				},
				PermissionsToUpdate: []storage.ACL{
					{ID: "42", Permission: storage.ReadPermission, Type: storage.WildcardACLType, Scope: storage.WildcardACLScope, IDFilter: []string{"n1", "n2"}},
					{ID: "43", Permission: storage.WritePermission, Type: storage.ACLTypeOf("bar"), Scope: storage.ACLScopeOf([]string{"n3"})},
				},
				PermissionsToDelete: []string{"42", "101"},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			nil,
		),
	)

	// edges

	// Create edges, merge graphs
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type: "foo",
				Key:  "bar",
				AssociationsToAdd: []storage2.TypeAndKey{
					{Type: "bar", Key: "baz"},
					{Type: "baz", Key: "quz"},
				},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g9", 0},
			[]expectedEntQueryResult{
				{"bar", "baz", "2", "", "g1", 0},
				{"baz", "quz", "3", "", "g2", 0},
			},
			[2]string{"g9", "g1"},
			[2]string{"g2", "g1"},
		),
	)

	// Create edge to something already in the same graph
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type: "foo",
				Key:  "bar",
				AssociationsToAdd: []storage2.TypeAndKey{
					{Type: "bar", Key: "baz"},
				},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			[]expectedEntQueryResult{
				{"bar", "baz", "2", "", "g1", 0},
			},
		),
	)

	// Delete edges, TODO: partition graph
	runCase(
		t,
		getTestCaseForEntityUpdate(
			storage.EntityUpdateCriteria{
				Type: "foo",
				Key:  "bar",
				AssociationsToDelete: []storage2.TypeAndKey{
					{Type: "bar", Key: "baz"},
				},
			},
			expectedEntQueryResult{"foo", "bar", "1", "", "g1", 0},
			[]expectedEntQueryResult{
				{"bar", "baz", "2", "", "g1", 0},
			},
		),
	)
}

func TestSqlConfiguratorStorage_LoadGraphForEntity(t *testing.T) {
	runFactory := func(networkID string, entityID storage2.TypeAndKey, loadCriteria storage.EntityLoadCriteria) func(store storage.ConfiguratorStorage) (interface{}, error) {
		return func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.LoadGraphForEntity(networkID, entityID, loadCriteria)
		}
	}

	expectedFooBar := expectedEntQueryResult{"foo", "bar", "foobar", "", "g1", 0}
	expectedBarBaz := expectedEntQueryResult{"bar", "baz", "barbaz", "p1", "g1", 1}
	expectedBazQuz := expectedEntQueryResult{"baz", "quz", "bazquz", "p2", "g1", 2}

	assocQueryArgs := []driver.Value{"barbaz", "bazquz", "foobar", "barbaz", "bazquz", "foobar"}

	// load a linked list of 3 nodes
	linkedList := &testCase{
		run: runFactory("network", storage2.TypeAndKey{Type: "foo", Key: "bar"}, storage.EntityLoadCriteria{}),

		setup: func(m sqlmock.Sqlmock) {
			expectBasicEntityQueries(m, expectedFooBar)
			expectBulkEntityQuery(m, []driver.Value{"g1"}, expectedFooBar, expectedBarBaz, expectedBazQuz)
			// foobar -> barbaz -> bazquz
			expectAssocQuery(m, assocQueryArgs, "foobar", "barbaz", "barbaz", "bazquz")
		},

		expectedResult: storage.EntityGraph{
			Entities: []storage.NetworkEntity{
				{
					Type: "bar", Key: "baz",
					PhysicalID: "p1", GraphID: "g1",
					Associations:       []storage2.TypeAndKey{{"baz", "quz"}},
					ParentAssociations: []storage2.TypeAndKey{{"foo", "bar"}},
					Version:            1,
				},
				{
					Type: "baz", Key: "quz",
					PhysicalID: "p2", GraphID: "g1",
					ParentAssociations: []storage2.TypeAndKey{{"bar", "baz"}},
					Version:            2,
				},
				{
					Type: "foo", Key: "bar",
					GraphID:      "g1",
					Associations: []storage2.TypeAndKey{{"bar", "baz"}},
				},
			},
			RootEntities: []storage2.TypeAndKey{{"foo", "bar"}},
			Edges: []storage.GraphEdge{
				{From: storage2.TypeAndKey{Type: "bar", Key: "baz"}, To: storage2.TypeAndKey{Type: "baz", Key: "quz"}},
				{From: storage2.TypeAndKey{Type: "foo", Key: "bar"}, To: storage2.TypeAndKey{Type: "bar", Key: "baz"}},
			},
		},
	}

	// load a simple tree of 3 nodes
	tree := &testCase{
		run: runFactory("network", storage2.TypeAndKey{Type: "baz", Key: "quz"}, storage.EntityLoadCriteria{}),

		setup: func(m sqlmock.Sqlmock) {
			expectBasicEntityQueries(m, expectedBazQuz)
			expectBulkEntityQuery(m, []driver.Value{"g1"}, expectedFooBar, expectedBarBaz, expectedBazQuz)
			// foobar -> barbaz; foobar -> bazquz
			expectAssocQuery(m, assocQueryArgs, "foobar", "barbaz", "foobar", "bazquz")
		},

		expectedResult: storage.EntityGraph{
			Entities: []storage.NetworkEntity{
				{
					Type: "bar", Key: "baz",
					PhysicalID: "p1", GraphID: "g1",
					ParentAssociations: []storage2.TypeAndKey{{"foo", "bar"}},
					Version:            1,
				},
				{
					Type: "baz", Key: "quz",
					PhysicalID: "p2", GraphID: "g1",
					ParentAssociations: []storage2.TypeAndKey{{"foo", "bar"}},
					Version:            2,
				},
				{
					Type: "foo", Key: "bar",
					GraphID:      "g1",
					Associations: []storage2.TypeAndKey{{"bar", "baz"}, {"baz", "quz"}},
				},
			},
			RootEntities: []storage2.TypeAndKey{{"foo", "bar"}},
			Edges: []storage.GraphEdge{
				{From: storage2.TypeAndKey{Type: "foo", Key: "bar"}, To: storage2.TypeAndKey{Type: "bar", Key: "baz"}},
				{From: storage2.TypeAndKey{Type: "foo", Key: "bar"}, To: storage2.TypeAndKey{Type: "baz", Key: "quz"}},
			},
		},
	}

	// load an upside-down tree
	upsideDownTree := &testCase{
		run: runFactory("network", storage2.TypeAndKey{Type: "foo", Key: "bar"}, storage.EntityLoadCriteria{}),

		setup: func(m sqlmock.Sqlmock) {
			expectBasicEntityQueries(m, expectedFooBar)
			expectBulkEntityQuery(m, []driver.Value{"g1"}, expectedFooBar, expectedBarBaz, expectedBazQuz)
			// barbaz -> foobar; bazquz -> foobar
			expectAssocQuery(m, assocQueryArgs, "barbaz", "foobar", "bazquz", "foobar")
		},

		expectedResult: storage.EntityGraph{
			Entities: []storage.NetworkEntity{
				{
					Type: "bar", Key: "baz",
					PhysicalID: "p1", GraphID: "g1",
					Associations: []storage2.TypeAndKey{{"foo", "bar"}},
					Version:      1,
				},
				{
					Type: "baz", Key: "quz",
					PhysicalID: "p2", GraphID: "g1",
					Associations: []storage2.TypeAndKey{{"foo", "bar"}},
					Version:      2,
				},
				{
					Type: "foo", Key: "bar",
					GraphID:            "g1",
					ParentAssociations: []storage2.TypeAndKey{{"bar", "baz"}, {"baz", "quz"}},
				},
			},
			RootEntities: []storage2.TypeAndKey{{"bar", "baz"}, {"baz", "quz"}},
			Edges: []storage.GraphEdge{
				{From: storage2.TypeAndKey{Type: "bar", Key: "baz"}, To: storage2.TypeAndKey{Type: "foo", Key: "bar"}},
				{From: storage2.TypeAndKey{Type: "baz", Key: "quz"}, To: storage2.TypeAndKey{Type: "foo", Key: "bar"}},
			},
		},
	}

	// load a graph with a cycle
	withCycle := &testCase{
		run: runFactory("network", storage2.TypeAndKey{Type: "foo", Key: "bar"}, storage.EntityLoadCriteria{}),

		setup: func(m sqlmock.Sqlmock) {
			expectBasicEntityQueries(m, expectedFooBar)
			expectBulkEntityQuery(m, []driver.Value{"g1"}, expectedFooBar, expectedBarBaz, expectedBazQuz)
			// foobar -> barbaz; barbaz <-> bazquz
			expectAssocQuery(m, assocQueryArgs, "foobar", "barbaz", "barbaz", "bazquz", "bazquz", "barbaz")
		},

		expectedResult: storage.EntityGraph{
			Entities: []storage.NetworkEntity{
				{
					Type: "bar", Key: "baz",
					PhysicalID: "p1", GraphID: "g1",
					Associations:       []storage2.TypeAndKey{{"baz", "quz"}},
					ParentAssociations: []storage2.TypeAndKey{{"baz", "quz"}, {"foo", "bar"}},
					Version:            1,
				},
				{
					Type: "baz", Key: "quz",
					PhysicalID: "p2", GraphID: "g1",
					Associations:       []storage2.TypeAndKey{{"bar", "baz"}},
					ParentAssociations: []storage2.TypeAndKey{{"bar", "baz"}},
					Version:            2,
				},
				{
					Type: "foo", Key: "bar",
					GraphID:      "g1",
					Associations: []storage2.TypeAndKey{{"bar", "baz"}},
				},
			},
			RootEntities: []storage2.TypeAndKey{{"foo", "bar"}},
			Edges: []storage.GraphEdge{
				{From: storage2.TypeAndKey{Type: "bar", Key: "baz"}, To: storage2.TypeAndKey{Type: "baz", Key: "quz"}},
				{From: storage2.TypeAndKey{Type: "baz", Key: "quz"}, To: storage2.TypeAndKey{Type: "bar", Key: "baz"}},
				{From: storage2.TypeAndKey{Type: "foo", Key: "bar"}, To: storage2.TypeAndKey{Type: "bar", Key: "baz"}},
			},
		},
	}

	// load a ring
	ring := &testCase{
		run: runFactory("network", storage2.TypeAndKey{Type: "foo", Key: "bar"}, storage.EntityLoadCriteria{}),

		setup: func(m sqlmock.Sqlmock) {
			expectBasicEntityQueries(m, expectedFooBar)
			expectBulkEntityQuery(m, []driver.Value{"g1"}, expectedFooBar, expectedBarBaz, expectedBazQuz)
			// foobar -> barbaz -> bazquz -> foobar -> ...
			expectAssocQuery(m, assocQueryArgs, "foobar", "barbaz", "barbaz", "bazquz", "bazquz", "foobar")
		},

		expectedError: errors.New("graph does not have root nodes because it is a ring"),
	}

	runCase(t, linkedList)
	runCase(t, tree)
	runCase(t, upsideDownTree)
	runCase(t, withCycle)
	runCase(t, ring)
}

type testCase struct {
	// setup mock expectations. Transaction start is expected on the mock
	// generically
	setup func(m sqlmock.Sqlmock)

	// run the test case
	run func(store storage.ConfiguratorStorage) (interface{}, error)

	expectedError      error
	matchErrorInstance bool
	expectedResult     interface{}
}

func runCase(t *testing.T, test *testCase) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening stub DB conn: %s", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Printf("error closing stub DB: %s", err)
		}
	}()

	mock.ExpectBegin()
	test.setup(mock)

	factory := storage.NewSQLConfiguratorStorageFactory(db, &mockIDGenerator{})
	store, err := factory.StartTransaction(context.Background(), nil)
	assert.NoError(t, err)
	actual, err := test.run(store)

	if test.expectedError != nil {
		if test.matchErrorInstance {
			assert.True(t, err == test.expectedError)
		}
		assert.EqualError(t, err, test.expectedError.Error())
	} else {
		assert.NoError(t, err)
	}

	if test.expectedResult != nil {
		assert.Equal(t, test.expectedResult, actual)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

// this is a lot more coupled to implementation than I like for test code
// but the alternative of db test fixtures and manually querying/dumping the db
// isn't much better, if we want true unit tests for the storage impl
// for now, this is just the happy path
// TODO: add hooks for non-happy path test cases
func getTestCaseForEntityUpdate(
	update storage.EntityUpdateCriteria,
	entToUpdate expectedEntQueryResult,
	expectedEdgeLoads []expectedEntQueryResult,
	expectedGraphMerges ...[2]string,
) *testCase {
	expectedResult := storage.NetworkEntity{
		Type:    entToUpdate.entType,
		Key:     entToUpdate.key,
		GraphID: entToUpdate.graphID,
		Version: entToUpdate.version + 1,

		Name:        stringValue(update.NewName),
		Description: stringValue(update.NewDescription),
		Config:      bytesVal(update.NewConfig),
	}
	if update.NewPhysicalID != nil {
		expectedResult.PhysicalID = *update.NewPhysicalID
	} else {
		expectedResult.PhysicalID = entToUpdate.physicalID
	}

	delPermIdsSet := funk.Map(update.PermissionsToDelete, func(i string) (string, bool) { return i, true }).(map[string]bool)
	permID := 1
	for _, perm := range update.PermissionsToCreate {
		perm.ID = fmt.Sprintf("%d", permID)
		expectedResult.Permissions = append(expectedResult.Permissions, perm)
		permID++
	}
	for _, perm := range update.PermissionsToUpdate {
		if _, wasDeleted := delPermIdsSet[perm.ID]; !wasDeleted {
			expectedResult.Permissions = append(expectedResult.Permissions, perm)
		}
	}

	edgeLoadsByTk := funk.Map(
		expectedEdgeLoads,
		func(e expectedEntQueryResult) (storage2.TypeAndKey, expectedEntQueryResult) {
			return storage2.TypeAndKey{Type: e.entType, Key: e.key}, e
		},
	).(map[storage2.TypeAndKey]expectedEntQueryResult)

	if !funk.IsEmpty(update.AssociationsToAdd) {
		expectedResult.Associations = append(expectedResult.Associations, update.AssociationsToAdd...)
	}

	if !funk.IsEmpty(expectedGraphMerges) {
		expectedResult.GraphID = expectedGraphMerges[0][1]
	}

	return &testCase{
		setup: func(m sqlmock.Sqlmock) {
			// Basic fields
			expectBasicEntityQueries(m, entToUpdate)
			updateWithArgs := []driver.Value{}
			if update.NewName != nil {
				updateWithArgs = append(updateWithArgs, update.NewName)
			}
			if update.NewDescription != nil {
				updateWithArgs = append(updateWithArgs, update.NewDescription)
			}
			if update.NewPhysicalID != nil {
				updateWithArgs = append(updateWithArgs, update.NewPhysicalID)
			}
			if update.NewConfig != nil {
				updateWithArgs = append(updateWithArgs, update.NewConfig)
			}
			updateWithArgs = append(updateWithArgs, entToUpdate.pk)

			m.ExpectExec("UPDATE cfg_entities").WithArgs(updateWithArgs...).WillReturnResult(mockResult)

			// Permissions
			if !funk.IsEmpty(update.PermissionsToCreate) {
				expectPermissionCreation(m, entToUpdate.pk, 1, update.PermissionsToCreate...)
			}
			if !funk.IsEmpty(update.PermissionsToUpdate) {
				m.ExpectQuery(`SELECT COUNT\(\*\) FROM cfg_acls`).WithArgs(funk.Map(update.PermissionsToUpdate, func(acl storage.ACL) driver.Value { return acl.ID }).([]driver.Value)...).
					WillReturnRows(
						sqlmock.NewRows([]string{"count"}).
							AddRow(len(update.PermissionsToUpdate)),
					)
				expectPermissionUpdates(m, entToUpdate.pk, update.PermissionsToUpdate...)
			}
			if !funk.IsEmpty(update.PermissionsToDelete) {
				expectPermissionDeletes(m, update.PermissionsToDelete...)
			}

			// Graph
			if !funk.IsEmpty(update.AssociationsToAdd) {
				expectEdgeQueries(m, update.AssociationsToAdd, edgeLoadsByTk)
				expectEdgeInsertions(m, assocsToEdges(entToUpdate.pk, update.AssociationsToAdd, edgeLoadsByTk))
				if !funk.IsEmpty(expectedGraphMerges) {
					expectMergeGraphs(m, expectedGraphMerges)
				}
			}

			if !funk.IsEmpty(update.AssociationsToDelete) {
				expectEdgeQueries(m, update.AssociationsToDelete, edgeLoadsByTk)
				expectEdgeDeletions(m, assocsToEdges(entToUpdate.pk, update.AssociationsToDelete, edgeLoadsByTk))
			}
		},
		run: func(store storage.ConfiguratorStorage) (interface{}, error) {
			return store.UpdateEntity("network", update)
		},
		expectedResult: expectedResult,
	}
}

type expectedEntQueryResult struct {
	entType, key, pk, physicalID, graphID string
	version                               uint64
}

func expectBasicEntityQueries(m sqlmock.Sqlmock, expectations ...expectedEntQueryResult) {
	stmt := m.ExpectPrepare("SELECT .* FROM cfg_entities").WillBeClosed()
	for _, expect := range expectations {
		stmt.ExpectQuery().WithArgs("network", expect.key, expect.entType).WillReturnRows(expectedEntQueriesToRows(expect))
	}
}

func expectBulkEntityQuery(m sqlmock.Sqlmock, queryArgs []driver.Value, expectations ...expectedEntQueryResult) {
	m.ExpectQuery("SELECT .* FROM cfg_entities").WithArgs(queryArgs...).
		WillReturnRows(expectedEntQueriesToRows(expectations...))
}

func expectedEntQueriesToRows(expectations ...expectedEntQueryResult) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"pk", "key", "type", "physical_id", "version", "graph_id"})
	for _, expect := range expectations {
		rowValues := make([]driver.Value, 0, 6)
		rowValues = append(rowValues, expect.pk, expect.key, expect.entType)
		if expect.physicalID == "" {
			rowValues = append(rowValues, nil)
		} else {
			rowValues = append(rowValues, expect.physicalID)
		}
		rowValues = append(rowValues, expect.version, expect.graphID)
		rows.AddRow(rowValues...)
	}
	return rows
}

func expectAssocQuery(m sqlmock.Sqlmock, queryArgs []driver.Value, assocPks ...string) {
	rows := sqlmock.NewRows([]string{"from_pk", "to_pk"})
	for i := 0; i < len(assocPks); i += 2 {
		rows.AddRow(assocPks[i], assocPks[i+1])
	}
	m.ExpectQuery("SELECT assoc.from_pk, assoc.to_pk FROM cfg_assocs").WithArgs(queryArgs...).WillReturnRows(rows)
}

// [(old graph ID, new graph ID)]
func expectMergeGraphs(m sqlmock.Sqlmock, graphIDChanges [][2]string) {
	mergeStmt := m.ExpectPrepare("UPDATE cfg_entities").WillBeClosed()
	for _, delta := range graphIDChanges {
		mergeStmt.ExpectExec().WithArgs(delta[1], delta[0]).WillReturnResult(mockResult)
	}
}

func expectEdgeQueries(m sqlmock.Sqlmock, assocs []storage2.TypeAndKey, edgeLoadsByTk map[storage2.TypeAndKey]expectedEntQueryResult) {
	expectedLoads := funk.Map(
		assocs,
		func(tk storage2.TypeAndKey) expectedEntQueryResult { return edgeLoadsByTk[tk] },
	).([]expectedEntQueryResult)
	expectBasicEntityQueries(m, expectedLoads...)
}

// [(from_pk, to_pk)]
func expectEdgeInsertions(m sqlmock.Sqlmock, newEdges [][2]string) {
	edgeStmt := m.ExpectPrepare("INSERT INTO cfg_assocs").WillBeClosed()
	for _, edge := range newEdges {
		edgeStmt.ExpectExec().WithArgs(edge[0], edge[1]).WillReturnResult(mockResult)
	}
}

func expectEdgeDeletions(m sqlmock.Sqlmock, edges [][2]string) {
	edgeStmt := m.ExpectPrepare("DELETE FROM cfg_assocs").WillBeClosed()
	for _, edge := range edges {
		edgeStmt.ExpectExec().WithArgs(edge[0], edge[1]).WillReturnResult(mockResult)
	}
}

func expectPermissionCreation(m sqlmock.Sqlmock, entPk string, startId int, perms ...storage.ACL) {
	aclStmt := m.ExpectPrepare("INSERT INTO cfg_acls").WillBeClosed()
	for _, perm := range perms {
		exp := getExpectedACLInsert(entPk, &startId, perm)
		aclStmt.ExpectExec().WithArgs(exp.id, exp.entPk, exp.scope, exp.perm, exp.aclType, exp.filter).WillReturnResult(mockResult)
		startId++
	}
}

func expectPermissionUpdates(m sqlmock.Sqlmock, entPk string, perms ...storage.ACL) {
	stmt := m.ExpectPrepare("UPDATE cfg_acls").WillBeClosed()
	for _, perm := range perms {
		exp := getExpectedACLInsert(entPk, nil, perm)
		stmt.ExpectExec().WithArgs(exp.scope, exp.perm, exp.aclType, exp.filter, exp.id).WillReturnResult(mockResult)
	}
}

func expectPermissionDeletes(m sqlmock.Sqlmock, permIDs ...string) {
	args := make([]driver.Value, 0, len(permIDs))
	funk.ConvertSlice(permIDs, &args)
	m.ExpectExec("DELETE FROM cfg_acls").WithArgs(args...).WillReturnResult(mockResult)
}

type expectedACLInsert struct {
	id, entPk, scope, perm, aclType, filter driver.Value
}

func getExpectedACLInsert(entPk string, idOverride *int, perm storage.ACL) expectedACLInsert {
	var scope, typeVal, filter driver.Value
	if perm.Scope.Wildcard == storage.WildcardAll {
		scope = "*"
	} else {
		scope = strings.Join(perm.Scope.NetworkIDs, ",")
	}

	if perm.Type.Wildcard == storage.WildcardAll {
		typeVal = "*"
	} else {
		typeVal = perm.Type.EntityType
	}

	if funk.IsEmpty(perm.IDFilter) {
		filter = nil
	} else {
		filter = strings.Join(perm.IDFilter, ",")
	}

	ret := expectedACLInsert{entPk: entPk, perm: perm.Permission, aclType: typeVal, scope: scope, filter: filter}
	if idOverride == nil {
		ret.id = perm.ID
	} else {
		ret.id = fmt.Sprintf("%d", *idOverride)
	}
	return ret
}

func assocsToEdges(entPk string, assocs []storage2.TypeAndKey, edgeLoadsByTk map[storage2.TypeAndKey]expectedEntQueryResult) [][2]string {
	return funk.Map(
		assocs,
		func(tk storage2.TypeAndKey) [2]string {
			return [2]string{entPk, edgeLoadsByTk[tk].pk}
		},
	).([][2]string)
}

func stringPointer(val string) *string {
	return &val
}

func stringValue(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func bytesPointer(val []byte) *[]byte {
	return &val
}

func bytesVal(val *[]byte) []byte {
	if val == nil {
		return nil
	}
	return *val
}
