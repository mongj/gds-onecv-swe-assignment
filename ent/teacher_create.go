// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
)

// TeacherCreate is the builder for creating a Teacher entity.
type TeacherCreate struct {
	config
	mutation *TeacherMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEmail sets the "email" field.
func (tc *TeacherCreate) SetEmail(s string) *TeacherCreate {
	tc.mutation.SetEmail(s)
	return tc
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (tc *TeacherCreate) AddStudentIDs(ids ...int) *TeacherCreate {
	tc.mutation.AddStudentIDs(ids...)
	return tc
}

// AddStudents adds the "students" edges to the Student entity.
func (tc *TeacherCreate) AddStudents(s ...*Student) *TeacherCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddStudentIDs(ids...)
}

// Mutation returns the TeacherMutation object of the builder.
func (tc *TeacherCreate) Mutation() *TeacherMutation {
	return tc.mutation
}

// Save creates the Teacher in the database.
func (tc *TeacherCreate) Save(ctx context.Context) (*Teacher, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeacherCreate) SaveX(ctx context.Context) *Teacher {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeacherCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeacherCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeacherCreate) check() error {
	if _, ok := tc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Teacher.email"`)}
	}
	if v, ok := tc.mutation.Email(); ok {
		if err := teacher.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Teacher.email": %w`, err)}
		}
	}
	return nil
}

func (tc *TeacherCreate) sqlSave(ctx context.Context) (*Teacher, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TeacherCreate) createSpec() (*Teacher, *sqlgraph.CreateSpec) {
	var (
		_node = &Teacher{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(teacher.Table, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Email(); ok {
		_spec.SetField(teacher.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if nodes := tc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   teacher.StudentsTable,
			Columns: teacher.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Teacher.Create().
//		SetEmail(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeacherUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (tc *TeacherCreate) OnConflict(opts ...sql.ConflictOption) *TeacherUpsertOne {
	tc.conflict = opts
	return &TeacherUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Teacher.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TeacherCreate) OnConflictColumns(columns ...string) *TeacherUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TeacherUpsertOne{
		create: tc,
	}
}

type (
	// TeacherUpsertOne is the builder for "upsert"-ing
	//  one Teacher node.
	TeacherUpsertOne struct {
		create *TeacherCreate
	}

	// TeacherUpsert is the "OnConflict" setter.
	TeacherUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *TeacherUpsert) SetEmail(v string) *TeacherUpsert {
	u.Set(teacher.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *TeacherUpsert) UpdateEmail() *TeacherUpsert {
	u.SetExcluded(teacher.FieldEmail)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Teacher.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TeacherUpsertOne) UpdateNewValues() *TeacherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Teacher.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TeacherUpsertOne) Ignore() *TeacherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeacherUpsertOne) DoNothing() *TeacherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeacherCreate.OnConflict
// documentation for more info.
func (u *TeacherUpsertOne) Update(set func(*TeacherUpsert)) *TeacherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeacherUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *TeacherUpsertOne) SetEmail(v string) *TeacherUpsertOne {
	return u.Update(func(s *TeacherUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *TeacherUpsertOne) UpdateEmail() *TeacherUpsertOne {
	return u.Update(func(s *TeacherUpsert) {
		s.UpdateEmail()
	})
}

// Exec executes the query.
func (u *TeacherUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeacherCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeacherUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeacherUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeacherUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeacherCreateBulk is the builder for creating many Teacher entities in bulk.
type TeacherCreateBulk struct {
	config
	err      error
	builders []*TeacherCreate
	conflict []sql.ConflictOption
}

// Save creates the Teacher entities in the database.
func (tcb *TeacherCreateBulk) Save(ctx context.Context) ([]*Teacher, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Teacher, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeacherMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeacherCreateBulk) SaveX(ctx context.Context) []*Teacher {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeacherCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeacherCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Teacher.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeacherUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (tcb *TeacherCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeacherUpsertBulk {
	tcb.conflict = opts
	return &TeacherUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Teacher.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TeacherCreateBulk) OnConflictColumns(columns ...string) *TeacherUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TeacherUpsertBulk{
		create: tcb,
	}
}

// TeacherUpsertBulk is the builder for "upsert"-ing
// a bulk of Teacher nodes.
type TeacherUpsertBulk struct {
	create *TeacherCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Teacher.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TeacherUpsertBulk) UpdateNewValues() *TeacherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Teacher.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TeacherUpsertBulk) Ignore() *TeacherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeacherUpsertBulk) DoNothing() *TeacherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeacherCreateBulk.OnConflict
// documentation for more info.
func (u *TeacherUpsertBulk) Update(set func(*TeacherUpsert)) *TeacherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeacherUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *TeacherUpsertBulk) SetEmail(v string) *TeacherUpsertBulk {
	return u.Update(func(s *TeacherUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *TeacherUpsertBulk) UpdateEmail() *TeacherUpsertBulk {
	return u.Update(func(s *TeacherUpsert) {
		s.UpdateEmail()
	})
}

// Exec executes the query.
func (u *TeacherUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TeacherCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeacherCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeacherUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
