package sso

var _ IdentityProvider = (*IdentityProviderFake)(nil)

// IdentityProviderFake represents in memory implementation of an external
// authentication service.
type IdentityProviderFake struct {
	authURL     string
	accessToken string
}

// GetAuthorizationURL retrieves the URL where user can sign in and obtain
// authorization code.
func (i IdentityProviderFake) GetAuthorizationURL() string {
	return i.authURL
}

// RequestAccessToken retrieves access token given authorization code.
func (i IdentityProviderFake) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return i.accessToken, nil
}

// NewIdentityProviderFake creates fake IdentityProvider.
func NewIdentityProviderFake(authURL string, accessToken string) IdentityProviderFake {
	return IdentityProviderFake{
		authURL:     authURL,
		accessToken: accessToken,
	}
}
