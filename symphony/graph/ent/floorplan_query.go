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
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/symphony/graph/ent/file"
	"github.com/facebookincubator/symphony/graph/ent/floorplan"
	"github.com/facebookincubator/symphony/graph/ent/floorplanreferencepoint"
	"github.com/facebookincubator/symphony/graph/ent/floorplanscale"
	"github.com/facebookincubator/symphony/graph/ent/location"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// FloorPlanQuery is the builder for querying FloorPlan entities.
type FloorPlanQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.FloorPlan
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (fpq *FloorPlanQuery) Where(ps ...predicate.FloorPlan) *FloorPlanQuery {
	fpq.predicates = append(fpq.predicates, ps...)
	return fpq
}

// Limit adds a limit step to the query.
func (fpq *FloorPlanQuery) Limit(limit int) *FloorPlanQuery {
	fpq.limit = &limit
	return fpq
}

// Offset adds an offset step to the query.
func (fpq *FloorPlanQuery) Offset(offset int) *FloorPlanQuery {
	fpq.offset = &offset
	return fpq
}

// Order adds an order step to the query.
func (fpq *FloorPlanQuery) Order(o ...Order) *FloorPlanQuery {
	fpq.order = append(fpq.order, o...)
	return fpq
}

// QueryLocation chains the current query on the location edge.
func (fpq *FloorPlanQuery) QueryLocation() *LocationQuery {
	query := &LocationQuery{config: fpq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(floorplan.Table, floorplan.FieldID, fpq.sqlQuery()),
		sqlgraph.To(location.Table, location.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, floorplan.LocationTable, floorplan.LocationColumn),
	)
	query.sql = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
	return query
}

// QueryReferencePoint chains the current query on the reference_point edge.
func (fpq *FloorPlanQuery) QueryReferencePoint() *FloorPlanReferencePointQuery {
	query := &FloorPlanReferencePointQuery{config: fpq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(floorplan.Table, floorplan.FieldID, fpq.sqlQuery()),
		sqlgraph.To(floorplanreferencepoint.Table, floorplanreferencepoint.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, floorplan.ReferencePointTable, floorplan.ReferencePointColumn),
	)
	query.sql = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
	return query
}

// QueryScale chains the current query on the scale edge.
func (fpq *FloorPlanQuery) QueryScale() *FloorPlanScaleQuery {
	query := &FloorPlanScaleQuery{config: fpq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(floorplan.Table, floorplan.FieldID, fpq.sqlQuery()),
		sqlgraph.To(floorplanscale.Table, floorplanscale.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, floorplan.ScaleTable, floorplan.ScaleColumn),
	)
	query.sql = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
	return query
}

// QueryImage chains the current query on the image edge.
func (fpq *FloorPlanQuery) QueryImage() *FileQuery {
	query := &FileQuery{config: fpq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(floorplan.Table, floorplan.FieldID, fpq.sqlQuery()),
		sqlgraph.To(file.Table, file.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, floorplan.ImageTable, floorplan.ImageColumn),
	)
	query.sql = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
	return query
}

// First returns the first FloorPlan entity in the query. Returns *ErrNotFound when no floorplan was found.
func (fpq *FloorPlanQuery) First(ctx context.Context) (*FloorPlan, error) {
	fps, err := fpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fps) == 0 {
		return nil, &ErrNotFound{floorplan.Label}
	}
	return fps[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fpq *FloorPlanQuery) FirstX(ctx context.Context) *FloorPlan {
	fp, err := fpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return fp
}

