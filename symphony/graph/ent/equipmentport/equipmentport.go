// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package equipmentport

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/symphony/graph/ent/schema"
)

const (
	// Label holds the string label denoting the equipmentport type in the database.
	Label = "equipment_port"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time vertex property in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time vertex property in the database.
	FieldUpdateTime = "update_time"

	// Table holds the table name of the equipmentport in the database.
	Table = "equipment_ports"
	// DefinitionTable is the table the holds the definition relation/edge.
	DefinitionTable = "equipment_ports"
	// DefinitionInverseTable is the table name for the EquipmentPortDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentportdefinition" package.
	DefinitionInverseTable = "equipment_port_definitions"
	// DefinitionColumn is the table column denoting the definition relation/edge.
	DefinitionColumn = "definition_id"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "equipment_ports"
	// ParentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	ParentInverseTable = "equipment"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// LinkTable is the table the holds the link relation/edge.
	LinkTable = "equipment_ports"
	// LinkInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	LinkInverseTable = "links"
	// LinkColumn is the table column denoting the link relation/edge.
	LinkColumn = "link_id"
	// PropertiesTable is the table the holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "equipment_port_id"
)

// Columns holds all SQL columns are equipmentport fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
}

var (
	mixin       = schema.EquipmentPort{}.Mixin()
	mixinFields = [...][]ent.Field{
		mixin[0].Fields(),
	}
	fields = schema.EquipmentPort{}.Fields()

	// descCreateTime is the schema descriptor for create_time field.
	descCreateTime = mixinFields[0][0].Descriptor()
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime = descCreateTime.Default.(func() time.Time)

	// descUpdateTime is the schema descriptor for update_time field.
	descUpdateTime = mixinFields[0][1].Descriptor()
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime = descUpdateTime.Default.(func() time.Time)
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime = descUpdateTime.UpdateDefault.(func() time.Time)
)
