package payload

import (
	"testing"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
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
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			tokenPayload := emailPayload.GetTokenPayload()
			mdtest.Equal(t, testCase.expectedEmail, tokenPayload["email"])
		})
	}
}

func TestEmailFactory_FromTokenPayload(t *testing.T) {
	testCases := []struct {
		name           string
		tokenPayload   fw.TokenPayload
		expectedHasErr bool
		expectedEmail  string
	}{
		{
			name: "user has email",
			tokenPayload: fw.TokenPayload{
				"id":    "alpha",
				"email": "alpha@example.com",
			},
			expectedHasErr: false,
			expectedEmail:  "alpha@example.com",
		},
		{
			name: "user email not found",
			tokenPayload: fw.TokenPayload{
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
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			gotUser := emailPayload.GetUser()
			mdtest.Equal(t, testCase.expectedEmail, gotUser.Email)
		})
	}
}
