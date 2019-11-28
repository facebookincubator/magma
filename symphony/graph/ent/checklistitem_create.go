// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/checklistitem"
)

// CheckListItemCreate is the builder for creating a CheckListItem entity.
type CheckListItemCreate struct {
	config
	title       *string
	_type       *string
	index       *int
	checked     *bool
	string_val  *string
	enum_values *string
	help_text   *string
	work_order  map[string]struct{}
}

// SetTitle sets the title field.
func (clic *CheckListItemCreate) SetTitle(s string) *CheckListItemCreate {
	clic.title = &s
	return clic
}

// SetType sets the type field.
func (clic *CheckListItemCreate) SetType(s string) *CheckListItemCreate {
	clic._type = &s
	return clic
}

// SetIndex sets the index field.
func (clic *CheckListItemCreate) SetIndex(i int) *CheckListItemCreate {
	clic.index = &i
	return clic
}

// SetNillableIndex sets the index field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableIndex(i *int) *CheckListItemCreate {
	if i != nil {
		clic.SetIndex(*i)
	}
	return clic
}

// SetChecked sets the checked field.
func (clic *CheckListItemCreate) SetChecked(b bool) *CheckListItemCreate {
	clic.checked = &b
	return clic
}

// SetNillableChecked sets the checked field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableChecked(b *bool) *CheckListItemCreate {
	if b != nil {
		clic.SetChecked(*b)
	}
	return clic
}

// SetStringVal sets the string_val field.
func (clic *CheckListItemCreate) SetStringVal(s string) *CheckListItemCreate {
	clic.string_val = &s
	return clic
}

// SetNillableStringVal sets the string_val field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableStringVal(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetStringVal(*s)
	}
	return clic
}

// SetEnumValues sets the enum_values field.
func (clic *CheckListItemCreate) SetEnumValues(s string) *CheckListItemCreate {
	clic.enum_values = &s
	return clic
}

// SetNillableEnumValues sets the enum_values field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableEnumValues(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetEnumValues(*s)
	}
	return clic
}

// SetHelpText sets the help_text field.
func (clic *CheckListItemCreate) SetHelpText(s string) *CheckListItemCreate {
	clic.help_text = &s
	return clic
}

// SetNillableHelpText sets the help_text field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableHelpText(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetHelpText(*s)
	}
	return clic
}

// SetWorkOrderID sets the work_order edge to WorkOrder by id.
func (clic *CheckListItemCreate) SetWorkOrderID(id string) *CheckListItemCreate {
	if clic.work_order == nil {
		clic.work_order = make(map[string]struct{})
	}
	clic.work_order[id] = struct{}{}
	return clic
}

// SetNillableWorkOrderID sets the work_order edge to WorkOrder by id if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableWorkOrderID(id *string) *CheckListItemCreate {
	if id != nil {
		clic = clic.SetWorkOrderID(*id)
	}
	return clic
}

// SetWorkOrder sets the work_order edge to WorkOrder.
func (clic *CheckListItemCreate) SetWorkOrder(w *WorkOrder) *CheckListItemCreate {
	return clic.SetWorkOrderID(w.ID)
}

// Save creates the CheckListItem in the database.
func (clic *CheckListItemCreate) Save(ctx context.Context) (*CheckListItem, error) {
	if clic.title == nil {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if clic._type == nil {
		return nil, errors.New("ent: missing required field \"type\"")
	}
	if len(clic.work_order) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"work_order\"")
	}
	return clic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (clic *CheckListItemCreate) SaveX(ctx context.Context) *CheckListItem {
	v, err := clic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (clic *CheckListItemCreate) sqlSave(ctx context.Context) (*CheckListItem, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(clic.driver.Dialect())
		cli     = &CheckListItem{config: clic.config}
	)
	tx, err := clic.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(checklistitem.Table).Default()
	if value := clic.title; value != nil {
		insert.Set(checklistitem.FieldTitle, *value)
		cli.Title = *value
	}
	if value := clic._type; value != nil {
		insert.Set(checklistitem.FieldType, *value)
		cli.Type = *value
	}
	if value := clic.index; value != nil {
		insert.Set(checklistitem.FieldIndex, *value)
		cli.Index = *value
	}
	if value := clic.checked; value != nil {
		insert.Set(checklistitem.FieldChecked, *value)
		cli.Checked = *value
	}
	if value := clic.string_val; value != nil {
		insert.Set(checklistitem.FieldStringVal, *value)
		cli.StringVal = *value
	}
	if value := clic.enum_values; value != nil {
		insert.Set(checklistitem.FieldEnumValues, *value)
		cli.EnumValues = *value
	}
	if value := clic.help_text; value != nil {
		insert.Set(checklistitem.FieldHelpText, *value)
		cli.HelpText = value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(checklistitem.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	cli.ID = strconv.FormatInt(id, 10)
	if len(clic.work_order) > 0 {
		for eid := range clic.work_order {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(checklistitem.WorkOrderTable).
				Set(checklistitem.WorkOrderColumn, eid).
				Where(sql.EQ(checklistitem.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return cli, nil
}
