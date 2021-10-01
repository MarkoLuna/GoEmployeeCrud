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
	//   201: employeeData
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	//
	// Information of the Employee created.
	// swagger:response employeeData
	router.HandleFunc("/api/employee/", employeeController.CreateEmployee).Methods("POST")

	// swagger:route GET /api/employee/ getEmployees
	//
	// Get Employee list.
	//
	// responses:
	//   200: employeeData
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	//
	// Information of the Employee created.
	// swagger:response employeeData
	router.HandleFunc("/api/employee/", employeeController.GetEmployees).Methods("GET")

	// swagger:route GET /api/employee/{employeeId} getEmployeeById
	//
	// Get Employe by Id.
	//
	// responses:
	//   200: employeeData
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	//
	// Information of the Employee created.
	// swagger:response employeeData
	router.HandleFunc("/api/employee/{employeeId}", employeeController.GetEmployeeById).Methods("GET")

	// swagger:route PUT /api/employee/{employeeId} updateEmployeeById
	//
	// Update Employe by Id.
	//
	// responses:
	//   200: employeeData
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	//
	// Information of the Employee created.
	// swagger:response employeeData
	router.HandleFunc("/api/employee/{employeeId}", employeeController.UpdateEmployee).Methods("PUT")

	// swagger:route DELETE /api/employee/{employeeId} deleteEmployeeById
	//
	// Delete Employe by Id.
	//
	// responses:
	//   200: employeeData
	// Consumes:
	//     - application/json
	//
	// Produces:
	//     - application/json
	//
	// Information of the Employee created.
	// swagger:response employeeData
	router.HandleFunc("/api/employee/{employeeId}", employeeController.DeleteEmployee).Methods("DELETE")
}

func RegisterHealthcheckRoute(router *mux.Router) {

	// swagger:route GET /healthcheck/ healthcheck
	//
	// Check health of the service.
	//
	// Produces:
	//     - text/plain
	router.HandleFunc("/healthcheck/", controllers.HealthCheckHandler).Methods("GET")
}
