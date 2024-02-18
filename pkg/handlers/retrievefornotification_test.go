package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/api"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/testutil"
	"github.com/stretchr/testify/require"
)

func TestRetrieveForNotification(t *testing.T) {
	database.InitTest(t)
	defer database.Client.Close()

	var teacher1 = "teacher1@gmail.com"

	var registeredStudents = []*ent.Student{
		{Email: "student1@gmail.com"},
		{Email: "student2@gmail.com"},
	}

	var unregisterdStudents = []*ent.Student{
		{Email: "student3@gmail.com"},
	}

	// table tests
	tests := testutil.TestTable{
		"no student mentioned": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "notification": "hello world"
			}`,
				teacher1)),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: retrieveForNotificationResponse{Recipients: []string{registeredStudents[0].Email, registeredStudents[1].Email}},
			ExpectIsError:    false,
		},
		"mention one registered student": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "notification": "hello world @%s"
			}`,
				teacher1, registeredStudents[0].Email)),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: retrieveForNotificationResponse{Recipients: []string{registeredStudents[0].Email, registeredStudents[1].Email}},
			ExpectIsError:    false,
		},
		"mention all students": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "notification": "hello world @%s @%s @%s"
			}`,
				teacher1, registeredStudents[0].Email, registeredStudents[1].Email, unregisterdStudents[0].Email)),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: retrieveForNotificationResponse{Recipients: []string{registeredStudents[0].Email, registeredStudents[1].Email, unregisterdStudents[0].Email}},
			ExpectIsError:    false,
		},
		"mention only unregistered student": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "notification": "hello world @%s"
			}`,
				teacher1, unregisterdStudents[0].Email)),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: retrieveForNotificationResponse{Recipients: []string{registeredStudents[0].Email, registeredStudents[1].Email, unregisterdStudents[0].Email}},
			ExpectIsError:    false,
		},
		"include email without @mention": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "notification": "hello world %s"
			}`,
				teacher1, unregisterdStudents[0].Email)),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: retrieveForNotificationResponse{Recipients: []string{registeredStudents[0].Email, registeredStudents[1].Email}},
			ExpectIsError:    false,
		},
		"multiple @mentions": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "notification": "hello world @%s @%s @%s"
			}`,
				teacher1, unregisterdStudents[0].Email, unregisterdStudents[0].Email, unregisterdStudents[0].Email)),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: retrieveForNotificationResponse{Recipients: []string{registeredStudents[0].Email, registeredStudents[1].Email, unregisterdStudents[0].Email}},
			ExpectIsError:    false,
		},
		"malformed JSON request body": {
			JSONBody:           []byte(`{:}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.JSONParseErrorStr,
		},
		"missing teacher in request": {
			JSONBody: []byte(`
			{
			  "notification": "hello world"
			}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing notification in request": {
			JSONBody: []byte(`
			{
			  "teacher": "anyteacher@gmail.com"
			}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing both teacher and notification #1": {
			JSONBody: []byte(`
			{
			  "teacher": "",
			  "notification": ""
			}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing both teacher and notification #2": {
			JSONBody:           []byte(`{ }`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
	}

	// create teacher
	entTeacher, err := database.CreateTeacher(database.Client, teacher1)
	if err != nil {
		t.Fatal(err)
	}

	// create students
	entStudents, err := database.BulkCreateStudents(database.Client, registeredStudents)
	if err != nil {
		t.Fatal(err)
	}
	_, err = database.BulkCreateStudents(database.Client, unregisterdStudents)
	if err != nil {
		t.Fatal(err)
	}

	// add relationships
	for _, student := range entStudents {
		err = database.AddRelationshipToTeacher(database.Client, student, entTeacher)
		if err != nil {
			t.Fatal(err)
		}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bodyReader := bytes.NewReader(test.JSONBody)
			req, err := http.NewRequest(http.MethodPost, "/retrievefornotification", bodyReader)
			if err != nil {
				t.Fatal(err)
			}
			rr := testutil.ExecuteRequest(req, RetrieveForNotifications)
			testutil.CheckResponseCode(t, test.ExpectedStatus, rr.Code)
			if test.ExpectIsError {
				testutil.CheckErrorMessage(t, test.ExpectErrorMessage, rr.Body.String())
			} else {
				var res retrieveForNotificationResponse
				err := json.Unmarshal(rr.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, test.ExpectedResponse, res)
			}
		})
	}
}
