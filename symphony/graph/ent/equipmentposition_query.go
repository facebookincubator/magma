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
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentposition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentpositiondefinition"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// EquipmentPositionQuery is the builder for querying EquipmentPosition entities.
type EquipmentPositionQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.EquipmentPosition
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (epq *EquipmentPositionQuery) Where(ps ...predicate.EquipmentPosition) *EquipmentPositionQuery {
	epq.predicates = append(epq.predicates, ps...)
	return epq
}

// Limit adds a limit step to the query.
func (epq *EquipmentPositionQuery) Limit(limit int) *EquipmentPositionQuery {
	epq.limit = &limit
	return epq
}

// Offset adds an offset step to the query.
func (epq *EquipmentPositionQuery) Offset(offset int) *EquipmentPositionQuery {
	epq.offset = &offset
	return epq
}

// Order adds an order step to the query.
func (epq *EquipmentPositionQuery) Order(o ...Order) *EquipmentPositionQuery {
	epq.order = append(epq.order, o...)
	return epq
}

// QueryDefinition chains the current query on the definition edge.
func (epq *EquipmentPositionQuery) QueryDefinition() *EquipmentPositionDefinitionQuery {
	query := &EquipmentPositionDefinitionQuery{config: epq.config}

	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(equipmentpositiondefinition.Table)
	t2 := epq.sqlQuery()
	t2.Select(t2.C(equipmentposition.DefinitionColumn))
	query.sql = builder.Select(t1.Columns(equipmentpositiondefinition.Columns...)...).
		From(t1).
		Join(t2).
		On(t1.C(equipmentpositiondefinition.FieldID), t2.C(equipmentposition.DefinitionColumn))
	return query
}

// QueryParent chains the current query on the parent edge.
func (epq *EquipmentPositionQuery) QueryParent() *EquipmentQuery {
	query := &EquipmentQuery{config: epq.config}

	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(equipment.Table)
	t2 := epq.sqlQuery()
	t2.Select(t2.C(equipmentposition.ParentColumn))
	query.sql = builder.Select(t1.Columns(equipment.Columns...)...).
		From(t1).
		Join(t2).
		On(t1.C(equipment.FieldID), t2.C(equipmentposition.ParentColumn))
	return query
}

// QueryAttachment chains the current query on the attachment edge.
func (epq *EquipmentPositionQuery) QueryAttachment() *EquipmentQuery {
	query := &EquipmentQuery{config: epq.config}

	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(equipment.Table)
	t2 := epq.sqlQuery()
	t2.Select(t2.C(equipmentposition.FieldID))
	query.sql = builder.Select().
		From(t1).
		Join(t2).
		On(t1.C(equipmentposition.AttachmentColumn), t2.C(equipmentposition.FieldID))
	return query
}

