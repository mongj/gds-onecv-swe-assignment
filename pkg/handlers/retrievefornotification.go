package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chi/render"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/api"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
)

type retrieveForNotificationRequest struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}

type retrieveForNotificationResponse struct {
	Recipients []string `json:"recipients"`
}

// RegisterStudent registers a student
func RetrieveForNotifications(w http.ResponseWriter, r *http.Request) {
	client := database.Client
	var err error

	// Read JSON body in request into a new registerRequest object
	decoder := json.NewDecoder(r.Body)
	var data retrieveForNotificationRequest
	err = decoder.Decode(&data)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.BuildError(api.WrapErrorJSONRequestBody(err)))
		return
	}

	// Check for missing input fields in request body
	if len(data.Teacher) == 0 || len(data.Notification) == 0 {
		err = api.ErrMissingInput(
			"teacher email and notification message",
			fmt.Sprintf("teacher: %s, notification: %s", data.Teacher, data.Notification),
		)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.BuildError(err))
		return
	}

	// Extract mentioned students from notification
	re, err := regexp.Compile(`@(\w+@\w+\.\w+)`) // regex only matches emails which are @mentioned, not all emails
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to extract mentioned students from notification")))
		return
	}

	m := re.FindAllStringSubmatch(data.Notification, -1)

	// Extract first capturing group from each match, which excludes the @ symbol in front
	matches := make([]string, len(m))
	for i, match := range m {
		matches[i] = match[1]
	}

	// Retrieve students registered under the teacher
	students, err := client.Teacher.
		Query().
		Where(teacher.Email(data.Teacher)).
		QueryStudents().
		All(context.Background())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to retrieve students for teacher")))
		return
	}

	registeredStudents := make([]string, len(students))
	for i, s := range students {
		registeredStudents[i] = s.Email
	}

	// Retrieve list of suspended students
	suspendedStudents, err := client.Student.
		Query().
		Where(student.IsSuspended(true)).
		All(context.Background())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.BuildError(api.WrapError(err, "failed to retrieve suspended students")))
		return
	}

	// Combined slice of registered students and mentioned students
	recipients := append(registeredStudents, matches...)

	// Deduplicate recipients and remove suspended students from list
	recipientMap := make(map[string]bool)
	deduplicatedRecipients := []string{}

	for _, s := range suspendedStudents {
		recipientMap[s.Email] = true
	}

	for _, r := range recipients {
		if !recipientMap[r] {
			recipientMap[r] = true
			deduplicatedRecipients = append(deduplicatedRecipients, r)
		}
	}

	render.JSON(w, r, retrieveForNotificationResponse{Recipients: deduplicatedRecipients})
}
