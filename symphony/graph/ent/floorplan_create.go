// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/floorplan"
)

// FloorPlanCreate is the builder for creating a FloorPlan entity.
type FloorPlanCreate struct {
	config
	create_time     *time.Time
	update_time     *time.Time
	name            *string
	location        map[string]struct{}
	reference_point map[string]struct{}
	scale           map[string]struct{}
	image           map[string]struct{}
}

// SetCreateTime sets the create_time field.
func (fpc *FloorPlanCreate) SetCreateTime(t time.Time) *FloorPlanCreate {
	fpc.create_time = &t
	return fpc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (fpc *FloorPlanCreate) SetNillableCreateTime(t *time.Time) *FloorPlanCreate {
	if t != nil {
		fpc.SetCreateTime(*t)
	}
	return fpc
}

// SetUpdateTime sets the update_time field.
func (fpc *FloorPlanCreate) SetUpdateTime(t time.Time) *FloorPlanCreate {
	fpc.update_time = &t
	return fpc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (fpc *FloorPlanCreate) SetNillableUpdateTime(t *time.Time) *FloorPlanCreate {
	if t != nil {
		fpc.SetUpdateTime(*t)
	}
	return fpc
}

// SetName sets the name field.
func (fpc *FloorPlanCreate) SetName(s string) *FloorPlanCreate {
	fpc.name = &s
	return fpc
}

// SetLocationID sets the location edge to Location by id.
func (fpc *FloorPlanCreate) SetLocationID(id string) *FloorPlanCreate {
	if fpc.location == nil {
		fpc.location = make(map[string]struct{})
	}
	fpc.location[id] = struct{}{}
	return fpc
}

// SetNillableLocationID sets the location edge to Location by id if the given value is not nil.
func (fpc *FloorPlanCreate) SetNillableLocationID(id *string) *FloorPlanCreate {
	if id != nil {
		fpc = fpc.SetLocationID(*id)
	}
	return fpc
}

// SetLocation sets the location edge to Location.
func (fpc *FloorPlanCreate) SetLocation(l *Location) *FloorPlanCreate {
	return fpc.SetLocationID(l.ID)
}

// SetReferencePointID sets the reference_point edge to FloorPlanReferencePoint by id.
func (fpc *FloorPlanCreate) SetReferencePointID(id string) *FloorPlanCreate {
	if fpc.reference_point == nil {
		fpc.reference_point = make(map[string]struct{})
	}
	fpc.reference_point[id] = struct{}{}
	return fpc
}

// SetNillableReferencePointID sets the reference_point edge to FloorPlanReferencePoint by id if the given value is not nil.
func (fpc *FloorPlanCreate) SetNillableReferencePointID(id *string) *FloorPlanCreate {
	if id != nil {
		fpc = fpc.SetReferencePointID(*id)
	}
	return fpc
}

// SetReferencePoint sets the reference_point edge to FloorPlanReferencePoint.
func (fpc *FloorPlanCreate) SetReferencePoint(f *FloorPlanReferencePoint) *FloorPlanCreate {
	return fpc.SetReferencePointID(f.ID)
}

// SetScaleID sets the scale edge to FloorPlanScale by id.
func (fpc *FloorPlanCreate) SetScaleID(id string) *FloorPlanCreate {
	if fpc.scale == nil {
		fpc.scale = make(map[string]struct{})
	}
	fpc.scale[id] = struct{}{}
	return fpc
}

// SetNillableScaleID sets the scale edge to FloorPlanScale by id if the given value is not nil.
func (fpc *FloorPlanCreate) SetNillableScaleID(id *string) *FloorPlanCreate {
	if id != nil {
		fpc = fpc.SetScaleID(*id)
	}
	return fpc
}

// SetScale sets the scale edge to FloorPlanScale.
func (fpc *FloorPlanCreate) SetScale(f *FloorPlanScale) *FloorPlanCreate {
	return fpc.SetScaleID(f.ID)
}

// SetImageID sets the image edge to File by id.
func (fpc *FloorPlanCreate) SetImageID(id string) *FloorPlanCreate {
	if fpc.image == nil {
		fpc.image = make(map[string]struct{})
	}
	fpc.image[id] = struct{}{}
	return fpc
}

// SetNillableImageID sets the image edge to File by id if the given value is not nil.
func (fpc *FloorPlanCreate) SetNillableImageID(id *string) *FloorPlanCreate {
	if id != nil {
		fpc = fpc.SetImageID(*id)
	}
	return fpc
}

// SetImage sets the image edge to File.
func (fpc *FloorPlanCreate) SetImage(f *File) *FloorPlanCreate {
	return fpc.SetImageID(f.ID)
}

// Save creates the FloorPlan in the database.
func (fpc *FloorPlanCreate) Save(ctx context.Context) (*FloorPlan, error) {
	if fpc.create_time == nil {
		v := floorplan.DefaultCreateTime()
		fpc.create_time = &v
	}
	if fpc.update_time == nil {
		v := floorplan.DefaultUpdateTime()
		fpc.update_time = &v
	}
	if fpc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(fpc.location) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"location\"")
	}
	if len(fpc.reference_point) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"reference_point\"")
	}
	if len(fpc.scale) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"scale\"")
	}
	if len(fpc.image) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"image\"")
	}
	return fpc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (fpc *FloorPlanCreate) SaveX(ctx context.Context) *FloorPlan {
	v, err := fpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fpc *FloorPlanCreate) sqlSave(ctx context.Context) (*FloorPlan, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(fpc.driver.Dialect())
		fp      = &FloorPlan{config: fpc.config}
	)
	tx, err := fpc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(floorplan.Table).Default()
	if value := fpc.create_time; value != nil {
		insert.Set(floorplan.FieldCreateTime, *value)
		fp.CreateTime = *value
	}
	if value := fpc.update_time; value != nil {
		insert.Set(floorplan.FieldUpdateTime, *value)
		fp.UpdateTime = *value
	}
	if value := fpc.name; value != nil {
		insert.Set(floorplan.FieldName, *value)
		fp.Name = *value
	}
	id, err := insertLastID(ctx, tx, insert.Returning(floorplan.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	fp.ID = strconv.FormatInt(id, 10)
	if len(fpc.location) > 0 {
		for eid := range fpc.location {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(floorplan.LocationTable).
				Set(floorplan.LocationColumn, eid).
				Where(sql.EQ(floorplan.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(fpc.reference_point) > 0 {
		for eid := range fpc.reference_point {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(floorplan.ReferencePointTable).
				Set(floorplan.ReferencePointColumn, eid).
				Where(sql.EQ(floorplan.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(fpc.scale) > 0 {
		for eid := range fpc.scale {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(floorplan.ScaleTable).
				Set(floorplan.ScaleColumn, eid).
				Where(sql.EQ(floorplan.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(fpc.image) > 0 {
		for eid := range fpc.image {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(floorplan.ImageTable).
				Set(floorplan.ImageColumn, eid).
				Where(sql.EQ(floorplan.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return fp, nil
}
