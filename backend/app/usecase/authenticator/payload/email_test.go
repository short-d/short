package payload

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/short/backend/app/entity"
)

func TestEmailFactory_FromUser(t *testing.T) {
	testCases := []struct {
		name           string
		user           entity.User
		expectedHasErr bool
		expectedEmail  string
	}{
		{
			name: "user has email",
			user: entity.User{
				ID:    "alpha",
				Email: "alpha@example.com",
			},
			expectedHasErr: false,
			expectedEmail:  "alpha@example.com",
		},
		{
			name: "user email not found",
			user: entity.User{
				ID: "alpha",
			},
			expectedHasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			emailPayloadFactory := NewEmailFactory()
			emailPayload, err := emailPayloadFactory.FromUser(testCase.user)
			if testCase.expectedHasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
			tokenPayload := emailPayload.GetTokenPayload()

			assert.Equal(t, testCase.expectedEmail, tokenPayload["email"])
		})
	}
}

func TestEmailFactory_FromTokenPayload(t *testing.T) {
	testCases := []struct {
		name           string
		tokenPayload   crypto.TokenPayload
		expectedHasErr bool
		expectedEmail  string
	}{
		{
			name: "user has email",
			tokenPayload: crypto.TokenPayload{
				"id":    "alpha",
				"email": "alpha@example.com",
			},
			expectedHasErr: false,
			expectedEmail:  "alpha@example.com",
		},
		{
			name: "user email not found",
			tokenPayload: crypto.TokenPayload{
				"id": "alpha",
			},
			expectedHasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			emailPayloadFactory := NewEmailFactory()
			emailPayload, err := emailPayloadFactory.FromTokenPayload(testCase.tokenPayload)
			if testCase.expectedHasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
			gotUser := emailPayload.GetUser()
			assert.Equal(t, testCase.expectedEmail, gotUser.Email)
		})
	}
}
