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

func TestRegister(t *testing.T) {
	database.InitTest(t)
	defer database.Client.Close()

	var (
		teacher1 = "teacherken@gmail.com"
		student1 = "studentjon@gmail.com"
		student2 = "studenthon@gmail.com"
	)

	// table tests
	tests := testutil.TestTable{
		"success": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "students": ["%s","%s"]
			}`,
				teacher1, student1, student2)),
			ExpectedStatus:   http.StatusNoContent,
			ExpectedResponse: interface{}(nil),
			ExpectIsError:    false,
		},
		"malformed JSON request body": {
			JSONBody:           []byte(`{:}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.JSONParseErrorStr,
		},
		"missing teacher in request": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "",
			  "students": ["%s","%s"]
			}`,
				student1, student2)),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing student in request": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "teacher": "%s",
			  "students": []
			}`,
				teacher1)),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing both teacher and student #1": {
			JSONBody: []byte(`
			{
			  "teacher": "",
			  "students": []
			}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing both teacher and student #2": {
			JSONBody:           []byte(`{ }`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bodyReader := bytes.NewReader(test.JSONBody)
			req, err := http.NewRequest(http.MethodPost, "/register", bodyReader)
			if err != nil {
				t.Fatal(err)
			}
			rr := testutil.ExecuteRequest(req, RegisterStudent)
			testutil.CheckResponseCode(t, test.ExpectedStatus, rr.Code)
			if test.ExpectIsError {
				testutil.CheckErrorMessage(t, test.ExpectErrorMessage, rr.Body.String())
			} else {
				var res interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, test.ExpectedResponse, res)
			}

			// for success case, check if the teacher and students are created
			if name == "success" {
				_, err := database.GetTeacherByEmail(database.Client, teacher1)
				if err != nil {
					t.Fatal(err)
				}

				students, err := database.GetStudentsByTeacher(database.Client, &ent.Teacher{Email: teacher1})
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, 2, len(students))
				for _, s := range students {
					if s.Email != student1 && s.Email != student2 {
						t.Fatalf("expected student %s or %s, got %s", student1, student2, s.Email)
					}
				}
			}
		})
	}
}
