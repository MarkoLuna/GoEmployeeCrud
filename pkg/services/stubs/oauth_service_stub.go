package stubs

import (
	"net/http"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/dto"
)

type OAuthServiceStub struct {
	AuthenticatedPath string
}

func NewOAuthServiceStub() OAuthServiceStub {
	return OAuthServiceStub{AuthenticatedPath: "/api/employee"}
}

func (eSrv OAuthServiceStub) HandleTokenGeneration(clientId string, clientSecret string, userId string) (dto.JWTResponse, error) {

	access := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9" + "." + "eyJhdWQiOiJjNmNlY2U1MyIsImV4cCI6MTY0Mjc5MTUzNiwic3ViIjoiMDAwMDAwIn0" + "." + "SA49Q2UZvzf7dgZmvzNTaBjF1aYP821iXZje2pxK1KgvjZlQNOmQQ1B1duxfDkXeIWUfbFi2dkzlXx4GcWOVeg"
	refresh := "ZJVLYTRINZUTZJNLMY01MZLLLWFJMMMTYMM3Y2YZNTQ1MWM2"

	jWTResponse := dto.JWTResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    int64(120), // 2 min
		Scope:        "all",
		TokenType:    "Bearer",
	}

	return jWTResponse, nil
}

func (oauthService OAuthServiceStub) ParseToken(accessToken string) (map[string]string, error) {

	dataMap := make(map[string]string)
	dataMap["subject"] = "000000"
	dataMap["audience"] = "client"
	dataMap["id"] = ""
	dataMap["issuer"] = ""

	return dataMap, nil
}

func (oauthService OAuthServiceStub) IsAuthenticated(req *http.Request) (bool, error) {

	return true, nil
}

func (oauthService OAuthServiceStub) ValidateToken(r *http.Request) (bool, error) {
	return true, nil
}
