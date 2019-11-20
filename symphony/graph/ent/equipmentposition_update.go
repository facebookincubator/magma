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
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentposition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentpositiondefinition"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// EquipmentPositionUpdate is the builder for updating EquipmentPosition entities.
type EquipmentPositionUpdate struct {
	config

	update_time       *time.Time
	definition        map[string]struct{}
	parent            map[string]struct{}
	attachment        map[string]struct{}
	clearedDefinition bool
	clearedParent     bool
	clearedAttachment bool
	predicates        []predicate.EquipmentPosition
}

// Where adds a new predicate for the builder.
func (epu *EquipmentPositionUpdate) Where(ps ...predicate.EquipmentPosition) *EquipmentPositionUpdate {
	epu.predicates = append(epu.predicates, ps...)
	return epu
}

// SetDefinitionID sets the definition edge to EquipmentPositionDefinition by id.
func (epu *EquipmentPositionUpdate) SetDefinitionID(id string) *EquipmentPositionUpdate {
	if epu.definition == nil {
		epu.definition = make(map[string]struct{})
	}
	epu.definition[id] = struct{}{}
	return epu
}

// SetDefinition sets the definition edge to EquipmentPositionDefinition.
func (epu *EquipmentPositionUpdate) SetDefinition(e *EquipmentPositionDefinition) *EquipmentPositionUpdate {
	return epu.SetDefinitionID(e.ID)
}

// SetParentID sets the parent edge to Equipment by id.
func (epu *EquipmentPositionUpdate) SetParentID(id string) *EquipmentPositionUpdate {
	if epu.parent == nil {
		epu.parent = make(map[string]struct{})
	}
	epu.parent[id] = struct{}{}
	return epu
}

// SetNillableParentID sets the parent edge to Equipment by id if the given value is not nil.
func (epu *EquipmentPositionUpdate) SetNillableParentID(id *string) *EquipmentPositionUpdate {
	if id != nil {
		epu = epu.SetParentID(*id)
	}
	return epu
}

// SetParent sets the parent edge to Equipment.
func (epu *EquipmentPositionUpdate) SetParent(e *Equipment) *EquipmentPositionUpdate {
	return epu.SetParentID(e.ID)
}

// SetAttachmentID sets the attachment edge to Equipment by id.
func (epu *EquipmentPositionUpdate) SetAttachmentID(id string) *EquipmentPositionUpdate {
	if epu.attachment == nil {
		epu.attachment = make(map[string]struct{})
	}
	epu.attachment[id] = struct{}{}
	return epu
}

// SetNillableAttachmentID sets the attachment edge to Equipment by id if the given value is not nil.
func (epu *EquipmentPositionUpdate) SetNillableAttachmentID(id *string) *EquipmentPositionUpdate {
	if id != nil {
		epu = epu.SetAttachmentID(*id)
	}
	return epu
}

// SetAttachment sets the attachment edge to Equipment.
func (epu *EquipmentPositionUpdate) SetAttachment(e *Equipment) *EquipmentPositionUpdate {
	return epu.SetAttachmentID(e.ID)
}

// ClearDefinition clears the definition edge to EquipmentPositionDefinition.
func (epu *EquipmentPositionUpdate) ClearDefinition() *EquipmentPositionUpdate {
	epu.clearedDefinition = true
	return epu
}

// ClearParent clears the parent edge to Equipment.
func (epu *EquipmentPositionUpdate) ClearParent() *EquipmentPositionUpdate {
	epu.clearedParent = true
	return epu
}

