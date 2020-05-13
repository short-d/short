package sso

import (
	"errors"

	"github.com/short-d/short/backend/app/usecase/authenticator"
)

// SingleSignOn enables sign in through external identity providers, such as
// Github, Facebook, and Google.
type SingleSignOn struct {
	identityProvider IdentityProvider
	account          Account
	accountLinker    AccountLinker
	authenticator    authenticator.Authenticator
}

// SignIn generates access token for a user using authorization code obtained
// from external identity provider.
func (o SingleSignOn) SignIn(authorizationCode string) (string, error) {
	if len(authorizationCode) < 1 {
		return "", errors.New("authorizationCode can't be empty")
	}

	accessToken, err := o.identityProvider.RequestAccessToken(authorizationCode)
	if err != nil {
		return "", err
	}

	ssoUser, err := o.account.GetSingleSignOnUser(accessToken)
	if err != nil {
		return "", err
	}

	isLinked, err := o.accountLinker.IsAccountLinked(ssoUser)
	if err != nil {
		return "", err
	}

	if !isLinked {
		err = o.accountLinker.CreateAndLinkAccount(ssoUser)
		if err != nil {
			return "", err
		}
	}

	user, err := o.accountLinker.GetShortUser(ssoUser)
	if err != nil {
		return "", err
	}
	return o.authenticator.GenerateToken(user)
}

func (o SingleSignOn) IsSignedIn(authToken string) bool {
	return o.authenticator.IsSignedIn(authToken)
}

func (o SingleSignOn) GetSignInLink() string {
	return o.identityProvider.GetAuthorizationURL()
}

type Factory struct {
	authenticator authenticator.Authenticator
}

func (s Factory) NewSingleSignOn(
	identityProvider IdentityProvider,
	account Account,
	accountLinker AccountLinker,
) SingleSignOn {
	return SingleSignOn{
		identityProvider: identityProvider,
		account:          account,
		accountLinker:    accountLinker,
		authenticator:    s.authenticator,
	}
}

func NewFactory(
	authenticator authenticator.Authenticator,
) Factory {
	return Factory{
		authenticator: authenticator,
	}
}
