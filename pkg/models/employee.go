package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/constants"
)

// Employee
//
// This model represents the Employee data structure for this application.
//
// swagger:model employeeData
type Employee struct {
	// The id for this employee
	//
	// required: false
	// example: 6204037c-30e6-408b-8aaa-dd8219860b4b
	Id string

	// The firstname for employee
	// required: true
	// min length: 1
	FirstName string `json:"firstName" validate:"required"`

	// The lastname for employee
	// required: true
	// min length: 1
	LastName string `json:"lastName" validate:"required"`

	// The second last name for employee
	// required: true
	// min length: 1
	SecondLastName string `json:"secondLastName" validate:"required"`

	// The date of birth for employee
	// required: true
	// example: 1994-04-25T12:00:00Z
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`

	// The date of employment for employee
	// required: true
	// example: 1994-04-25T12:00:00Z
	DateOfEmployment time.Time `json:"dateOfEmployment" validate:"required"`

	// The status for employee
	// required: true
	Status constants.EmployeeStatus `json:"status" validate:"EmployeeStatusValid"`
}

func (e Employee) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
