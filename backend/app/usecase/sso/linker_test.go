// +build !integration all

package sso

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestLinker_IsAccountLinked(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		keys              []string
		users             []entity.User
		mappingUserIDs    []string
		mappingSSOUserIDs []string
		ssoUser           entity.SSOUser
		expectedIsLinked  bool
	}{
		{
			name:              "account not linked",
			keys:              []string{},
			mappingUserIDs:    []string{},
			mappingSSOUserIDs: []string{},
			ssoUser: entity.SSOUser{
				ID:    "alpha",
				Email: "alpha@example.com",
				Name:  "Alpha User",
			},
			expectedIsLinked: false,
		},
		{
			name: "account already linked",
			keys: []string{},
			mappingUserIDs: []string{
				"beta",
			},
			mappingSSOUserIDs: []string{
				"alpha",
			},
			ssoUser: entity.SSOUser{
				ID:    "alpha",
				Email: "alpha@example.com",
				Name:  "Alpha User",
			},
			expectedIsLinked: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			userRepo := repository.NewUserFake([]entity.User{})
			linkerFactory := NewAccountLinkerFactory(keyGen, &userRepo)
			ssoMap, err := repository.NewsSSOMapFake(testCase.mappingSSOUserIDs, testCase.mappingUserIDs)
			assert.Equal(t, nil, err)

			linker := linkerFactory.NewAccountLinker(&ssoMap)
			isLinked, err := linker.IsAccountLinked(testCase.ssoUser)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedIsLinked, isLinked)
		})
	}
}

func TestLinker_LinkAccount(t *testing.T) {
	testCases := []struct {
		name              string
		key               string
		mappingUserIDs    []string
		mappingSSOUserIDs []string
		users             []entity.User
		ssoUser           entity.SSOUser
		user              entity.User
		expectedIDExist   bool
	}{
		{
			name:              "account exists not linked",
			key:               "alpha",
			mappingUserIDs:    []string{},
			mappingSSOUserIDs: []string{},
			users: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			ssoUser: entity.SSOUser{
				ID:    "gama",
				Email: "alpha@example.com",
			},
			user: entity.User{
				ID:    "alpha",
				Email: "alpha@example.com",
			},
			expectedIDExist: false,
		},
		{
			name:              "create new account",
			key:               "alpha",
			mappingUserIDs:    []string{},
			mappingSSOUserIDs: []string{},
			users:             []entity.User{},
			ssoUser: entity.SSOUser{
				ID:    "gama",
				Email: "alpha@example.com",
			},
			user: entity.User{
				ID:    "alpha",
				Email: "alpha@example.com",
			},
			expectedIDExist: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{
				keygen.Key(testCase.key),
			})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			userRepo := repository.NewUserFake(testCase.users)
			linkerFactory := NewAccountLinkerFactory(keyGen, &userRepo)
			ssoMap, err := repository.NewsSSOMapFake(testCase.mappingSSOUserIDs, testCase.mappingUserIDs)
			assert.Equal(t, nil, err)

			linker := linkerFactory.NewAccountLinker(&ssoMap)

			gotIsRelationExist := ssoMap.IsRelationExist(testCase.ssoUser.ID, testCase.user.ID)
			assert.Equal(t, testCase.expectedIDExist, gotIsRelationExist)

			err = linker.CreateAndLinkAccount(testCase.ssoUser)
			assert.Equal(t, nil, err)

			gotIsRelationExist = ssoMap.IsRelationExist(testCase.ssoUser.ID, testCase.user.ID)
			assert.Equal(t, true, gotIsRelationExist)

			gotIsIDExist := userRepo.IsUserIDExist(testCase.user.ID)
			assert.Equal(t, true, gotIsIDExist)
		})
	}
}
