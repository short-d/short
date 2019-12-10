package google

import (
	"fmt"
	"net/http"
	"net/url"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const (
	authorizationAPI = "https://accounts.google.com/o/oauth2/v2/auth"
	accessTokenAPI   = "https://www.googleapis.com/oauth2/v4/token"
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
	scope := "profile"
	includeGrantedScopes := "true"
	responseType := "code"
	clientID := g.clientID
	redirectURI, err := url.QueryUnescape(g.redirectURI)
	if err != nil {
		return ""
	}
	fmt.Println(redirectURI)
	fmt.Println(g.redirectURI)
	u, err := url.Parse(authorizationAPI)
	if err != nil {
		return ""
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
	grantType := "authorization_code"
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI, err := url.QueryUnescape(g.redirectURI)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(accessTokenAPI)
	if err != nil {
		return "", err
	}

	query := u.Query()
	query.Set("code", authorizationCode)
	query.Set("client_id", clientID)
	query.Set("client_secret", clientSecret)
	query.Set("redirect_uri", redirectURI)
	query.Set("grant_type", grantType)
	u.RawQuery = query.Encode()

	body := url.Values{}
	body.Set("code", authorizationCode)
	body.Set("client_id", clientID)
	body.Set("client_secret", clientSecret)
	body.Set("redirect_uri", redirectURI)
	body.Set("grant_type", grantType)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	apiRes := accessTokenResponse{}
	err = g.httpRequest.JSON(http.MethodPost, u.String(), headers, body.Encode(), &apiRes)
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
