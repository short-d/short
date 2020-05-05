package facebook

import (
	"net/http"
	"net/url"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/app/usecase/external"
)

// More info here: https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow

const (
	fbAuthorizationAPI = "https://www.facebook.com/v4.0/dialog/oauth"
	fbAccessTokenAPI   = "https://graph.facebook.com/v4.0/oauth/access_token"
	fbScopes           = "public_profile,email"
	fbResponseType     = "code"
)

var _ external.IdentityProvider = (*IdentityProvider)(nil)

// IdentityProvider represents Facebook OAuth service.
type IdentityProvider struct {
	clientID     string
	clientSecret string
	httpRequest  webreq.HTTP
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
	type fbAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

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
	err = g.httpRequest.JSON(http.MethodPost, u.String(), headers, body.Encode(), &apiRes)

	if err != nil {
		return "", err
	}

	return apiRes.AccessToken, nil
}

// NewIdentityProvider initializes Facebook OAuth service.
func NewIdentityProvider(
	httpRequest webreq.HTTP,
	clientID string,
	clientSecret string,
	redirectURI string,
) IdentityProvider {
	return IdentityProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		httpRequest:  httpRequest,
		redirectURI:  redirectURI,
	}
}
