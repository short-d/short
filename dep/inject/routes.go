package inject

import (
	"short/app/adapter/account"
	"short/app/adapter/oauth"
	"short/app/adapter/routing"
	"short/app/usecase/auth"
	"short/app/usecase/url"
	"short/fw"
)

type WwwRoot string

func ShortRoutes(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot WwwRoot,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubOAuth oauth.Github,
	githubAccount account.Github,
	authenticator auth.Authenticator,
) []fw.Route {
	return routing.NewShort(
		logger,
		tracer,
		string(wwwRoot),
		timer,
		urlRetriever,
		githubOAuth,
		githubAccount,
		authenticator,
	)
}
