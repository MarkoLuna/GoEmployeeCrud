package app

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	enableCORS(app.Router)
	enableOAuth(app.Router, app.OAuthService)
}

func (app *Application) Address() string {
	port := utils.GetEnv("SERVER_PORT", "8080")
	host := utils.GetEnv("SERVER_HOST", "0.0.0.0")

	return host + ":" + port
}

func (app *Application) HandleRoutes() {
	http.Handle("/", app.Router)
}

func enableOAuth(router *mux.Router, oauthService services.OAuthService) {
	router.PathPrefix("/api/employee/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ok, err := oauthService.ValidateToken(req)
		if !ok || err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// f.ServeHTTP(w, r)
		// w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)

	// router.Use(middlewareCors)
}

/*
func validateToken(f http.HandlerFunc, app *Application, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ok, err := app.OAuthService.ValidateToken(r)
		if !ok || err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		f.ServeHTTP(w, r)
	})
}
*/

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
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
