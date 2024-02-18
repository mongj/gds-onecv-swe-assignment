package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/api"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
)

type registerRequest struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}

// RegisterStudent registers a student
func RegisterStudent(w http.ResponseWriter, r *http.Request) {
	client := database.Client
	var err error

	// Read JSON body in request into a new registerRequest object
	decoder := json.NewDecoder(r.Body)
	var data registerRequest
	err = decoder.Decode(&data)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.BuildError(api.WrapErrorJSONRequestBody(err)))
		return
	}

	// Check for missing input fields in request body
	var numTeacher int
	if len(data.Teacher) > 0 {
		numTeacher = 1
	} else {
		numTeacher = 0
	}
	numStudents := len(data.Students)
	if numTeacher == 0 || numStudents == 0 {
		err = api.ErrMissingInput(
			"1 teacher and at least 1 student",
			fmt.Sprintf("%d teacher and %d student(s)", numTeacher, numStudents),
		)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.BuildError(err))
		return
	}

	// Create teacher if not exists
	teacherId, err := client.Teacher.
		Create().
		SetEmail(data.Teacher).
		OnConflictColumns(teacher.FieldEmail).
		UpdateNewValues().
		ID(context.Background())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to create teacher")))
		return
	}

	// Create students if not exists
	err = client.Student.
		MapCreateBulk(
			data.Students,
			func(s *ent.StudentCreate, i int) {
				s.SetEmail(data.Students[i]).AddTeacherIDs(teacherId)
			}).
		OnConflictColumns(student.FieldEmail).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to create student(s)")))
		return
	}

	// Return 204 No Content on success as specified in user story
	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
}
