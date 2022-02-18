package config

import (
	"net/http"
	"strings"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/gorilla/mux"
)

var (
	DefaultSkippedPaths = [...]string{
		"/healthcheck/",
		"/oauth/token",
	}
)

type AuthConfig struct {
	EnableAuth   bool
	SkippedPaths []string
	OAuthService services.OAuthService
}

func NewAuthConfig(router *mux.Router, enableAuth bool, skippedPaths []string, authService services.OAuthService) {
	if enableAuth {
		authConfig := AuthConfig{EnableAuth: enableAuth, SkippedPaths: skippedPaths, OAuthService: authService}
		router.Use(authConfig.AuthMiddleware)
	}
}

func (authConfig AuthConfig) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			if !authConfig.isSkippedPath(req.URL.Path) {
				ok, err := authConfig.OAuthService.ValidateToken(req)
				if !ok || err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, req)
		})
}

func (authConfig AuthConfig) isSkippedPath(path string) bool {
	for i := 0; i < len(authConfig.SkippedPaths); i++ {
		if strings.HasPrefix(path, authConfig.SkippedPaths[i]) {
			return true
		}
	}

	return false
}
