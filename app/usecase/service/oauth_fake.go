package service

var _ OAuth = (*OAuthFake)(nil)

type OAuthFake struct {
	authUrl     string
	accessToken string
}

func (O OAuthFake) GetAuthorizationURL() string {
	return O.authUrl
}

func (O OAuthFake) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return O.accessToken, nil
}

func NewOAuthFake(authUrl string, accessToken string) OAuthFake {
	return OAuthFake{
		authUrl:     authUrl,
		accessToken: accessToken,
	}
}
