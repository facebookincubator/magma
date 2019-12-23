// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"github.com/AlekSi/pointer"
	"testing"

	"github.com/facebookincubator/symphony/graph/ent/property"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/graph/viewer/viewertest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func pointerToServiceStatus(status models.ServiceStatus) *models.ServiceStatus {
	return &status
}

func TestAddServiceWithProperties(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr := r.Mutation()

	serviceTypeStrValue := "Foo"
	serviceStrPropType := models.PropertyTypeInput{
		Name:        "service_str_prop",
		Type:        "string",
		StringValue: &serviceTypeStrValue,
	}
	servicePropTypeInput := []*models.PropertyTypeInput{&serviceStrPropType}

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "Internet Access", HasCustomer: false, Properties: servicePropTypeInput})
	require.NoError(t, err)

	propertyType, err := serviceType.QueryPropertyTypes().Only(ctx)
	require.NoError(t, err)

	serviceStrValue := "Bar"
	serviceStrProp := models.PropertyInput{PropertyTypeID: propertyType.ID, StringValue: &serviceStrValue}

	servicePropInput := []*models.PropertyInput{&serviceStrProp}

	service, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Kent building, room 201",
		ServiceTypeID: serviceType.ID,
		Properties:    servicePropInput,
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)

	fetchedProperty, err := service.QueryProperties().Only(ctx)
	require.NoError(t, err)

	assert.Equal(t, fetchedProperty.StringVal, serviceStrValue)
}

func TestAddServiceWithExternalIdUnique(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr := r.Mutation()

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:        "Internet Access",
		HasCustomer: false,
		Properties:  []*models.PropertyTypeInput{},
	})
	require.NoError(t, err)

	externalID1 := "S121"
	externalID2 := "S122"

	s1, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Kent building, room 201",
		ServiceTypeID: serviceType.ID,
		ExternalID:    &externalID1,
		Properties:    []*models.PropertyInput{},
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	assert.Equal(t, *s1.ExternalID, externalID1)

	_, err = mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Kent building, room 202",
		ServiceTypeID: serviceType.ID,
		ExternalID:    &externalID1,
		Properties:    []*models.PropertyInput{},
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.Error(t, err)

	s2, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Kent building, room 203",
		ServiceTypeID: serviceType.ID,
		ExternalID:    &externalID2,
		Properties:    []*models.PropertyInput{},
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	assert.Equal(t, *s2.ExternalID, externalID2)
}

func TestAddServiceWithCustomer(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:        "Internet Access",
		HasCustomer: true,
	})
	require.NoError(t, err)

	customerID := "S3213"
	customer, err := mr.AddCustomer(ctx, models.AddCustomerInput{Name: "Donald", ExternalID: &customerID})
	require.NoError(t, err)

	s, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Kent building, room 201",
		ServiceTypeID: serviceType.ID,
		CustomerID:    &customer.ID,
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)

	fetchedService, err := qr.Service(ctx, s.ID)
	require.NoError(t, err)

	customer = fetchedService.QueryCustomer().OnlyX(ctx)

	assert.Equal(t, customer.Name, "Donald")
	assert.Equal(t, *customer.ExternalID, customerID)
}

func TestServiceTopologyReturnsCorrectLinksAndEquipment(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr := r.Mutation()

	locType, _ := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "Room",
	})

	eqt, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "Router",
		Ports: []*models.EquipmentPortInput{
			{Name: "typ1_p1"},
			{Name: "typ1_p2"},
		},
	})

	loc, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "Room2",
		Type: locType.ID,
	})

	eq1, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router1",
		Type:     eqt.ID,
		Location: &loc.ID,
	})

	eq2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router2",
		Type:     eqt.ID,
		Location: &loc.ID,
	})

	eq3, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router3",
		Type:     eqt.ID,
		Location: &loc.ID,
	})

	equipmentType := r.client.EquipmentType.GetX(ctx, eqt.ID)
	defs := equipmentType.QueryPortDefinitions().AllX(ctx)

	l1, _ := mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: eq1.ID, Port: defs[0].ID},
			{Equipment: eq2.ID, Port: defs[0].ID},
		},
	})
	l2, _ := mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: eq2.ID, Port: defs[1].ID},
			{Equipment: eq3.ID, Port: defs[1].ID},
		},
	})

	st, _ := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "Internet Access", HasCustomer: false})

	s, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:                "Internet Access Room 2",
		ServiceTypeID:       st.ID,
		TerminationPointIds: []string{eq1.ID},
		Status:              pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	_, err = mr.AddServiceLink(ctx, s.ID, l1.ID)
	require.NoError(t, err)
	_, err = mr.AddServiceLink(ctx, s.ID, l2.ID)
	require.NoError(t, err)

	res, err := r.Service().Topology(ctx, s)
	require.NoError(t, err)

	require.Len(t, res.Nodes, 3)
	require.Len(t, res.Links, 2)
}

