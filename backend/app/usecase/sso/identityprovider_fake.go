package sso

var _ IdentityProvider = (*IdentityProviderFake)(nil)

// IdentityProviderFake represents in memory implementation of an external
// authentication service.
type IdentityProviderFake struct {
	authShortLink string
	accessToken   string
}

// GetAuthorizationShortLink retrieves the ShortLink where user can sign in and obtain
// authorization code.
func (i IdentityProviderFake) GetAuthorizationShortLink() string {
	return i.authShortLink
}

// RequestAccessToken retrieves access token given authorization code.
func (i IdentityProviderFake) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return i.accessToken, nil
}

// NewIdentityProviderFake creates fake IdentityProvider.
func NewIdentityProviderFake(authShortLink string, accessToken string) IdentityProviderFake {
	return IdentityProviderFake{
		authShortLink: authShortLink,
		accessToken:   accessToken,
	}
}
