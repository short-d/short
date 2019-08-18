package routing

import (
	"short/app/adapter/oauth"
	"short/app/usecase/auth"
	"short/app/usecase/url"
	"short/fw"
)

func NewShort(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	github oauth.Github,
	authenticator auth.Authenticator,
) []fw.Route {
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewGithubSignIn(logger, tracer, github, authenticator),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalUrl(logger, tracer, urlRetriever, timer),
		},
		{
			Method:      "GET",
			MatchPrefix: true,
			Path:        "/",
			Handle:      NewServeFile(logger, tracer, wwwRoot),
		},
	}
}
