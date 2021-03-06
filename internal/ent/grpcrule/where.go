// Code generated by entc, DO NOT EDIT.

package grpcrule

import (
	"go-gateway/internal/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ServiceID applies equality check predicate on the "service_id" field. It's identical to ServiceIDEQ.
func ServiceID(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceID), v))
	})
}

// Port applies equality check predicate on the "Port" field. It's identical to PortEQ.
func Port(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPort), v))
	})
}

// HeaderTransfor applies equality check predicate on the "header_transfor" field. It's identical to HeaderTransforEQ.
func HeaderTransfor(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeaderTransfor), v))
	})
}

// ServiceIDEQ applies the EQ predicate on the "service_id" field.
func ServiceIDEQ(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceID), v))
	})
}

// ServiceIDNEQ applies the NEQ predicate on the "service_id" field.
func ServiceIDNEQ(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServiceID), v))
	})
}

// ServiceIDIn applies the In predicate on the "service_id" field.
func ServiceIDIn(vs ...int64) predicate.GrpcRule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GrpcRule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldServiceID), v...))
	})
}

// ServiceIDNotIn applies the NotIn predicate on the "service_id" field.
func ServiceIDNotIn(vs ...int64) predicate.GrpcRule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GrpcRule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldServiceID), v...))
	})
}

// ServiceIDGT applies the GT predicate on the "service_id" field.
func ServiceIDGT(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServiceID), v))
	})
}

// ServiceIDGTE applies the GTE predicate on the "service_id" field.
func ServiceIDGTE(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServiceID), v))
	})
}

// ServiceIDLT applies the LT predicate on the "service_id" field.
func ServiceIDLT(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServiceID), v))
	})
}

// ServiceIDLTE applies the LTE predicate on the "service_id" field.
func ServiceIDLTE(v int64) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServiceID), v))
	})
}

// PortEQ applies the EQ predicate on the "Port" field.
func PortEQ(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPort), v))
	})
}

// PortNEQ applies the NEQ predicate on the "Port" field.
func PortNEQ(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPort), v))
	})
}

// PortIn applies the In predicate on the "Port" field.
func PortIn(vs ...int) predicate.GrpcRule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GrpcRule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPort), v...))
	})
}

// PortNotIn applies the NotIn predicate on the "Port" field.
func PortNotIn(vs ...int) predicate.GrpcRule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GrpcRule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPort), v...))
	})
}

// PortGT applies the GT predicate on the "Port" field.
func PortGT(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPort), v))
	})
}

// PortGTE applies the GTE predicate on the "Port" field.
func PortGTE(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPort), v))
	})
}

// PortLT applies the LT predicate on the "Port" field.
func PortLT(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPort), v))
	})
}

// PortLTE applies the LTE predicate on the "Port" field.
func PortLTE(v int) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPort), v))
	})
}

// HeaderTransforEQ applies the EQ predicate on the "header_transfor" field.
func HeaderTransforEQ(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforNEQ applies the NEQ predicate on the "header_transfor" field.
func HeaderTransforNEQ(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforIn applies the In predicate on the "header_transfor" field.
func HeaderTransforIn(vs ...string) predicate.GrpcRule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GrpcRule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHeaderTransfor), v...))
	})
}

// HeaderTransforNotIn applies the NotIn predicate on the "header_transfor" field.
func HeaderTransforNotIn(vs ...string) predicate.GrpcRule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GrpcRule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHeaderTransfor), v...))
	})
}

// HeaderTransforGT applies the GT predicate on the "header_transfor" field.
func HeaderTransforGT(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforGTE applies the GTE predicate on the "header_transfor" field.
func HeaderTransforGTE(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforLT applies the LT predicate on the "header_transfor" field.
func HeaderTransforLT(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforLTE applies the LTE predicate on the "header_transfor" field.
func HeaderTransforLTE(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforContains applies the Contains predicate on the "header_transfor" field.
func HeaderTransforContains(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforHasPrefix applies the HasPrefix predicate on the "header_transfor" field.
func HeaderTransforHasPrefix(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforHasSuffix applies the HasSuffix predicate on the "header_transfor" field.
func HeaderTransforHasSuffix(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforEqualFold applies the EqualFold predicate on the "header_transfor" field.
func HeaderTransforEqualFold(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHeaderTransfor), v))
	})
}

// HeaderTransforContainsFold applies the ContainsFold predicate on the "header_transfor" field.
func HeaderTransforContainsFold(v string) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHeaderTransfor), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.GrpcRule) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.GrpcRule) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
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
func Not(p predicate.GrpcRule) predicate.GrpcRule {
	return predicate.GrpcRule(func(s *sql.Selector) {
		p(s.Not())
	})
}
