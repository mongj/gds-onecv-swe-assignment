// Code generated by ent, DO NOT EDIT.

package student

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mongj/gds-onecv-swe-assignment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEmail, v))
}

// IsSuspended applies equality check predicate on the "is_suspended" field. It's identical to IsSuspendedEQ.
func IsSuspended(v bool) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldIsSuspended, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldEmail, v))
}

// IsSuspendedEQ applies the EQ predicate on the "is_suspended" field.
func IsSuspendedEQ(v bool) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldIsSuspended, v))
}

// IsSuspendedNEQ applies the NEQ predicate on the "is_suspended" field.
func IsSuspendedNEQ(v bool) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldIsSuspended, v))
}

// HasTeachers applies the HasEdge predicate on the "teachers" edge.
func HasTeachers() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TeachersTable, TeachersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeachersWith applies the HasEdge predicate on the "teachers" edge with a given conditions (other predicates).
func HasTeachersWith(preds ...predicate.Teacher) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := newTeachersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(sql.NotPredicates(p))
}
