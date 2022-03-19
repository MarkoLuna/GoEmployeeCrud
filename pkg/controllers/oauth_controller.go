package controllers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"

	"github.com/labstack/echo/v4"
)

type OAuthController struct {
	oauthServer   *server.Server
	oauthSevice   services.OAuthService
	clientService services.ClientService
	userService   services.UserService
}

func NewOAuthController(oauthServer *server.Server,
	oauthService services.OAuthService,
	clientService services.ClientService,
	userService services.UserService) OAuthController {
	ctrl := OAuthController{oauthServer, oauthService, clientService, userService}
	ctrl.Configure()
	return ctrl
}

func (ctrl OAuthController) Configure() {

	ctrl.oauthServer.SetAllowGetAccessRequest(true)
	ctrl.oauthServer.SetClientInfoHandler(server.ClientFormHandler)

	ctrl.oauthServer.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	ctrl.oauthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		re.SetHeader("error", err.Error())
		return
	})
}

func (ctrl OAuthController) TokenHandler(c echo.Context) error {
	auth, ok := ctrl.GetBasicAuth(c.Request().Header)
	if !ok {
		return c.String(http.StatusUnauthorized, "Unable to find the Authentication")
	}

	clientIdReq, clientSecretReq := ctrl.DecodeBasicAuth(auth)
	validClientCred, err := ctrl.clientService.IsValidClientCredentials(clientIdReq, clientSecretReq)
	if !validClientCred || err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	userNameReq := c.FormValue("username")
	passwordReq := c.FormValue("password")

	userId, err := ctrl.userService.GetUserId(userNameReq, passwordReq)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	jWTResponse, err := ctrl.oauthSevice.HandleTokenGeneration(clientIdReq, clientSecretReq, userId)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, jWTResponse)
}

func (ctrl OAuthController) GetUserInfo(c echo.Context) error {
	accessToken, ok := ctrl.GetBearerAuth(c.Request().Header)
	if !ok {
		return c.String(http.StatusUnauthorized, "Unable to find the Authentication")
	}

	log.Println("auth token: ", accessToken)
	claims, err := ctrl.oauthSevice.GetTokenClaims(accessToken)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, claims)
}

func (ctrl OAuthController) GetBearerAuth(headers http.Header) (string, bool) {
	return ctrl.GetAuthHeader(headers, "Bearer ")
}

func (ctrl OAuthController) GetBasicAuth(headers http.Header) (string, bool) {
	return ctrl.GetAuthHeader(headers, "Basic ")
}

func (ctrl OAuthController) GetAuthHeader(headers http.Header, prefix string) (string, bool) {
	auth := headers.Get("Authorization")
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}

	return token, token != ""
}

func (ctrl OAuthController) DecodeBasicAuth(auth string) (string, string) {
	authDecoded, _ := base64.StdEncoding.DecodeString(auth)
	authReq := strings.Split(string(authDecoded), ":")

	return authReq[0], authReq[1]
}
