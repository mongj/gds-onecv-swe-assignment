package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/api"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
)

type commonStudentsResponse struct {
	Students []string `json:"students"`
}

// ListCommonStudents lists students who registered under the given teacher(s)
func ListCommonStudents(w http.ResponseWriter, r *http.Request) {
	client := database.Client
	var err error

	teachers := r.URL.Query()["teacher"]

	// Find all the teacher nodes with the given email(s)
	// and run graph traversal to find students registered under them
	s, err := client.Teacher.
		Query().
		Where(teacher.EmailIn(teachers...)).
		QueryStudents().
		All(context.Background())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to list common students")))
		return
	}

	// Format response
	students := make([]string, len(s))
	for i, student := range s {
		students[i] = student.Email
	}

	// Return list of students
	render.JSON(w, r, commonStudentsResponse{Students: students})

	// If no teacher email is provided in the query params, or if no student is found,
	// an empty array is returned with status code 200.
}
