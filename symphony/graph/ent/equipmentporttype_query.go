// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentporttype"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
)

// EquipmentPortTypeQuery is the builder for querying EquipmentPortType entities.
type EquipmentPortTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.EquipmentPortType
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (eptq *EquipmentPortTypeQuery) Where(ps ...predicate.EquipmentPortType) *EquipmentPortTypeQuery {
	eptq.predicates = append(eptq.predicates, ps...)
	return eptq
}

// Limit adds a limit step to the query.
func (eptq *EquipmentPortTypeQuery) Limit(limit int) *EquipmentPortTypeQuery {
	eptq.limit = &limit
	return eptq
}

// Offset adds an offset step to the query.
func (eptq *EquipmentPortTypeQuery) Offset(offset int) *EquipmentPortTypeQuery {
	eptq.offset = &offset
	return eptq
}

// Order adds an order step to the query.
func (eptq *EquipmentPortTypeQuery) Order(o ...Order) *EquipmentPortTypeQuery {
	eptq.order = append(eptq.order, o...)
	return eptq
}

// QueryPropertyTypes chains the current query on the property_types edge.
func (eptq *EquipmentPortTypeQuery) QueryPropertyTypes() *PropertyTypeQuery {
	query := &PropertyTypeQuery{config: eptq.config}

	builder := sql.Dialect(eptq.driver.Dialect())
	t1 := builder.Table(propertytype.Table)
	t2 := eptq.sqlQuery()
	t2.Select(t2.C(equipmentporttype.FieldID))
	query.sql = builder.Select().
		From(t1).
		Join(t2).
		On(t1.C(equipmentporttype.PropertyTypesColumn), t2.C(equipmentporttype.FieldID))
	return query
}

// QueryLinkPropertyTypes chains the current query on the link_property_types edge.
func (eptq *EquipmentPortTypeQuery) QueryLinkPropertyTypes() *PropertyTypeQuery {
	query := &PropertyTypeQuery{config: eptq.config}

	builder := sql.Dialect(eptq.driver.Dialect())
	t1 := builder.Table(propertytype.Table)
	t2 := eptq.sqlQuery()
	t2.Select(t2.C(equipmentporttype.FieldID))
	query.sql = builder.Select().
		From(t1).
		Join(t2).
		On(t1.C(equipmentporttype.LinkPropertyTypesColumn), t2.C(equipmentporttype.FieldID))
	return query
}

// QueryPortDefinitions chains the current query on the port_definitions edge.
func (eptq *EquipmentPortTypeQuery) QueryPortDefinitions() *EquipmentPortDefinitionQuery {
	query := &EquipmentPortDefinitionQuery{config: eptq.config}

	builder := sql.Dialect(eptq.driver.Dialect())
	t1 := builder.Table(equipmentportdefinition.Table)
	t2 := eptq.sqlQuery()
	t2.Select(t2.C(equipmentporttype.FieldID))
	query.sql = builder.Select().
		From(t1).
		Join(t2).
		On(t1.C(equipmentporttype.PortDefinitionsColumn), t2.C(equipmentporttype.FieldID))
	return query
}

// First returns the first EquipmentPortType entity in the query. Returns *ErrNotFound when no equipmentporttype was found.
func (eptq *EquipmentPortTypeQuery) First(ctx context.Context) (*EquipmentPortType, error) {
	epts, err := eptq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(epts) == 0 {
		return nil, &ErrNotFound{equipmentporttype.Label}
	}
	return epts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) FirstX(ctx context.Context) *EquipmentPortType {
	ept, err := eptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ept
}