func TestServiceTopologyWithSlots(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr := r.Mutation()

	locType, _ := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "Room",
	})

	router, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "Router",

		Positions: []*models.EquipmentPositionInput{
			{Name: "slot1"},
		},
	})

	card, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "Card",
		Ports: []*models.EquipmentPortInput{
			{Name: "port1"},
		},
	})

	loc, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "Room2",
		Type: locType.ID,
	})

	posDefs := router.QueryPositionDefinitions().AllX(ctx)

	router1, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router1",
		Type:     router.ID,
		Location: &loc.ID,
	})

	card1, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:               "Card1",
		Type:               card.ID,
		Parent:             &router1.ID,
		PositionDefinition: &posDefs[0].ID,
	})

	router2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router2",
		Type:     router.ID,
		Location: &loc.ID,
	})

	card2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:               "Card2",
		Type:               card.ID,
		Parent:             &router2.ID,
		PositionDefinition: &posDefs[0].ID,
	})

	portDefs := card.QueryPortDefinitions().AllX(ctx)

	l, _ := mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: card1.ID, Port: portDefs[0].ID},
			{Equipment: card2.ID, Port: portDefs[0].ID},
		},
	})

	st, _ := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "Internet Access", HasCustomer: false})

	s, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:                "Internet Access Room 2",
		ServiceTypeID:       st.ID,
		TerminationPointIds: []string{router1.ID},
		Status:              pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	_, err = mr.AddServiceLink(ctx, s.ID, l.ID)
	require.NoError(t, err)

	res, err := r.Service().Topology(ctx, s)
	require.NoError(t, err)

	require.Len(t, res.Nodes, 2)
	require.Len(t, res.Links, 1)

	source, err := res.Links[0].Source.Node(ctx)
	require.NoError(t, err)
	require.Contains(t, []string{router1.ID, router2.ID}, source.ID)
	target, err := res.Links[0].Target.Node(ctx)
	require.NoError(t, err)
	require.Contains(t, []string{router1.ID, router2.ID}, target.ID)
}

func TestEditService(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "service_type_name",
	})
	require.NoError(t, err)
	require.Equal(t, "service_type_name", serviceType.Name)

	service, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "service_name",
		ServiceTypeID: serviceType.ID,
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)

	newService, err := mr.EditService(ctx, models.ServiceEditData{
		ID:   service.ID,
		Name: pointer.ToString("new_service_name"),
	})
	require.NoError(t, err)
	require.Equal(t, "new_service_name", newService.Name)

	fetchedService, _ := qr.Service(ctx, service.ID)
	require.Equal(t, newService.Name, fetchedService.Name)
}

func TestEditServiceWithExternalID(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "service_type_name",
	})
	require.NoError(t, err)
	require.Equal(t, "service_type_name", serviceType.Name)

	service, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "service_name",
		ServiceTypeID: serviceType.ID,
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	fetchedService, _ := qr.Service(ctx, service.ID)
	require.Nil(t, fetchedService.ExternalID)

	externalID1 := "externalID1"
	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:         service.ID,
		Name:       pointer.ToString(service.Name),
		ExternalID: &externalID1,
	})
	require.NoError(t, err)
	fetchedService, _ = qr.Service(ctx, service.ID)
	require.Equal(t, externalID1, *fetchedService.ExternalID)

	externalID2 := "externalID2"
	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:         service.ID,
		Name:       pointer.ToString(service.Name),
		ExternalID: &externalID2,
	})
	require.NoError(t, err)
	fetchedService, _ = qr.Service(ctx, service.ID)
	require.Equal(t, externalID2, *fetchedService.ExternalID)
}

