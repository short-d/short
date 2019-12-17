package service

// OAuth enables Short to obtain limited access to user account of external
// identity provider, such as Github, Facebook, and Google.
type OAuth interface {
	GetAuthorizationURL() string
	RequestAccessToken(authorizationCode string) (accessToken string, err error)
}
