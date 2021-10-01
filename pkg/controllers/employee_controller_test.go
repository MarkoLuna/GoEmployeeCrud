package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/constants"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/models"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeController_GetEmployeesEmployees(t *testing.T) {

	employeeRepository := repositories.NewEmployeeRepositoryStub()
	employeeController := NewEmployeeController(employeeRepository)

	req, err := http.NewRequest("GET", "/api/employee/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employeeController.GetEmployees)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.NotEqual(t, 0, len(rr.Body.String()), "handler returned unexpected body: got empty")

	employeesSlice := make([]models.Employee, 0)
	err = json.Unmarshal(rr.Body.Bytes(), &employeesSlice)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(employeesSlice), "handler returned unexpected body: got empty")

	employee1 := employeesSlice[0]
	employee2 := employeesSlice[1]

	assert.NotNil(t, employee1)
	assert.NotNil(t, employee2)

	fmt.Println(employee1)
	fmt.Println(employee2)

	assert.Equal(t, "1", employee1.Id, "id employee returned is wrong")
	assert.Equal(t, "2", employee2.Id, "id employee returned is wrong")
}

func TestEmployeeController_CreateEmployeeEmployee(t *testing.T) {

	employeeRepository := repositories.NewEmployeeRepositoryStub()
	employeeController := NewEmployeeController(employeeRepository)

	var employee models.Employee
	employee.FirstName = "Marcos"
	employee.LastName = "Luna"
	employee.SecondLastName = "Valdez"
	employee.DateOfBirth = time.Date(1994, time.April, 25, 8, 0, 0, 0, time.UTC)
	employee.DateOfEmployment = time.Now().UTC()
	employee.Status = constants.ACTIVE

	jsonStr, _ := json.Marshal(employee)
	req, err := http.NewRequest("POST", "/api/employee/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employeeController.CreateEmployee)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "handler returned wrong status code")
	assert.NotEqual(t, 0, len(rr.Body.String()), "handler returned unexpected body: got empty")

	employeeResponse := models.Employee{}
	err = json.Unmarshal(rr.Body.Bytes(), &employeeResponse)

	assert.NoError(t, err)
	assert.NotNil(t, employeeResponse)
	fmt.Println(employeeResponse)

	assert.Equal(t, employee.FirstName, employeeResponse.FirstName, "FirstName employee returned is wrong")
}

func TestEmployeeController_GetEmployeeByIdEmployee(t *testing.T) {

	employeeRepository := repositories.NewEmployeeRepositoryStub()
	employeeController := NewEmployeeController(employeeRepository)

	req, err := http.NewRequest("GET", "/api/employee/1", nil)
	assert.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{
		"employeeId": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employeeController.GetEmployeeById)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.NotEqual(t, 0, len(rr.Body.String()), "handler returned unexpected body: got empty")

	employeeResponse := models.Employee{}
	err = json.Unmarshal(rr.Body.Bytes(), &employeeResponse)

	assert.NoError(t, err)

	assert.NotNil(t, employeeResponse)
	fmt.Println(employeeResponse)

	assert.Equal(t, "1", employeeResponse.Id, "id employee returned is wrong")
}

func TestEmployeeController_UpdateEmployee(t *testing.T) {

	employeeRepository := repositories.NewEmployeeRepositoryStub()
	employeeController := NewEmployeeController(employeeRepository)

	var employee models.Employee
	employee.Id = "1"
	employee.FirstName = "Marcos"
	employee.LastName = "Luna"
	employee.SecondLastName = "Valdez"
	employee.DateOfBirth = time.Date(1994, time.April, 25, 8, 0, 0, 0, time.UTC)
	employee.DateOfEmployment = time.Now().UTC()
	employee.Status = constants.ACTIVE

	jsonStr, _ := json.Marshal(employee)
	req, err := http.NewRequest("PUT", "/api/employee/1", bytes.NewBuffer(jsonStr))
	assert.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{
		"employeeId": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employeeController.UpdateEmployee)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.NotEqual(t, 0, len(rr.Body.String()), "handler returned unexpected body: got empty")

	employeeResponse := models.Employee{}
	err = json.Unmarshal(rr.Body.Bytes(), &employeeResponse)

	assert.NoError(t, err)

	assert.NotNil(t, employeeResponse)
	fmt.Println(employeeResponse)

	assert.Equal(t, "1", employeeResponse.Id, "id employee returned is wrong")
}

func TestEmployeeController_DeleteEmployee(t *testing.T) {

	employeeRepository := repositories.NewEmployeeRepositoryStub()
	employeeController := NewEmployeeController(employeeRepository)

	req, err := http.NewRequest("DELETE", "/api/employee/1", nil)
	assert.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{
		"employeeId": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employeeController.DeleteEmployee)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
}
