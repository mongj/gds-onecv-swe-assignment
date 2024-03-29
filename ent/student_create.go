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

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEmail sets the "email" field.
func (sc *StudentCreate) SetEmail(s string) *StudentCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetIsSuspended sets the "is_suspended" field.
func (sc *StudentCreate) SetIsSuspended(b bool) *StudentCreate {
	sc.mutation.SetIsSuspended(b)
	return sc
}

// SetNillableIsSuspended sets the "is_suspended" field if the given value is not nil.
func (sc *StudentCreate) SetNillableIsSuspended(b *bool) *StudentCreate {
	if b != nil {
		sc.SetIsSuspended(*b)
	}
	return sc
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (sc *StudentCreate) AddTeacherIDs(ids ...int) *StudentCreate {
	sc.mutation.AddTeacherIDs(ids...)
	return sc
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (sc *StudentCreate) AddTeachers(t ...*Teacher) *StudentCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTeacherIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StudentCreate) defaults() {
	if _, ok := sc.mutation.IsSuspended(); !ok {
		v := student.DefaultIsSuspended
		sc.mutation.SetIsSuspended(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Student.email"`)}
	}
	if v, ok := sc.mutation.Email(); ok {
		if err := student.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Student.email": %w`, err)}
		}
	}
	if _, ok := sc.mutation.IsSuspended(); !ok {
		return &ValidationError{Name: "is_suspended", err: errors.New(`ent: missing required field "Student.is_suspended"`)}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(student.Table, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Email(); ok {
		_spec.SetField(student.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := sc.mutation.IsSuspended(); ok {
		_spec.SetField(student.FieldIsSuspended, field.TypeBool, value)
		_node.IsSuspended = value
	}
	if nodes := sc.mutation.TeachersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Student.Create().
//		SetEmail(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StudentUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (sc *StudentCreate) OnConflict(opts ...sql.ConflictOption) *StudentUpsertOne {
	sc.conflict = opts
	return &StudentUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *StudentCreate) OnConflictColumns(columns ...string) *StudentUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StudentUpsertOne{
		create: sc,
	}
}

type (
	// StudentUpsertOne is the builder for "upsert"-ing
	//  one Student node.
	StudentUpsertOne struct {
		create *StudentCreate
	}

	// StudentUpsert is the "OnConflict" setter.
	StudentUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *StudentUpsert) SetEmail(v string) *StudentUpsert {
	u.Set(student.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *StudentUpsert) UpdateEmail() *StudentUpsert {
	u.SetExcluded(student.FieldEmail)
	return u
}

// SetIsSuspended sets the "is_suspended" field.
func (u *StudentUpsert) SetIsSuspended(v bool) *StudentUpsert {
	u.Set(student.FieldIsSuspended, v)
	return u
}

// UpdateIsSuspended sets the "is_suspended" field to the value that was provided on create.
func (u *StudentUpsert) UpdateIsSuspended() *StudentUpsert {
	u.SetExcluded(student.FieldIsSuspended)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *StudentUpsertOne) UpdateNewValues() *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Student.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StudentUpsertOne) Ignore() *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StudentUpsertOne) DoNothing() *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StudentCreate.OnConflict
// documentation for more info.
func (u *StudentUpsertOne) Update(set func(*StudentUpsert)) *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StudentUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *StudentUpsertOne) SetEmail(v string) *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *StudentUpsertOne) UpdateEmail() *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateEmail()
	})
}

// SetIsSuspended sets the "is_suspended" field.
func (u *StudentUpsertOne) SetIsSuspended(v bool) *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.SetIsSuspended(v)
	})
}

// UpdateIsSuspended sets the "is_suspended" field to the value that was provided on create.
func (u *StudentUpsertOne) UpdateIsSuspended() *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateIsSuspended()
	})
}

// Exec executes the query.
func (u *StudentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StudentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StudentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StudentUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StudentUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	err      error
	builders []*StudentCreate
	conflict []sql.ConflictOption
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Student.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StudentUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (scb *StudentCreateBulk) OnConflict(opts ...sql.ConflictOption) *StudentUpsertBulk {
	scb.conflict = opts
	return &StudentUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *StudentCreateBulk) OnConflictColumns(columns ...string) *StudentUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StudentUpsertBulk{
		create: scb,
	}
}

// StudentUpsertBulk is the builder for "upsert"-ing
// a bulk of Student nodes.
type StudentUpsertBulk struct {
	create *StudentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *StudentUpsertBulk) UpdateNewValues() *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StudentUpsertBulk) Ignore() *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StudentUpsertBulk) DoNothing() *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StudentCreateBulk.OnConflict
// documentation for more info.
func (u *StudentUpsertBulk) Update(set func(*StudentUpsert)) *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StudentUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *StudentUpsertBulk) SetEmail(v string) *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *StudentUpsertBulk) UpdateEmail() *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateEmail()
	})
}

// SetIsSuspended sets the "is_suspended" field.
func (u *StudentUpsertBulk) SetIsSuspended(v bool) *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.SetIsSuspended(v)
	})
}

// UpdateIsSuspended sets the "is_suspended" field to the value that was provided on create.
func (u *StudentUpsertBulk) UpdateIsSuspended() *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateIsSuspended()
	})
}

// Exec executes the query.
func (u *StudentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StudentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StudentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StudentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
