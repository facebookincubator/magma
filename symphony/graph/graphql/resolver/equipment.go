// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/facebookincubator/symphony/graph/ent/file"

	"github.com/facebookincubator/symphony/graph/ent/service"

	"github.com/pkg/errors"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/graphql/models"
)

type equipmentPortResolver struct{}

func (equipmentPortResolver) Definition(ctx context.Context, obj *ent.EquipmentPort) (*ent.EquipmentPortDefinition, error) {
	return obj.QueryDefinition().Only(ctx)
}

func (equipmentPortResolver) ParentEquipment(ctx context.Context, obj *ent.EquipmentPort) (*ent.Equipment, error) {
	return obj.QueryParent().Only(ctx)
}

func (equipmentPortResolver) Link(ctx context.Context, obj *ent.EquipmentPort) (*ent.Link, error) {
	l, err := obj.QueryLink().Only(ctx)
	return l, ent.MaskNotFound(err)
}

func (equipmentPortResolver) Properties(ctx context.Context, obj *ent.EquipmentPort) ([]*ent.Property, error) {
	return obj.QueryProperties().All(ctx)
}

type equipmentPositionResolver struct{}

func (equipmentPositionResolver) Definition(ctx context.Context, obj *ent.EquipmentPosition) (*ent.EquipmentPositionDefinition, error) {
	return obj.QueryDefinition().Only(ctx)
}

func (equipmentPositionResolver) ParentEquipment(ctx context.Context, obj *ent.EquipmentPosition) (*ent.Equipment, error) {
	return obj.QueryParent().Only(ctx)
}

func (equipmentPositionResolver) AttachedEquipment(ctx context.Context, obj *ent.EquipmentPosition) (*ent.Equipment, error) {
	e, err := obj.QueryAttachment().Only(ctx)
	return e, ent.MaskNotFound(err)
}

type equipmentPortDefinitionResolver struct{}

func (equipmentPortDefinitionResolver) PortType(ctx context.Context, obj *ent.EquipmentPortDefinition) (*ent.EquipmentPortType, error) {
	l, err := obj.QueryEquipmentPortType().Only(ctx)
	return l, ent.MaskNotFound(err)
}

type equipmentPortTypeResolver struct{}

func (equipmentPortTypeResolver) PropertyTypes(ctx context.Context, obj *ent.EquipmentPortType) ([]*ent.PropertyType, error) {
	return obj.QueryPropertyTypes().All(ctx)
}

func (equipmentPortTypeResolver) LinkPropertyTypes(ctx context.Context, obj *ent.EquipmentPortType) ([]*ent.PropertyType, error) {
	return obj.QueryLinkPropertyTypes().All(ctx)
}

func (equipmentPortTypeResolver) NumberOfPortDefinitions(ctx context.Context, obj *ent.EquipmentPortType) (int, error) {
	return obj.QueryPortDefinitions().Count(ctx)
}

type equipmentTypeResolver struct{}

func (equipmentTypeResolver) Category(ctx context.Context, obj *ent.EquipmentType) (*string, error) {
	c, err := obj.QueryCategory().Only(ctx)
	if c != nil {
		return &c.Name, err
	}
	return nil, ent.MaskNotFound(err)
}

func (equipmentTypeResolver) PositionDefinitions(ctx context.Context, obj *ent.EquipmentType) ([]*ent.EquipmentPositionDefinition, error) {
	return obj.QueryPositionDefinitions().All(ctx)
}

func (equipmentTypeResolver) PortDefinitions(ctx context.Context, obj *ent.EquipmentType) ([]*ent.EquipmentPortDefinition, error) {
	return obj.QueryPortDefinitions().All(ctx)
}

func (equipmentTypeResolver) PropertyTypes(ctx context.Context, obj *ent.EquipmentType) ([]*ent.PropertyType, error) {
	return obj.QueryPropertyTypes().All(ctx)
}

func (equipmentTypeResolver) Equipments(ctx context.Context, obj *ent.EquipmentType) ([]*ent.Equipment, error) {
	return obj.QueryEquipment().All(ctx)
}

func (equipmentTypeResolver) NumberOfEquipment(ctx context.Context, obj *ent.EquipmentType) (int, error) {
	return obj.QueryEquipment().Count(ctx)
}

type equipmentResolver struct{ resolver }

func (equipmentResolver) ParentLocation(ctx context.Context, obj *ent.Equipment) (*ent.Location, error) {
	l, err := obj.QueryLocation().Only(ctx)
	return l, ent.MaskNotFound(err)
}

