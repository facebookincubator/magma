// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Equipment is the model entity for the Equipment schema.
type Equipment struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// FutureState holds the value of the "future_state" field.
	FutureState string `json:"future_state,omitempty"`
	// DeviceID holds the value of the "device_id" field.
	DeviceID string `json:"device_id,omitempty"`
}

// FromRows scans the sql response data into Equipment.
func (e *Equipment) FromRows(rows *sql.Rows) error {
	var scane struct {
		ID          int
		CreateTime  sql.NullTime
		UpdateTime  sql.NullTime
		Name        sql.NullString
		FutureState sql.NullString
		DeviceID    sql.NullString
	}
	// the order here should be the same as in the `equipment.Columns`.
	if err := rows.Scan(
		&scane.ID,
		&scane.CreateTime,
		&scane.UpdateTime,
		&scane.Name,
		&scane.FutureState,
		&scane.DeviceID,
	); err != nil {
		return err
	}
	e.ID = strconv.Itoa(scane.ID)
	e.CreateTime = scane.CreateTime.Time
	e.UpdateTime = scane.UpdateTime.Time
	e.Name = scane.Name.String
	e.FutureState = scane.FutureState.String
	e.DeviceID = scane.DeviceID.String
	return nil
}

// QueryType queries the type edge of the Equipment.
func (e *Equipment) QueryType() *EquipmentTypeQuery {
	return (&EquipmentClient{e.config}).QueryType(e)
}

// QueryLocation queries the location edge of the Equipment.
func (e *Equipment) QueryLocation() *LocationQuery {
	return (&EquipmentClient{e.config}).QueryLocation(e)
}

// QueryParentPosition queries the parent_position edge of the Equipment.
func (e *Equipment) QueryParentPosition() *EquipmentPositionQuery {
	return (&EquipmentClient{e.config}).QueryParentPosition(e)
}

// QueryPositions queries the positions edge of the Equipment.
func (e *Equipment) QueryPositions() *EquipmentPositionQuery {
	return (&EquipmentClient{e.config}).QueryPositions(e)
}

// QueryPorts queries the ports edge of the Equipment.
func (e *Equipment) QueryPorts() *EquipmentPortQuery {
	return (&EquipmentClient{e.config}).QueryPorts(e)
}

// QueryWorkOrder queries the work_order edge of the Equipment.
func (e *Equipment) QueryWorkOrder() *WorkOrderQuery {
	return (&EquipmentClient{e.config}).QueryWorkOrder(e)
}

// QueryProperties queries the properties edge of the Equipment.
func (e *Equipment) QueryProperties() *PropertyQuery {
	return (&EquipmentClient{e.config}).QueryProperties(e)
}

// QueryService queries the service edge of the Equipment.
func (e *Equipment) QueryService() *ServiceQuery {
	return (&EquipmentClient{e.config}).QueryService(e)
}

// QueryFiles queries the files edge of the Equipment.
func (e *Equipment) QueryFiles() *FileQuery {
	return (&EquipmentClient{e.config}).QueryFiles(e)
}

// Update returns a builder for updating this Equipment.
// Note that, you need to call Equipment.Unwrap() before calling this method, if this Equipment
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Equipment) Update() *EquipmentUpdateOne {
	return (&EquipmentClient{e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Equipment) Unwrap() *Equipment {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Equipment is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Equipment) String() string {
	var builder strings.Builder
	builder.WriteString("Equipment(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(e.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(e.Name)
	builder.WriteString(", future_state=")
	builder.WriteString(e.FutureState)
	builder.WriteString(", device_id=")
	builder.WriteString(e.DeviceID)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (e *Equipment) id() int {
	id, _ := strconv.Atoi(e.ID)
	return id
}

// EquipmentSlice is a parsable slice of Equipment.
type EquipmentSlice []*Equipment

// FromRows scans the sql response data into EquipmentSlice.
func (e *EquipmentSlice) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scane := &Equipment{}
		if err := scane.FromRows(rows); err != nil {
			return err
		}
		*e = append(*e, scane)
	}
	return nil
}

func (e EquipmentSlice) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
