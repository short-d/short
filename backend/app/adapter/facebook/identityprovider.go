package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

// More info here: https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow

// TODO(byliuyang): rewrite this file
const (
	fbAuthorizationAPI = "https://www.facebook.com/v4.0/dialog/oauth"
	fbAccessTokenAPI   = "https://graph.facebook.com/v4.0/oauth/access_token"
	fbScopes           = "public_profile,email"
	fbResponseType     = "code"
)

type fbAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

var _ service.IdentityProvider = (*IdentityProvider)(nil)

// IdentityProvider represents Facebook OAuth service.
type IdentityProvider struct {
	clientID     string
	clientSecret string
	http         fw.HTTPRequest
	redirectURI  string
}

// GetAuthorizationURL retrieves the  URL of Facebook sign in page.
func (g IdentityProvider) GetAuthorizationURL() string {
	escapedScope := url.QueryEscape(fbScopes)
	clientID := g.clientID
	redirectURI := g.redirectURI
	responseType := fbResponseType

	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=%s",
		fbAuthorizationAPI, clientID, redirectURI, escapedScope, responseType)
}

// RequestAccessToken retrieves access token of user's Facebook account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.redirectURI

	u := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		fbAccessTokenAPI, clientID, clientSecret, authorizationCode, redirectURI)

	headers := map[string]string{}

	apiRes := fbAccessTokenResponse{}
	err = g.http.JSON(http.MethodGet, u, headers, "", &apiRes)

	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

// NewIdentityProvider initializes Facebook OAuth service.
func NewIdentityProvider(
	http fw.HTTPRequest,
	clientID string,
	clientSecret string,
	redirectURI string,
) IdentityProvider {
	return IdentityProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		http:         http,
		redirectURI:  url.QueryEscape(redirectURI),
	}
}
