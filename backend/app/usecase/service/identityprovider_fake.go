package service

var _ IdentityProvider = (*IdentityProviderFake)(nil)

type IdentityProviderFake struct {
	authUrl     string
	accessToken string
}

func (i IdentityProviderFake) GetAuthorizationURL() string {
	return i.authUrl
}

func (i IdentityProviderFake) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return i.accessToken, nil
}

func NewIdentityProviderFake(authUrl string, accessToken string) IdentityProviderFake {
	return IdentityProviderFake{
		authUrl:     authUrl,
		accessToken: accessToken,
	}
}
