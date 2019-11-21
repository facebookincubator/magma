// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/location"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/project"
	"github.com/facebookincubator/symphony/graph/ent/projecttype"
	"github.com/facebookincubator/symphony/graph/ent/property"
	"github.com/facebookincubator/symphony/graph/ent/workorder"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config

	update_time       *time.Time
	name              *string
	description       *string
	cleardescription  bool
	creator           *string
	clearcreator      bool
	_type             map[string]struct{}
	location          map[string]struct{}
	work_orders       map[string]struct{}
	properties        map[string]struct{}
	clearedType       bool
	clearedLocation   bool
	removedWorkOrders map[string]struct{}
	removedProperties map[string]struct{}
	predicates        []predicate.Project
}

// Where adds a new predicate for the builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.name = &s
	return pu
}

// SetDescription sets the description field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.description = &s
	return pu
}

// SetNillableDescription sets the description field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDescription(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of description.
func (pu *ProjectUpdate) ClearDescription() *ProjectUpdate {
	pu.description = nil
	pu.cleardescription = true
	return pu
}

// SetCreator sets the creator field.
func (pu *ProjectUpdate) SetCreator(s string) *ProjectUpdate {
	pu.creator = &s
	return pu
}

// SetNillableCreator sets the creator field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCreator(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetCreator(*s)
	}
	return pu
}

// ClearCreator clears the value of creator.
func (pu *ProjectUpdate) ClearCreator() *ProjectUpdate {
	pu.creator = nil
	pu.clearcreator = true
	return pu
}

// SetTypeID sets the type edge to ProjectType by id.
func (pu *ProjectUpdate) SetTypeID(id string) *ProjectUpdate {
	if pu._type == nil {
		pu._type = make(map[string]struct{})
	}
	pu._type[id] = struct{}{}
	return pu
}

// SetType sets the type edge to ProjectType.
func (pu *ProjectUpdate) SetType(p *ProjectType) *ProjectUpdate {
	return pu.SetTypeID(p.ID)
}

// SetLocationID sets the location edge to Location by id.
func (pu *ProjectUpdate) SetLocationID(id string) *ProjectUpdate {
	if pu.location == nil {
		pu.location = make(map[string]struct{})
	}
	pu.location[id] = struct{}{}
	return pu
}

// SetNillableLocationID sets the location edge to Location by id if the given value is not nil.
func (pu *ProjectUpdate) SetNillableLocationID(id *string) *ProjectUpdate {
	if id != nil {
		pu = pu.SetLocationID(*id)
	}
	return pu
}

// SetLocation sets the location edge to Location.
func (pu *ProjectUpdate) SetLocation(l *Location) *ProjectUpdate {
	return pu.SetLocationID(l.ID)
}

// AddWorkOrderIDs adds the work_orders edge to WorkOrder by ids.
func (pu *ProjectUpdate) AddWorkOrderIDs(ids ...string) *ProjectUpdate {
	if pu.work_orders == nil {
		pu.work_orders = make(map[string]struct{})
	}
	for i := range ids {
		pu.work_orders[ids[i]] = struct{}{}
	}
	return pu
}

// AddWorkOrders adds the work_orders edges to WorkOrder.
func (pu *ProjectUpdate) AddWorkOrders(w ...*WorkOrder) *ProjectUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pu.AddWorkOrderIDs(ids...)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (pu *ProjectUpdate) AddPropertyIDs(ids ...string) *ProjectUpdate {
	if pu.properties == nil {
		pu.properties = make(map[string]struct{})
	}
	for i := range ids {
		pu.properties[ids[i]] = struct{}{}
	}
	return pu
}

// AddProperties adds the properties edges to Property.
func (pu *ProjectUpdate) AddProperties(p ...*Property) *ProjectUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPropertyIDs(ids...)
}

// ClearType clears the type edge to ProjectType.
func (pu *ProjectUpdate) ClearType() *ProjectUpdate {
	pu.clearedType = true
	return pu
}

// ClearLocation clears the location edge to Location.
func (pu *ProjectUpdate) ClearLocation() *ProjectUpdate {
	pu.clearedLocation = true
	return pu
}

