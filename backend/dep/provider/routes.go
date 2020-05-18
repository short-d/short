package provider

import (
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/adapter/routing"
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
	featureDecisionMakerFactory feature.DecisionMakerFactory,
	githubSSO github.SingleSignOn,
	facebookSSO facebook.SingleSignOn,
	googleSSO google.SingleSignOn,
) []router.Route {
	return routing.NewShort(
		instrumentationFactory,
		string(webFrontendURL),
		timer,
		urlRetriever,
		featureDecisionMakerFactory,
		githubSSO,
		facebookSSO,
		googleSSO,
	)
}
