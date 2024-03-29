// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/mongj/gds-onecv-swe-assignment/ent/schema"
	"github.com/mongj/gds-onecv-swe-assignment/ent/student"
	"github.com/mongj/gds-onecv-swe-assignment/ent/teacher"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescEmail is the schema descriptor for email field.
	studentDescEmail := studentFields[0].Descriptor()
	// student.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	student.EmailValidator = studentDescEmail.Validators[0].(func(string) error)
	// studentDescIsSuspended is the schema descriptor for is_suspended field.
	studentDescIsSuspended := studentFields[1].Descriptor()
	// student.DefaultIsSuspended holds the default value on creation for the is_suspended field.
	student.DefaultIsSuspended = studentDescIsSuspended.Default.(bool)
	teacherFields := schema.Teacher{}.Fields()
	_ = teacherFields
	// teacherDescEmail is the schema descriptor for email field.
	teacherDescEmail := teacherFields[0].Descriptor()
	// teacher.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	teacher.EmailValidator = teacherDescEmail.Validators[0].(func(string) error)
}
