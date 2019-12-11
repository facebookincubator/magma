// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolverutil

import (
	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentport"
	"github.com/facebookincubator/symphony/graph/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentporttype"
	"github.com/facebookincubator/symphony/graph/ent/equipmentposition"
	"github.com/facebookincubator/symphony/graph/ent/equipmenttype"
	"github.com/facebookincubator/symphony/graph/ent/link"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/property"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
	"github.com/facebookincubator/symphony/graph/ent/service"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/pkg/errors"
)

func handleLinkFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.FilterType == models.LinkFilterTypeLinkFutureStatus {
		return stateFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func stateFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		p := link.FutureStateIn(filter.IDSet...)
		if find(filter.IDSet, models.FutureStateInstall.String()) {
			p = link.Or(p, link.FutureStateIsNil())
		}
		return q.Where(p), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func handleLinkLocationFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.FilterType == models.LinkFilterTypeLocationInst {
		return linkLocationFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func linkLocationFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		var ps []predicate.Link
		for _, lid := range filter.IDSet {
			ps = append(ps, link.HasPortsWith(GetPortLocationPredicate(lid, filter.MaxDepth)))
		}
		return q.Where(link.Or(ps...)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func handleLinkEquipmentFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.FilterType == models.LinkFilterTypeEquipmentType {
		return linkEquipmentTypeFilter(q, filter)
	} else if filter.FilterType == models.LinkFilterTypeEquipmentInst {
		return linkEquipmentFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func linkEquipmentTypeFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(link.HasPortsWith(equipmentport.HasParentWith(equipment.HasTypeWith(equipmenttype.IDIn(filter.IDSet...))))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

// BuildEquipmentAncestorFilter returns a joined predicate for equipment ancestors
func BuildEquipmentAncestorFilter(equipmentIDs []string, depth int, maxDepth int) predicate.Equipment {
	if depth >= maxDepth {
		return equipment.IDIn(equipmentIDs...)
	}

	return equipment.Or(
		equipment.IDIn(equipmentIDs...),
		equipment.HasParentPositionWith(
			equipmentposition.HasParentWith(BuildEquipmentAncestorFilter(equipmentIDs, depth+1, maxDepth))))
}

func linkEquipmentFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(link.HasPortsWith(
			equipmentport.HasParentWith(BuildEquipmentAncestorFilter(filter.IDSet, 1, *filter.MaxDepth)))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func handleLinkServiceFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.FilterType == models.LinkFilterTypeServiceInst {
		return linkServiceFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func linkServiceFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(link.HasServiceWith(service.IDIn(filter.IDSet...))), nil
	} else if filter.Operator == models.FilterOperatorIsNotOneOf {
		return q.Where(link.Not(
			link.HasServiceWith(service.IDIn(filter.IDSet...)),
		),
		), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func find(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func handleLinkPropertyFilter(q *ent.LinkQuery, filter *models.LinkFilterInput) (*ent.LinkQuery, error) {
	p := filter.PropertyValue
	switch filter.Operator {
	case models.FilterOperatorIs:
		propPred, err := GetPropertyPredicate(*p)
		if err != nil {
			return nil, err
		}

		propTypePred, err := GetPropertyTypePredicate(*p)
		if err != nil {
			return nil, err
		}
		q = q.Where(link.Or(
			link.HasPropertiesWith(
				property.And(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.Type(p.Type.String()),
					),
					propPred,
				),
			),
			link.And(
				link.HasPortsWith(equipmentport.HasDefinitionWith(equipmentportdefinition.HasEquipmentPortTypeWith(equipmentporttype.HasLinkPropertyTypesWith(
					propertytype.Name(p.Name),
					propertytype.Type(p.Type.String()),
					propTypePred,
				)))),
				link.Not(link.HasPropertiesWith(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.Type(p.Type.String()),
					)),
				)),
		),
		)
		return q, nil
	case models.FilterOperatorDateLessThan, models.FilterOperatorDateGreaterThan:
		propPred, propTypePred, err := GetDatePropertyPred(*p, filter.Operator)
		if err != nil {
			return nil, err
		}
		q = q.Where(link.Or(
			link.HasPropertiesWith(
				property.And(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.Type(p.Type.String()),
					),
					propPred,
				),
			),
			link.And(
				link.HasPortsWith(equipmentport.HasDefinitionWith(equipmentportdefinition.HasEquipmentPortTypeWith(equipmentporttype.HasLinkPropertyTypesWith(
					propertytype.Name(p.Name),
					propertytype.Type(p.Type.String()),
					propTypePred,
				)))),
				link.Not(link.HasPortsWith(equipmentport.HasPropertiesWith(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.Type(p.Type.String()),
					)),
				)))))
		return q, nil
	default:
		return nil, errors.Errorf("operator %q not supported", filter.Operator)
	}
}
