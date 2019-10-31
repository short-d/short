package provider

import (
	"short/app/adapter/facebook"
	"short/app/adapter/github"
	"short/app/adapter/oauth"
	"short/app/adapter/routing"
	"short/app/usecase/auth"
	"short/app/usecase/service"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
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
	githubOAuth oauth.Github,
	githubAPI github.API,
	authenticator auth.Authenticator,
	accountService service.Account,
	facebookOAuth oauth.Facebook,
	facebookAPI facebook.API,
) []fw.Route {
	observability := routing.Observability{
		Logger: logger,
		Tracer: tracer,
	}
	gh := routing.Github{
		OAuth: githubOAuth,
		API:   githubAPI,
	}
	fb := routing.Facebook{
		OAuth: facebookOAuth,
		API:   facebookAPI,
	}
	return routing.NewShort(
		observability,
		string(webFrontendURL),
		timer,
		urlRetriever,
		gh,
		fb,
		authenticator,
		accountService,
	)
}
