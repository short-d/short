package provider

import (
	"short/app/adapter/github"
	"short/app/adapter/oauth"
	"short/app/adapter/routing"
	"short/app/usecase/auth"
	"short/app/usecase/service"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

type WwwRoot string

func ShortRoutes(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot WwwRoot,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubOAuth oauth.Github,
	githubAPI github.API,
	authenticator auth.Authenticator,
	accountService service.Account,
) []fw.Route {
	return routing.NewShort(
		logger,
		tracer,
		string(wwwRoot),
		timer,
		urlRetriever,
		githubOAuth,
		githubAPI,
		authenticator,
		accountService,
	)
}
