// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package checklistitem

import (
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.CheckListItem {
	return predicate.CheckListItem(
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
func IDNotIn(ids ...string) predicate.CheckListItem {
	return predicate.CheckListItem(
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
func IDGT(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldTitle), v))
		},
	)
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldType), v))
		},
	)
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldIndex), v))
		},
	)
}

// Checked applies equality check predicate on the "checked" field. It's identical to CheckedEQ.
func Checked(v bool) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldChecked), v))
		},
	)
}

// StringVal applies equality check predicate on the "string_val" field. It's identical to StringValEQ.
func StringVal(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldStringVal), v))
		},
	)
}

// EnumValues applies equality check predicate on the "enum_values" field. It's identical to EnumValuesEQ.
func EnumValues(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldEnumValues), v))
		},
	)
}

// HelpText applies equality check predicate on the "help_text" field. It's identical to HelpTextEQ.
func HelpText(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldHelpText), v))
		},
	)
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldTitle), v))
		},
	)
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldTitle), v))
		},
	)
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldTitle), v...))
		},
	)
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldTitle), v...))
		},
	)
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldTitle), v))
		},
	)
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldTitle), v))
		},
	)
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldTitle), v))
		},
	)
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldTitle), v))
		},
	)
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldTitle), v))
		},
	)
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldTitle), v))
		},
	)
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldTitle), v))
		},
	)
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldTitle), v))
		},
	)
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldTitle), v))
		},
	)
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldType), v))
		},
	)
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldType), v))
		},
	)
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldType), v...))
		},
	)
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldType), v...))
		},
	)
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldType), v))
		},
	)
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldType), v))
		},
	)
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldType), v))
		},
	)
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldType), v))
		},
	)
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldType), v))
		},
	)
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldType), v))
		},
	)
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldType), v))
		},
	)
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldType), v))
		},
	)
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldType), v))
		},
	)
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldIndex), v))
		},
	)
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldIndex), v))
		},
	)
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldIndex), v...))
		},
	)
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldIndex), v...))
		},
	)
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldIndex), v))
		},
	)
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldIndex), v))
		},
	)
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldIndex), v))
		},
	)
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldIndex), v))
		},
	)
}

// IndexIsNil applies the IsNil predicate on the "index" field.
func IndexIsNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldIndex)))
		},
	)
}

// IndexNotNil applies the NotNil predicate on the "index" field.
func IndexNotNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldIndex)))
		},
	)
}

// CheckedEQ applies the EQ predicate on the "checked" field.
func CheckedEQ(v bool) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldChecked), v))
		},
	)
}

// CheckedNEQ applies the NEQ predicate on the "checked" field.
func CheckedNEQ(v bool) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldChecked), v))
		},
	)
}

// CheckedIsNil applies the IsNil predicate on the "checked" field.
func CheckedIsNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldChecked)))
		},
	)
}

// CheckedNotNil applies the NotNil predicate on the "checked" field.
func CheckedNotNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldChecked)))
		},
	)
}

// StringValEQ applies the EQ predicate on the "string_val" field.
func StringValEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldStringVal), v))
		},
	)
}

// StringValNEQ applies the NEQ predicate on the "string_val" field.
func StringValNEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldStringVal), v))
		},
	)
}

// StringValIn applies the In predicate on the "string_val" field.
func StringValIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldStringVal), v...))
		},
	)
}

// StringValNotIn applies the NotIn predicate on the "string_val" field.
func StringValNotIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldStringVal), v...))
		},
	)
}

// StringValGT applies the GT predicate on the "string_val" field.
func StringValGT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldStringVal), v))
		},
	)
}

// StringValGTE applies the GTE predicate on the "string_val" field.
func StringValGTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldStringVal), v))
		},
	)
}

// StringValLT applies the LT predicate on the "string_val" field.
func StringValLT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldStringVal), v))
		},
	)
}

// StringValLTE applies the LTE predicate on the "string_val" field.
func StringValLTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldStringVal), v))
		},
	)
}

// StringValContains applies the Contains predicate on the "string_val" field.
func StringValContains(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldStringVal), v))
		},
	)
}

// StringValHasPrefix applies the HasPrefix predicate on the "string_val" field.
func StringValHasPrefix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldStringVal), v))
		},
	)
}

// StringValHasSuffix applies the HasSuffix predicate on the "string_val" field.
func StringValHasSuffix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldStringVal), v))
		},
	)
}

// StringValIsNil applies the IsNil predicate on the "string_val" field.
func StringValIsNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldStringVal)))
		},
	)
}

