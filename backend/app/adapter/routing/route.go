package routing

import (
	netURL "net/url"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/adapter/routing/analytics"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/feature"
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
) []router.Route {
	frontendURL, err := netURL.Parse(webFrontendURL)
	if err != nil {
		panic(err)
	}
	return []router.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewSSOSignIn(
				sso.SingleSignOn(githubSSO),
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewSSOSignInCallback(
				sso.SingleSignOn(githubSSO),
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in",
			Handle: NewSSOSignIn(
				sso.SingleSignOn(facebookSSO),
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/facebook/sign-in/callback",
			Handle: NewSSOSignInCallback(
				sso.SingleSignOn(facebookSSO),
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in",
			Handle: NewSSOSignIn(
				sso.SingleSignOn(googleSSO),
				webFrontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/oauth/google/sign-in/callback",
			Handle: NewSSOSignInCallback(
				sso.SingleSignOn(googleSSO),
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewLongLink(
				instrumentationFactory,
				shortLinkRetriever,
				timer,
				*frontendURL,
			),
		},
		{
			Method: "GET",
			Path:   "/features/:featureID",
			Handle: FeatureHandle(
				instrumentationFactory,
				featureDecisionMakerFactory,
				authenticator,
			),
		},
		{
			Method: "GET",
			Path:   "/analytics/track/:event",
			Handle: analytics.TrackHandle(instrumentationFactory),
		},
		{
			Method: "POST",
			Path:   "/api/search",
			Handle: SearchHandle(),
		},
	}
}
