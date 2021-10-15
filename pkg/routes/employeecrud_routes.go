package routes

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterEmployeeStoreRoutes(router *mux.Router, employeeController *controllers.EmployeeController) {

	// swagger:route POST /api/employee/ createEmployee
	//
	// Create a new Employee.
	//
	// responses:
	//   201: employeeData Employee created successfully
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	router.HandleFunc("/api/employee/", employeeController.CreateEmployee).Methods("POST")

	// swagger:route GET /api/employee/ getEmployees
	//
	// Get Employee list.
	//
	// responses:
	//   200: employeeData Employee information
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	router.HandleFunc("/api/employee/", employeeController.GetEmployees).Methods("GET")

	// swagger:route GET /api/employee/{EmployeeId} getEmployeeById
	//
	// Get Employe by Id.
	//
	// responses:
	//   200: employeeData Employee information
	// responses:
	//   404: empty Employee not found
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	router.HandleFunc("/api/employee/{employeeId}", employeeController.GetEmployeeById).Methods("GET")

	// swagger:route PUT /api/employee/{EmployeeId} updateEmployeeById
	//
	// Update Employe by Id.
	//
	// responses:
	//   200: employeeData Employee updated successfully
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	router.HandleFunc("/api/employee/{employeeId}", employeeController.UpdateEmployee).Methods("PUT")

	// swagger:route DELETE /api/employee/{EmployeeId} deleteEmployeeById
	//
	// Delete Employee by Id.
	//
	// responses:
	//   200: empty Employee deleted successfully
	// responses:
	//   404: empty Employee not found
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	router.HandleFunc("/api/employee/{employeeId}", employeeController.DeleteEmployee).Methods("DELETE")
}

func RegisterHealthcheckRoute(router *mux.Router) {

	// swagger:route GET /healthcheck/ healthcheck
	//
	// Check health of the service.
	//
	// Produces:
	//     - text/plain
	// responses:
	//   200: empty Service is up
	router.HandleFunc("/healthcheck/", controllers.HealthCheckHandler).Methods("GET")
}
