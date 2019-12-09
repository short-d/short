package google

import (
	"fmt"
	"net/http"
	"net/url"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const (
	authorizationAPI     = "https://accounts.google.com/o/oauth2/v2/auth?"
	accessTokenAPI       = "https://www.googleapis.com/oauth2/v3/token"
	grantType            = "authorization_code"
	scope                = "https://www.googleapis.com/auth/userinfo.email"
	accessType           = "offline"
	includeGrantedScopes = "true"
	responseType         = "code"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

var _ service.IdentityProvider = (*IdentityProvider)(nil)

// IdentityProvider represents Google OAuth service.
type IdentityProvider struct {
	clientID     string
	clientSecret string
	httpRequest         fw.HTTPRequest
	redirectURI  string
}

// GetAuthorizationURL retrieves the URL of Google sign in page.
func (g IdentityProvider) GetAuthorizationURL() string {
	clientID := g.clientID
	redirectURI := g.redirectURI
	return fmt.Sprintf("%s&client_id=%s&redirect_uri=%s&scope=%s&access_type=%s&include_granted_scopes=%s&response_type=%s",
		authorizationAPI, clientID, redirectURI, scope, accessType, includeGrantedScopes, responseType)
}

// RequestAccessToken retrieves access token of user's Google account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (string, error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.redirectURI

	u := fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=%s",
		accessTokenAPI, authorizationCode, clientID, clientSecret, redirectURI, grantType)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	apiRes := accessTokenResponse{}
	err := g.http.JSON(http.MethodPost, u, headers, "", &apiRes)
	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

// NewIdentityProvider initializes Google OAuth service.
func NewIdentityProvider(http fw.HTTPRequest, clientID string, clientSecret string, redirectURI string) IdentityProvider {
	return IdentityProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		http:         http,
		redirectURI:  url.QueryEscape(redirectURI),
	}
}
