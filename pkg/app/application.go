package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/app/config"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/routes"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
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

func (app *Application) LoadConfiguration() {
	app.loadEnvValues()
	routes.RegisterHealthcheckRoute(app.Router)
	routes.RegisterEmployeeStoreRoutes(app.Router, &app.EmployeeController)
	routes.RegisterOAuthRoutes(app.Router, &app.OAuthController)
	config.EnableCORS(app.Router)
	config.NewAuthConfig(app.Router, true, config.DefaultSkippedPaths[:], app.OAuthService)
}

func (app *Application) loadEnvValues() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		} else {
			log.Println("app environment values loaded successfully")
		}
	} else {
		log.Println("Unable to find the env file for load app environment values")
	}
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
