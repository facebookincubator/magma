// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/location"
	"github.com/facebookincubator/symphony/graph/ent/locationtype"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/property"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/graph/resolverutil"

	"github.com/pkg/errors"
)

func handleLocationFilter(q *ent.LocationQuery, filter *models.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.FilterType == models.LocationFilterTypeLocationInst {
		return resolverutil.LocationFilterPredicate(q, filter)
	} else if filter.FilterType == models.LocationFilterTypeLocationInstHasEquipment {
		return locationHasEquipmentFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func locationHasEquipmentFilter(q *ent.LocationQuery, filter *models.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.Operator == models.FilterOperatorIs {
		var pp predicate.Location
		if *filter.BoolValue {
			pp = location.HasEquipment()
		} else {
			pp = location.Not(location.HasEquipment())
		}
		return q.Where(pp), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func handleLocationTypeFilter(q *ent.LocationQuery, filter *models.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.FilterType == models.LocationFilterTypeLocationType {
		return locationLocationTypeFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func locationLocationTypeFilter(q *ent.LocationQuery, filter *models.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(location.HasTypeWith(locationtype.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func handleLocationPropertyFilter(q *ent.LocationQuery, filter *models.LocationFilterInput) (*ent.LocationQuery, error) {
	p := filter.PropertyValue
	switch filter.Operator {
	case models.FilterOperatorIs:
		q = q.Where(
			location.HasPropertiesWith(
				property.HasTypeWith(
					propertytype.Name(p.Name),
					propertytype.Type(p.Type.String()),
				),
			),
		)
		pred, err := resolverutil.GetPropertyPredicate(*p)
		if err != nil {
			return nil, err
		}
		if pred != nil {
			q = q.Where(location.HasPropertiesWith(pred))
		}
		return q, nil
	default:
		return nil, errors.Errorf("operator %q not supported", filter.Operator)
	}
}
