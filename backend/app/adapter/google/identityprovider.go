package google

import (
	"fmt"
	"net/http"
	"net/url"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const (
	authorizationAPI     = "accounts.google.com"
	accessTokenAPI       = "www.googleapis.com"
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
	httpRequest  fw.HTTPRequest
	redirectURI  string
}

// GetAuthorizationURL retrieves the URL of Google sign in page.
func (g IdentityProvider) GetAuthorizationURL() string {
	clientID := g.clientID
	redirectURI := g.redirectURI
	u := &url.URL{
		Scheme: "https",
		Host:   authorizationAPI,
		Path:   "o/oauth2/v2/auth",
		RawQuery: fmt.Sprintf("&client_id=%s&redirect_uri=%s&scope=%s&access_type=%s&include_granted_scopes=%s&response_type=%s",
			clientID, redirectURI, scope, accessType, includeGrantedScopes, responseType),
	}
	return u.String()
}

// RequestAccessToken retrieves access token of user's Google account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (string, error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.redirectURI

	u := &url.URL{
		Scheme: "https",
		Host:   accessTokenAPI,
		Path:   "oauth2/v3/token",
		RawQuery: fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=%s",
			authorizationCode, clientID, clientSecret, redirectURI, grantType),
	}
	headers := map[string]string{}

	apiRes := accessTokenResponse{}
	err := g.httpRequest.JSON(http.MethodPost, u.String(), headers, "", &apiRes)
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
		httpRequest:  http,
		redirectURI:  url.QueryEscape(redirectURI),
	}
}
