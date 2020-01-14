package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/facebook"
	"github.com/short-d/short/app/adapter/github"
	"github.com/short-d/short/app/adapter/google"
	"github.com/short-d/short/app/adapter/routing"
	"github.com/short-d/short/app/usecase/account"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/url"
)

// WebFrontendURL represents the URL of the web frontend
type WebFrontendURL string

// NewShortRoutes creates HTTP routes for Short API with WwwRoot to uniquely identify WwwRoot during dependency injection.
func NewShortRoutes(
	logger fw.Logger,
	tracer fw.Tracer,
	webFrontendURL WebFrontendURL,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubAPI github.API,
	facebookAPI facebook.API,
	googleAPI google.API,
	authenticator auth.Authenticator,
	accountProvider account.Provider,
) []fw.Route {
	observability := routing.Observability{
		Logger: logger,
		Tracer: tracer,
	}

	return routing.NewShort(
		observability,
		string(webFrontendURL),
		timer,
		urlRetriever,
		githubAPI,
		facebookAPI,
		googleAPI,
		authenticator,
		accountProvider,
	)
}
