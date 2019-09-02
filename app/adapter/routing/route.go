package routing

import (
	"short/app/adapter/github"
	"short/app/adapter/oauth"
	"short/app/usecase/auth"
	"short/app/usecase/service"
	"short/app/usecase/signin"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

func NewShort(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubOAuth oauth.Github,
	githubAPI github.API,
	authenticator auth.Authenticator,
	accountService service.Account,
) []fw.Route {
	githubSignIn := signin.NewOAuth(githubOAuth, githubAPI, accountService, authenticator)
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewGithubSignIn(logger, tracer, githubOAuth, authenticator),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewGithubSignInCallback(logger, tracer, githubSignIn),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalURL(logger, tracer, urlRetriever, timer),
		},
		{
			Method:      "GET",
			MatchPrefix: true,
			Path:        "/",
			Handle:      NewServeFile(logger, tracer, wwwRoot),
		},
	}
}