// First returns the first EquipmentPosition entity in the query. Returns *ErrNotFound when no equipmentposition was found.
func (epq *EquipmentPositionQuery) First(ctx context.Context) (*EquipmentPosition, error) {
	eps, err := epq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(eps) == 0 {
		return nil, &ErrNotFound{equipmentposition.Label}
	}
	return eps[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (epq *EquipmentPositionQuery) FirstX(ctx context.Context) *EquipmentPosition {
	ep, err := epq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ep
}

// FirstID returns the first EquipmentPosition id in the query. Returns *ErrNotFound when no id was found.
func (epq *EquipmentPositionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = epq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{equipmentposition.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (epq *EquipmentPositionQuery) FirstXID(ctx context.Context) string {
	id, err := epq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only EquipmentPosition entity in the query, returns an error if not exactly one entity was returned.
func (epq *EquipmentPositionQuery) Only(ctx context.Context) (*EquipmentPosition, error) {
	eps, err := epq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(eps) {
	case 1:
		return eps[0], nil
	case 0:
		return nil, &ErrNotFound{equipmentposition.Label}
	default:
		return nil, &ErrNotSingular{equipmentposition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (epq *EquipmentPositionQuery) OnlyX(ctx context.Context) *EquipmentPosition {
	ep, err := epq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ep
}

// OnlyID returns the only EquipmentPosition id in the query, returns an error if not exactly one id was returned.
func (epq *EquipmentPositionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = epq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{equipmentposition.Label}
	default:
		err = &ErrNotSingular{equipmentposition.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (epq *EquipmentPositionQuery) OnlyXID(ctx context.Context) string {
	id, err := epq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentPositions.
func (epq *EquipmentPositionQuery) All(ctx context.Context) ([]*EquipmentPosition, error) {
	return epq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (epq *EquipmentPositionQuery) AllX(ctx context.Context) []*EquipmentPosition {
	eps, err := epq.All(ctx)
	if err != nil {
		panic(err)
	}
	return eps
}

// IDs executes the query and returns a list of EquipmentPosition ids.
func (epq *EquipmentPositionQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := epq.Select(equipmentposition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (epq *EquipmentPositionQuery) IDsX(ctx context.Context) []string {
	ids, err := epq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (epq *EquipmentPositionQuery) Count(ctx context.Context) (int, error) {
	return epq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (epq *EquipmentPositionQuery) CountX(ctx context.Context) int {
	count, err := epq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (epq *EquipmentPositionQuery) Exist(ctx context.Context) (bool, error) {
	return epq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (epq *EquipmentPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := epq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (epq *EquipmentPositionQuery) Clone() *EquipmentPositionQuery {
	return &EquipmentPositionQuery{
		config:     epq.config,
		limit:      epq.limit,
		offset:     epq.offset,
		order:      append([]Order{}, epq.order...),
		unique:     append([]string{}, epq.unique...),
		predicates: append([]predicate.EquipmentPosition{}, epq.predicates...),
		// clone intermediate queries.
		sql: epq.sql.Clone(),
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
//	client.EquipmentPosition.Query().
//		GroupBy(equipmentposition.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (epq *EquipmentPositionQuery) GroupBy(field string, fields ...string) *EquipmentPositionGroupBy {
	group := &EquipmentPositionGroupBy{config: epq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = epq.sqlQuery()
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
//	client.EquipmentPosition.Query().
//		Select(equipmentposition.FieldCreateTime).
//		Scan(ctx, &v)
//
func (epq *EquipmentPositionQuery) Select(field string, fields ...string) *EquipmentPositionSelect {
	selector := &EquipmentPositionSelect{config: epq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = epq.sqlQuery()
	return selector
}

func (epq *EquipmentPositionQuery) sqlAll(ctx context.Context) ([]*EquipmentPosition, error) {
	rows := &sql.Rows{}
	selector := epq.sqlQuery()
	if unique := epq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := epq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var eps EquipmentPositions
	if err := eps.FromRows(rows); err != nil {
		return nil, err
	}
	eps.config(epq.config)
	return eps, nil
}

func (epq *EquipmentPositionQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := epq.sqlQuery()
	unique := []string{equipmentposition.FieldID}
	if len(epq.unique) > 0 {
		unique = epq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := epq.driver.Query(ctx, query, args, rows); err != nil {
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

func (epq *EquipmentPositionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := epq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (epq *EquipmentPositionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(equipmentposition.Table)
	selector := builder.Select(t1.Columns(equipmentposition.Columns...)...).From(t1)
	if epq.sql != nil {
		selector = epq.sql
		selector.Select(selector.Columns(equipmentposition.Columns...)...)
	}
	for _, p := range epq.predicates {
		p(selector)
	}
	for _, p := range epq.order {
		p(selector)
	}
	if offset := epq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := epq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentPositionGroupBy is the builder for group-by EquipmentPosition entities.
type EquipmentPositionGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (epgb *EquipmentPositionGroupBy) Aggregate(fns ...Aggregate) *EquipmentPositionGroupBy {
	epgb.fns = append(epgb.fns, fns...)
	return epgb
}

// Scan applies the group-by query and scan the result into the given value.
func (epgb *EquipmentPositionGroupBy) Scan(ctx context.Context, v interface{}) error {
	return epgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (epgb *EquipmentPositionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := epgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (epgb *EquipmentPositionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(epgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := epgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (epgb *EquipmentPositionGroupBy) StringsX(ctx context.Context) []string {
	v, err := epgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (epgb *EquipmentPositionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(epgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := epgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (epgb *EquipmentPositionGroupBy) IntsX(ctx context.Context) []int {
	v, err := epgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (epgb *EquipmentPositionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(epgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := epgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (epgb *EquipmentPositionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := epgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (epgb *EquipmentPositionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(epgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := epgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (epgb *EquipmentPositionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := epgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (epgb *EquipmentPositionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := epgb.sqlQuery().Query()
	if err := epgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (epgb *EquipmentPositionGroupBy) sqlQuery() *sql.Selector {
	selector := epgb.sql
	columns := make([]string, 0, len(epgb.fields)+len(epgb.fns))
	columns = append(columns, epgb.fields...)
	for _, fn := range epgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(epgb.fields...)
}

// EquipmentPositionSelect is the builder for select fields of EquipmentPosition entities.
type EquipmentPositionSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (eps *EquipmentPositionSelect) Scan(ctx context.Context, v interface{}) error {
	return eps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (eps *EquipmentPositionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := eps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (eps *EquipmentPositionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(eps.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := eps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (eps *EquipmentPositionSelect) StringsX(ctx context.Context) []string {
	v, err := eps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (eps *EquipmentPositionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(eps.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := eps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (eps *EquipmentPositionSelect) IntsX(ctx context.Context) []int {
	v, err := eps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (eps *EquipmentPositionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(eps.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := eps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (eps *EquipmentPositionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := eps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (eps *EquipmentPositionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(eps.fields) > 1 {
		return nil, errors.New("ent: EquipmentPositionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := eps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (eps *EquipmentPositionSelect) BoolsX(ctx context.Context) []bool {
	v, err := eps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (eps *EquipmentPositionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := eps.sqlQuery().Query()
	if err := eps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eps *EquipmentPositionSelect) sqlQuery() sql.Querier {
	view := "equipmentposition_view"
	return sql.Dialect(eps.driver.Dialect()).
		Select(eps.fields...).From(eps.sql.As(view))
}