// ClearAttachment clears the attachment edge to Equipment.
func (epu *EquipmentPositionUpdate) ClearAttachment() *EquipmentPositionUpdate {
	epu.clearedAttachment = true
	return epu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (epu *EquipmentPositionUpdate) Save(ctx context.Context) (int, error) {
	if epu.update_time == nil {
		v := equipmentposition.UpdateDefaultUpdateTime()
		epu.update_time = &v
	}
	if len(epu.definition) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"definition\"")
	}
	if epu.clearedDefinition && epu.definition == nil {
		return 0, errors.New("ent: clearing a unique edge \"definition\"")
	}
	if len(epu.parent) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	if len(epu.attachment) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"attachment\"")
	}
	return epu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (epu *EquipmentPositionUpdate) SaveX(ctx context.Context) int {
	affected, err := epu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (epu *EquipmentPositionUpdate) Exec(ctx context.Context) error {
	_, err := epu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epu *EquipmentPositionUpdate) ExecX(ctx context.Context) {
	if err := epu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (epu *EquipmentPositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(epu.driver.Dialect())
		selector = builder.Select(equipmentposition.FieldID).From(builder.Table(equipmentposition.Table))
	)
	for _, p := range epu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = epu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := epu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(equipmentposition.Table).Where(sql.InInts(equipmentposition.FieldID, ids...))
	)
	if value := epu.update_time; value != nil {
		updater.Set(equipmentposition.FieldUpdateTime, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if epu.clearedDefinition {
		query, args := builder.Update(equipmentposition.DefinitionTable).
			SetNull(equipmentposition.DefinitionColumn).
			Where(sql.InInts(equipmentpositiondefinition.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(epu.definition) > 0 {
		for eid := range epu.definition {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(equipmentposition.DefinitionTable).
				Set(equipmentposition.DefinitionColumn, eid).
				Where(sql.InInts(equipmentposition.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if epu.clearedParent {
		query, args := builder.Update(equipmentposition.ParentTable).
			SetNull(equipmentposition.ParentColumn).
			Where(sql.InInts(equipment.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(epu.parent) > 0 {
		for eid := range epu.parent {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(equipmentposition.ParentTable).
				Set(equipmentposition.ParentColumn, eid).
				Where(sql.InInts(equipmentposition.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if epu.clearedAttachment {
		query, args := builder.Update(equipmentposition.AttachmentTable).
			SetNull(equipmentposition.AttachmentColumn).
			Where(sql.InInts(equipment.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(epu.attachment) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(epu.attachment)[0])
			if serr != nil {
				return 0, rollback(tx, err)
			}
			query, args := builder.Update(equipmentposition.AttachmentTable).
				Set(equipmentposition.AttachmentColumn, id).
				Where(sql.EQ(equipment.FieldID, eid).And().IsNull(equipmentposition.AttachmentColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(epu.attachment) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"attachment\" %v already connected to a different \"EquipmentPosition\"", keys(epu.attachment))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// EquipmentPositionUpdateOne is the builder for updating a single EquipmentPosition entity.
type EquipmentPositionUpdateOne struct {
	config
	id string

	update_time       *time.Time
	definition        map[string]struct{}
	parent            map[string]struct{}
	attachment        map[string]struct{}
	clearedDefinition bool
	clearedParent     bool
	clearedAttachment bool
}

// SetDefinitionID sets the definition edge to EquipmentPositionDefinition by id.
func (epuo *EquipmentPositionUpdateOne) SetDefinitionID(id string) *EquipmentPositionUpdateOne {
	if epuo.definition == nil {
		epuo.definition = make(map[string]struct{})
	}
	epuo.definition[id] = struct{}{}
	return epuo
}

// SetDefinition sets the definition edge to EquipmentPositionDefinition.
func (epuo *EquipmentPositionUpdateOne) SetDefinition(e *EquipmentPositionDefinition) *EquipmentPositionUpdateOne {
	return epuo.SetDefinitionID(e.ID)
}

// SetParentID sets the parent edge to Equipment by id.
func (epuo *EquipmentPositionUpdateOne) SetParentID(id string) *EquipmentPositionUpdateOne {
	if epuo.parent == nil {
		epuo.parent = make(map[string]struct{})
	}
	epuo.parent[id] = struct{}{}
	return epuo
}

// SetNillableParentID sets the parent edge to Equipment by id if the given value is not nil.
func (epuo *EquipmentPositionUpdateOne) SetNillableParentID(id *string) *EquipmentPositionUpdateOne {
	if id != nil {
		epuo = epuo.SetParentID(*id)
	}
	return epuo
}

// SetParent sets the parent edge to Equipment.
func (epuo *EquipmentPositionUpdateOne) SetParent(e *Equipment) *EquipmentPositionUpdateOne {
	return epuo.SetParentID(e.ID)
}

// SetAttachmentID sets the attachment edge to Equipment by id.
func (epuo *EquipmentPositionUpdateOne) SetAttachmentID(id string) *EquipmentPositionUpdateOne {
	if epuo.attachment == nil {
		epuo.attachment = make(map[string]struct{})
	}
	epuo.attachment[id] = struct{}{}
	return epuo
}

// SetNillableAttachmentID sets the attachment edge to Equipment by id if the given value is not nil.
func (epuo *EquipmentPositionUpdateOne) SetNillableAttachmentID(id *string) *EquipmentPositionUpdateOne {
	if id != nil {
		epuo = epuo.SetAttachmentID(*id)
	}
	return epuo
}

// SetAttachment sets the attachment edge to Equipment.
func (epuo *EquipmentPositionUpdateOne) SetAttachment(e *Equipment) *EquipmentPositionUpdateOne {
	return epuo.SetAttachmentID(e.ID)
}

// ClearDefinition clears the definition edge to EquipmentPositionDefinition.
func (epuo *EquipmentPositionUpdateOne) ClearDefinition() *EquipmentPositionUpdateOne {
	epuo.clearedDefinition = true
	return epuo
}

// ClearParent clears the parent edge to Equipment.
func (epuo *EquipmentPositionUpdateOne) ClearParent() *EquipmentPositionUpdateOne {
	epuo.clearedParent = true
	return epuo
}

// ClearAttachment clears the attachment edge to Equipment.
func (epuo *EquipmentPositionUpdateOne) ClearAttachment() *EquipmentPositionUpdateOne {
	epuo.clearedAttachment = true
	return epuo
}

// Save executes the query and returns the updated entity.
func (epuo *EquipmentPositionUpdateOne) Save(ctx context.Context) (*EquipmentPosition, error) {
	if epuo.update_time == nil {
		v := equipmentposition.UpdateDefaultUpdateTime()
		epuo.update_time = &v
	}
	if len(epuo.definition) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"definition\"")
	}
	if epuo.clearedDefinition && epuo.definition == nil {
		return nil, errors.New("ent: clearing a unique edge \"definition\"")
	}
	if len(epuo.parent) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	if len(epuo.attachment) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"attachment\"")
	}
	return epuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (epuo *EquipmentPositionUpdateOne) SaveX(ctx context.Context) *EquipmentPosition {
	ep, err := epuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ep
}

// Exec executes the query on the entity.
func (epuo *EquipmentPositionUpdateOne) Exec(ctx context.Context) error {
	_, err := epuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epuo *EquipmentPositionUpdateOne) ExecX(ctx context.Context) {
	if err := epuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (epuo *EquipmentPositionUpdateOne) sqlSave(ctx context.Context) (ep *EquipmentPosition, err error) {
	var (
		builder  = sql.Dialect(epuo.driver.Dialect())
		selector = builder.Select(equipmentposition.Columns...).From(builder.Table(equipmentposition.Table))
	)
	equipmentposition.ID(epuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = epuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		ep = &EquipmentPosition{config: epuo.config}
		if err := ep.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into EquipmentPosition: %v", err)
		}
		id = ep.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("EquipmentPosition with id: %v", epuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one EquipmentPosition with the same id: %v", epuo.id)
	}

	tx, err := epuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(equipmentposition.Table).Where(sql.InInts(equipmentposition.FieldID, ids...))
	)
	if value := epuo.update_time; value != nil {
		updater.Set(equipmentposition.FieldUpdateTime, *value)
		ep.UpdateTime = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if epuo.clearedDefinition {
		query, args := builder.Update(equipmentposition.DefinitionTable).
			SetNull(equipmentposition.DefinitionColumn).
			Where(sql.InInts(equipmentpositiondefinition.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(epuo.definition) > 0 {
		for eid := range epuo.definition {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(equipmentposition.DefinitionTable).
				Set(equipmentposition.DefinitionColumn, eid).
				Where(sql.InInts(equipmentposition.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if epuo.clearedParent {
		query, args := builder.Update(equipmentposition.ParentTable).
			SetNull(equipmentposition.ParentColumn).
			Where(sql.InInts(equipment.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(epuo.parent) > 0 {
		for eid := range epuo.parent {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(equipmentposition.ParentTable).
				Set(equipmentposition.ParentColumn, eid).
				Where(sql.InInts(equipmentposition.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if epuo.clearedAttachment {
		query, args := builder.Update(equipmentposition.AttachmentTable).
			SetNull(equipmentposition.AttachmentColumn).
			Where(sql.InInts(equipment.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(epuo.attachment) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(epuo.attachment)[0])
			if serr != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(equipmentposition.AttachmentTable).
				Set(equipmentposition.AttachmentColumn, id).
				Where(sql.EQ(equipment.FieldID, eid).And().IsNull(equipmentposition.AttachmentColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(epuo.attachment) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"attachment\" %v already connected to a different \"EquipmentPosition\"", keys(epuo.attachment))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return ep, nil
}
