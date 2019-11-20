package routing

import (
	netURL "net/url"
	"short/app/adapter/facebook"
	"short/app/adapter/github"
	"short/app/usecase/auth"
	"short/app/usecase/service"
	"short/app/usecase/sso"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

// Observability represents a set of tools to improves observability of the
// system.
type Observability struct {
	Logger fw.Logger
	Tracer fw.Tracer
}

func NewShort(
	observability Observability,
	webFrontendURL string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubAPI github.API,
	facebookAPI facebook.API,
	authenticator auth.Authenticator,
	accountService service.Account,
) []fw.Route {
	githubSignIn := sso.NewSingleSignOn(
		githubAPI.IdentityProvider,
		githubAPI.Account,
		accountService,
		authenticator,
	)
	facebookSignIn := sso.NewSingleSignOn(
		facebookAPI.IdentityProvider,
		facebookAPI.Account,
		accountService,
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
