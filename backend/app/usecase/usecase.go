package usecase

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/service"
)

// UseCase represents all the business logic for Short.
type UseCase struct {
	logger           fw.Logger
	timer            fw.Timer
	authenticator    auth.Authenticator
	githubIDProvider service.IdentityProvider
}

// RequestGithubSignIn directs user to Github sign in screen.
func (u UseCase) RequestGithubSignIn(authToken string, presenter Presenter) {
	u.requestSSOSignIn(authToken, u.githubIDProvider, presenter)
}

func (u UseCase) requestSSOSignIn(
	authToken string,
	identityProvider service.IdentityProvider,
	presenter Presenter,
) {
	if u.authenticator.IsSignedIn(authToken) {
		presenter.ShowHome()
		return
	}
	signInLink := identityProvider.GetAuthorizationURL()
	presenter.ShowExternalPage(signInLink)
}

// GithubIDProvider provides Github authentication service.
type GithubIDProvider service.IdentityProvider

// NewUseCase creates UseCase.
func NewUseCase(
	logger fw.Logger,
	timer fw.Timer,
	authenticator auth.Authenticator,
	githubIDProvider GithubIDProvider,
) UseCase {
	return UseCase{
		logger:           logger,
		timer:            timer,
		authenticator:    authenticator,
		githubIDProvider: githubIDProvider,
	}
}
