package provider

import (
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/usecase/sso"
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
	req webreq.HTTP,
	clientID GoogleClientID,
	clientSecret GoogleClientSecret,
	redirectURI GoogleRedirectURI,
) google.IdentityProvider {
	return google.NewIdentityProvider(req, string(clientID), string(clientSecret), string(redirectURI))
}

// NewGoogleAccountLinker creates GoogleAccountLinker.
func NewGoogleAccountLinker(
	factory sso.AccountLinkerFactory,
	googleSSORepo sqldb.GoogleSSOSql,
) google.AccountLinker {
	return google.AccountLinker(factory.NewAccountLinker(googleSSORepo))
}

// NewGoogleSSO creates GoogleSingleSignOn.
func NewGoogleSSO(
	ssoFactory sso.Factory,
	identityProvider google.IdentityProvider,
	account google.Account,
	linker google.AccountLinker,
) google.SingleSignOn {
	return google.SingleSignOn(
		ssoFactory.NewSingleSignOn(
			identityProvider,
			account,
			sso.AccountLinker(linker),
			google.NewInstrumentationFactory()),
	)
}

// GoogleAPIKey represents the credential for Google APIs.
type GoogleAPIKey string

// NewSafeBrowsing creates new SafeBrowsing with GoogleAPIKey to uniquely
// identify apiKey during dependency injection.
func NewSafeBrowsing(
	apiKey GoogleAPIKey,
	httpRequest webreq.HTTP,
) google.SafeBrowsing {
	return google.NewSafeBrowsing(string(apiKey), httpRequest)
}
