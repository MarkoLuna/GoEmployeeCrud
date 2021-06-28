package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/routes"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	"github.com/gorilla/mux"
)

type Application struct {
	Router             *mux.Router
	DbConnection       *sql.DB
	EmployeeRepository repositories.EmployeeRepository
	EmployeeController controllers.EmployeeController
}

func (app *Application) RegisterRoutes() {
	routes.RegisterHealthcheckRoute(app.Router)
	routes.RegisterEmployeeStoreRoutes(app.Router, &app.EmployeeController)
}

func (app *Application) StartServer() {
	http.Handle("/", app.Router)
	port := utils.GetEnv("SERVER_PORT", "8080")
	host := utils.GetEnv("SERVER_HOST", "0.0.0.0")
	log.Println("Starting server on port:", port)

	log.Fatal(http.ListenAndServe(host+":"+port, app.Router))
}
