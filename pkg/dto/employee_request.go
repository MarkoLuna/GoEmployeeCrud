package dto

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/constants"
)

type EmployeeRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	SecondLastName string `json:"secondLastName" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	DateOfEmployment time.Time `json:"dateOfEmployment" validate:"required"`
	Status constants.EmployeeStatus `json:"status" validate:"EmployeeStatusValid"`
}

func (e EmployeeRequest) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
