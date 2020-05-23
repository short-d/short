package google

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/usecase/sso"
)

const (
	authorizationAPI = "https://accounts.google.com/o/oauth2/v2/auth"
	accessTokenAPI   = "https://www.googleapis.com/oauth2/v4/token"
)

var _ sso.IdentityProvider = (*IdentityProvider)(nil)

type scope = string

const (
	email   = "email"
	profile = "profile"
)

// IdentityProvider represents Google OAuth service.
type IdentityProvider struct {
	clientID     string
	clientSecret string
	httpRequest  webreq.HTTP
	redirectURI  string
}

// GetAuthorizationShortLink retrieves the URL of Google sign in page.
func (g IdentityProvider) GetAuthorizationShortLink() string {
	clientID := g.clientID
	redirectURI := g.redirectURI

	u, err := url.Parse(authorizationAPI)
	if err != nil {
		return ""
	}

	scopes := []scope{email, profile}

	query := u.Query()
	query.Set("client_id", clientID)
	query.Set("redirect_uri", redirectURI)
	query.Set("scope", strings.Join(scopes, " "))
	query.Set("include_granted_scopes", "true")
	query.Set("response_type", "code")
	u.RawQuery = query.Encode()

	return u.String()
}

// RequestAccessToken retrieves access token of user's Google account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (string, error) {
	grantType := "authorization_code"
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.redirectURI

	u, err := url.Parse(accessTokenAPI)
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Set("code", authorizationCode)
	query.Set("client_id", clientID)
	query.Set("client_secret", clientSecret)
	query.Set("redirect_uri", redirectURI)
	query.Set("grant_type", grantType)
	u.RawQuery = query.Encode()

	apiRes := accessTokenResponse{}
	err = g.httpRequest.JSON(http.MethodPost, u.String(), map[string]string{}, "", &apiRes)
	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// NewIdentityProvider initializes Google OAuth service.
func NewIdentityProvider(http webreq.HTTP, clientID string, clientSecret string, redirectURI string) IdentityProvider {
	return IdentityProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		httpRequest:  http,
		redirectURI:  redirectURI,
	}
}