func (equipmentResolver) ParentPosition(ctx context.Context, obj *ent.Equipment) (*ent.EquipmentPosition, error) {
	p, err := obj.QueryParentPosition().Only(ctx)
	return p, ent.MaskNotFound(err)
}

func (equipmentResolver) EquipmentType(ctx context.Context, obj *ent.Equipment) (*ent.EquipmentType, error) {
	return obj.QueryType().Only(ctx)
}

func (equipmentResolver) Positions(ctx context.Context, obj *ent.Equipment) ([]*ent.EquipmentPosition, error) {
	return obj.QueryPositions().All(ctx)
}

func (equipmentResolver) Ports(ctx context.Context, obj *ent.Equipment) ([]*ent.EquipmentPort, error) {
	return obj.QueryPorts().All(ctx)
}

func (equipmentResolver) Properties(ctx context.Context, obj *ent.Equipment) ([]*ent.Property, error) {
	return obj.QueryProperties().All(ctx)
}

func (equipmentResolver) FutureState(ctx context.Context, obj *ent.Equipment) (*models.FutureState, error) {
	fs := models.FutureState(obj.FutureState)
	return &fs, nil
}

func (equipmentResolver) WorkOrder(ctx context.Context, obj *ent.Equipment) (*ent.WorkOrder, error) {
	wo, err := obj.QueryWorkOrder().Only(ctx)
	return wo, ent.MaskNotFound(err)
}

func (equipmentResolver) Images(ctx context.Context, obj *ent.Equipment) ([]*ent.File, error) {
	return obj.QueryFiles().Where(file.Type(models.FileTypeImage.String())).All(ctx)
}

func (equipmentResolver) Files(ctx context.Context, obj *ent.Equipment) ([]*ent.File, error) {
	return obj.QueryFiles().Where(file.Type(models.FileTypeFile.String())).All(ctx)
}

func (equipmentResolver) PositionHierarchy(ctx context.Context, eq *ent.Equipment) ([]*ent.EquipmentPosition, error) {
	var positions []*ent.EquipmentPosition
	ppos, err := eq.QueryParentPosition().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "querying parent position")
	}
	for ppos != nil {
		positions = append([]*ent.EquipmentPosition{ppos}, positions...)
		p, err := ppos.QueryParent().QueryParentPosition().Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, errors.Wrap(err, "querying parent position")
		}

		ppos = p
	}
	return positions, nil
}

func (r equipmentResolver) LocationHierarchy(ctx context.Context, eq *ent.Equipment) ([]*ent.Location, error) {
	ph, err := r.PositionHierarchy(ctx, eq)
	if err != nil {
		return nil, err
	}
	var locs []*ent.Location
	var pl *ent.Location
	if len(ph) > 0 {
		pl = ph[0].QueryParent().QueryLocation().OnlyX(ctx)
	} else {
		pl = eq.QueryLocation().OnlyX(ctx)
	}
	for pl != nil {
		locs = append([]*ent.Location{pl}, locs...)
		l, err := pl.QueryParent().Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		pl = l
	}
	return locs, nil
}

func uriFromDeviceID(deviceID string, hostname string) string {
	slices := strings.Split(deviceID, ".")
	return "https://" + hostname + "/magma/v1/networks/" + slices[1] + "/gateways/" + slices[0] + "/status"
}

func (r equipmentResolver) Device(ctx context.Context, eq *ent.Equipment) (*models.Device, error) {
	var dev models.Device

	if eq.DeviceID == "" {
		return nil, nil
	}

	if r.orc8rClient == nil {
		return nil, errors.New("orc8r client was not provided")
	}

	uri := uriFromDeviceID(eq.DeviceID, r.orc8rClient.Hostname)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.resolver.orc8rClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var jsonResult struct {
		CheckinTime int64 `json:"checkin_time"`
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonResult)
	if err != nil {
		return nil, err
	}

	if jsonResult.CheckinTime != 0 {
		up := jsonResult.CheckinTime > time.Now().Add(3*time.Minute).Unix()
		dev.Up = &up
	}

	return &dev, nil
}

func (equipmentResolver) Services(ctx context.Context, obj *ent.Equipment) ([]*ent.Service, error) {
	services, err := obj.QueryService().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querying services where equipment is termination point")
	}

	ids := make([]string, len(services))
	for i, svc := range services {
		ids[i] = svc.ID
	}

	linkServices, err := obj.QueryPorts().QueryLink().QueryService().Where(service.Not(service.IDIn(ids...))).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querying services where equipment connected to link of service")
	}

	services = append(services, linkServices...)

	return services, nil
}
