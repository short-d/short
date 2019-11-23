package sso

import (
	"encoding/json"
	"short/app/entity"
	"short/app/usecase/account"
	"short/app/usecase/auth"
	"short/app/usecase/repository"
	"short/app/usecase/service"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestSingleSignOn_SignIn(t *testing.T) {
	testCases := []struct {
		name              string
		authorizationCode string
		ssoUser           entity.SSOUser
		users             []entity.User
		hasErr            bool
	}{
		{
			name:              "empty authorization code",
			authorizationCode: "",
			hasErr:            true,
		},
		{
			name:              "account found",
			authorizationCode: "authorized",
			ssoUser: entity.SSOUser{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			users: []entity.User{
				{Email: "alpha@example.com"},
			},
			hasErr: false,
		},
		{
			name:              "account not exist",
			authorizationCode: "authorized",
			ssoUser: entity.SSOUser{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			users:  []entity.User{},
			hasErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			identityProvider := service.NewIdentityProviderFake("http://localhost/sign-in", "")
			profileService := service.NewSSOAccountFake(testCase.ssoUser)
			fakeUserRepo := repository.NewUserFake(testCase.users)
			fakeTimer := mdtest.NewTimerFake(time.Now())
			accountProvider := account.NewProvider(&fakeUserRepo, fakeTimer)

			now := time.Now()
			authenticator := auth.NewAuthenticatorFake(now, time.Minute)

			singleSignOn := NewSingleSignOn(identityProvider, profileService, accountProvider, authenticator)
			gotAuthToken, err := singleSignOn.SignIn(testCase.authorizationCode)
			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}

			values := map[string]string{
				"email":     testCase.ssoUser.Email,
				"issued_at": now.Format(time.RFC3339Nano),
			}

			buf, err := json.Marshal(values)
			mdtest.Equal(t, nil, err)

			expAuthToken := string(buf)
			mdtest.Equal(t, expAuthToken, gotAuthToken)
		})
	}
}
