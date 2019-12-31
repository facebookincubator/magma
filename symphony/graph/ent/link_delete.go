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
	"github.com/facebookincubator/symphony/graph/ent/link"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// LinkDelete is the builder for deleting a Link entity.
type LinkDelete struct {
	config
	predicates []predicate.Link
}

// Where adds a new predicate to the delete builder.
func (ld *LinkDelete) Where(ps ...predicate.Link) *LinkDelete {
	ld.predicates = append(ld.predicates, ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LinkDelete) Exec(ctx context.Context) (int, error) {
	return ld.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LinkDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LinkDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: link.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: link.FieldID,
			},
		},
	}
	if ps := ld.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ld.driver, spec)
}

// LinkDeleteOne is the builder for deleting a single Link entity.
type LinkDeleteOne struct {
	ld *LinkDelete
}

// Exec executes the deletion query.
func (ldo *LinkDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{link.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LinkDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}
