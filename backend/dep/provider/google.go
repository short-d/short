package provider

import (
	"short/app/adapter/google"

	"github.com/byliuyang/app/fw"
)

// GoogleClientID represents client ID used for Github OAuth.
type GoogleClientID string

// GoogleClientSecret represents client secret used for Github OAuth.
type GoogleClientSecret string

// GoogleRedirectURI represents redirect URL for facebook single sign on.
type GoogleRedirectURI string

// NewGoogleIdentityProvider creates a new Google OAuth client with
// GoogleClientID and GoogleClientSecret to uniquely identify clientID and
// clientSecret during dependency injection.
func NewGoogleIdentityProvider(
	req fw.HTTPRequest,
	clientID GoogleClientID,
	clientSecret GoogleClientSecret,
	redirectURI GoogleRedirectURI,
) google.IdentityProvider {
	return google.NewIdentityProvider(req, string(clientID), string(clientSecret), string(redirectURI))
}
