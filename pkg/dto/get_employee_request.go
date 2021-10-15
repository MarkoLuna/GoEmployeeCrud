package dto

import (
	"encoding/json"
	"log"
)

// Employee path params
//
// This model represents the Employee path params for this application.
//
// swagger:parameters getEmployeeById updateEmployeeById deleteEmployeeById
type GetEmployeeRequest struct {

	// The id for the employee
	//
	// required: true
	// in: path
	// example: 6204037c-30e6-408b-8aaa-dd8219860b4b
	EmployeeId string
}

func (e GetEmployeeRequest) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
