package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/testutil"
	"github.com/stretchr/testify/require"
)

func TestListCommonStudents(t *testing.T) {
	database.InitTest(t)
	defer database.Client.Close()

	var teachers = []*ent.Teacher{
		{Email: "teacher1@gmail.com"},
		{Email: "teacher2@gmail.com"},
		{Email: "teachernostudent@gmail.com"},
	}

	var students = []*ent.Student{
		{Email: "student1@gmail.com"},
		{Email: "student2@gmail.com"},
		{Email: "studentcommon@gmail.com"},
	}

	// table tests
	tests := testutil.TestTable{
		"no teacher provided": {
			URL:              "/commonstudents?teacher=",
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: CommonStudentsResponse{Students: []string{}},
		},
		"teacher 1 only": {
			URL: fmt.Sprintf("/commonstudents?teacher=%s",
				url.QueryEscape((teachers[0].Email)),
			),
			ExpectedStatus: http.StatusOK,
			ExpectedResponse: CommonStudentsResponse{
				Students: []string{students[0].Email, students[2].Email},
			},
		},
		"teacher 2 only": {
			URL: fmt.Sprintf(
				"/commonstudents?teacher=%s",
				url.QueryEscape((teachers[1].Email)),
			),
			ExpectedStatus: http.StatusOK,
			ExpectedResponse: CommonStudentsResponse{
				Students: []string{students[1].Email, students[2].Email},
			},
		},
		"teacher 1 and 2 (1 common student)": {
			URL: fmt.Sprintf(
				"/commonstudents?teacher=%s&teacher=%s",
				url.QueryEscape((teachers[0].Email)),
				url.QueryEscape((teachers[1].Email)),
			),
			ExpectedStatus: http.StatusOK,
			ExpectedResponse: CommonStudentsResponse{
				Students: []string{students[2].Email},
			},
		},
		"teacher 1, 2, 3 (no common student)": {
			URL: fmt.Sprintf("/commonstudents?teacher=%s&teacher=%s&teacher=%s",
				url.QueryEscape((teachers[0].Email)),
				url.QueryEscape((teachers[1].Email)),
				url.QueryEscape((teachers[2].Email)),
			),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: CommonStudentsResponse{Students: []string{}},
		},
	}

	// create teachers
	entTeachers, err := database.BulkCreateTeachers(database.Client, teachers)
	if err != nil {
		t.Fatal(err)
	}

	// create students
	entStudents, err := database.BulkCreateStudents(database.Client, students)
	if err != nil {
		t.Fatal(err)
	}

	// add relationships
	for _, student := range []*ent.Student{entStudents[0], entStudents[2]} {
		err = database.AddRelationshipToTeacher(database.Client, student, entTeachers[0])
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, student := range []*ent.Student{entStudents[1], entStudents[2]} {
		err = database.AddRelationshipToTeacher(database.Client, student, entTeachers[1])
		if err != nil {
			t.Fatal(err)
		}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bodyReader := bytes.NewReader(test.JSONBody)
			req, err := http.NewRequest(http.MethodGet, test.URL, bodyReader)
			if err != nil {
				t.Fatal(err)
			}
			rr := testutil.ExecuteRequest(req, ListCommonStudents)
			testutil.CheckResponseCode(t, test.ExpectedStatus, rr.Code)
			if test.ExpectIsError {
				testutil.CheckErrorMessage(t, test.ExpectErrorMessage, rr.Body.String())
			} else {
				var res CommonStudentsResponse
				err := json.Unmarshal(rr.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, test.ExpectedResponse, res)
			}
		})
	}
}
