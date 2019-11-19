package service

// IdentityProvider represents external service that verifies the user's
// identity.
type IdentityProvider interface {
	GetAuthorizationURL() string
	RequestAccessToken(authorizationCode string) (accessToken string, err error)
}
