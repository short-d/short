package routing

import (
	"short/app/adapter/request"
	"short/app/usecase/url"
	"short/fw"
)

func NewShort(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot string,
	urlRetriever url.Retriever,
	req request.Http,
	githubClientId string,
	githubClientSecret string,
) []fw.Route {
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewGithubSignIn(logger, tracer, req, githubClientId, githubClientSecret),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalUrl(logger, tracer, urlRetriever),
		},
		{
			Method:      "GET",
			MatchPrefix: true,
			Path:        "/",
			Handle:      NewServeFile(logger, tracer, wwwRoot),
		},
	}
}
