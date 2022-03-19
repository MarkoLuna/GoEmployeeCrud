package config

import (
	"errors"
	"net/http"
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
			ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {

				accessToken, ok := GetBearerAuth(c.Request().Header)
				if !ok {
					return nil, errors.New("invalid token")
				}
				token, err := authConfig.OAuthService.ParseToken(accessToken)
				if err != nil {
					return nil, err
				}

				if !token.Valid {
					return nil, errors.New("invalid token")
				}
				return token, nil
			},
		}

		echoInstance.Use(middleware.JWTWithConfig(defaultJWTConfig))
	}
}

func GetBearerAuth(headers http.Header) (string, bool) {
	return GetAuthHeader(headers, "Bearer ")
}

func GetBasicAuth(headers http.Header) (string, bool) {
	return GetAuthHeader(headers, "Basic ")
}

func GetAuthHeader(headers http.Header, prefix string) (string, bool) {
	auth := headers.Get("Authorization")
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}

	return token, token != ""
}

func (authConfig AuthConfig) isSkippedPath(path string) bool {
	for i := 0; i < len(authConfig.SkippedPaths); i++ {
		if strings.HasPrefix(path, authConfig.SkippedPaths[i]) {
			return true
		}
	}

	return false
}
