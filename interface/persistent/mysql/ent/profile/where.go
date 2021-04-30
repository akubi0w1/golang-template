// Code generated by entc, DO NOT EDIT.

package profile

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// AvatarURL applies equality check predicate on the "avatar_url" field. It's identical to AvatarURLEQ.
func AvatarURL(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatarURL), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AvatarURLEQ applies the EQ predicate on the "avatar_url" field.
func AvatarURLEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLNEQ applies the NEQ predicate on the "avatar_url" field.
func AvatarURLNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLIn applies the In predicate on the "avatar_url" field.
func AvatarURLIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvatarURL), v...))
	})
}

// AvatarURLNotIn applies the NotIn predicate on the "avatar_url" field.
func AvatarURLNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvatarURL), v...))
	})
}

// AvatarURLGT applies the GT predicate on the "avatar_url" field.
func AvatarURLGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLGTE applies the GTE predicate on the "avatar_url" field.
func AvatarURLGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLLT applies the LT predicate on the "avatar_url" field.
func AvatarURLLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLLTE applies the LTE predicate on the "avatar_url" field.
func AvatarURLLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLContains applies the Contains predicate on the "avatar_url" field.
func AvatarURLContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLHasPrefix applies the HasPrefix predicate on the "avatar_url" field.
func AvatarURLHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLHasSuffix applies the HasSuffix predicate on the "avatar_url" field.
func AvatarURLHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLEqualFold applies the EqualFold predicate on the "avatar_url" field.
func AvatarURLEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatarURL), v))
	})
}

// AvatarURLContainsFold applies the ContainsFold predicate on the "avatar_url" field.
func AvatarURLContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatarURL), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		p(s.Not())
	})
}
