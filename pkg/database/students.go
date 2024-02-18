package database

import (
	"context"
	"fmt"

	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
)

// CreateStudent creates a student
func CreateStudent(client *ent.Client, s *ent.Student) (*ent.Student, error) {
	return client.Student.
		Create().
		SetEmail(s.Email).
		Save(context.Background())
}

// BulkCreateStudents creates students in bulk
func BulkCreateStudents(client *ent.Client, students []*ent.Student) ([]*ent.Student, error) {
	s, err := client.Student.
		MapCreateBulk(students, func(s *ent.StudentCreate, i int) {
			s.SetEmail(students[i].Email)
		}).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to create students: %w", err)
	}
	return s, nil
}

// AddRelationshipToTeacher adds a many-many relationship between a student and a teacher
func AddRelationshipToTeacher(client *ent.Client, s *ent.Student, t *ent.Teacher) error {
	err := client.Student.
		UpdateOne(s).
		AddTeachers(t).
		Exec(context.Background())

	if err != nil {
		return fmt.Errorf("failed to add relationship: %w", err)
	}
	return err
}

// GetStudentByEmail gets a student by email
func GetStudentByEmail(client *ent.Client, email string) (*ent.Student, error) {
	return client.Student.
		Query().
		Where(student.Email(email)).
		Only(context.Background())
}

// GetStudentsByTeacher gets students by teacher
func GetStudentsByTeacher(client *ent.Client, t *ent.Teacher) ([]*ent.Student, error) {
	students, err := client.Teacher.
		Query().
		Where(teacher.Email(t.Email)).
		QueryStudents().
		All(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to get students: %w", err)
	}
	return students, nil
}
