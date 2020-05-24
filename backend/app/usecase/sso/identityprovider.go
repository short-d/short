package sso

// IdentityProvider represents external service that verifies the user's
// identity.
type IdentityProvider interface {
	GetAuthorizationShortLink() string
	RequestAccessToken(authorizationCode string) (accessToken string, err error)
}
