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

func TestURLUpdaterPersist_UpdateURL(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	alias := "boGp9w35"
	validNewAlias := "eBJRJJty"
	validNewLongLink := "https://httpbin.org/get?p1=v1"

	testCases := []struct {
		name          string
		alias         *string
		availableKeys []keygen.Key
		urls          urlMap
		user          entity.User
		urlUpdate     entity.URL
		relationUsers []entity.User
		relationURLs  []entity.URL
		expHasErr     bool
		expectedURL   entity.URL
	}{
		{
			name:  "successfully update existing long link",
			alias: &alias,
			urls: urlMap{
				"boGp9w35": entity.URL{
					Alias:       "boGp9w35",
					OriginalURL: "https://httpbin.org",
				},
			},
			user: entity.User{
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.URL{
				OriginalURL: validNewLongLink,
			},
			expHasErr: false,
			expectedURL: entity.URL{
				Alias:       "boGp9w35",
				OriginalURL: validNewLongLink,
			},
		},
		{
			name:  "alias doesn't exist",
			alias: &validNewAlias,
			urls: urlMap{
				"boGp9w35zzzz": entity.URL{
					Alias:       "boGp9w35zzzz",
					OriginalURL: "https://httpbin.org",
				},
			},
			user: entity.User{
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.URL{
				OriginalURL: validNewLongLink,
			},
			expHasErr:   true,
			expectedURL: entity.URL{},
		},
		{
			name:  "long link is invalid",
			alias: &alias,
			urls: urlMap{
				"boGp9w35": entity.URL{
					Alias:       "boGp9w35",
					OriginalURL: "https://httpbin.org",
				},
			},
			user: entity.User{
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.URL{
				OriginalURL: "aaaaaaaaaaaaaaaaaaa",
			},
			expHasErr:   true,
			expectedURL: entity.URL{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			blockedHash := map[string]bool{}
			tm := timer.NewStub(now)
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
			blacklist := risk.NewBlackListFake(blockedHash)
			riskDetector := risk.NewDetector(blacklist)
			updater := NewUpdaterPersist(
				&urlRepo,
				&userURLRepo,
				keyGen,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)

			url, err := updater.UpdateURL(*testCase.alias, testCase.urlUpdate, testCase.user)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)

				_, err = urlRepo.GetByAlias(testCase.expectedURL.Alias)
				assert.NotEqual(t, nil, err)

				isExist := userURLRepo.IsRelationExist(testCase.user, testCase.expectedURL)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedURL.OriginalURL, url.OriginalURL)
			assert.Equal(t, testCase.expectedURL.Alias, url.Alias)
			assert.Equal(t, testCase.expectedURL.CreatedAt, url.CreatedAt)
			assert.NotEqual(t, testCase.expectedURL.UpdatedAt, url.UpdatedAt)
		})
	}
}
