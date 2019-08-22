package signin

import (
	"encoding/json"
	"errors"
	"short/app/entity"
	"short/app/usecase/auth"
	"short/app/usecase/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOAuth_SignIn(t *testing.T) {
	testCases := []struct {
		name              string
		authorizationCode string
		userProfile       entity.UserProfile
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
			userProfile: entity.UserProfile{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			isAccountExist: true,
			hasErr:         false,
		},
		{
			name:              "account not exist",
			authorizationCode: "authorized",
			userProfile: entity.UserProfile{
				Email: "alpha@example.com",
				Name:  "Alpha",
			},
			isAccountExist: false,
			hasErr:         false,
		},
		{
			name:              "check account existence error",
			authorizationCode: "authorized",
			userProfile: entity.UserProfile{
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
			oauthService := service.NewOAuthFake("http://localhost/sign-in", "")
			profileService := service.NewProfileFake(testCase.userProfile)
			accountService := service.NewAccountFake(
				testCase.isAccountExist,
				testCase.isAccountExistErr,
				testCase.createAccountErr,
			)

			now := time.Now()
			authenticator := auth.NewAuthenticatorFake(now, time.Minute)

			oauth := NewOAuth(oauthService, profileService, accountService, authenticator)
			gotAuthToken, err := oauth.SignIn(testCase.authorizationCode)
			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}

			values := map[string]string{
				"email":     testCase.userProfile.Email,
				"issued_at": now.Format(time.RFC3339Nano),
			}

			buf, err := json.Marshal(values)
			assert.Nil(t, err)

			expAuthToken := string(buf)
			assert.Equal(t, expAuthToken, gotAuthToken)
		})
	}
}
