package dto

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/constants"
)

// Employee Request
//
// This model represents the Employee data structure for this application.
//
// swagger:parameters createEmployee updateEmployeeById
type EmployeeRequest struct {
	// The firstname for employee
	// required: true
	// example: Marcos
	// in: body
	FirstName string `json:"firstName" validate:"required"`

	// The lastname for employee
	// required: true
	// example: Luna
	// in: body
	LastName string `json:"lastName" validate:"required"`

	// The second last name for employee
	// required: true
	// example: Valdez
	// in: body
	SecondLastName string `json:"secondLastName" validate:"required"`

	// The date of birth for employee
	// required: true
	// example: 1994-04-25T12:00:00Z
	// in: body
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`

	// The date of employment for employee
	// required: true
	// example: 1994-04-25T12:00:00Z
	// in: body
	DateOfEmployment time.Time `json:"dateOfEmployment" validate:"required"`

	// The status for employee
	// required: true
	// example: Active, Inactive
	// in: body
	Status constants.EmployeeStatus `json:"status" validate:"EmployeeStatusValid"`
}

func (e EmployeeRequest) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
