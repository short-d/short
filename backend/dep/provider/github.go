package provider

import (
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/usecase/sso"
)

// GithubClientID represents client ID used for Github OAuth.
type GithubClientID string

// GithubClientSecret represents client secret used for Github OAuth.
type GithubClientSecret string

// NewGithubIdentityProvider creates a new Github OAuth client with
// GithubClientID and GithubClientSecret to uniquely identify clientID and
// clientSecret during dependency injection.
func NewGithubIdentityProvider(
	req webreq.HTTP,
	clientID GithubClientID,
	clientSecret GithubClientSecret,
) github.IdentityProvider {
	return github.NewIdentityProvider(req, string(clientID), string(clientSecret))
}

// NewGithubAccountLinker creates GithubAccountLinker.
func NewGithubAccountLinker(
	factory sso.AccountLinkerFactory,
	ssoMap sqldb.GithubSSOSql,
) github.AccountLinker {
	return github.AccountLinker(factory.NewAccountLinker(ssoMap))
}

// NewGithubSSO creates GithubSingleSignOn.
func NewGithubSSO(
	ssoFactory sso.Factory,
	accountLinker github.AccountLinker,
	identityProvider github.IdentityProvider,
	account github.Account,
) github.SingleSignOn {
	return github.SingleSignOn(
		ssoFactory.NewSingleSignOn(
			identityProvider,
			account,
			sso.AccountLinker(accountLinker),
			github.NewInstrumentationFactory(),
		),
	)
}
