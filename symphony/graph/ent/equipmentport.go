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
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentport"
	"github.com/facebookincubator/symphony/graph/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/graph/ent/link"
)

// EquipmentPort is the model entity for the EquipmentPort schema.
type EquipmentPort struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentPortQuery when eager-loading is set.
	Edges         EquipmentPortEdges `json:"edges"`
	parent_id     *string
	definition_id *string
	link_id       *string
}

// EquipmentPortEdges holds the relations/edges for other nodes in the graph.
type EquipmentPortEdges struct {
	// Definition holds the value of the definition edge.
	Definition *EquipmentPortDefinition
	// Parent holds the value of the parent edge.
	Parent *Equipment
	// Link holds the value of the link edge.
	Link *Link
	// Properties holds the value of the properties edge.
	Properties []*Property
	// Endpoints holds the value of the endpoints edge.
	Endpoints []*ServiceEndpoint
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// DefinitionErr returns the Definition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentPortEdges) DefinitionErr() (*EquipmentPortDefinition, error) {
	if e.loadedTypes[0] {
		if e.Definition == nil {
			// The edge definition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipmentportdefinition.Label}
		}
		return e.Definition, nil
	}
	return nil, &NotLoadedError{edge: "definition"}
}

// ParentErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentPortEdges) ParentErr() (*Equipment, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// LinkErr returns the Link value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentPortEdges) LinkErr() (*Link, error) {
	if e.loadedTypes[2] {
		if e.Link == nil {
			// The edge link was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: link.Label}
		}
		return e.Link, nil
	}
	return nil, &NotLoadedError{edge: "link"}
}

// PropertiesErr returns the Properties value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentPortEdges) PropertiesErr() ([]*Property, error) {
	if e.loadedTypes[3] {
		return e.Properties, nil
	}
	return nil, &NotLoadedError{edge: "properties"}
}

// EndpointsErr returns the Endpoints value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentPortEdges) EndpointsErr() ([]*ServiceEndpoint, error) {
	if e.loadedTypes[4] {
		return e.Endpoints, nil
	}
	return nil, &NotLoadedError{edge: "endpoints"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentPort) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // create_time
		&sql.NullTime{},  // update_time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*EquipmentPort) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // parent_id
		&sql.NullInt64{}, // definition_id
		&sql.NullInt64{}, // link_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentPort fields.
func (ep *EquipmentPort) assignValues(values ...interface{}) error {
	if m, n := len(values), len(equipmentport.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ep.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		ep.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		ep.UpdateTime = value.Time
	}
	values = values[2:]
	if len(values) == len(equipmentport.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field parent_id", value)
		} else if value.Valid {
			ep.parent_id = new(string)
			*ep.parent_id = strconv.FormatInt(value.Int64, 10)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field definition_id", value)
		} else if value.Valid {
			ep.definition_id = new(string)
			*ep.definition_id = strconv.FormatInt(value.Int64, 10)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field link_id", value)
		} else if value.Valid {
			ep.link_id = new(string)
			*ep.link_id = strconv.FormatInt(value.Int64, 10)
		}
	}
	return nil
}

// QueryDefinition queries the definition edge of the EquipmentPort.
func (ep *EquipmentPort) QueryDefinition() *EquipmentPortDefinitionQuery {
	return (&EquipmentPortClient{ep.config}).QueryDefinition(ep)
}

// QueryParent queries the parent edge of the EquipmentPort.
func (ep *EquipmentPort) QueryParent() *EquipmentQuery {
	return (&EquipmentPortClient{ep.config}).QueryParent(ep)
}

// QueryLink queries the link edge of the EquipmentPort.
func (ep *EquipmentPort) QueryLink() *LinkQuery {
	return (&EquipmentPortClient{ep.config}).QueryLink(ep)
}

// QueryProperties queries the properties edge of the EquipmentPort.
func (ep *EquipmentPort) QueryProperties() *PropertyQuery {
	return (&EquipmentPortClient{ep.config}).QueryProperties(ep)
}

// QueryEndpoints queries the endpoints edge of the EquipmentPort.
func (ep *EquipmentPort) QueryEndpoints() *ServiceEndpointQuery {
	return (&EquipmentPortClient{ep.config}).QueryEndpoints(ep)
}

// Update returns a builder for updating this EquipmentPort.
// Note that, you need to call EquipmentPort.Unwrap() before calling this method, if this EquipmentPort
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EquipmentPort) Update() *EquipmentPortUpdateOne {
	return (&EquipmentPortClient{ep.config}).UpdateOne(ep)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ep *EquipmentPort) Unwrap() *EquipmentPort {
	tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentPort is not a transactional entity")
	}
	ep.config.driver = tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EquipmentPort) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentPort(")
	builder.WriteString(fmt.Sprintf("id=%v", ep.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ep.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ep.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (ep *EquipmentPort) id() int {
	id, _ := strconv.Atoi(ep.ID)
	return id
}

// EquipmentPorts is a parsable slice of EquipmentPort.
type EquipmentPorts []*EquipmentPort

func (ep EquipmentPorts) config(cfg config) {
	for _i := range ep {
		ep[_i].config = cfg
	}
}
