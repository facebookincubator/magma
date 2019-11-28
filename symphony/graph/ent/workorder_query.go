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
	"github.com/facebookincubator/symphony/graph/ent/checklistitem"
	"github.com/facebookincubator/symphony/graph/ent/comment"
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/file"
	"github.com/facebookincubator/symphony/graph/ent/link"
	"github.com/facebookincubator/symphony/graph/ent/location"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/project"
	"github.com/facebookincubator/symphony/graph/ent/property"
	"github.com/facebookincubator/symphony/graph/ent/technician"
	"github.com/facebookincubator/symphony/graph/ent/workorder"
	"github.com/facebookincubator/symphony/graph/ent/workordertype"
)

// WorkOrderQuery is the builder for querying WorkOrder entities.
type WorkOrderQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.WorkOrder
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (woq *WorkOrderQuery) Where(ps ...predicate.WorkOrder) *WorkOrderQuery {
	woq.predicates = append(woq.predicates, ps...)
	return woq
}

// Limit adds a limit step to the query.
func (woq *WorkOrderQuery) Limit(limit int) *WorkOrderQuery {
	woq.limit = &limit
	return woq
}

// Offset adds an offset step to the query.
func (woq *WorkOrderQuery) Offset(offset int) *WorkOrderQuery {
	woq.offset = &offset
	return woq
}

// Order adds an order step to the query.
func (woq *WorkOrderQuery) Order(o ...Order) *WorkOrderQuery {
	woq.order = append(woq.order, o...)
	return woq
}

