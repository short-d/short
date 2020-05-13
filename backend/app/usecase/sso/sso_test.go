// +build !integration all

package sso

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestSingleSignOn_SignIn(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		authorizationCode string
		profileSSOUser    entity.SSOUser
		usersMapping      []entity.User
		ssoUsersMapping   []entity.SSOUser
		user              entity.User
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
			profileSSOUser: entity.SSOUser{
				ID:    "random_sso_id",
				Email: "alpha@example.com",
			},
			usersMapping: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			ssoUsersMapping: []entity.SSOUser{
				{
					ID:    "random_sso_id",
					Email: "alpha@example.com",
				},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: false,
		},
		{
			name:              "account not exist",
			authorizationCode: "authorized",
			profileSSOUser: entity.SSOUser{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			usersMapping:    []entity.User{},
			ssoUsersMapping: []entity.SSOUser{},
			hasErr:          true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			identityProvider := NewIdentityProviderFake("http://localhost/sign-in", "")
			profileService := NewAccountFake(testCase.profileSSOUser)

			now := time.Now()
			auth := authenticator.NewAuthenticatorFake(now, time.Minute)

			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			userRepo := repository.NewUserFake(testCase.usersMapping)
			linkerFactory := NewAccountLinkerFactory(keyGen, &userRepo)
			ssoMap, err := repository.NewsSSOMapFake(testCase.ssoUsersMapping, testCase.usersMapping)
			assert.Equal(t, nil, err)

			linker := linkerFactory.NewAccountLinker(&ssoMap)
			factory := NewFactory(auth)

			singleSignOn := factory.NewSingleSignOn(identityProvider, profileService, linker)
			gotAuthToken, err := singleSignOn.SignIn(testCase.authorizationCode)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			values := map[string]string{
				"id":        testCase.user.ID,
				"issued_at": now.Format(time.RFC3339Nano),
			}

			buf, err := json.Marshal(values)
			assert.Equal(t, nil, err)

			expAuthToken := string(buf)
			assert.Equal(t, expAuthToken, gotAuthToken)
		})
	}
}
