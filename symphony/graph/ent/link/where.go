// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package link

import (
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i], _ = strconv.Atoi(ids[i])
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i], _ = strconv.Atoi(ids[i])
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreateTime), v))
		},
	)
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdateTime), v))
		},
	)
}

// FutureState applies equality check predicate on the "future_state" field. It's identical to FutureStateEQ.
func FutureState(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldFutureState), v))
		},
	)
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreateTime), v))
		},
	)
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldCreateTime), v))
		},
	)
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldCreateTime), v...))
		},
	)
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
		},
	)
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldCreateTime), v))
		},
	)
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldCreateTime), v))
		},
	)
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldCreateTime), v))
		},
	)
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldCreateTime), v))
		},
	)
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdateTime), v))
		},
	)
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
		},
	)
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldUpdateTime), v...))
		},
	)
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
		},
	)
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldUpdateTime), v))
		},
	)
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldUpdateTime), v))
		},
	)
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldUpdateTime), v))
		},
	)
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldUpdateTime), v))
		},
	)
}

// FutureStateEQ applies the EQ predicate on the "future_state" field.
func FutureStateEQ(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateNEQ applies the NEQ predicate on the "future_state" field.
func FutureStateNEQ(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateIn applies the In predicate on the "future_state" field.
func FutureStateIn(vs ...string) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldFutureState), v...))
		},
	)
}

// FutureStateNotIn applies the NotIn predicate on the "future_state" field.
func FutureStateNotIn(vs ...string) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldFutureState), v...))
		},
	)
}

// FutureStateGT applies the GT predicate on the "future_state" field.
func FutureStateGT(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateGTE applies the GTE predicate on the "future_state" field.
func FutureStateGTE(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateLT applies the LT predicate on the "future_state" field.
func FutureStateLT(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateLTE applies the LTE predicate on the "future_state" field.
func FutureStateLTE(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateContains applies the Contains predicate on the "future_state" field.
func FutureStateContains(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateHasPrefix applies the HasPrefix predicate on the "future_state" field.
func FutureStateHasPrefix(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateHasSuffix applies the HasSuffix predicate on the "future_state" field.
func FutureStateHasSuffix(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateIsNil applies the IsNil predicate on the "future_state" field.
func FutureStateIsNil() predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldFutureState)))
		},
	)
}

// FutureStateNotNil applies the NotNil predicate on the "future_state" field.
func FutureStateNotNil() predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldFutureState)))
		},
	)
}

// FutureStateEqualFold applies the EqualFold predicate on the "future_state" field.
func FutureStateEqualFold(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldFutureState), v))
		},
	)
}

// FutureStateContainsFold applies the ContainsFold predicate on the "future_state" field.
func FutureStateContainsFold(v string) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldFutureState), v))
		},
	)
}

// HasPorts applies the HasEdge predicate on the "ports" edge.
func HasPorts() predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(PortsTable, FieldID),
				sql.Edge(sql.O2M, true, PortsTable, PortsColumn),
			)
			sql.HasNeighbors(s, step)
		},
	)
}

// HasPortsWith applies the HasEdge predicate on the "ports" edge with a given conditions (other predicates).
func HasPortsWith(preds ...predicate.EquipmentPort) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(PortsColumn).From(builder.Table(PortsTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
	)
}

// HasWorkOrder applies the HasEdge predicate on the "work_order" edge.
func HasWorkOrder() predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(WorkOrderTable, FieldID),
				sql.Edge(sql.M2O, false, WorkOrderTable, WorkOrderColumn),
			)
			sql.HasNeighbors(s, step)
		},
	)
}

// HasWorkOrderWith applies the HasEdge predicate on the "work_order" edge with a given conditions (other predicates).
func HasWorkOrderWith(preds ...predicate.WorkOrder) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(FieldID).From(builder.Table(WorkOrderInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(WorkOrderColumn), t2))
		},
	)
}

// HasProperties applies the HasEdge predicate on the "properties" edge.
func HasProperties() predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(PropertiesTable, FieldID),
				sql.Edge(sql.O2M, false, PropertiesTable, PropertiesColumn),
			)
			sql.HasNeighbors(s, step)
		},
	)
}

// HasPropertiesWith applies the HasEdge predicate on the "properties" edge with a given conditions (other predicates).
func HasPropertiesWith(preds ...predicate.Property) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(PropertiesColumn).From(builder.Table(PropertiesTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
	)
}

// HasService applies the HasEdge predicate on the "service" edge.
func HasService() predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(ServiceTable, FieldID),
				sql.Edge(sql.M2M, true, ServiceTable, ServicePrimaryKey...),
			)
			sql.HasNeighbors(s, step)
		},
	)
}

// HasServiceWith applies the HasEdge predicate on the "service" edge with a given conditions (other predicates).
func HasServiceWith(preds ...predicate.Service) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Table(ServiceInverseTable)
			t3 := builder.Table(ServiceTable)
			t4 := builder.Select(t3.C(ServicePrimaryKey[1])).
				From(t3).
				Join(t2).
				On(t3.C(ServicePrimaryKey[0]), t2.C(FieldID))
			t5 := builder.Select().From(t2)
			for _, p := range preds {
				p(t5)
			}
			t4.FromSelect(t5)
			s.Where(sql.In(t1.C(FieldID), t4))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Link) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
