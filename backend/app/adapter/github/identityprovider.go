package github

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/usecase/external"
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

var _ external.IdentityProvider = (*IdentityProvider)(nil)

// IdentityProvider represents Github OAuth service.
type IdentityProvider struct {
	clientID     string
	clientSecret string
	http         webreq.HTTP
}

// GetAuthorizationURL retrieves the URL of Github sign in page.
func (g IdentityProvider) GetAuthorizationURL() string {
	scopes := strings.Join([]string{
		readUserProfileScope,
	}, " ")
	escapedScope := url.QueryEscape(scopes)
	clientID := g.clientID
	return fmt.Sprintf("%s?client_id=%s&scope=%s", authorizationAPI, clientID, escapedScope)
}

// RequestAccessToken retrieves access token of user's Github account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	clientID := g.clientID
	clientSecret := g.clientSecret

	body := url.Values{}
	body.Set("client_id", clientID)
	body.Set("client_secret", clientSecret)
	body.Set("code", authorizationCode)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	apiRes := accessTokenResponse{}
	err = g.http.JSON(http.MethodPost, accessTokenAPI, headers, body.Encode(), &apiRes)
	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

// NewIdentityProvider initializes Github OAuth service.
func NewIdentityProvider(http webreq.HTTP, clientID string, clientSecret string) IdentityProvider {
	return IdentityProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		http:         http,
	}
}
