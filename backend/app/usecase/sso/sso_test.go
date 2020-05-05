// +build !integration all

package sso

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/account"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/external"
	"github.com/short-d/short/app/usecase/repository"
)

func TestSingleSignOn_SignIn(t *testing.T) {
	t.Parallel()

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
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			identityProvider := external.NewIdentityProviderFake("http://localhost/sign-in", "")
			profileService := external.NewSSOAccountFake(testCase.ssoUser)
			fakeUserRepo := repository.NewUserFake(testCase.users)
			tm := timer.NewStub(time.Now())
			accountProvider := account.NewProvider(&fakeUserRepo, tm)

			now := time.Now()
			auth := authenticator.NewAuthenticatorFake(now, time.Minute)

			singleSignOn := NewSingleSignOn(identityProvider, profileService, accountProvider, auth)
			gotAuthToken, err := singleSignOn.SignIn(testCase.authorizationCode)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			values := map[string]string{
				"email":     testCase.ssoUser.Email,
				"issued_at": now.Format(time.RFC3339Nano),
			}

			buf, err := json.Marshal(values)
			assert.Equal(t, nil, err)

			expAuthToken := string(buf)
			assert.Equal(t, expAuthToken, gotAuthToken)
		})
	}
}
