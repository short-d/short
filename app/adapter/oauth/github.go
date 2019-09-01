package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"short/fw"
)

const (
	authorizationAPI     = "https://github.com/login/oauth/authorize"
	accessTokenAPI       = "https://github.com/login/oauth/access_token"
	readUserProfileScope = "read:user"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type Github struct {
	clientID     string
	clientSecret string
	http         fw.HTTPRequest
}

func (g Github) GetAuthorizationURL() string {
	scopes := strings.Join([]string{
		readUserProfileScope,
	}, " ")
	escapedScope := url.QueryEscape(scopes)
	clientID := g.clientID
	return fmt.Sprintf("%s?client_id=%s&scope=%s", authorizationAPI, clientID, escapedScope)
}

func (g Github) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	body := fmt.Sprintf("client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, authorizationCode)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	apiRes := accessTokenResponse{}
	err = g.http.JSON(http.MethodPost, accessTokenAPI, headers, body, &apiRes)
	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

func NewGithub(http fw.HTTPRequest, clientID string, clientSecret string) Github {
	return Github{
		clientID:     clientID,
		clientSecret: clientSecret,
		http:         http,
	}
}