func TestEditServiceWithCustomer(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "service_type_name",
	})
	require.NoError(t, err)
	require.Equal(t, "service_type_name", serviceType.Name)

	service, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "service_name",
		ServiceTypeID: serviceType.ID,
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)

	fetchedService, _ := qr.Service(ctx, service.ID)
	exist := fetchedService.QueryCustomer().ExistX(ctx)
	require.Equal(t, false, exist)

	donald, err := mr.AddCustomer(ctx, models.AddCustomerInput{
		Name: "Donald Duck",
	})
	require.NoError(t, err)

	dafi, err := mr.AddCustomer(ctx, models.AddCustomerInput{
		Name: "Dafi Duck",
	})
	require.NoError(t, err)

	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:         service.ID,
		Name:       pointer.ToString(service.Name),
		CustomerID: &donald.ID,
	})
	require.NoError(t, err)
	fetchedService, _ = qr.Service(ctx, service.ID)
	fetchedCustomer := fetchedService.QueryCustomer().OnlyX(ctx)
	require.Equal(t, donald.ID, fetchedCustomer.ID)

	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:         service.ID,
		Name:       pointer.ToString(service.Name),
		CustomerID: &dafi.ID,
	})
	require.NoError(t, err)
	fetchedService, _ = qr.Service(ctx, service.ID)
	fetchedCustomer = fetchedService.QueryCustomer().OnlyX(ctx)
	require.Equal(t, dafi.ID, fetchedCustomer.ID)
}

func TestEditServiceWithProperties(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	pTypes := models.PropertyTypeInput{
		Name: "str_prop",
		Type: "string",
	}

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:       "type_name_1",
		Properties: []*models.PropertyTypeInput{&pTypes},
	})
	require.NoError(t, err)
	pTypeID := serviceType.QueryPropertyTypes().OnlyXID(ctx)

	strValue := "Foo"
	strProp := models.PropertyInput{
		PropertyTypeID: pTypeID,
		StringValue:    &strValue,
	}
	strValue2 := "Bar"
	strProp2 := models.PropertyInput{
		PropertyTypeID: pTypeID,
		StringValue:    &strValue2,
	}

	service, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "inst_name_1",
		ServiceTypeID: serviceType.ID,
		Properties:    []*models.PropertyInput{&strProp, &strProp2},
		Status:        pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	fetchedService, _ := qr.Service(ctx, service.ID)
	fetchedProps, _ := fetchedService.QueryProperties().All(ctx)

	// Property[] -> PropertyInput[]
	var propInputClone []*models.PropertyInput
	for _, v := range fetchedProps {
		var strValue = v.StringVal + "-2"
		propInput := &models.PropertyInput{
			ID:             &v.ID,
			PropertyTypeID: v.QueryType().OnlyXID(ctx),
			StringValue:    &strValue,
		}
		propInputClone = append(propInputClone, propInput)
	}

	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:         service.ID,
		Name:       pointer.ToString("service_name_1"),
		Properties: propInputClone,
	})
	require.NoError(t, err, "Editing service")

	newFetchedService, err := qr.Service(ctx, service.ID)
	require.NoError(t, err)
	existA := newFetchedService.QueryProperties().Where(property.StringVal("Foo-2")).ExistX(ctx)
	require.NoError(t, err)
	require.True(t, existA, "Property with the new name should exist on service")
	existB := newFetchedService.QueryProperties().Where(property.StringVal("Bar-2")).ExistX(ctx)
	require.NoError(t, err)
	require.True(t, existB, "Property with the new name should exist on service")
	existC := newFetchedService.QueryProperties().Where(property.StringVal("Bar")).ExistX(ctx)
	require.NoError(t, err)
	require.False(t, existC, "Property with the old name should not exist on service")
}

