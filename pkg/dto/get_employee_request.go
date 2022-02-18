package dto

import (
	"encoding/json"
	"log"
)

type GetEmployeeRequest struct {
	EmployeeId string
}

func (e GetEmployeeRequest) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
