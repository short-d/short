package service

type OAuth interface {
	GetAuthorizationUrl(scopes []string) string
	RequestAccessToken(authorizationCode string) (accessToken string, tokenType string, err error)
}
