package routing

import (
	"short/app/adapter/account"
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
	githubOAuth oauth.Github,
	githubAccount account.Github,
	authenticator auth.Authenticator,
) []fw.Route {
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewGithubSignIn(logger, tracer, githubOAuth, authenticator),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewGithubSignInCallback(logger, tracer, githubOAuth, githubAccount, authenticator),
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
