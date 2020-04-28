package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/facebook"
	"github.com/short-d/short/app/adapter/github"
	"github.com/short-d/short/app/adapter/google"
	"github.com/short-d/short/app/adapter/request"
	"github.com/short-d/short/app/adapter/routing"
	"github.com/short-d/short/app/usecase/account"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/feature"
	"github.com/short-d/short/app/usecase/url"
)

// WebFrontendURL represents the URL of the web frontend
type WebFrontendURL string

// NewShortRoutes creates HTTP routes for Short API with WwwRoot to uniquely identify WwwRoot during dependency injection.
func NewShortRoutes(
	instrumentationFactory request.InstrumentationFactory,
	webFrontendURL WebFrontendURL,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubAPI github.API,
	facebookAPI facebook.API,
	googleAPI google.API,
	featureDecisionMakerFactory feature.DecisionMakerFactory,
	authenticator authenticator.Authenticator,
	accountProvider account.Provider,
) []fw.Route {
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
