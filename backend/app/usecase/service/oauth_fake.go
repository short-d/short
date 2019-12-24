package service

var _ OAuth = (*OAuthFake)(nil)

// OAuthFake represents in memory implementation of an external oauth provider.
type OAuthFake struct {
	authUrl     string
	accessToken string
}

// GetAuthorizationURL retrieves authorization url of the oauth provider.
func (O OAuthFake) GetAuthorizationURL() string {
	return O.authUrl
}

// RequestAccessToken obtains access token of an user given scoped
// authorizationCode.
func (O OAuthFake) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return O.accessToken, nil
}

// NewOAuthFake creates fake OAuth provider.
func NewOAuthFake(authUrl string, accessToken string) OAuthFake {
	return OAuthFake{
		authUrl:     authUrl,
		accessToken: accessToken,
	}
}
