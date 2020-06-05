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
		alias         string
		availableKeys []keygen.Key
		urls          urlMap
		user          entity.User
		urlUpdate     entity.ShortLink
		relationUsers []entity.User
		relationURLs  []entity.ShortLink
		expHasErr     bool
		expectedURL   entity.ShortLink
	}{
		{
			name:  "successfully update existing long link",
			alias: alias,
			urls: urlMap{
				"boGp9w35": entity.ShortLink{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.ShortLink{
				LongLink: validNewLongLink,
			},
			relationUsers: []entity.User{
				{ID: "1"},
			},
			relationURLs: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr: false,
			expectedURL: entity.ShortLink{
				Alias:    "boGp9w35",
				LongLink: validNewLongLink,
			},
		},
		{
			name:  "alias doesn't exist",
			alias: validNewAlias,
			urls: urlMap{
				"boGp9w35zzzz": entity.ShortLink{
					Alias:     "boGp9w35zzzz",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.ShortLink{
				LongLink: validNewLongLink,
			},
			relationUsers: []entity.User{
				{ID: "1"},
			},
			relationURLs: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr:   true,
			expectedURL: entity.ShortLink{},
		},
		{
			name:  "long link is invalid",
			alias: alias,
			urls: urlMap{
				"boGp9w35": entity.ShortLink{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.ShortLink{
				LongLink: "aaaaaaaaaaaaaaaaaaa",
			},
			relationUsers: []entity.User{
				{ID: "1"},
			},
			relationURLs: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr:   true,
			expectedURL: entity.ShortLink{},
		},
		{
			name:  "short link is not owned by the user",
			alias: alias,
			urls: urlMap{
				"boGp9w35": entity.ShortLink{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.ShortLink{
				LongLink: "https://google.com/",
			},
			relationUsers: []entity.User{
				{ID: "2"},
			},
			relationURLs: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr:   true,
			expectedURL: entity.ShortLink{},
		},
		{
			name:  "malicious url update",
			alias: alias,
			user: entity.User{
				Email: "gopher@golang.org",
			},
			urlUpdate: entity.ShortLink{
				LongLink: "http://malware.wicar.org/data/ms14_064_ole_not_xp.html",
			},
			expHasErr:   true,
			expectedURL: entity.ShortLink{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			blockedHash := map[string]bool{
				"http://malware.wicar.org/data/ms14_064_ole_not_xp.html": false,
			}
			tm := timer.NewStub(now)
			urlRepo := repository.NewShortLinkFake(testCase.urls)
			userURLRepo := repository.NewUserShortLinkRepoFake(
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

			url, err := updater.UpdateURL(testCase.alias, testCase.urlUpdate, testCase.user)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)

				_, err = urlRepo.GetShortLinkByAlias(testCase.expectedURL.Alias)
				assert.NotEqual(t, nil, err)

				isExist := userURLRepo.IsRelationExist(testCase.user, testCase.expectedURL)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedURL.LongLink, url.LongLink)
			assert.Equal(t, testCase.expectedURL.Alias, url.Alias)
			assert.Equal(t, testCase.expectedURL.CreatedAt, url.CreatedAt)
			assert.Equal(t, true, url.UpdatedAt.After(now))
			assert.Equal(t, true, userURLRepo.IsRelationExist(testCase.user, url))
		})
	}
}
