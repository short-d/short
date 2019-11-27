package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repository"
	"short/app/usecase/validator"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestURLCreatorPersist_CreateURL(t *testing.T) {
	now := time.Now()

	alias := "220uFicCJj"
	longAlias := "an-alias-cannot-be-used-to-specify-default-arguments"

	testCases := []struct {
		name            string
		urls            urlMap
		alias           *string
		availableKeys   []string
		user            entity.User
		url             entity.URL
		userURLRelation map[string]string
		isPublic        bool
		expHasErr       bool
		expectedURL     entity.URL
	}{
		{
			name: "alias exists",
			urls: urlMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias: &alias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			url:       entity.URL{},
			isPublic:  false,
			expHasErr: true,
		},
		{
			name: "alias too long",
			urls: urlMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias: &longAlias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			url: entity.URL{
				OriginalURL: "https://www.google.com",
			},
			expHasErr: true,
		},
		{
			name:  "create alias successfully",
			urls:  urlMap{},
			alias: &alias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "https://www.google.com",
				ExpireAt:    &now,
			},
			isPublic:  false,
			expHasErr: false,
			expectedURL: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "https://www.google.com",
				ExpireAt:    &now,
			},
		},
		{
			name: "automatically generate alias",
			urls: urlMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			availableKeys: []string{
				"test",
			},
			alias: nil,
			user: entity.User{
				Email: "alpha@example.com",
			},
			url: entity.URL{
				OriginalURL: "https://www.google.com",
			},
			expHasErr: false,
			expectedURL: entity.URL{
				Alias:       "test",
				OriginalURL: "https://www.google.com",
			},
		},
		{
			name: "no available key",
			urls: urlMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			availableKeys: []string{},
			alias:         nil,
			user: entity.User{
				Email: "alpha@example.com",
			},
			url: entity.URL{
				OriginalURL: "https://www.google.com",
			},
			expHasErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			urlRepo := repository.NewURLFake(testCase.urls)
			userURLRepo := repository.NewUserURLRepoFake(testCase.userURLRelation)
			keyGen := keygen.NewFake(testCase.availableKeys)
			longLinkValidator := validator.NewLongLink()
			aliasValidator := validator.NewCustomAlias()

			creator := NewCreatorPersist(
				&urlRepo,
				&userURLRepo,
				&keyGen,
				longLinkValidator,
				aliasValidator,
			)

			url, err := creator.CreateURL(testCase.url, testCase.alias, testCase.user, testCase.isPublic)
			if testCase.expHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, url)
		})
	}
}
