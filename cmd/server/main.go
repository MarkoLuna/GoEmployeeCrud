package main

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/app"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/config"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/controllers"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services/impl"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/gorilla/mux"
)

var (
	App = app.Application{}
)

func main() {
	ConfigureApp()
	defer App.DbConnection.Close()
	App.Run()
}

func ConfigureApp() {
	App.Router = mux.NewRouter()
	if App.DbConnection == nil {
		App.DbConnection = config.GetDB()
	}

	if App.EmployeeRepository == nil {
		App.EmployeeRepository = repositories.NewEmployeeRepository(App.DbConnection, true)
	}

	App.EmployeeService = services.NewEmployeeService(App.EmployeeRepository)
	App.EmployeeController = controllers.NewEmployeeController(App.EmployeeService)

	App.ClientService = services.NewClientService()
	App.UserService = services.NewUserService()

	if App.OAuthService == nil {
		App.OAuthService = impl.NewOAuthService()
	}

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	App.OAuthController = controllers.NewOAuthController(oauthServer, App.OAuthService, App.ClientService, App.UserService)

	App.RegisterRoutes()
}
