package routes

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterEmployeeStoreRoutes(router *mux.Router, employeeController *controllers.EmployeeController) {

	router.HandleFunc("/api/employee/", employeeController.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/", employeeController.GetEmployees).Methods("GET")
	router.HandleFunc("/api/employee/{employeeId}", employeeController.GetEmployeeById).Methods("GET")
	router.HandleFunc("/api/employee/{employeeId}", employeeController.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/{employeeId}", employeeController.DeleteEmployee).Methods("DELETE")
}

func RegisterHealthcheckRoute(router *mux.Router) {
	router.HandleFunc("/healthcheck/", controllers.HealthCheckHandler).Methods("GET")
}

func RegisterOAuthRoutes(router *mux.Router, oauthController *controllers.OAuthController) {
	router.HandleFunc("/oauth/token", oauthController.TokenHandler).Methods("POST")
	router.HandleFunc("/oauth/userinfo", oauthController.GetUserInfo).Methods("POST")
}
