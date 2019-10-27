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

// Observability represents a set of tools to improves observability of the
// system.
type Observability struct {
	Logger fw.Logger
	Tracer fw.Tracer
}

// Github groups Github oauth and public APIs together.
type Github struct {
	OAuth oauth.Github
	API   github.API
}

func NewShort(
	observability Observability,
	webFrontendURL string,
	timer fw.Timer,
	urlRetriever url.Retriever,
	github Github,
	authenticator auth.Authenticator,
	accountService service.Account,
) []fw.Route {
	githubSignIn := signin.NewOAuth(github.OAuth, github.API, accountService, authenticator)
	frontendURL, err := netURL.Parse(webFrontendURL)
	if err != nil {
		panic(err)
	}
	logger := observability.Logger
	tracer := observability.Tracer
	return []fw.Route{
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in",
			Handle: NewGithubSignIn(logger, tracer, github.OAuth, authenticator, webFrontendURL),
		},
		{
			Method: "GET",
			Path:   "/oauth/github/sign-in/callback",
			Handle: NewGithubSignInCallback(logger, tracer, githubSignIn, *frontendURL),
		},
		{
			Method: "GET",
			Path:   "/r/:alias",
			Handle: NewOriginalURL(logger, tracer, urlRetriever, timer, *frontendURL),
		},
		{
			Method: "GET",
			Path:   "/health",
			Handle: NewHealthURL(logger),
		},
	}
}
