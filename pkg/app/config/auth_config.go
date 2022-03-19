package config

import (
	"strings"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DefaultSkippedPaths = [...]string{
		"/healthcheck/",
		"/oauth/token",
	}
)

var (
	signingKey = utils.GetEnv("OAUTH_SIGNING_KEY", "00000000")
)

type AuthConfig struct {
	EnableAuth   bool
	SkippedPaths []string
	OAuthService services.OAuthService
}

func NewAuthConfig(echoInstance *echo.Echo, enableAuth bool, skippedPaths []string, authService services.OAuthService) {
	if enableAuth {
		authConfig := AuthConfig{EnableAuth: enableAuth, SkippedPaths: skippedPaths, OAuthService: authService}

		defaultJWTConfig := middleware.JWTConfig{
			SigningKey: []byte(signingKey),
			// oauth skipper returns false which processes the middleware.
			Skipper: func(e echo.Context) bool {
				return authConfig.isSkippedPath(e.Request().URL.Path)
			},
			SigningMethod: middleware.AlgorithmHS256,
			TokenLookup:   "header:" + echo.HeaderAuthorization,
		}

		echoInstance.Use(middleware.JWTWithConfig(defaultJWTConfig))
	}
}

func (authConfig AuthConfig) isSkippedPath(path string) bool {
	for i := 0; i < len(authConfig.SkippedPaths); i++ {
		if strings.HasPrefix(path, authConfig.SkippedPaths[i]) {
			return true
		}
	}

	return false
}
