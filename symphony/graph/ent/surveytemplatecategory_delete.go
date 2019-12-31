// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/surveytemplatecategory"
)

// SurveyTemplateCategoryDelete is the builder for deleting a SurveyTemplateCategory entity.
type SurveyTemplateCategoryDelete struct {
	config
	predicates []predicate.SurveyTemplateCategory
}

// Where adds a new predicate to the delete builder.
func (stcd *SurveyTemplateCategoryDelete) Where(ps ...predicate.SurveyTemplateCategory) *SurveyTemplateCategoryDelete {
	stcd.predicates = append(stcd.predicates, ps...)
	return stcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (stcd *SurveyTemplateCategoryDelete) Exec(ctx context.Context) (int, error) {
	return stcd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (stcd *SurveyTemplateCategoryDelete) ExecX(ctx context.Context) int {
	n, err := stcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (stcd *SurveyTemplateCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: surveytemplatecategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: surveytemplatecategory.FieldID,
			},
		},
	}
	if ps := stcd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, stcd.driver, spec)
}

// SurveyTemplateCategoryDeleteOne is the builder for deleting a single SurveyTemplateCategory entity.
type SurveyTemplateCategoryDeleteOne struct {
	stcd *SurveyTemplateCategoryDelete
}

// Exec executes the deletion query.
func (stcdo *SurveyTemplateCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := stcdo.stcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{surveytemplatecategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stcdo *SurveyTemplateCategoryDeleteOne) ExecX(ctx context.Context) {
	stcdo.stcd.ExecX(ctx)
}
