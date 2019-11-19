package service

var _ IdentityProvider = (*IdentityProviderFake)(nil)

// IdentityProviderFake represents in memory implementation of an identity
// provider.
type IdentityProviderFake struct {
	authURL     string
	accessToken string
}

// GetAuthorizationURL gets the URL where user can sign in and obtain
// authorization code.
func (i IdentityProviderFake) GetAuthorizationURL() string {
	return i.authURL
}

// RequestAccessToken gets access token given authorization code.
func (i IdentityProviderFake) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return i.accessToken, nil
}

// NewIdentityProviderFake creates fake IdentityProvider
func NewIdentityProviderFake(authURL string, accessToken string) IdentityProviderFake {
	return IdentityProviderFake{
		authURL:     authURL,
		accessToken: accessToken,
	}
}
