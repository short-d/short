package sso

import (
	"encoding/json"
	"errors"
	"short/app/entity"
	"short/app/usecase/auth"
	"short/app/usecase/service"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestSingleSignOn_SignIn(t *testing.T) {
	testCases := []struct {
		name              string
		authorizationCode string
		ssoUser       entity.SSOUser
		isAccountExist    bool
		isAccountExistErr error
		createAccountErr  error
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
			isAccountExist: true,
			hasErr:         false,
		},
		{
			name:              "account not exist",
			authorizationCode: "authorized",
			ssoUser: entity.SSOUser{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			isAccountExist: false,
			hasErr:         false,
		},
		{
			name:              "check account existence error",
			authorizationCode: "authorized",
			ssoUser: entity.SSOUser{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			isAccountExist:    false,
			isAccountExistErr: errors.New("error"),
			hasErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			identityProvider := service.NewIdentityProviderFake("http://localhost/sign-in", "")
			profileService := service.NewSSOAccountFake(testCase.ssoUser)
			accountService := service.NewAccountFake(
				testCase.isAccountExist,
				testCase.isAccountExistErr,
				testCase.createAccountErr,
			)

			now := time.Now()
			authenticator := auth.NewAuthenticatorFake(now, time.Minute)

			singleSignOn := NewSingleSignOn(identityProvider, profileService, accountService, authenticator)
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
