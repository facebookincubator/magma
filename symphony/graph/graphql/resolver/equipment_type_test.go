// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// nolint: goconst
package resolver

import (
	"sort"
	"testing"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentpositiondefinition"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/graph/viewer/viewertest"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddEquipmentTypesSameName(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr := r.Mutation()
	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "example_type_name",
	})
	require.NoError(t, err)
	assert.Equal(t, "example_type_name", equipmentType.Name)
	_, err = mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "example_type_name",
	})
	require.Error(t, err)
}

func TestQueryEquipmentTypes(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()
	for _, suffix := range []string{"a", "b"} {
		_, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
			Name:     "example_type_" + suffix,
			Category: pointer.ToString("example_type"),
		})
		require.NoError(t, err)
	}
	types, _ := qr.EquipmentTypes(ctx, nil, nil, nil, nil)
	require.Len(t, types.Edges, 2)

	var (
		names      = make([]string, len(types.Edges))
		categories = make([]*ent.EquipmentCategory, len(types.Edges))
	)
	for i, v := range types.Edges {
		names[i] = v.Node.Name
		category, err := v.Node.QueryCategory().Only(ctx)
		require.NoError(t, err)
		categories[i] = category
		require.Equal(t, "example_type", category.Name)
	}
	require.Len(t, categories, 2)
	assert.Equal(t, categories[0].ID, categories[1].ID)
	sort.Strings(names)
	assert.Equal(t, names, []string{"example_type_a", "example_type_b"})
}

func TestAddEquipmentTypeWithPositions(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	position1 := models.EquipmentPositionInput{
		Name: "Position 1",
	}
	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:      "equipment_type_name_1",
		Positions: []*models.EquipmentPositionInput{&position1},
	})
	require.NoError(t, err)
	fetchedEquipmentType, err := qr.EquipmentType(ctx, equipmentType.ID)
	require.NoError(t, err)

	require.Equal(t, equipmentType.ID, fetchedEquipmentType.ID, "Verifying saved equipment type vs fetched equipmenttype : ID")
	require.Equal(t, equipmentType.Name, fetchedEquipmentType.Name, "Verifying saved equipment type  vs fetched equipment type : Name")
	require.Equal(t, equipmentType.QueryPositionDefinitions().OnlyXID(ctx), fetchedEquipmentType.QueryPositionDefinitions().OnlyXID(ctx), "Verifying saved equipment type  vs fetched equipment type: position definition")
}

func TestAddEquipmentTypeWithProperties(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr, qr, etr := r.Mutation(), r.Query(), r.EquipmentType()

	strValue := "Foo"
	index := 5
	ptype := models.PropertyTypeInput{
		Name:        "str_prop",
		Type:        "string",
		Index:       &index,
		StringValue: &strValue,
	}
	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:       "example_type_a",
		Properties: []*models.PropertyTypeInput{&ptype},
	})
	require.NoError(t, err)

	fetchedEquipmentType, _ := qr.EquipmentType(ctx, equipmentType.ID)
	fetchedPropertyTypes, _ := etr.PropertyTypes(ctx, fetchedEquipmentType)
	require.Len(t, fetchedPropertyTypes, 1)
	assert.Equal(t, fetchedPropertyTypes[0].Name, "str_prop")
	assert.Equal(t, fetchedPropertyTypes[0].Type, "string")
	assert.Equal(t, fetchedPropertyTypes[0].Index, 5)
}

func TestAddEquipmentTypeWithoutPositionNames(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr := r.Mutation()

	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "equipment_type_name_1",
	})
	require.NoError(t, err)
	positions, err := equipmentType.QueryPositionDefinitions().All(ctx)
	require.NoError(t, err)
	assert.Len(t, positions, 0)
}

func TestAddEquipmentTypeWithPorts(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr, qr := r.Mutation(), r.Query()

	visibleLabel := "Eth1"
	bandwidth := "10/100/1000BASE-T"
	portDef := models.EquipmentPortInput{
		Name:         "Port 1",
		Type:         "Eth",
		VisibleLabel: &visibleLabel,
		Bandwidth:    &bandwidth,
	}

	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:  "example_type_a",
		Ports: []*models.EquipmentPortInput{&portDef},
	})
	require.NoError(t, err)
	fetchedEquipmentType, _ := qr.EquipmentType(ctx, equipmentType.ID)
	ports := fetchedEquipmentType.QueryPortDefinitions().AllX(ctx)
	require.Len(t, ports, 1)

	assert.Equal(t, ports[0].Name, "Port 1")
	assert.Equal(t, ports[0].Type, "Eth")
	assert.Equal(t, ports[0].VisibilityLabel, visibleLabel)
	assert.Equal(t, ports[0].Bandwidth, bandwidth)
}

