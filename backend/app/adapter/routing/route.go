package routing

import (
	netURL "net/url"
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
	webFrontendURL string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubOAuth oauth.Github,
	githubAPI github.API,
	authenticator auth.Authenticator,
	accountService service.Account,
) []fw.Route {
	githubSignIn := signin.NewOAuth(githubOAuth, githubAPI, accountService, authenticator)
	frontendURL, err := netURL.Parse(webFrontendURL)
	if err != nil {
		panic(err)
	}
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewGithubSignIn(logger, tracer, githubOAuth, authenticator, frontendURL),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewGithubSignInCallback(logger, tracer, githubSignIn, frontendURL),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalURL(logger, tracer, urlRetriever, timer, frontendURL),
		},
	}
}
