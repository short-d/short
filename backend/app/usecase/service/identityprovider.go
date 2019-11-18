package service

type IdentityProvider interface {
	GetAuthorizationURL() string
	RequestAccessToken(authorizationCode string) (accessToken string, err error)
}
