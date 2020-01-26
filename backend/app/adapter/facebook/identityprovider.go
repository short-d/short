package facebook

import (
	"net/http"
	"net/url"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/service"
)

// More info here: https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow

const (
	fbAuthorizationAPI = "https://www.facebook.com/v4.0/dialog/oauth"
	fbAccessTokenAPI   = "https://graph.facebook.com/v4.0/oauth/access_token"
	fbScopes           = "public_profile,email"
	fbResponseType     = "code"
)

var _ service.IdentityProvider = (*IdentityProvider)(nil)

// IdentityProvider represents Facebook OAuth service.
type IdentityProvider struct {
	clientID     string
	clientSecret string
	http         fw.HTTPRequest
	redirectURI  string
}

// GetAuthorizationURL retrieves the URL of Facebook sign in page.
func (g IdentityProvider) GetAuthorizationURL() string {
	clientID := g.clientID
	redirectURI := g.redirectURI
	responseType := fbResponseType
	scope := fbScopes

	u, err := url.Parse(fbAuthorizationAPI)
	if err != nil {
		return ""
	}

	query := u.Query()
	query.Set("client_id", clientID)
	query.Set("redirect_uri", redirectURI)
	query.Set("scope", scope)
	query.Set("response_type", responseType)
	u.RawQuery = query.Encode()

	return u.String()
}

// RequestAccessToken retrieves access token of user's Facebook account using
// authorization code.
func (g IdentityProvider) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	clientID := g.clientID
	clientSecret := g.clientSecret
	redirectURI := g.redirectURI

	u, err := url.Parse(fbAccessTokenAPI)
	if err != nil {
		return "", err
	}

	query := u.Query()
	query.Set("client_id", clientID)
	query.Set("redirect_uri", redirectURI)
	query.Set("client_secret", clientSecret)
	query.Set("code", authorizationCode)
	u.RawQuery = query.Encode()

	body := url.Values{}
	body.Set("client_id", clientID)
	body.Set("redirect_uri", redirectURI)
	body.Set("client_secret", clientSecret)
	body.Set("code", authorizationCode)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	apiRes := fbAccessTokenResponse{}
	err = g.http.JSON(http.MethodPost, u.String(), headers, body.Encode(), &apiRes)

	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

type fbAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
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
		redirectURI:  redirectURI,
	}
}
