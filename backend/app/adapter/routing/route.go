package routing

import (
	"net/url"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/adapter/routing/analytics"
	"github.com/short-d/short/backend/app/adapter/routing/handle"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/feature"
	"github.com/short-d/short/backend/app/usecase/search"
	"github.com/short-d/short/backend/app/usecase/shortlink"
	"github.com/short-d/short/backend/app/usecase/sso"
)

type shortRouter struct {
	routes []router.Route
}

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

	var shortRouter shortRouter

	shortRouter.addRoute("GET", "/oauth/github/sign-in", handle.NewSSOSignIn(
		sso.SingleSignOn(githubSSO),
		webFrontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/github/sign-in/callback", handle.NewSSOSignInCallback(
		sso.SingleSignOn(githubSSO),
		*frontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/github/sign-in/callback", handle.NewSSOSignInCallback(
		sso.SingleSignOn(githubSSO),
		*frontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/github/sign-in/callback", handle.NewSSOSignInCallback(
		sso.SingleSignOn(githubSSO),
		*frontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/facebook/sign-in", handle.NewSSOSignIn(
		sso.SingleSignOn(facebookSSO),
		webFrontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/facebook/sign-in/callback", handle.NewSSOSignInCallback(
		sso.SingleSignOn(facebookSSO),
		*frontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/google/sign-in", handle.NewSSOSignIn(
		sso.SingleSignOn(googleSSO),
		webFrontendURL,
	))
	shortRouter.addRoute("GET", "/oauth/google/sign-in/callback", handle.NewSSOSignInCallback(
		sso.SingleSignOn(googleSSO),
		*frontendURL,
	))
	shortRouter.addRoute("GET", "/r/:alias", handle.NewLongLink(
		instrumentationFactory,
		shortLinkRetriever,
		timer,
		*frontendURL,
	))
	shortRouter.addRoute("GET", "/features/:featureID", handle.Feature(
		instrumentationFactory,
		featureDecisionMakerFactory,
		authenticator,
	))
	shortRouter.addRoute("GET", "/analytics/track/:event", analytics.TrackHandle(
		instrumentationFactory,
	))
	shortRouter.addRoute("POST", "/api/search", handle.Search(
		search,
	))

	return shortRouter.routes
}

func (s *shortRouter) addRoute(method, path string, handle router.Handle) []router.Route {
	s.routes = append(s.routes, router.Route{
		Method: method,
		Path:   path,
		Handle: handle,
	})
	return s.routes
}