func TestEditServiceWithTerminationPoints(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr, qr := r.Mutation(), r.Query()

	locType, err := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "loc_type_name",
	})
	require.NoError(t, err)

	location, err := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "loc_inst_name",
		Type: locType.ID,
	})
	require.NoError(t, err)

	eqType, err := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "eq_type_name",
	})
	require.NoError(t, err)

	eq1, err := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst_name_1",
		Type:     eqType.ID,
		Location: &location.ID,
	})
	require.NoError(t, err)

	eq2, err := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst_name_2",
		Type:     eqType.ID,
		Location: &location.ID,
	})
	require.NoError(t, err)

	eq3, err := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst_name_3",
		Type:     eqType.ID,
		Location: &location.ID,
	})
	require.NoError(t, err)

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "service_type_name",
	})
	require.NoError(t, err)
	require.Equal(t, "service_type_name", serviceType.Name)

	service, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:                "service_name",
		ServiceTypeID:       serviceType.ID,
		TerminationPointIds: []string{eq1.ID, eq2.ID},
		Status:              pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)

	fetchedService, _ := qr.Service(ctx, service.ID)
	terminationPoints := fetchedService.QueryTerminationPoints().IDsX(ctx)
	require.Len(t, terminationPoints, 2)
	require.NotContains(t, terminationPoints, eq3.ID)

	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:                  service.ID,
		Name:                pointer.ToString(service.Name),
		TerminationPointIds: []string{eq2.ID, eq3.ID},
	})
	require.NoError(t, err)
	fetchedService, _ = qr.Service(ctx, service.ID)
	terminationPoints = fetchedService.QueryTerminationPoints().IDsX(ctx)
	require.Len(t, terminationPoints, 2)
	require.Contains(t, terminationPoints, eq3.ID)
	require.NotContains(t, terminationPoints, eq1.ID)

	_, err = mr.EditService(ctx, models.ServiceEditData{
		ID:                  service.ID,
		Name:                pointer.ToString(service.Name),
		TerminationPointIds: []string{},
	})
	require.NoError(t, err)
	exist := fetchedService.QueryTerminationPoints().ExistX(ctx)
	require.False(t, exist)
}

func TestServicesOfEquipment(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	mr := r.Mutation()

	locType, _ := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "Room",
	})

	eqt, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "Router",
		Ports: []*models.EquipmentPortInput{
			{Name: "typ1_p1"},
			{Name: "typ1_p2"},
		},
	})

	loc, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "Room2",
		Type: locType.ID,
	})

	eq1, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router1",
		Type:     eqt.ID,
		Location: &loc.ID,
	})

	eq2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router2",
		Type:     eqt.ID,
		Location: &loc.ID,
	})

	eq3, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "Router3",
		Type:     eqt.ID,
		Location: &loc.ID,
	})

	equipmentType := r.client.EquipmentType.GetX(ctx, eqt.ID)
	defs := equipmentType.QueryPortDefinitions().AllX(ctx)

	l1, _ := mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: eq1.ID, Port: defs[0].ID},
			{Equipment: eq2.ID, Port: defs[0].ID},
		},
	})
	l2, _ := mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: eq2.ID, Port: defs[1].ID},
			{Equipment: eq3.ID, Port: defs[1].ID},
		},
	})

	st, _ := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "Internet Access", HasCustomer: false})

	_, err = mr.AddService(ctx, models.ServiceCreateData{
		Name:                "Internet Access Room 2a",
		ServiceTypeID:       st.ID,
		TerminationPointIds: []string{eq1.ID},
		Status:              pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)

	s2, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:                "Internet Access Room 2b",
		ServiceTypeID:       st.ID,
		TerminationPointIds: []string{eq1.ID},
		Status:              pointerToServiceStatus(models.ServiceStatusPending),
	})
	require.NoError(t, err)
	_, err = mr.AddServiceLink(ctx, s2.ID, l1.ID)
	require.NoError(t, err)
	_, err = mr.AddServiceLink(ctx, s2.ID, l2.ID)
	require.NoError(t, err)

	eq1Services, err := r.Equipment().Services(ctx, eq1)
	require.NoError(t, err)
	require.Len(t, eq1Services, 2)

	eq2Services, err := r.Equipment().Services(ctx, eq2)
	require.NoError(t, err)
	require.Len(t, eq2Services, 1)
	require.Equal(t, s2.ID, eq2Services[0].ID)
}
