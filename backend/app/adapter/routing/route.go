package routing

import (
	netURL "net/url"

	"github.com/short-d/short/app/adapter/facebook"
	"github.com/short-d/short/app/adapter/github"
	"github.com/short-d/short/app/adapter/google"
	"github.com/short-d/short/app/usecase/account"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/sso"
	"github.com/short-d/short/app/usecase/url"

	"github.com/short-d/app/fw"
)

// Observability represents a set of metrics data producers which improve the observability of the
// system, such as logger and tracer.
type Observability struct {
	Logger fw.Logger
	Tracer fw.Tracer
}

// NewShort creates HTTP routing table.
func NewShort(
	observability Observability,
	webFrontendURL string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubAPI github.API,
	facebookAPI facebook.API,
	googleAPI google.API,
	authenticator auth.Authenticator,
	accountProvider account.Provider,
) []fw.Route {
	githubSignIn := sso.NewSingleSignOn(
		githubAPI.IdentityProvider,
		githubAPI.Account,
		accountProvider,
		authenticator,
	)
	facebookSignIn := sso.NewSingleSignOn(
		facebookAPI.IdentityProvider,
		facebookAPI.Account,
		accountProvider,
		authenticator,
	)
	googleSignIn := sso.NewSingleSignOn(
		googleAPI.IdentityProvider,
		googleAPI.Account,
		accountProvider,
		authenticator,
	)
	frontendURL, err := netURL.Parse(webFrontendURL)
	if err != nil {
		panic(err)
	}
	logger := observability.Logger
	tracer := observability.Tracer
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewSSOSignIn(
				logger,
				tracer,
				githubAPI.IdentityProvider,
				authenticator,
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewSSOSignInCallback(
				logger,
				tracer,
				githubSignIn,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in",
			Handle: NewSSOSignIn(
				logger,
				tracer,
				facebookAPI.IdentityProvider,
				authenticator,
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in/callback",
			Handle: NewSSOSignInCallback(
				logger,
				tracer,
				facebookSignIn,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in",
			Handle: NewSSOSignIn(
				logger,
				tracer,
				googleAPI.IdentityProvider,
				authenticator,
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in/callback",
			Handle: NewSSOSignInCallback(
				logger,
				tracer,
				googleSignIn,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalURL(
				logger,
				tracer,
				urlRetriever,
				timer,
				*frontendURL,
			),
		},
	}
}