// FirstID returns the first FloorPlan id in the query. Returns *ErrNotFound when no id was found.
func (fpq *FloorPlanQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{floorplan.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (fpq *FloorPlanQuery) FirstXID(ctx context.Context) string {
	id, err := fpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only FloorPlan entity in the query, returns an error if not exactly one entity was returned.
func (fpq *FloorPlanQuery) Only(ctx context.Context) (*FloorPlan, error) {
	fps, err := fpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fps) {
	case 1:
		return fps[0], nil
	case 0:
		return nil, &ErrNotFound{floorplan.Label}
	default:
		return nil, &ErrNotSingular{floorplan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fpq *FloorPlanQuery) OnlyX(ctx context.Context) *FloorPlan {
	fp, err := fpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return fp
}

// OnlyID returns the only FloorPlan id in the query, returns an error if not exactly one id was returned.
func (fpq *FloorPlanQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{floorplan.Label}
	default:
		err = &ErrNotSingular{floorplan.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (fpq *FloorPlanQuery) OnlyXID(ctx context.Context) string {
	id, err := fpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FloorPlans.
func (fpq *FloorPlanQuery) All(ctx context.Context) ([]*FloorPlan, error) {
	return fpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fpq *FloorPlanQuery) AllX(ctx context.Context) []*FloorPlan {
	fps, err := fpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fps
}

// IDs executes the query and returns a list of FloorPlan ids.
func (fpq *FloorPlanQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := fpq.Select(floorplan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fpq *FloorPlanQuery) IDsX(ctx context.Context) []string {
	ids, err := fpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fpq *FloorPlanQuery) Count(ctx context.Context) (int, error) {
	return fpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fpq *FloorPlanQuery) CountX(ctx context.Context) int {
	count, err := fpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fpq *FloorPlanQuery) Exist(ctx context.Context) (bool, error) {
	return fpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fpq *FloorPlanQuery) ExistX(ctx context.Context) bool {
	exist, err := fpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fpq *FloorPlanQuery) Clone() *FloorPlanQuery {
	return &FloorPlanQuery{
		config:     fpq.config,
		limit:      fpq.limit,
		offset:     fpq.offset,
		order:      append([]Order{}, fpq.order...),
		unique:     append([]string{}, fpq.unique...),
		predicates: append([]predicate.FloorPlan{}, fpq.predicates...),
		// clone intermediate query.
		sql: fpq.sql.Clone(),
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
//	client.FloorPlan.Query().
//		GroupBy(floorplan.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fpq *FloorPlanQuery) GroupBy(field string, fields ...string) *FloorPlanGroupBy {
	group := &FloorPlanGroupBy{config: fpq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = fpq.sqlQuery()
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
//	client.FloorPlan.Query().
//		Select(floorplan.FieldCreateTime).
//		Scan(ctx, &v)
//
func (fpq *FloorPlanQuery) Select(field string, fields ...string) *FloorPlanSelect {
	selector := &FloorPlanSelect{config: fpq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = fpq.sqlQuery()
	return selector
}

func (fpq *FloorPlanQuery) sqlAll(ctx context.Context) ([]*FloorPlan, error) {
	rows := &sql.Rows{}
	selector := fpq.sqlQuery()
	if unique := fpq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := fpq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var fps FloorPlans
	if err := fps.FromRows(rows); err != nil {
		return nil, err
	}
	fps.config(fpq.config)
	return fps, nil
}

func (fpq *FloorPlanQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := fpq.sqlQuery()
	unique := []string{floorplan.FieldID}
	if len(fpq.unique) > 0 {
		unique = fpq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := fpq.driver.Query(ctx, query, args, rows); err != nil {
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

func (fpq *FloorPlanQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fpq *FloorPlanQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(fpq.driver.Dialect())
	t1 := builder.Table(floorplan.Table)
	selector := builder.Select(t1.Columns(floorplan.Columns...)...).From(t1)
	if fpq.sql != nil {
		selector = fpq.sql
		selector.Select(selector.Columns(floorplan.Columns...)...)
	}
	for _, p := range fpq.predicates {
		p(selector)
	}
	for _, p := range fpq.order {
		p(selector)
	}
	if offset := fpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FloorPlanGroupBy is the builder for group-by FloorPlan entities.
type FloorPlanGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fpgb *FloorPlanGroupBy) Aggregate(fns ...Aggregate) *FloorPlanGroupBy {
	fpgb.fns = append(fpgb.fns, fns...)
	return fpgb
}

// Scan applies the group-by query and scan the result into the given value.
func (fpgb *FloorPlanGroupBy) Scan(ctx context.Context, v interface{}) error {
	return fpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fpgb *FloorPlanGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (fpgb *FloorPlanGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fpgb *FloorPlanGroupBy) StringsX(ctx context.Context) []string {
	v, err := fpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (fpgb *FloorPlanGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fpgb *FloorPlanGroupBy) IntsX(ctx context.Context) []int {
	v, err := fpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (fpgb *FloorPlanGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fpgb *FloorPlanGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (fpgb *FloorPlanGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fpgb *FloorPlanGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fpgb *FloorPlanGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fpgb.sqlQuery().Query()
	if err := fpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fpgb *FloorPlanGroupBy) sqlQuery() *sql.Selector {
	selector := fpgb.sql
	columns := make([]string, 0, len(fpgb.fields)+len(fpgb.fns))
	columns = append(columns, fpgb.fields...)
	for _, fn := range fpgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(fpgb.fields...)
}

// FloorPlanSelect is the builder for select fields of FloorPlan entities.
type FloorPlanSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (fps *FloorPlanSelect) Scan(ctx context.Context, v interface{}) error {
	return fps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fps *FloorPlanSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fps *FloorPlanSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FloorPlanSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fps *FloorPlanSelect) StringsX(ctx context.Context) []string {
	v, err := fps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fps *FloorPlanSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FloorPlanSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fps *FloorPlanSelect) IntsX(ctx context.Context) []int {
	v, err := fps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fps *FloorPlanSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FloorPlanSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fps *FloorPlanSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fps *FloorPlanSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FloorPlanSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fps *FloorPlanSelect) BoolsX(ctx context.Context) []bool {
	v, err := fps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fps *FloorPlanSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fps.sqlQuery().Query()
	if err := fps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fps *FloorPlanSelect) sqlQuery() sql.Querier {
	view := "floorplan_view"
	return sql.Dialect(fps.driver.Dialect()).
		Select(fps.fields...).From(fps.sql.As(view))
}