// QueryType chains the current query on the type edge.
func (woq *WorkOrderQuery) QueryType() *WorkOrderTypeQuery {
	query := &WorkOrderTypeQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(workordertype.Table, workordertype.FieldID),
		sql.Edge(sql.M2O, false, workorder.TypeTable, workorder.TypeColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryEquipment chains the current query on the equipment edge.
func (woq *WorkOrderQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(equipment.Table, equipment.FieldID),
		sql.Edge(sql.O2M, true, workorder.EquipmentTable, workorder.EquipmentColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryLinks chains the current query on the links edge.
func (woq *WorkOrderQuery) QueryLinks() *LinkQuery {
	query := &LinkQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(link.Table, link.FieldID),
		sql.Edge(sql.O2M, true, workorder.LinksTable, workorder.LinksColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryFiles chains the current query on the files edge.
func (woq *WorkOrderQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(file.Table, file.FieldID),
		sql.Edge(sql.O2M, false, workorder.FilesTable, workorder.FilesColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryLocation chains the current query on the location edge.
func (woq *WorkOrderQuery) QueryLocation() *LocationQuery {
	query := &LocationQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(location.Table, location.FieldID),
		sql.Edge(sql.M2O, false, workorder.LocationTable, workorder.LocationColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryComments chains the current query on the comments edge.
func (woq *WorkOrderQuery) QueryComments() *CommentQuery {
	query := &CommentQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(comment.Table, comment.FieldID),
		sql.Edge(sql.O2M, false, workorder.CommentsTable, workorder.CommentsColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryProperties chains the current query on the properties edge.
func (woq *WorkOrderQuery) QueryProperties() *PropertyQuery {
	query := &PropertyQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(property.Table, property.FieldID),
		sql.Edge(sql.O2M, false, workorder.PropertiesTable, workorder.PropertiesColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryCheckListItems chains the current query on the check_list_items edge.
func (woq *WorkOrderQuery) QueryCheckListItems() *CheckListItemQuery {
	query := &CheckListItemQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(checklistitem.Table, checklistitem.FieldID),
		sql.Edge(sql.O2M, false, workorder.CheckListItemsTable, workorder.CheckListItemsColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryTechnician chains the current query on the technician edge.
func (woq *WorkOrderQuery) QueryTechnician() *TechnicianQuery {
	query := &TechnicianQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(technician.Table, technician.FieldID),
		sql.Edge(sql.M2O, false, workorder.TechnicianTable, workorder.TechnicianColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// QueryProject chains the current query on the project edge.
func (woq *WorkOrderQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: woq.config}
	step := sql.NewStep(
		sql.From(workorder.Table, workorder.FieldID, woq.sqlQuery()),
		sql.To(project.Table, project.FieldID),
		sql.Edge(sql.M2O, true, workorder.ProjectTable, workorder.ProjectColumn),
	)
	query.sql = sql.SetNeighbors(woq.driver.Dialect(), step)
	return query
}

// First returns the first WorkOrder entity in the query. Returns *ErrNotFound when no workorder was found.
func (woq *WorkOrderQuery) First(ctx context.Context) (*WorkOrder, error) {
	wos, err := woq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(wos) == 0 {
		return nil, &ErrNotFound{workorder.Label}
	}
	return wos[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (woq *WorkOrderQuery) FirstX(ctx context.Context) *WorkOrder {
	wo, err := woq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return wo
}

// FirstID returns the first WorkOrder id in the query. Returns *ErrNotFound when no id was found.
func (woq *WorkOrderQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = woq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{workorder.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (woq *WorkOrderQuery) FirstXID(ctx context.Context) string {
	id, err := woq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WorkOrder entity in the query, returns an error if not exactly one entity was returned.
func (woq *WorkOrderQuery) Only(ctx context.Context) (*WorkOrder, error) {
	wos, err := woq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(wos) {
	case 1:
		return wos[0], nil
	case 0:
		return nil, &ErrNotFound{workorder.Label}
	default:
		return nil, &ErrNotSingular{workorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (woq *WorkOrderQuery) OnlyX(ctx context.Context) *WorkOrder {
	wo, err := woq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return wo
}

// OnlyID returns the only WorkOrder id in the query, returns an error if not exactly one id was returned.
func (woq *WorkOrderQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = woq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{workorder.Label}
	default:
		err = &ErrNotSingular{workorder.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (woq *WorkOrderQuery) OnlyXID(ctx context.Context) string {
	id, err := woq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkOrders.
func (woq *WorkOrderQuery) All(ctx context.Context) ([]*WorkOrder, error) {
	return woq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (woq *WorkOrderQuery) AllX(ctx context.Context) []*WorkOrder {
	wos, err := woq.All(ctx)
	if err != nil {
		panic(err)
	}
	return wos
}

// IDs executes the query and returns a list of WorkOrder ids.
func (woq *WorkOrderQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := woq.Select(workorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (woq *WorkOrderQuery) IDsX(ctx context.Context) []string {
	ids, err := woq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (woq *WorkOrderQuery) Count(ctx context.Context) (int, error) {
	return woq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (woq *WorkOrderQuery) CountX(ctx context.Context) int {
	count, err := woq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (woq *WorkOrderQuery) Exist(ctx context.Context) (bool, error) {
	return woq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (woq *WorkOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := woq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (woq *WorkOrderQuery) Clone() *WorkOrderQuery {
	return &WorkOrderQuery{
		config:     woq.config,
		limit:      woq.limit,
		offset:     woq.offset,
		order:      append([]Order{}, woq.order...),
		unique:     append([]string{}, woq.unique...),
		predicates: append([]predicate.WorkOrder{}, woq.predicates...),
		// clone intermediate queries.
		sql: woq.sql.Clone(),
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
//	client.WorkOrder.Query().
//		GroupBy(workorder.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (woq *WorkOrderQuery) GroupBy(field string, fields ...string) *WorkOrderGroupBy {
	group := &WorkOrderGroupBy{config: woq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = woq.sqlQuery()
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
//	client.WorkOrder.Query().
//		Select(workorder.FieldCreateTime).
//		Scan(ctx, &v)
//
func (woq *WorkOrderQuery) Select(field string, fields ...string) *WorkOrderSelect {
	selector := &WorkOrderSelect{config: woq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = woq.sqlQuery()
	return selector
}

func (woq *WorkOrderQuery) sqlAll(ctx context.Context) ([]*WorkOrder, error) {
	rows := &sql.Rows{}
	selector := woq.sqlQuery()
	if unique := woq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := woq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var wos WorkOrders
	if err := wos.FromRows(rows); err != nil {
		return nil, err
	}
	wos.config(woq.config)
	return wos, nil
}

func (woq *WorkOrderQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := woq.sqlQuery()
	unique := []string{workorder.FieldID}
	if len(woq.unique) > 0 {
		unique = woq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := woq.driver.Query(ctx, query, args, rows); err != nil {
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

func (woq *WorkOrderQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := woq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (woq *WorkOrderQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(woq.driver.Dialect())
	t1 := builder.Table(workorder.Table)
	selector := builder.Select(t1.Columns(workorder.Columns...)...).From(t1)
	if woq.sql != nil {
		selector = woq.sql
		selector.Select(selector.Columns(workorder.Columns...)...)
	}
	for _, p := range woq.predicates {
		p(selector)
	}
	for _, p := range woq.order {
		p(selector)
	}
	if offset := woq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := woq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkOrderGroupBy is the builder for group-by WorkOrder entities.
type WorkOrderGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wogb *WorkOrderGroupBy) Aggregate(fns ...Aggregate) *WorkOrderGroupBy {
	wogb.fns = append(wogb.fns, fns...)
	return wogb
}

// Scan applies the group-by query and scan the result into the given value.
func (wogb *WorkOrderGroupBy) Scan(ctx context.Context, v interface{}) error {
	return wogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wogb *WorkOrderGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wogb *WorkOrderGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wogb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wogb *WorkOrderGroupBy) StringsX(ctx context.Context) []string {
	v, err := wogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wogb *WorkOrderGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wogb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wogb *WorkOrderGroupBy) IntsX(ctx context.Context) []int {
	v, err := wogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wogb *WorkOrderGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wogb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wogb *WorkOrderGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wogb *WorkOrderGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wogb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wogb *WorkOrderGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wogb *WorkOrderGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wogb.sqlQuery().Query()
	if err := wogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wogb *WorkOrderGroupBy) sqlQuery() *sql.Selector {
	selector := wogb.sql
	columns := make([]string, 0, len(wogb.fields)+len(wogb.fns))
	columns = append(columns, wogb.fields...)
	for _, fn := range wogb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(wogb.fields...)
}

// WorkOrderSelect is the builder for select fields of WorkOrder entities.
type WorkOrderSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (wos *WorkOrderSelect) Scan(ctx context.Context, v interface{}) error {
	return wos.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wos *WorkOrderSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wos.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wos *WorkOrderSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wos.fields) > 1 {
		return nil, errors.New("ent: WorkOrderSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wos *WorkOrderSelect) StringsX(ctx context.Context) []string {
	v, err := wos.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wos *WorkOrderSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wos.fields) > 1 {
		return nil, errors.New("ent: WorkOrderSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wos *WorkOrderSelect) IntsX(ctx context.Context) []int {
	v, err := wos.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wos *WorkOrderSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wos.fields) > 1 {
		return nil, errors.New("ent: WorkOrderSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wos *WorkOrderSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wos.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wos *WorkOrderSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wos.fields) > 1 {
		return nil, errors.New("ent: WorkOrderSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wos *WorkOrderSelect) BoolsX(ctx context.Context) []bool {
	v, err := wos.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wos *WorkOrderSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wos.sqlQuery().Query()
	if err := wos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wos *WorkOrderSelect) sqlQuery() sql.Querier {
	view := "workorder_view"
	return sql.Dialect(wos.driver.Dialect()).
		Select(wos.fields...).From(wos.sql.As(view))
}
