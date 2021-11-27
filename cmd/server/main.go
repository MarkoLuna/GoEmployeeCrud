package main

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/app"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/config"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
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

	App.ClientService = services.NewClientService()
	App.UserService = services.NewUserService()
	App.OAuthService = services.NewOAuthService()

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	/*
		manager := manage.NewDefaultManager()
		// token memory store
		manager.MustTokenStorage(store.NewMemoryTokenStore())

		// client memory store
		clientStore := store.NewClientStore()
		/*
			clientStore.Set("employee_client", &models.Client{
				ID:     "employee_client",
				Secret: "employee_secret",
				Domain: "http://localhost:8080",
			})
	*/
	/*
		clientStore.Set("000000", &models.Client{
			ID:     "000000",
			Secret: "999999",
			Domain: "http://localhost:8080",
		})
		manager.MapClientStorage(clientStore)
	*/

	/*
		oauthServer.SetAllowGetAccessRequest(true)
		oauthServer.SetClientInfoHandler(server.ClientFormHandler)

		oauthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			log.Println("Internal Error:", err.Error())
			return
		})

		oauthServer.SetResponseErrorHandler(func(re *errors.Response) {
			log.Println("Response Error:", re.Error.Error())
		})
	*/

	App.OAuthController = controllers.NewOAuthController(oauthServer, App.OAuthService, App.ClientService, App.UserService)

	App.RegisterRoutes()

	defer App.DbConnection.Close()

	App.Run()
}
