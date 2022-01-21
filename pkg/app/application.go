package app

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/app/config"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/routes"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	"github.com/gorilla/mux"
)

type Application struct {
	Router             *mux.Router
	DbConnection       *sql.DB
	EmployeeService    services.EmployeeService
	ClientService      services.ClientService
	UserService        services.UserService
	OAuthService       services.OAuthService
	EmployeeRepository repositories.EmployeeRepository
	EmployeeController controllers.EmployeeController
	OAuthController    controllers.OAuthController
}

func (app *Application) RegisterRoutes() {
	routes.RegisterHealthcheckRoute(app.Router)
	routes.RegisterEmployeeStoreRoutes(app.Router, &app.EmployeeController)
	routes.RegisterOAuthRoutes(app.Router, &app.OAuthController)
	config.EnableCORS(app.Router)
	config.NewAuthConfig(app.Router, true, nil, app.OAuthService)
}

func (app *Application) Address() string {
	port := utils.GetEnv("SERVER_PORT", "8080")
	host := utils.GetEnv("SERVER_HOST", "0.0.0.0")

	return host + ":" + port
}

func (app *Application) HandleRoutes() {
	http.Handle("/", app.Router)
}

func (app *Application) StartServer() {
	app.HandleRoutes()
	address := app.Address()
	log.Println("Starting server on:", address)

	log.Fatal(http.ListenAndServe(address, app.Router))
}

func (app *Application) StartSecureServer() {
	app.HandleRoutes()
	address := app.Address()
	log.Println("Starting server on:", address)

	path := "/Users/marcos.luna/go-projects/GoEmployeeCrud"
	certFile := utils.GetEnv("SERVER_SSL_CERT_FILE_PATH", path+"/resources/ssl/cert.pem")
	keyFile := utils.GetEnv("SERVER_SSL_KEY_FILE_PATH", path+"/resources/ssl/key.pem")
	log.Fatal(http.ListenAndServeTLS(address, certFile, keyFile, app.Router))
}

func (app *Application) Run() {
	server_ssl_enabled := utils.GetEnv("SERVER_SSL_ENABLED", "false")
	ssl_enabled, _ := strconv.ParseBool(server_ssl_enabled)
	if ssl_enabled {
		app.StartSecureServer()
	} else {
		app.StartServer()
	}
}
