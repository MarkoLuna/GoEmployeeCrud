package routes

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterEmployeeStoreRoutes(echoInstance *echo.Echo, employeeController *controllers.EmployeeController) {

	echoInstance.POST("/api/employee/", employeeController.CreateEmployee)
	echoInstance.GET("/api/employee/", employeeController.GetEmployees)
	echoInstance.GET("/api/employee/:employeeId", employeeController.GetEmployeeById)
	echoInstance.PUT("/api/employee/:employeeId", employeeController.UpdateEmployee)
	echoInstance.DELETE("/api/employee/:employeeId", employeeController.DeleteEmployee)
}

func RegisterHealthcheckRoute(echoInstance *echo.Echo) {
	echoInstance.GET("/healthcheck/", controllers.HealthCheckHandler)
}

func RegisterOAuthRoutes(echoInstance *echo.Echo, oauthController *controllers.OAuthController) {
	echoInstance.POST("/oauth/token", oauthController.TokenHandler)
	echoInstance.GET("/oauth/userinfo", oauthController.GetUserInfo)
}
