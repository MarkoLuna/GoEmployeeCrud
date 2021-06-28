package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/constants"
)

type Employee struct {
	Id               int64
	FirstName        string                   `json:"firstName" validate:"required"`
	LastName         string                   `json:"lastName" validate:"required"`
	SecondLastName   string                   `json:"secondLastName" validate:"required"`
	DateOfBirth      time.Time                `json:"dateOfBirth" validate:"required"`
	DateOfEmployment time.Time                `json:"dateOfEmployment" validate:"required"`
	Status           constants.EmployeeStatus `json:"status" validate:"EmployeeStatusValid"`
}

func (e Employee) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
