package routing

import (
	"net/url"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/adapter/routing/handle"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/feature"
	"github.com/short-d/short/backend/app/usecase/search"
	"github.com/short-d/short/backend/app/usecase/shortlink"
	"github.com/short-d/short/backend/app/usecase/sso"
)

// NewShort creates HTTP routing table.
func NewShort(
	instrumentationFactory request.InstrumentationFactory,
	webFrontendURL string,
	timer timer.Timer,
	shortLinkRetriever shortlink.Retriever,
	featureDecisionMakerFactory feature.DecisionMakerFactory,
	githubSSO github.SingleSignOn,
	facebookSSO facebook.SingleSignOn,
	googleSSO google.SingleSignOn,
	authenticator authenticator.Authenticator,
	search search.Search,
) []router.Route {
	frontendURL, err := url.Parse(webFrontendURL)
	if err != nil {
		panic(err)
	}
	return []router.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: handle.SSOSignIn(
				sso.SingleSignOn(githubSSO),
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: handle.SSOSignInCallback(
				sso.SingleSignOn(githubSSO),
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in",
			Handle: handle.SSOSignIn(
				sso.SingleSignOn(facebookSSO),
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in/callback",
			Handle: handle.SSOSignInCallback(
				sso.SingleSignOn(facebookSSO),
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in",
			Handle: handle.SSOSignIn(
				sso.SingleSignOn(googleSSO),
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in/callback",
			Handle: handle.SSOSignInCallback(
				sso.SingleSignOn(googleSSO),
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: handle.LongLink(
				instrumentationFactory,
				shortLinkRetriever,
				timer,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/features/:featureID",
			Handle: handle.Feature(
				instrumentationFactory,
				featureDecisionMakerFactory,
				authenticator,
			),
		},
		{
			Method: "GET",
			Path:   "/analytics/track/:event",
			Handle: handle.Track(instrumentationFactory),
		},
		{
			Method: "POST",
			Path:   "/api/search",
			Handle: handle.Search(
				search,
				authenticator,
			),
		},
	}
}
