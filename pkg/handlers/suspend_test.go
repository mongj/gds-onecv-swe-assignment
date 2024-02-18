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

func TestSuspend(t *testing.T) {
	database.InitTest(t)
	defer database.Client.Close()

	var (
		// both students to be suspended but only student1 is registered
		student1 = "studentjon@gmail.com"
		student2 = "studenthon@gmail.com"
	)

	// table tests
	tests := testutil.TestTable{
		"success": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "student": "%s"
			}`,
				student1)),
			ExpectedStatus: http.StatusNoContent,
			ExpectedBody:   "null\n",
			ExpectIsError:  false,
		},
		"unregistered student": {
			JSONBody: []byte(fmt.Sprintf(`
			{
			  "student": "%s"
			}`,
				student2)),
			ExpectedStatus:     http.StatusNotFound,
			ExpectIsError:      true,
			ExpectErrorMessage: api.StudentNotFoundErrorStr,
		},
		"malformed JSON request body": {
			JSONBody:           []byte(`{:}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.JSONParseErrorStr,
		},
		"missing student in request #1": {
			JSONBody: []byte(`
			{
			  "student": ""
			}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
		"missing student in request #2": {
			JSONBody:           []byte(`{}`),
			ExpectedStatus:     http.StatusBadRequest,
			ExpectIsError:      true,
			ExpectErrorMessage: api.MissingInputErrorStr,
		},
	}

	// register student1
	_, err := database.CreateStudent(database.Client, &ent.Student{Email: student1})
	if err != nil {
		t.Fatal(err)
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bodyReader := bytes.NewReader(test.JSONBody)
			req, err := http.NewRequest(http.MethodPost, "/suspend", bodyReader)
			if err != nil {
				t.Fatal(err)
			}
			rr := testutil.ExecuteRequest(req, SuspendStudent)
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

			// for success case, check that the student is suspended
			if name == "success" {
				s, err := database.GetStudentByEmail(database.Client, student1)
				if err != nil {
					t.Fatal(err)
				}

				require.Equal(t, true, s.IsSuspended)
			}
		})
	}
}