// StringValNotNil applies the NotNil predicate on the "string_val" field.
func StringValNotNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldStringVal)))
		},
	)
}

// StringValEqualFold applies the EqualFold predicate on the "string_val" field.
func StringValEqualFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldStringVal), v))
		},
	)
}

// StringValContainsFold applies the ContainsFold predicate on the "string_val" field.
func StringValContainsFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldStringVal), v))
		},
	)
}

// EnumValuesEQ applies the EQ predicate on the "enum_values" field.
func EnumValuesEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesNEQ applies the NEQ predicate on the "enum_values" field.
func EnumValuesNEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesIn applies the In predicate on the "enum_values" field.
func EnumValuesIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldEnumValues), v...))
		},
	)
}

// EnumValuesNotIn applies the NotIn predicate on the "enum_values" field.
func EnumValuesNotIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldEnumValues), v...))
		},
	)
}

// EnumValuesGT applies the GT predicate on the "enum_values" field.
func EnumValuesGT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesGTE applies the GTE predicate on the "enum_values" field.
func EnumValuesGTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesLT applies the LT predicate on the "enum_values" field.
func EnumValuesLT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesLTE applies the LTE predicate on the "enum_values" field.
func EnumValuesLTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesContains applies the Contains predicate on the "enum_values" field.
func EnumValuesContains(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesHasPrefix applies the HasPrefix predicate on the "enum_values" field.
func EnumValuesHasPrefix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesHasSuffix applies the HasSuffix predicate on the "enum_values" field.
func EnumValuesHasSuffix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesIsNil applies the IsNil predicate on the "enum_values" field.
func EnumValuesIsNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldEnumValues)))
		},
	)
}

// EnumValuesNotNil applies the NotNil predicate on the "enum_values" field.
func EnumValuesNotNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldEnumValues)))
		},
	)
}

// EnumValuesEqualFold applies the EqualFold predicate on the "enum_values" field.
func EnumValuesEqualFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldEnumValues), v))
		},
	)
}

// EnumValuesContainsFold applies the ContainsFold predicate on the "enum_values" field.
func EnumValuesContainsFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldEnumValues), v))
		},
	)
}

// HelpTextEQ applies the EQ predicate on the "help_text" field.
func HelpTextEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextNEQ applies the NEQ predicate on the "help_text" field.
func HelpTextNEQ(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextIn applies the In predicate on the "help_text" field.
func HelpTextIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldHelpText), v...))
		},
	)
}

// HelpTextNotIn applies the NotIn predicate on the "help_text" field.
func HelpTextNotIn(vs ...string) predicate.CheckListItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldHelpText), v...))
		},
	)
}

// HelpTextGT applies the GT predicate on the "help_text" field.
func HelpTextGT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextGTE applies the GTE predicate on the "help_text" field.
func HelpTextGTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextLT applies the LT predicate on the "help_text" field.
func HelpTextLT(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextLTE applies the LTE predicate on the "help_text" field.
func HelpTextLTE(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextContains applies the Contains predicate on the "help_text" field.
func HelpTextContains(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextHasPrefix applies the HasPrefix predicate on the "help_text" field.
func HelpTextHasPrefix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextHasSuffix applies the HasSuffix predicate on the "help_text" field.
func HelpTextHasSuffix(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextIsNil applies the IsNil predicate on the "help_text" field.
func HelpTextIsNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldHelpText)))
		},
	)
}

// HelpTextNotNil applies the NotNil predicate on the "help_text" field.
func HelpTextNotNil() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldHelpText)))
		},
	)
}

// HelpTextEqualFold applies the EqualFold predicate on the "help_text" field.
func HelpTextEqualFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldHelpText), v))
		},
	)
}

// HelpTextContainsFold applies the ContainsFold predicate on the "help_text" field.
func HelpTextContainsFold(v string) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldHelpText), v))
		},
	)
}

// HasWorkOrder applies the HasEdge predicate on the "work_order" edge.
func HasWorkOrder() predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(WorkOrderColumn)))
		},
	)
}

// HasWorkOrderWith applies the HasEdge predicate on the "work_order" edge with a given conditions (other predicates).
func HasWorkOrderWith(preds ...predicate.WorkOrder) predicate.CheckListItem {
	return predicate.CheckListItem(
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

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.CheckListItem) predicate.CheckListItem {
	return predicate.CheckListItem(
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
func Or(predicates ...predicate.CheckListItem) predicate.CheckListItem {
	return predicate.CheckListItem(
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
func Not(p predicate.CheckListItem) predicate.CheckListItem {
	return predicate.CheckListItem(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
