package controllers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
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

/*
func (ctrl OAuthController) AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	err := ctrl.oauthServer.HandleAuthorizeRequest(w, r)
	if err != nil {
		log.Printf("an error '%s' was not expected when authorize ... ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
*/

func (ctrl OAuthController) TokenHandler(w http.ResponseWriter, r *http.Request) {
	auth, ok := ctrl.GetBasicAuth(r)
	if !ok {
		http.Error(w, "Unable to find the Authentication", http.StatusUnauthorized)
		return
	}

	clientIdReq, clientSecretReq := ctrl.DecodeBasicAuth(auth)
	validClientCred, err := ctrl.clientService.IsValidClientCredentials(clientIdReq, clientSecretReq)
	if !validClientCred || err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	userNameReq := r.FormValue("username")
	passwordReq := r.FormValue("password")

	userId, err := ctrl.userService.GetUserId(userNameReq, passwordReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	jWTResponse, err := ctrl.oauthSevice.HandleTokenGeneration(clientIdReq, clientSecretReq, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res, _ := json.Marshal(jWTResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ctrl OAuthController) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	accessToken, ok := ctrl.GetBearerAuth(r)
	if !ok {
		http.Error(w, "Unable to find the Authentication Token", http.StatusUnauthorized)
		return
	}

	log.Println("auth token: ", accessToken)
	claims, err := ctrl.oauthSevice.ParseToken(accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res, _ := json.Marshal(claims)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ctrl OAuthController) GetBearerAuth(r *http.Request) (string, bool) {
	return ctrl.GetAuthHeader(r, "Bearer ")
}

func (ctrl OAuthController) GetBasicAuth(r *http.Request) (string, bool) {
	return ctrl.GetAuthHeader(r, "Basic ")
}

func (ctrl OAuthController) GetAuthHeader(r *http.Request, prefix string) (string, bool) {
	auth := r.Header.Get("Authorization")
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
