package provider

import (
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/usecase/sso"
)

// FacebookClientID represents client ID used for Facebook OAuth.
type FacebookClientID string

// FacebookClientSecret represents client secret used for Facebook OAuth.
type FacebookClientSecret string

// FacebookRedirectURI represents redirect URL for facebook single sign on.
type FacebookRedirectURI string

// NewFacebookIdentityProvider creates a new Facebook OAuth client with
// FacebookClientID and FacebookClientSecret to uniquely identify clientID and
// clientSecret during dependency injection.
func NewFacebookIdentityProvider(
	req webreq.HTTP,
	clientID FacebookClientID,
	clientSecret FacebookClientSecret,
	redirectURI FacebookRedirectURI,
) facebook.IdentityProvider {
	return facebook.NewIdentityProvider(req, string(clientID), string(clientSecret), string(redirectURI))
}

func NewFacebookSSO(
	ssoFactory sso.Factory,
	accountLinkerFactory sso.AccountLinkerFactory,
	identityProvider facebook.IdentityProvider,
	account facebook.Account,
	facebookSSORepo sqldb.FacebookSSOSql,
) facebook.SingleSignOn {
	linker := accountLinkerFactory.NewAccountLinker(facebookSSORepo)
	return facebook.SingleSignOn(ssoFactory.NewSingleSignOn(identityProvider, account, linker))
}
