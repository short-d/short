package provider

import (
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/adapter/routing"
	"github.com/short-d/short/backend/app/usecase/account"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/feature"
	"github.com/short-d/short/backend/app/usecase/url"
)

// WebFrontendURL represents the URL of the web frontend
type WebFrontendURL string

// NewShortRoutes creates HTTP routes for Short API with WwwRoot to uniquely identify WwwRoot during dependency injection.
func NewShortRoutes(
	instrumentationFactory request.InstrumentationFactory,
	webFrontendURL WebFrontendURL,
	timer timer.Timer,
	urlRetriever url.Retriever,
	githubAPI github.API,
	facebookAPI facebook.API,
	googleAPI google.API,
	featureDecisionMakerFactory feature.DecisionMakerFactory,
	authenticator authenticator.Authenticator,
	accountProvider account.Provider,
) []router.Route {
	return routing.NewShort(
		instrumentationFactory,
		string(webFrontendURL),
		timer,
		urlRetriever,
		githubAPI,
		facebookAPI,
		googleAPI,
		featureDecisionMakerFactory,
		authenticator,
		accountProvider,
	)
}