func TestRemoveEquipmentTypeWithExistingEquipments(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr, qr := r.Mutation(), r.Query()

	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "example_type_a",
	})
	require.NoError(t, err)

	locationType, err := mr.AddLocationType(ctx, models.AddLocationTypeInput{Name: "location_type_name_1"})
	require.NoError(t, err)

	location, err := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "location_name_1",
		Type: locationType.ID,
	})
	require.NoError(t, err)

	_, err = mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "equipment_name_1",
		Type:     equipmentType.ID,
		Location: &location.ID,
	})
	require.NoError(t, err)

	_, err = mr.RemoveEquipmentType(ctx, equipmentType.ID)
	require.Error(t, err)

	fetchedEquipmentType, err := qr.EquipmentType(ctx, equipmentType.ID)
	require.NoError(t, err)
	assert.Equal(t, fetchedEquipmentType.ID, equipmentType.ID)
}

func TestRemoveEquipmentType(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr, qr := r.Mutation(), r.Query()

	visibleLabel := "Eth1"
	bandwidth := "10/100/1000BASE-T"
	portDef := models.EquipmentPortInput{
		Name:         "Port 1",
		Type:         "Eth",
		VisibleLabel: &visibleLabel,
		Bandwidth:    &bandwidth,
	}
	strValue := "Foo"
	strPropType := models.PropertyTypeInput{
		Name:        "str_prop",
		Type:        models.PropertyKindString,
		StringValue: &strValue,
	}
	position1 := models.EquipmentPositionInput{
		Name: "Position 1",
	}

	equipmentType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:       "example_type_a",
		Positions:  []*models.EquipmentPositionInput{&position1},
		Ports:      []*models.EquipmentPortInput{&portDef},
		Properties: []*models.PropertyTypeInput{&strPropType},
	})
	require.NoError(t, err)

	_, err = mr.RemoveEquipmentType(ctx, equipmentType.ID)
	require.NoError(t, err)

	deletedEquipmentType, err := qr.EquipmentType(ctx, equipmentType.ID)
	require.NoError(t, err)
	assert.Nil(t, deletedEquipmentType)

	propertyTypes := equipmentType.QueryPropertyTypes().AllX(ctx)
	require.Len(t, propertyTypes, 0)

	for _, positionDefinition := range equipmentType.QueryPositionDefinitions().AllX(ctx) {
		fetchedPositionDefinition, err := qr.EquipmentPositionDefinition(ctx, positionDefinition.ID)
		require.Error(t, err)
		assert.Nil(t, fetchedPositionDefinition)
	}
}

func TestEditEquipmentType(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr, qr := r.Mutation(), r.Query()

	eqType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "example_type_name",
	})
	require.NoError(t, err)
	require.Equal(t, "example_type_name", eqType.Name)
	c, _ := eqType.QueryCategory().Only(ctx)
	require.Nil(t, c)

	category := "example_type"
	newType, err := mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:       eqType.ID,
		Name:     "example_type_name_edited",
		Category: &category,
	})
	require.NoError(t, err)
	require.Equal(t, "example_type_name_edited", newType.Name, "successfully edited equipment type name")
	c, _ = newType.QueryCategory().Only(ctx)
	require.Equal(t, category, c.Name)

	eqType, err = mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "example_type_name_2",
	})
	require.NoError(t, err)
	_, err = mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:   eqType.ID,
		Name: "example_type_name_edited",
	})
	require.Error(t, err, "duplicate names")

	types, err := qr.EquipmentTypes(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	require.Len(t, types.Edges, 2)

	typ, err := qr.EquipmentType(ctx, eqType.ID)
	require.NoError(t, err)
	require.Equal(t, "example_type_name_2", typ.Name)
}

func TestEditEquipmentTypeRemoveCategory(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr := r.Mutation()

	category := "example_type"
	eqType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:     "example_type_name",
		Category: &category,
	})
	require.NoError(t, err)
	require.Equal(t, "example_type_name", eqType.Name)
	c, _ := eqType.QueryCategory().Only(ctx)
	require.Equal(t, category, c.Name)

	newType, err := mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:   eqType.ID,
		Name: "example_type_name_edited",
	})
	require.NoError(t, err)
	require.Equal(t, "example_type_name_edited", newType.Name, "successfully edited equipment type name")
	c, _ = newType.QueryCategory().Only(ctx)
	require.Nil(t, c)
}

