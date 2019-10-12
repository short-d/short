package service

type OAuth interface {
	GetAuthorizationURL() string
	RequestAccessToken(authorizationCode string) (accessToken string, err error)
}