// FirstID returns the first EquipmentPortType id in the query. Returns *ErrNotFound when no id was found.
func (eptq *EquipmentPortTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eptq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{equipmentporttype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) FirstXID(ctx context.Context) string {
	id, err := eptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only EquipmentPortType entity in the query, returns an error if not exactly one entity was returned.
func (eptq *EquipmentPortTypeQuery) Only(ctx context.Context) (*EquipmentPortType, error) {
	epts, err := eptq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(epts) {
	case 1:
		return epts[0], nil
	case 0:
		return nil, &ErrNotFound{equipmentporttype.Label}
	default:
		return nil, &ErrNotSingular{equipmentporttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) OnlyX(ctx context.Context) *EquipmentPortType {
	ept, err := eptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ept
}

// OnlyID returns the only EquipmentPortType id in the query, returns an error if not exactly one id was returned.
func (eptq *EquipmentPortTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eptq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{equipmentporttype.Label}
	default:
		err = &ErrNotSingular{equipmentporttype.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) OnlyXID(ctx context.Context) string {
	id, err := eptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentPortTypes.
func (eptq *EquipmentPortTypeQuery) All(ctx context.Context) ([]*EquipmentPortType, error) {
	return eptq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) AllX(ctx context.Context) []*EquipmentPortType {
	epts, err := eptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return epts
}

// IDs executes the query and returns a list of EquipmentPortType ids.
func (eptq *EquipmentPortTypeQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := eptq.Select(equipmentporttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := eptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eptq *EquipmentPortTypeQuery) Count(ctx context.Context) (int, error) {
	return eptq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) CountX(ctx context.Context) int {
	count, err := eptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eptq *EquipmentPortTypeQuery) Exist(ctx context.Context) (bool, error) {
	return eptq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eptq *EquipmentPortTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := eptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eptq *EquipmentPortTypeQuery) Clone() *EquipmentPortTypeQuery {
	return &EquipmentPortTypeQuery{
		config:     eptq.config,
		limit:      eptq.limit,
		offset:     eptq.offset,
		order:      append([]Order{}, eptq.order...),
		unique:     append([]string{}, eptq.unique...),
		predicates: append([]predicate.EquipmentPortType{}, eptq.predicates...),
		// clone intermediate queries.
		sql: eptq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentPortType.Query().
//		GroupBy(equipmentporttype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (eptq *EquipmentPortTypeQuery) GroupBy(field string, fields ...string) *EquipmentPortTypeGroupBy {
	group := &EquipmentPortTypeGroupBy{config: eptq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = eptq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.EquipmentPortType.Query().
//		Select(equipmentporttype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (eptq *EquipmentPortTypeQuery) Select(field string, fields ...string) *EquipmentPortTypeSelect {
	selector := &EquipmentPortTypeSelect{config: eptq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = eptq.sqlQuery()
	return selector
}

func (eptq *EquipmentPortTypeQuery) sqlAll(ctx context.Context) ([]*EquipmentPortType, error) {
	rows := &sql.Rows{}
	selector := eptq.sqlQuery()
	if unique := eptq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := eptq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var epts EquipmentPortTypes
	if err := epts.FromRows(rows); err != nil {
		return nil, err
	}
	epts.config(eptq.config)
	return epts, nil
}

func (eptq *EquipmentPortTypeQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := eptq.sqlQuery()
	unique := []string{equipmentporttype.FieldID}
	if len(eptq.unique) > 0 {
		unique = eptq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := eptq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (eptq *EquipmentPortTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eptq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (eptq *EquipmentPortTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(eptq.driver.Dialect())
	t1 := builder.Table(equipmentporttype.Table)
	selector := builder.Select(t1.Columns(equipmentporttype.Columns...)...).From(t1)
	if eptq.sql != nil {
		selector = eptq.sql
		selector.Select(selector.Columns(equipmentporttype.Columns...)...)
	}
	for _, p := range eptq.predicates {
		p(selector)
	}
	for _, p := range eptq.order {
		p(selector)
	}
	if offset := eptq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eptq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentPortTypeGroupBy is the builder for group-by EquipmentPortType entities.
type EquipmentPortTypeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eptgb *EquipmentPortTypeGroupBy) Aggregate(fns ...Aggregate) *EquipmentPortTypeGroupBy {
	eptgb.fns = append(eptgb.fns, fns...)
	return eptgb
}

// Scan applies the group-by query and scan the result into the given value.
func (eptgb *EquipmentPortTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	return eptgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (eptgb *EquipmentPortTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := eptgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (eptgb *EquipmentPortTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(eptgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := eptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (eptgb *EquipmentPortTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := eptgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (eptgb *EquipmentPortTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(eptgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := eptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (eptgb *EquipmentPortTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := eptgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (eptgb *EquipmentPortTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(eptgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := eptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (eptgb *EquipmentPortTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := eptgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (eptgb *EquipmentPortTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(eptgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := eptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (eptgb *EquipmentPortTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := eptgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (eptgb *EquipmentPortTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := eptgb.sqlQuery().Query()
	if err := eptgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eptgb *EquipmentPortTypeGroupBy) sqlQuery() *sql.Selector {
	selector := eptgb.sql
	columns := make([]string, 0, len(eptgb.fields)+len(eptgb.fns))
	columns = append(columns, eptgb.fields...)
	for _, fn := range eptgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(eptgb.fields...)
}

// EquipmentPortTypeSelect is the builder for select fields of EquipmentPortType entities.
type EquipmentPortTypeSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (epts *EquipmentPortTypeSelect) Scan(ctx context.Context, v interface{}) error {
	return epts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (epts *EquipmentPortTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := epts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (epts *EquipmentPortTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(epts.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := epts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (epts *EquipmentPortTypeSelect) StringsX(ctx context.Context) []string {
	v, err := epts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (epts *EquipmentPortTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(epts.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := epts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (epts *EquipmentPortTypeSelect) IntsX(ctx context.Context) []int {
	v, err := epts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (epts *EquipmentPortTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(epts.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := epts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (epts *EquipmentPortTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := epts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (epts *EquipmentPortTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(epts.fields) > 1 {
		return nil, errors.New("ent: EquipmentPortTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := epts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (epts *EquipmentPortTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := epts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (epts *EquipmentPortTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := epts.sqlQuery().Query()
	if err := epts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (epts *EquipmentPortTypeSelect) sqlQuery() sql.Querier {
	view := "equipmentporttype_view"
	return sql.Dialect(epts.driver.Dialect()).
		Select(epts.fields...).From(epts.sql.As(view))
}
