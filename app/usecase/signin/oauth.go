package signin

import (
	"errors"
	"log"
	"short/app/usecase/auth"
	"short/app/usecase/service"
)

type OAuth struct {
	oauthService   service.OAuth
	profileService service.Profile
	accountService service.Account
	authenticator  auth.Authenticator
}

func (o OAuth) SignIn(authorizationCode string) (string, error) {
	if len(authorizationCode) < 1 {
		return "", errors.New("authorizationCode can't be empty")
	}

	accessToken, err := o.oauthService.RequestAccessToken(authorizationCode)
	log.Print("accessToken: ", accessToken)
	if err != nil {
		return "", err
	}

	userProfile, err := o.profileService.GetUserProfile(accessToken)
	log.Print("userProfile: ", userProfile)
	if err != nil {
		return "", err
	}

	email := userProfile.Email
	isExist, err := o.accountService.IsAccountExist(email)
	log.Print("isExist: ", isExist)
	if err != nil {
		return "", err
	}

	authToken, err := o.authenticator.GenerateToken(email)
	log.Print("authToken: ", authToken)
	if err != nil {
		return "", err
	}

	if isExist {
		return authToken, nil
	}

	err = o.accountService.CreateAccount(email, userProfile.Name)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

func NewOAuth(
	oauthService service.OAuth,
	profileService service.Profile,
	accountService service.Account,
	authenticator auth.Authenticator,
) OAuth {
	return OAuth{
		oauthService:   oauthService,
		profileService: profileService,
		accountService: accountService,
		authenticator:  authenticator,
	}
}
