// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mongj/gds-onecv-swe-assignment/ent/predicate"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetEmail sets the "email" field.
func (su *StudentUpdate) SetEmail(s string) *StudentUpdate {
	su.mutation.SetEmail(s)
	return su
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (su *StudentUpdate) SetNillableEmail(s *string) *StudentUpdate {
	if s != nil {
		su.SetEmail(*s)
	}
	return su
}

// SetIsSuspended sets the "is_suspended" field.
func (su *StudentUpdate) SetIsSuspended(b bool) *StudentUpdate {
	su.mutation.SetIsSuspended(b)
	return su
}

// SetNillableIsSuspended sets the "is_suspended" field if the given value is not nil.
func (su *StudentUpdate) SetNillableIsSuspended(b *bool) *StudentUpdate {
	if b != nil {
		su.SetIsSuspended(*b)
	}
	return su
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (su *StudentUpdate) AddTeacherIDs(ids ...int) *StudentUpdate {
	su.mutation.AddTeacherIDs(ids...)
	return su
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (su *StudentUpdate) AddTeachers(t ...*Teacher) *StudentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTeacherIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (su *StudentUpdate) ClearTeachers() *StudentUpdate {
	su.mutation.ClearTeachers()
	return su
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (su *StudentUpdate) RemoveTeacherIDs(ids ...int) *StudentUpdate {
	su.mutation.RemoveTeacherIDs(ids...)
	return su
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (su *StudentUpdate) RemoveTeachers(t ...*Teacher) *StudentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTeacherIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StudentUpdate) check() error {
	if v, ok := su.mutation.Email(); ok {
		if err := student.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Student.email": %w`, err)}
		}
	}
	return nil
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Email(); ok {
		_spec.SetField(student.FieldEmail, field.TypeString, value)
	}
	if value, ok := su.mutation.IsSuspended(); ok {
		_spec.SetField(student.FieldIsSuspended, field.TypeBool, value)
	}
	if su.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !su.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetEmail sets the "email" field.
func (suo *StudentUpdateOne) SetEmail(s string) *StudentUpdateOne {
	suo.mutation.SetEmail(s)
	return suo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableEmail(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetEmail(*s)
	}
	return suo
}

// SetIsSuspended sets the "is_suspended" field.
func (suo *StudentUpdateOne) SetIsSuspended(b bool) *StudentUpdateOne {
	suo.mutation.SetIsSuspended(b)
	return suo
}

// SetNillableIsSuspended sets the "is_suspended" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableIsSuspended(b *bool) *StudentUpdateOne {
	if b != nil {
		suo.SetIsSuspended(*b)
	}
	return suo
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (suo *StudentUpdateOne) AddTeacherIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.AddTeacherIDs(ids...)
	return suo
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (suo *StudentUpdateOne) AddTeachers(t ...*Teacher) *StudentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTeacherIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (suo *StudentUpdateOne) ClearTeachers() *StudentUpdateOne {
	suo.mutation.ClearTeachers()
	return suo
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (suo *StudentUpdateOne) RemoveTeacherIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.RemoveTeacherIDs(ids...)
	return suo
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (suo *StudentUpdateOne) RemoveTeachers(t ...*Teacher) *StudentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTeacherIDs(ids...)
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StudentUpdateOne) check() error {
	if v, ok := suo.mutation.Email(); ok {
		if err := student.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Student.email": %w`, err)}
		}
	}
	return nil
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Email(); ok {
		_spec.SetField(student.FieldEmail, field.TypeString, value)
	}
	if value, ok := suo.mutation.IsSuspended(); ok {
		_spec.SetField(student.FieldIsSuspended, field.TypeBool, value)
	}
	if suo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !suo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
