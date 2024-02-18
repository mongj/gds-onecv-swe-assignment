package seed

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/api"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
)

type seedResponse struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	client := database.Client
	var err error

	// clear all tables
	_, err = client.Teacher.Delete().Exec(context.Background())
	if err != nil {
		render.JSON(w, r, api.BuildError(err))
		return
	}
	_, err = client.Student.Delete().Exec(context.Background())
	if err != nil {
		render.JSON(w, r, api.BuildError(err))
		return
	}

	// create teachers
	t, err := createTeachers(client, teachers)
	if err != nil {
		render.JSON(w, r, api.BuildError(err))
		return
	}

	// create students
	s, err := createStudents(client, students)
	if err != nil {
		render.JSON(w, r, api.BuildError(err))
		return
	}

	// add relationships
	for _, student := range []*ent.Student{s[0], s[1], s[4], s[5]} {
		err = addRelationship(client, student, t[0])
		if err != nil {
			render.JSON(w, r, api.BuildError(err))
			return
		}
	}
	for _, student := range []*ent.Student{s[2], s[3], s[4], s[5]} {
		err = addRelationship(client, student, t[1])
		if err != nil {
			render.JSON(w, r, api.BuildError(err))
			return
		}
	}

	render.JSON(w, r, seedResponse{Message: "Completed seeding"})
}

func createTeachers(client *ent.Client, teachers []*ent.Teacher) ([]*ent.Teacher, error) {
	t, err := client.Teacher.
		MapCreateBulk(teachers, func(t *ent.TeacherCreate, i int) {
			t.SetEmail(teachers[i].Email)
		}).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to create teachers: %w", err)
	}
	return t, nil
}

func createStudents(client *ent.Client, students []*ent.Student) ([]*ent.Student, error) {
	s, err := client.Student.
		MapCreateBulk(students, func(s *ent.StudentCreate, i int) {
			s.SetEmail(students[i].Email)
		}).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to create teacher: %w", err)
	}
	return s, nil
}

func addRelationship(client *ent.Client, student *ent.Student, teacher *ent.Teacher) error {
	err := client.Student.
		UpdateOne(student).
		AddTeachers(teacher).
		Exec(context.Background())

	if err != nil {
		return fmt.Errorf("failed to add relationship: %w", err)
	}
	return err
}
