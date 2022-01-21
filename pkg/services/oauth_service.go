package services

import (
	"net/http"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/dto"
)

type OAuthService interface {
	HandleTokenGeneration(clientId string, clientSecret string, userId string) (dto.JWTResponse, error)
	ParseToken(accessToken string) (map[string]string, error)
	IsAuthenticated(req *http.Request) (bool, error)
	ValidateToken(r *http.Request) (bool, error)
}
