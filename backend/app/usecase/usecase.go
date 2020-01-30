package usecase

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/url"
)

type UseCase interface {
	RequestGithubSignIn(authToken string, presenter Presenter)
}

var _ UseCase = (*Short)(nil)

type Short struct {
	logger           fw.Logger
	timer            fw.Timer
	urlRetriever     url.Retriever
	authenticator    auth.Authenticator
	githubIDProvider service.IdentityProvider
}

func (s Short) RequestGithubSignIn(authToken string, presenter Presenter) {
	s.requestSSOSignIn(authToken, s.githubIDProvider, presenter)
}

func (s Short) requestSSOSignIn(
	authToken string,
	identityProvider service.IdentityProvider,
	presenter Presenter,
) {
	if s.authenticator.IsSignedIn(authToken) {
		presenter.ShowHome()
		return
	}
	signInLink := identityProvider.GetAuthorizationURL()
	presenter.ShowExternalPage(signInLink)
}

type GithubIDProvider service.IdentityProvider
type GithubSSOAccount service.SSOAccount

func NewShort(
	logger fw.Logger,
	timer fw.Timer,
	urlRetriever url.Retriever,
	authenticator auth.Authenticator,
	githubIDProvider GithubIDProvider,
) Short {
	return Short{
		logger:           logger,
		timer:            timer,
		urlRetriever:     urlRetriever,
		authenticator:    authenticator,
		githubIDProvider: githubIDProvider,
	}
}
