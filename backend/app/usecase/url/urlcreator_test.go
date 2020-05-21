// +build !integration all

package url

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

func TestURLCreatorPersist_CreateURL(t *testing.T) {
	t.Parallel()

	now := time.Now()
	utc := now.UTC()

	alias := "220uFicCJj"
	longAlias := "an-alias-cannot-be-used-to-specify-default-arguments"
	emptyAlias := ""

	testCases := []struct {
		name          string
		urls          urlMap
		alias         *string
		availableKeys []keygen.Key
		user          entity.User
		url           entity.URL
		relationUsers []entity.User
		relationURLs  []entity.URL
		isPublic      bool
		// TODO(issue#803): Check error types in tests.
		expHasErr   bool
		expectedURL entity.URL
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
			url: entity.URL{
				OriginalURL: "https://www.google.com",
			},
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
				CreatedAt:   &utc,
			},
		},
		{
			name: "automatically generate alias if null alias provided",
			urls: urlMap{},
			availableKeys: []keygen.Key{
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
				CreatedAt:   &utc,
			},
		},
		{
			name: "automatically generate alias if empty string alias provided",
			urls: urlMap{},
			availableKeys: []keygen.Key{
				"test",
			},
			alias: &emptyAlias,
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
				CreatedAt:   &utc,
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
			availableKeys: []keygen.Key{},
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
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			blockedURLs := map[string]bool{}
			blacklist := risk.NewBlackListFake(blockedURLs)
			urlRepo := repository.NewURLFake(testCase.urls)
			userURLRepo := repository.NewUserURLRepoFake(
				testCase.relationUsers,
				testCase.relationURLs,
			)
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)
			longLinkValidator := validator.NewLongLink()
			aliasValidator := validator.NewCustomAlias()
			tm := timer.NewStub(now)
			riskDetector := risk.NewDetector(blacklist)

			creator := NewCreatorPersist(
				&urlRepo,
				&userURLRepo,
				keyGen,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)

			_, err = urlRepo.GetByAlias(testCase.url.Alias)
			assert.NotEqual(t, nil, err)

			isExist := userURLRepo.IsRelationExist(testCase.user, testCase.url)
			assert.Equal(t, false, isExist)

			url, err := creator.CreateURL(testCase.url, testCase.alias, testCase.user, testCase.isPublic)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)

				_, err = urlRepo.GetByAlias(testCase.expectedURL.Alias)
				assert.NotEqual(t, nil, err)

				isExist := userURLRepo.IsRelationExist(testCase.user, testCase.expectedURL)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedURL, url)

			savedURL, err := urlRepo.GetByAlias(testCase.expectedURL.Alias)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedURL, savedURL)

			isExist = userURLRepo.IsRelationExist(testCase.user, testCase.expectedURL)
			assert.Equal(t, true, isExist)
		})
	}
}