// RemoveWorkOrderIDs removes the work_orders edge to WorkOrder by ids.
func (pu *ProjectUpdate) RemoveWorkOrderIDs(ids ...string) *ProjectUpdate {
	if pu.removedWorkOrders == nil {
		pu.removedWorkOrders = make(map[string]struct{})
	}
	for i := range ids {
		pu.removedWorkOrders[ids[i]] = struct{}{}
	}
	return pu
}

// RemoveWorkOrders removes work_orders edges to WorkOrder.
func (pu *ProjectUpdate) RemoveWorkOrders(w ...*WorkOrder) *ProjectUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pu.RemoveWorkOrderIDs(ids...)
}

// RemovePropertyIDs removes the properties edge to Property by ids.
func (pu *ProjectUpdate) RemovePropertyIDs(ids ...string) *ProjectUpdate {
	if pu.removedProperties == nil {
		pu.removedProperties = make(map[string]struct{})
	}
	for i := range ids {
		pu.removedProperties[ids[i]] = struct{}{}
	}
	return pu
}

// RemoveProperties removes properties edges to Property.
func (pu *ProjectUpdate) RemoveProperties(p ...*Property) *ProjectUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePropertyIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	if pu.update_time == nil {
		v := project.UpdateDefaultUpdateTime()
		pu.update_time = &v
	}
	if pu.name != nil {
		if err := project.NameValidator(*pu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(pu._type) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"type\"")
	}
	if pu.clearedType && pu._type == nil {
		return 0, errors.New("ent: clearing a unique edge \"type\"")
	}
	if len(pu.location) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"location\"")
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(pu.driver.Dialect())
		selector = builder.Select(project.FieldID).From(builder.Table(project.Table))
	)
	for _, p := range pu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = pu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := pu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(project.Table).Where(sql.InInts(project.FieldID, ids...))
	)
	if value := pu.update_time; value != nil {
		updater.Set(project.FieldUpdateTime, *value)
	}
	if value := pu.name; value != nil {
		updater.Set(project.FieldName, *value)
	}
	if value := pu.description; value != nil {
		updater.Set(project.FieldDescription, *value)
	}
	if pu.cleardescription {
		updater.SetNull(project.FieldDescription)
	}
	if value := pu.creator; value != nil {
		updater.Set(project.FieldCreator, *value)
	}
	if pu.clearcreator {
		updater.SetNull(project.FieldCreator)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if pu.clearedType {
		query, args := builder.Update(project.TypeTable).
			SetNull(project.TypeColumn).
			Where(sql.InInts(projecttype.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu._type) > 0 {
		for eid := range pu._type {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(project.TypeTable).
				Set(project.TypeColumn, eid).
				Where(sql.InInts(project.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if pu.clearedLocation {
		query, args := builder.Update(project.LocationTable).
			SetNull(project.LocationColumn).
			Where(sql.InInts(location.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.location) > 0 {
		for eid := range pu.location {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(project.LocationTable).
				Set(project.LocationColumn, eid).
				Where(sql.InInts(project.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if len(pu.removedWorkOrders) > 0 {
		eids := make([]int, len(pu.removedWorkOrders))
		for eid := range pu.removedWorkOrders {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(project.WorkOrdersTable).
			SetNull(project.WorkOrdersColumn).
			Where(sql.InInts(project.WorkOrdersColumn, ids...)).
			Where(sql.InInts(workorder.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.work_orders) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range pu.work_orders {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(workorder.FieldID, eid)
			}
			query, args := builder.Update(project.WorkOrdersTable).
				Set(project.WorkOrdersColumn, id).
				Where(sql.And(p, sql.IsNull(project.WorkOrdersColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(pu.work_orders) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"work_orders\" %v already connected to a different \"Project\"", keys(pu.work_orders))})
			}
		}
	}
	if len(pu.removedProperties) > 0 {
		eids := make([]int, len(pu.removedProperties))
		for eid := range pu.removedProperties {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(project.PropertiesTable).
			SetNull(project.PropertiesColumn).
			Where(sql.InInts(project.PropertiesColumn, ids...)).
			Where(sql.InInts(property.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.properties) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range pu.properties {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(property.FieldID, eid)
			}
			query, args := builder.Update(project.PropertiesTable).
				Set(project.PropertiesColumn, id).
				Where(sql.And(p, sql.IsNull(project.PropertiesColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(pu.properties) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"properties\" %v already connected to a different \"Project\"", keys(pu.properties))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	id string

	update_time       *time.Time
	name              *string
	description       *string
	cleardescription  bool
	creator           *string
	clearcreator      bool
	_type             map[string]struct{}
	location          map[string]struct{}
	work_orders       map[string]struct{}
	properties        map[string]struct{}
	clearedType       bool
	clearedLocation   bool
	removedWorkOrders map[string]struct{}
	removedProperties map[string]struct{}
}

// SetName sets the name field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.name = &s
	return puo
}

// SetDescription sets the description field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.description = &s
	return puo
}

// SetNillableDescription sets the description field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDescription(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of description.
func (puo *ProjectUpdateOne) ClearDescription() *ProjectUpdateOne {
	puo.description = nil
	puo.cleardescription = true
	return puo
}

// SetCreator sets the creator field.
func (puo *ProjectUpdateOne) SetCreator(s string) *ProjectUpdateOne {
	puo.creator = &s
	return puo
}

// SetNillableCreator sets the creator field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCreator(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetCreator(*s)
	}
	return puo
}

// ClearCreator clears the value of creator.
func (puo *ProjectUpdateOne) ClearCreator() *ProjectUpdateOne {
	puo.creator = nil
	puo.clearcreator = true
	return puo
}

// SetTypeID sets the type edge to ProjectType by id.
func (puo *ProjectUpdateOne) SetTypeID(id string) *ProjectUpdateOne {
	if puo._type == nil {
		puo._type = make(map[string]struct{})
	}
	puo._type[id] = struct{}{}
	return puo
}

// SetType sets the type edge to ProjectType.
func (puo *ProjectUpdateOne) SetType(p *ProjectType) *ProjectUpdateOne {
	return puo.SetTypeID(p.ID)
}

// SetLocationID sets the location edge to Location by id.
func (puo *ProjectUpdateOne) SetLocationID(id string) *ProjectUpdateOne {
	if puo.location == nil {
		puo.location = make(map[string]struct{})
	}
	puo.location[id] = struct{}{}
	return puo
}

// SetNillableLocationID sets the location edge to Location by id if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableLocationID(id *string) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetLocationID(*id)
	}
	return puo
}

// SetLocation sets the location edge to Location.
func (puo *ProjectUpdateOne) SetLocation(l *Location) *ProjectUpdateOne {
	return puo.SetLocationID(l.ID)
}

// AddWorkOrderIDs adds the work_orders edge to WorkOrder by ids.
func (puo *ProjectUpdateOne) AddWorkOrderIDs(ids ...string) *ProjectUpdateOne {
	if puo.work_orders == nil {
		puo.work_orders = make(map[string]struct{})
	}
	for i := range ids {
		puo.work_orders[ids[i]] = struct{}{}
	}
	return puo
}

// AddWorkOrders adds the work_orders edges to WorkOrder.
func (puo *ProjectUpdateOne) AddWorkOrders(w ...*WorkOrder) *ProjectUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return puo.AddWorkOrderIDs(ids...)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (puo *ProjectUpdateOne) AddPropertyIDs(ids ...string) *ProjectUpdateOne {
	if puo.properties == nil {
		puo.properties = make(map[string]struct{})
	}
	for i := range ids {
		puo.properties[ids[i]] = struct{}{}
	}
	return puo
}

// AddProperties adds the properties edges to Property.
func (puo *ProjectUpdateOne) AddProperties(p ...*Property) *ProjectUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPropertyIDs(ids...)
}

// ClearType clears the type edge to ProjectType.
func (puo *ProjectUpdateOne) ClearType() *ProjectUpdateOne {
	puo.clearedType = true
	return puo
}

// ClearLocation clears the location edge to Location.
func (puo *ProjectUpdateOne) ClearLocation() *ProjectUpdateOne {
	puo.clearedLocation = true
	return puo
}

// RemoveWorkOrderIDs removes the work_orders edge to WorkOrder by ids.
func (puo *ProjectUpdateOne) RemoveWorkOrderIDs(ids ...string) *ProjectUpdateOne {
	if puo.removedWorkOrders == nil {
		puo.removedWorkOrders = make(map[string]struct{})
	}
	for i := range ids {
		puo.removedWorkOrders[ids[i]] = struct{}{}
	}
	return puo
}

// RemoveWorkOrders removes work_orders edges to WorkOrder.
func (puo *ProjectUpdateOne) RemoveWorkOrders(w ...*WorkOrder) *ProjectUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return puo.RemoveWorkOrderIDs(ids...)
}

// RemovePropertyIDs removes the properties edge to Property by ids.
func (puo *ProjectUpdateOne) RemovePropertyIDs(ids ...string) *ProjectUpdateOne {
	if puo.removedProperties == nil {
		puo.removedProperties = make(map[string]struct{})
	}
	for i := range ids {
		puo.removedProperties[ids[i]] = struct{}{}
	}
	return puo
}

// RemoveProperties removes properties edges to Property.
func (puo *ProjectUpdateOne) RemoveProperties(p ...*Property) *ProjectUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePropertyIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	if puo.update_time == nil {
		v := project.UpdateDefaultUpdateTime()
		puo.update_time = &v
	}
	if puo.name != nil {
		if err := project.NameValidator(*puo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(puo._type) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"type\"")
	}
	if puo.clearedType && puo._type == nil {
		return nil, errors.New("ent: clearing a unique edge \"type\"")
	}
	if len(puo.location) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"location\"")
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (pr *Project, err error) {
	var (
		builder  = sql.Dialect(puo.driver.Dialect())
		selector = builder.Select(project.Columns...).From(builder.Table(project.Table))
	)
	project.ID(puo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = puo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		pr = &Project{config: puo.config}
		if err := pr.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Project: %v", err)
		}
		id = pr.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Project with id: %v", puo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Project with the same id: %v", puo.id)
	}

	tx, err := puo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(project.Table).Where(sql.InInts(project.FieldID, ids...))
	)
	if value := puo.update_time; value != nil {
		updater.Set(project.FieldUpdateTime, *value)
		pr.UpdateTime = *value
	}
	if value := puo.name; value != nil {
		updater.Set(project.FieldName, *value)
		pr.Name = *value
	}
	if value := puo.description; value != nil {
		updater.Set(project.FieldDescription, *value)
		pr.Description = value
	}
	if puo.cleardescription {
		pr.Description = nil
		updater.SetNull(project.FieldDescription)
	}
	if value := puo.creator; value != nil {
		updater.Set(project.FieldCreator, *value)
		pr.Creator = value
	}
	if puo.clearcreator {
		pr.Creator = nil
		updater.SetNull(project.FieldCreator)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if puo.clearedType {
		query, args := builder.Update(project.TypeTable).
			SetNull(project.TypeColumn).
			Where(sql.InInts(projecttype.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo._type) > 0 {
		for eid := range puo._type {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(project.TypeTable).
				Set(project.TypeColumn, eid).
				Where(sql.InInts(project.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if puo.clearedLocation {
		query, args := builder.Update(project.LocationTable).
			SetNull(project.LocationColumn).
			Where(sql.InInts(location.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.location) > 0 {
		for eid := range puo.location {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(project.LocationTable).
				Set(project.LocationColumn, eid).
				Where(sql.InInts(project.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(puo.removedWorkOrders) > 0 {
		eids := make([]int, len(puo.removedWorkOrders))
		for eid := range puo.removedWorkOrders {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(project.WorkOrdersTable).
			SetNull(project.WorkOrdersColumn).
			Where(sql.InInts(project.WorkOrdersColumn, ids...)).
			Where(sql.InInts(workorder.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.work_orders) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range puo.work_orders {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(workorder.FieldID, eid)
			}
			query, args := builder.Update(project.WorkOrdersTable).
				Set(project.WorkOrdersColumn, id).
				Where(sql.And(p, sql.IsNull(project.WorkOrdersColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(puo.work_orders) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"work_orders\" %v already connected to a different \"Project\"", keys(puo.work_orders))})
			}
		}
	}
	if len(puo.removedProperties) > 0 {
		eids := make([]int, len(puo.removedProperties))
		for eid := range puo.removedProperties {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(project.PropertiesTable).
			SetNull(project.PropertiesColumn).
			Where(sql.InInts(project.PropertiesColumn, ids...)).
			Where(sql.InInts(property.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.properties) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range puo.properties {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(property.FieldID, eid)
			}
			query, args := builder.Update(project.PropertiesTable).
				Set(project.PropertiesColumn, id).
				Where(sql.And(p, sql.IsNull(project.PropertiesColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(puo.properties) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"properties\" %v already connected to a different \"Project\"", keys(puo.properties))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return pr, nil
}