func TestEditEquipmentTypeWithProperties(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr := r.Mutation()

	strValue := "Foo"
	strPropType := models.PropertyTypeInput{
		Name:        "str_prop",
		Type:        models.PropertyKindString,
		StringValue: &strValue,
	}
	propTypeInput := []*models.PropertyTypeInput{&strPropType}
	eqType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:       "example_type_a",
		Properties: propTypeInput,
	})
	require.NoError(t, err)

	strProp := eqType.QueryPropertyTypes().Where(propertytype.Type("string")).OnlyX(ctx)
	strValue = "Foo - edited"
	intValue := 5
	strPropType = models.PropertyTypeInput{
		ID:          &strProp.ID,
		Name:        "str_prop_new",
		Type:        models.PropertyKindString,
		StringValue: &strValue,
	}
	intPropType := models.PropertyTypeInput{
		Name:     "int_prop",
		Type:     models.PropertyKindInt,
		IntValue: &intValue,
	}
	editedPropTypeInput := []*models.PropertyTypeInput{&strPropType, &intPropType}
	newType, err := mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:         eqType.ID,
		Name:       "example_type_a",
		Properties: editedPropTypeInput,
	})
	require.NoError(t, err)
	require.Equal(t, eqType.Name, newType.Name, "successfully edited equipment type name")

	strProp = eqType.QueryPropertyTypes().Where(propertytype.Type("string")).OnlyX(ctx)
	require.Equal(t, "str_prop_new", strProp.Name, "successfully edited prop type name")
	require.Equal(t, strValue, strProp.StringVal, "successfully edited prop type string value")

	intProp := eqType.QueryPropertyTypes().Where(propertytype.Type("int")).OnlyX(ctx)
	require.Equal(t, "int_prop", intProp.Name, "successfully edited prop type name")
	require.Equal(t, intValue, intProp.IntVal, "successfully edited prop type int value")

	intValue = 6
	intPropType = models.PropertyTypeInput{
		Name:     "int_prop",
		Type:     models.PropertyKindInt,
		IntValue: &intValue,
	}
	editedPropTypeInput = []*models.PropertyTypeInput{&intPropType}
	_, err = mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:         eqType.ID,
		Name:       "example_type_a",
		Properties: editedPropTypeInput,
	})
	require.Error(t, err, "duplicate property type names")
}

func TestEditEquipmentTypeWithPortsAndPositions(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	mr := r.Mutation()

	bandwidth := "b1"
	label := "v1"
	strPortType := models.EquipmentPortInput{
		Name:         "str_prop",
		Type:         "string",
		Bandwidth:    &bandwidth,
		VisibleLabel: &label,
	}
	posTypeA := models.EquipmentPositionInput{
		Name:         "str_prop",
		VisibleLabel: &label,
	}
	posTypeInput := []*models.EquipmentPositionInput{&posTypeA}
	portTypeInput := []*models.EquipmentPortInput{&strPortType}
	eqType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name:      "example_type_a",
		Positions: posTypeInput,
		Ports:     portTypeInput,
	})
	require.NoError(t, err)

	bandwidth = "b1 edited"
	label = "v1 edited"
	strPort := eqType.QueryPortDefinitions().OnlyX(ctx)
	strPortType = models.EquipmentPortInput{
		ID:           &strPort.ID,
		Name:         "str_port_edited",
		Type:         "string",
		Bandwidth:    &bandwidth,
		VisibleLabel: &label,
	}
	bandwidthInt := "b2 new"
	labelInt := "v2 new"
	intPortType := models.EquipmentPortInput{
		Name:         "int_port",
		Type:         "int",
		Bandwidth:    &bandwidthInt,
		VisibleLabel: &labelInt,
	}
	portTypeInput = []*models.EquipmentPortInput{&strPortType, &intPortType}

	strPos := eqType.QueryPositionDefinitions().OnlyX(ctx)
	posTypeA = models.EquipmentPositionInput{
		ID:           &strPos.ID,
		Name:         "str_pos_edited",
		VisibleLabel: &label,
	}
	posTypeB := models.EquipmentPositionInput{
		Name:         "str_pos_new",
		VisibleLabel: &label,
	}
	posTypeInput = []*models.EquipmentPositionInput{&posTypeA, &posTypeB}

	newType, err := mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:        eqType.ID,
		Name:      "example_type_a",
		Positions: posTypeInput,
		Ports:     portTypeInput,
	})
	require.NoError(t, err)
	require.Equal(t, eqType.Name, newType.Name, "successfully edited equipment type name")

	strPort = eqType.QueryPortDefinitions().Where(equipmentportdefinition.Type("string")).OnlyX(ctx)
	require.Equal(t, "str_port_edited", strPort.Name, "successfully edited prop type name")
	require.Equal(t, bandwidth, strPort.Bandwidth, "successfully edited prop type string value")
	require.Equal(t, label, strPort.VisibilityLabel, "successfully edited prop type string value")
	require.Equal(t, "string", strPort.Type, "successfully edited prop type string value")

	intPort := eqType.QueryPortDefinitions().Where(equipmentportdefinition.Type("int")).OnlyX(ctx)
	require.Equal(t, "int_port", intPort.Name, "successfully edited prop type name")
	require.Equal(t, bandwidthInt, intPort.Bandwidth, "successfully edited prop type string value")
	require.Equal(t, labelInt, intPort.VisibilityLabel, "successfully edited prop type string value")
	require.Equal(t, "int", intPort.Type, "successfully edited prop type string value")

	pos1 := eqType.QueryPositionDefinitions().Where(equipmentpositiondefinition.Name("str_pos_edited")).OnlyX(ctx)
	require.Equal(t, label, pos1.VisibilityLabel, "successfully edited prop type string value")

	pos2 := eqType.QueryPositionDefinitions().Where(equipmentpositiondefinition.Name("str_pos_new")).OnlyX(ctx)
	require.Equal(t, label, pos2.VisibilityLabel, "successfully edited prop type string value")
}
