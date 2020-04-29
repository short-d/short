package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/google"
)

// GoogleClientID represents client ID used for Google OAuth.
type GoogleClientID string

// GoogleClientSecret represents client secret used for Google OAuth.
type GoogleClientSecret string

// GoogleRedirectURI represents redirect URL for Google single sign on.
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

// GoogleAPIKey represents the credential for Google APIs.
type GoogleAPIKey string

// NewSafeBrowsing creates new SafeBrowsing with GoogleAPIKey to uniquely
// identify apiKey during dependency injection.
func NewSafeBrowsing(
	apiKey GoogleAPIKey,
	httpRequest fw.HTTPRequest,
) google.SafeBrowsing {
	return google.NewSafeBrowsing(string(apiKey), httpRequest)
}
