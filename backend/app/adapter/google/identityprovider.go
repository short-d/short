package google

import (
	"net/http"
	"net/url"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const (
	authorizationAPI     = "https://accounts.google.com/o/oauth2/v2/auth"
	accessTokenAPI       = "https://www.googleapis.com/oauth2/v4/token"
	grantType            = "authorization_code"
	scope                = "https://www.googleapis.com/auth/userinfo.email"
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
	includeGrantedScopes := "true"
	responseType := "code"

	u := &url.URL{
		Host:   authorizationAPI,
	}
	q := u.Query()
	q.Set("client_id", clientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("scope", scope)
	q.Set("include_granted_scopes", includeGrantedScopes)
	q.Set("response_type", responseType)
	u.RawQuery = q.Encode()

	return u.String()
}

// RequestAccessToken retrieves access token of user's Google account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (string, error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.redirectURI

	u := &url.URL{
		Host:   accessTokenAPI,
	}
	q := u.Query()
	q.Set("code", authorizationCode)
	q.Set("client_id", clientID)
	q.Set("client_secret", clientSecret)
	q.Set("redirect_uri", redirectURI)
	q.Set("grant_type", grantType)
	u.RawQuery = q.Encode()

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

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
