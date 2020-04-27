package routing

import (
	netURL "net/url"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/facebook"
	"github.com/short-d/short/app/adapter/github"
	"github.com/short-d/short/app/adapter/google"
	"github.com/short-d/short/app/adapter/request"
	"github.com/short-d/short/app/usecase/account"
	"github.com/short-d/short/app/usecase/authenticator"

	"github.com/short-d/short/app/usecase/feature"
	"github.com/short-d/short/app/usecase/sso"
	"github.com/short-d/short/app/usecase/url"
)

// Observability represents a set of metrics data producers which improve the observability of the
// system, such as logger and tracer.
type Observability struct {
	Logger fw.Logger
	Tracer fw.Tracer
}

// NewShort creates HTTP routing table.
func NewShort(
	instrumentationFactory request.InstrumentationFactory,
	webFrontendURL string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubAPI github.API,
	facebookAPI facebook.API,
	googleAPI google.API,
	featureDecisionFactory feature.DecisionFactory,
	auth authenticator.Authenticator,
	accountProvider account.Provider,
) []fw.Route {
	githubSignIn := sso.NewSingleSignOn(
		githubAPI.IdentityProvider,
		githubAPI.Account,
		accountProvider,
		auth,
	)
	facebookSignIn := sso.NewSingleSignOn(
		facebookAPI.IdentityProvider,
		facebookAPI.Account,
		accountProvider,
		auth,
	)
	googleSignIn := sso.NewSingleSignOn(
		googleAPI.IdentityProvider,
		googleAPI.Account,
		accountProvider,
		auth,
	)
	frontendURL, err := netURL.Parse(webFrontendURL)
	if err != nil {
		panic(err)
	}
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewSSOSignIn(
				githubAPI.IdentityProvider,
				auth,
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewSSOSignInCallback(
				githubSignIn,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in",
			Handle: NewSSOSignIn(
				facebookAPI.IdentityProvider,
				auth,
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in/callback",
			Handle: NewSSOSignInCallback(
				facebookSignIn,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in",
			Handle: NewSSOSignIn(
				googleAPI.IdentityProvider,
				auth,
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in/callback",
			Handle: NewSSOSignInCallback(
				googleSignIn,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalURL(
				instrumentationFactory,
				urlRetriever,
				timer,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/features/:featureID",
			Handle: FeatureHandle(instrumentationFactory, featureDecisionFactory),
		},
	}
}
