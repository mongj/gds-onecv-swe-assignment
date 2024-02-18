package database

import (
	"context"
	"fmt"

	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
)

// CreateTeacher creates a teacher
func CreateTeacher(client *ent.Client, email string) (*ent.Teacher, error) {
	return client.Teacher.
		Create().
		SetEmail(email).
		Save(context.Background())
}

// BulkCreateTeachers creates teachers in bulk
func BulkCreateTeachers(client *ent.Client, teachers []*ent.Teacher) ([]*ent.Teacher, error) {
	t, err := client.Teacher.
		MapCreateBulk(teachers, func(t *ent.TeacherCreate, i int) {
			t.SetEmail(teachers[i].Email)
		}).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to create teachers: %w", err)
	}
	return t, nil
}

// GetTeacherByEmail gets a teacher by email
func GetTeacherByEmail(client *ent.Client, email string) (*ent.Teacher, error) {
	return client.Teacher.
		Query().
		Where(teacher.Email(email)).
		Only(context.Background())
}
