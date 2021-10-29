package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/constants"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/dto"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEmployeeStoreRoutes(t *testing.T) {
	router := mux.NewRouter()

	employeeRepository := repositories.NewEmployeeRepositoryStub()
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	RegisterEmployeeStoreRoutes(router, &employeeController)

	var employee dto.EmployeeRequest
	employee.FirstName = "Marcos"
	employee.LastName = "Luna"
	employee.SecondLastName = "Valdez"
	employee.DateOfBirth = time.Date(1994, time.April, 25, 8, 0, 0, 0, time.UTC)
	employee.DateOfEmployment = time.Now().UTC()
	employee.Status = constants.ACTIVE

	jsonStr, _ := json.Marshal(employee)

	tables := []struct {
		method string
		path   string
		body   io.Reader
		status int
	}{
		{"GET", "/api/employee/", nil, http.StatusOK},
		{"POST", "/api/employee/", bytes.NewBuffer(jsonStr), http.StatusCreated},
		{"GET", "/api/employee/1", nil, http.StatusOK},
		{"PUT", "/api/employee/1", bytes.NewBuffer(jsonStr), http.StatusOK},
		{"DELETE", "/api/employee/1", nil, http.StatusOK},
		{"GET", "/api/employe/", nil, http.StatusNotFound},
		{"POST", "/api/employe/", bytes.NewBuffer(jsonStr), http.StatusNotFound},
		{"GET", "/api/employe/1", nil, http.StatusNotFound},
		{"PUT", "/api/employe/1", bytes.NewBuffer(jsonStr), http.StatusNotFound},
		{"DELETE", "/api/employe/1", nil, http.StatusNotFound},
	}

	for _, table := range tables {
		req, err := http.NewRequest(table.method, table.path, table.body)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, table.status, rr.Code, "handler returned wrong status code")
	}
}

func TestRegisterHealthcheckRoute(t *testing.T) {
	router := mux.NewRouter()

	RegisterHealthcheckRoute(router)

	tables := []struct {
		method   string
		path     string
		response string
		status   int
	}{
		{"GET", "/healthcheck/", `OK`, http.StatusOK},
		{"GET", "/healthchec/", "", http.StatusNotFound},
	}

	for _, table := range tables {
		req, err := http.NewRequest(table.method, table.path, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, table.status, rr.Code, "handler returned wrong status code")
		if table.response != "" {
			assert.Equal(t, table.response, rr.Body.String(), "handler returned unexpected body")
		}
	}
}
