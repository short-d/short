// +build !integration all

package account

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestProvider_IsAccountExist(t *testing.T) {
	t.Parallel()

	now := time.Now()

	testCases := []struct {
		name            string
		users           []entity.User
		userEmail       string
		expectedIsExist bool
	}{
		{
			name: "account exists",
			users: []entity.User{
				{Email: "alpha@example.com"},
			},
			userEmail:       "alpha@example.com",
			expectedIsExist: true,
		},
		{
			name:            "account not found",
			users:           []entity.User{},
			userEmail:       "alpha@example.com",
			expectedIsExist: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			fakeUserRepo := repository.NewUserFake(testCase.users)
			tm := timer.NewStub(now)
			accountProvider := NewProvider(&fakeUserRepo, tm)
			gotIsExist, err := accountProvider.IsAccountExist(testCase.userEmail)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedIsExist, gotIsExist)
		})
	}
}

func TestProvider_CreateAccount(t *testing.T) {
	t.Parallel()

	now := time.Now()

	testCases := []struct {
		name           string
		users          []entity.User
		email          string
		userName       string
		expectedHasErr bool
	}{
		{
			name:           "successfully created account",
			users:          []entity.User{},
			email:          "alpha@example.com",
			userName:       "Alpha",
			expectedHasErr: false,
		},
		{
			name: "account exists",
			users: []entity.User{
				{Email: "alpha@example.com"},
			},
			email:          "alpha@example.com",
			userName:       "Alpha",
			expectedHasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			fakeUserRepo := repository.NewUserFake(testCase.users)
			tm := timer.NewStub(now)
			accountProvider := NewProvider(&fakeUserRepo, tm)
			err := accountProvider.CreateAccount(testCase.email, testCase.userName)
			if testCase.expectedHasErr {
				assert.NotEqual(t, nil, err)
				isEmailExist, err := fakeUserRepo.IsEmailExist(testCase.email)
				assert.Equal(t, nil, err)
				assert.Equal(t, true, isEmailExist)
				return
			}
			assert.Equal(t, nil, err)
		})
	}
}
