package oauth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/byliuyang/app/fw"
)

// More info here: https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow

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

// Facebook represents FB OAuth configuration
type Facebook struct {
	clientID     string
	clientSecret string
	http         fw.HTTPRequest
	redirectURI  string
}

// GetAuthorizationURL returns an authorization URL to start FB authorization process
func (g Facebook) GetAuthorizationURL() string {
	escapedScope := url.QueryEscape(fbScopes)
	clientID := g.clientID
	redirectURI := g.RedirectURI()
	responseType := fbResponseType

	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=%s",
		fbAuthorizationAPI, clientID, redirectURI, escapedScope, responseType)
}

// RequestAccessToken performs a call to Facebook that converts OAuth code into access
func (g Facebook) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.RedirectURI()

	url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		fbAccessTokenAPI, clientID, clientSecret, authorizationCode, redirectURI)

	headers := map[string]string{}

	apiRes := fbAccessTokenResponse{}
	err = g.http.JSON(http.MethodGet, url, headers, "", &apiRes)

	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

// RedirectURI returns URL-escaped redirect_uri value
func (g Facebook) RedirectURI() (redirectURI string) {
	return url.QueryEscape(g.redirectURI)
}

// NewFacebook initializes Facebook object
func NewFacebook(http fw.HTTPRequest, clientID string, clientSecret string, redirectURI string) Facebook {
	return Facebook{
		clientID:     clientID,
		clientSecret: clientSecret,
		http:         http,
		redirectURI:  redirectURI,
	}
}
