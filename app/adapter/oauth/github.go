package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"short/fw"
	"strings"
)

const (
	authorizationApi = "https://github.com/login/oauth/authorize"
	accessTokenApi   = "https://github.com/login/oauth/access_token"
	userEmailScope   = "user:email"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type Github struct {
	clientId     string
	clientSecret string
	http         fw.HttpRequest
}

func (g Github) GetAuthorizationUrl() string {
	scopes := strings.Join([]string{
		userEmailScope,
	}, " ")
	escapedScope := url.QueryEscape(scopes)
	clientId := g.clientId
	return fmt.Sprintf("%s?client_id=%s&scope=%s", authorizationApi, clientId, escapedScope)
}

func (g Github) RequestAccessToken(authorizationCode string) (accessToken string, tokenType string, err error) {
	clientId := g.clientId
	clientSecret := g.clientSecret
	body := fmt.Sprintf("client_id=%s&client_secret=%s&code=%s", clientId, clientSecret, authorizationCode)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	apiRes := accessTokenResponse{}
	err = g.http.Json(http.MethodPost, accessTokenApi, headers, body, &apiRes)
	if err != nil {
		return "", "", err
	}

	return apiRes.AccessToken, apiRes.Scope, nil
}

func NewGithub(http fw.HttpRequest, clientId string, clientSecret string) Github {
	return Github{
		clientId:     clientId,
		clientSecret: clientSecret,
		http:         http,
	}
}
