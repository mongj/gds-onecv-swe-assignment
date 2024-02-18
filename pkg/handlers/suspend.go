package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/api"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
)

type suspendStudentRequest struct {
	Student string `json:"student"`
}

// ListCommonStudents lists students who registered under the given teacher(s)
func SuspendStudent(w http.ResponseWriter, r *http.Request) {
	client := database.Client
	var err error

	// Read JSON body in request into a new registerRequest object
	decoder := json.NewDecoder(r.Body)
	var data suspendStudentRequest
	err = decoder.Decode(&data)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.BuildError(api.WrapErrorJSONRequestBody(err)))
		return
	}

	// Check for missing input fields in request body
	if len(data.Student) == 0 {
		err = api.ErrMissingInput(
			"1 student email",
			"none",
		)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.BuildError(err))
		return
	}

	// Suspend student
	s, err := client.Student.
		Query().
		Where(student.Email(data.Student)).
		Only(context.Background())
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, api.BuildError(fmt.Errorf("%s: %s", api.StudentNotFoundErrorStr, data.Student)))
		default:
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to retrieve specified student")))
		}
		return
	}

	err = client.Student.
		UpdateOne(s).
		SetIsSuspended(true).
		Exec(context.Background())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to update suspension status for specified student")))
		return
	}

	// Return 204 No Content on success as specified in user story
	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
}
