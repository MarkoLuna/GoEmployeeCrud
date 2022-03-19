package services

import (
	"net/http"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/dto"
	"github.com/golang-jwt/jwt"
)

type OAuthService interface {
	HandleTokenGeneration(clientId string, clientSecret string, userId string) (dto.JWTResponse, error)
	ParseToken(accessToken string) (*jwt.Token, error)
	GetTokenClaims(accessToken string) (map[string]string, error)
	IsAuthenticated(req *http.Request) (bool, error)
	IsValidToken(accessToken string) (bool, error)
}
