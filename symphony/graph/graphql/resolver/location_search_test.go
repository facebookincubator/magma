// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"
	"testing"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/graph/viewer/viewertest"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/require"
)

type locationSearchDataModels struct {
	loc1     *ent.Location
	loc2     *ent.Location
	locType1 *ent.LocationType
	locType2 *ent.LocationType
}

// nolint: errcheck
func prepareLocationData(ctx context.Context, r *TestResolver, props []*models.PropertyInput) locationSearchDataModels {
	mr := r.Mutation()
	locType1, _ := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "loc_type1",
	})

	loc1, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "loc_inst1",
		Type: locType1.ID,
	})

	locType2, _ := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "loc_type2",
	})

	loc2, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name:   "loc_inst2",
		Type:   locType2.ID,
		Parent: pointer.ToString(loc1.ID),
	})

	equType, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "eq_type",
	})
	mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst",
		Type:     equType.ID,
		Location: &loc1.ID,
	})
	return locationSearchDataModels{
		loc1,
		loc2,
		locType1,
		locType2,
	}
}

func TestSearchLocationAncestors(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	data := prepareLocationData(ctx, r, nil)
	/*
		helper: data now is of type:
		 loc1 (loc_type1):
			eq_inst (eq_type)
			loc2 (loc_type2)
	*/
	qr := r.Query()
	limit := 100
	all, err := qr.LocationSearch(ctx, []*models.LocationFilterInput{}, &limit)
	require.NoError(t, err)
	require.Len(t, all.Locations, 2)
	require.Equal(t, all.Count, 2)
	maxDepth := 2
	f1 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationInst,
		Operator:   models.FilterOperatorIsOneOf,
		IDSet:      []string{data.loc1.ID},
		MaxDepth:   &maxDepth,
	}
	res, err := qr.LocationSearch(ctx, []*models.LocationFilterInput{&f1}, &limit)
	require.NoError(t, err)
	require.Len(t, res.Locations, 2)
	require.Equal(t, res.Count, 2)

	f2 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationInst,
		Operator:   models.FilterOperatorIsOneOf,
		IDSet:      []string{data.loc2.ID},
		MaxDepth:   &maxDepth,
	}
	res, err = qr.LocationSearch(ctx, []*models.LocationFilterInput{&f2}, &limit)
	require.NoError(t, err)
	require.Len(t, res.Locations, 1)
	require.Equal(t, res.Count, 1)
}

func TestSearchLocationByType(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	data := prepareLocationData(ctx, r, nil)
	/*
		helper: data now is of type:
		 loc1 (loc_type1):
			eq_inst (eq_type)
			loc2 (loc_type2)
	*/
	qr := r.Query()
	f1 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationType,
		Operator:   models.FilterOperatorIsOneOf,
		IDSet:      []string{data.locType2.ID},
	}
	res, err := qr.LocationSearch(ctx, []*models.LocationFilterInput{&f1}, pointer.ToInt(100))
	require.NoError(t, err)
	require.Len(t, res.Locations, 1)
	require.Equal(t, res.Count, 1)
}

func TestSearchLocationHasEquipment(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	prepareLocationData(ctx, r, nil)
	/*
		helper: data now is of type:
		 loc1 (loc_type1):
			eq_inst (eq_type)
			loc2 (loc_type2)
	*/
	qr := r.Query()
	f1 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationInstHasEquipment,
		Operator:   models.FilterOperatorIs,
		BoolValue:  pointer.ToBool(true),
	}
	res, err := qr.LocationSearch(ctx, []*models.LocationFilterInput{&f1}, pointer.ToInt(100))
	require.NoError(t, err)
	require.Len(t, res.Locations, 1)
	require.Equal(t, res.Count, 1)

	f2 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationInstHasEquipment,
		Operator:   models.FilterOperatorIs,
		BoolValue:  pointer.ToBool(false),
	}
	res, err = qr.LocationSearch(ctx, []*models.LocationFilterInput{&f2}, pointer.ToInt(100))
	require.NoError(t, err)
	require.Len(t, res.Locations, 1)
	require.Equal(t, res.Count, 1)
}

func TestSearchMultipleFilters(t *testing.T) {
	r, err := newTestResolver(t)
	require.NoError(t, err)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)

	data := prepareLocationData(ctx, r, nil)
	/*
		helper: data now is of type:
		 loc1 (loc_type1):
			eq_inst (eq_type)
			loc2 (loc_type2)
	*/
	qr := r.Query()
	f1 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationInst,
		Operator:   models.FilterOperatorIsOneOf,
		IDSet:      []string{data.loc1.ID},
		MaxDepth:   pointer.ToInt(2),
	}
	res, err := qr.LocationSearch(ctx, []*models.LocationFilterInput{&f1}, pointer.ToInt(100))
	require.NoError(t, err)
	require.Len(t, res.Locations, 2)
	require.Equal(t, res.Count, 2)

	f2 := models.LocationFilterInput{
		FilterType: models.LocationFilterTypeLocationType,
		Operator:   models.FilterOperatorIsOneOf,
		IDSet:      []string{data.locType2.ID},
	}
	res, err = qr.LocationSearch(ctx, []*models.LocationFilterInput{&f1, &f2}, pointer.ToInt(100))
	require.NoError(t, err)
	require.Len(t, res.Locations, 1)
	require.Equal(t, res.Count, 1)
}
