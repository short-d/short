package account

import (
	"short/app/entity"
	"short/app/usecase/repo"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestProvider_IsAccountExist(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			fakeUserRepo := repo.NewUserFake(testCase.users)
			fakeTimer := mdtest.NewTimerFake(now)
			accountProvider := NewProvider(&fakeUserRepo, fakeTimer)
			gotIsExist, err := accountProvider.IsAccountExist(testCase.userEmail)
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedIsExist, gotIsExist)
		})
	}
}

func TestProvider_CreateAccount(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			fakeUserRepo := repo.NewUserFake(testCase.users)
			fakeTimer := mdtest.NewTimerFake(now)
			accountProvider := NewProvider(&fakeUserRepo, fakeTimer)
			err := accountProvider.CreateAccount(testCase.email, testCase.userName)
			if testCase.expectedHasErr {
				mdtest.NotEqual(t, nil, err)
				isEmailExist, err := fakeUserRepo.IsEmailExist(testCase.email)
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, true, isEmailExist)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}
