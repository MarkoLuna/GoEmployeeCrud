package main

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/app"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/config"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/gorilla/mux"
)

var (
	App = app.Application{}
)

func main() {
	App.Router = mux.NewRouter()
	if App.DbConnection == nil {
		App.DbConnection = config.GetDB()
	}

	App.EmployeeRepository = repositories.NewEmployeeRepository(App.DbConnection)
	App.EmployeeService = services.NewEmployeeService(App.EmployeeRepository)
	App.EmployeeController = controllers.NewEmployeeController(App.EmployeeService)

	App.RegisterRoutes()

	defer App.DbConnection.Close()

	App.Run()
}
