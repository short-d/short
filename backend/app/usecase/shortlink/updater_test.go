// +build !integration all

package shortlink

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

func TestShortLinkUpdaterPersist_UpdateShortLink(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	alias := "boGp9w35"
	validNewAlias := "eBJRJJty"
	validNewLongLink := "https://httpbin.org/get?p1=v1"

	testCases := []struct {
		name               string
		alias              string
		availableKeys      []keygen.Key
		shortlinks         shortLinks
		user               entity.User
		update    entity.ShortLink
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		expectedHasErr          bool
		expectedShortLink  entity.ShortLink
	}{
		{
			name:  "successfully update existing long link",
			alias: alias,
			shortlinks: shortLinks{
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
			shortLinkUpdate: entity.ShortLink{
				LongLink: validNewLongLink,
			},
			relationUsers: []entity.User{
				{ID: "1"},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:    "boGp9w35",
				LongLink: validNewLongLink,
			},
		},
		{
			name:  "alias doesn't exist",
			alias: validNewAlias,
			shortlinks: shortLinks{
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
			shortLinkUpdate: entity.ShortLink{
				LongLink: validNewLongLink,
			},
			relationUsers: []entity.User{
				{ID: "1"},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr:         true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "long link is invalid",
			alias: alias,
			shortlinks: shortLinks{
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
			shortLinkUpdate: entity.ShortLink{
				LongLink: "aaaaaaaaaaaaaaaaaaa",
			},
			relationUsers: []entity.User{
				{ID: "1"},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr:         true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "short link is not owned by the user",
			alias: alias,
			shortlinks: shortLinks{
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
			shortLinkUpdate: entity.ShortLink{
				LongLink: "https://google.com/",
			},
			relationUsers: []entity.User{
				{ID: "2"},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:     "boGp9w35",
					LongLink:  "https://httpbin.org",
					UpdatedAt: &now,
				},
			},
			expHasErr:         true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "reject malicious long link",
			alias: alias,
			user: entity.User{
				Email: "gopher@golang.org",
			},
			shortLinkUpdate: entity.ShortLink{
				LongLink: "http://malware.wicar.org/data/ms14_064_ole_not_xp.html",
			},
			expHasErr:         true,
			expectedShortLink: entity.ShortLink{},
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
			shortLinkRepo := repository.NewShortLinkFake(testCase.shortlinks)
			userShortLinkRepo := repository.NewUserShortLinkRepoFake(
				testCase.relationUsers,
				testCase.relationShortLinks,
			)
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			longLinkValidator := validator.NewLongLink()
			aliasValidator := validator.NewCustomAlias()
			blacklist := risk.NewBlackListFake(blockedHash)
			riskDetector := risk.NewDetector(blacklist)
			updater := NewUpdaterPersist(
				&shortLinkRepo,
				&userShortLinkRepo,
				keyGen,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)

			shortLink, err := updater.UpdateShortLink(testCase.alias, testCase.shortLinkUpdate, testCase.user)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)

				_, err = shortLinkRepo.GetShortLinkByAlias(testCase.expectedShortLink.Alias)
				assert.NotEqual(t, nil, err)

				isExist := userShortLinkRepo.IsRelationExist(testCase.user, testCase.expectedShortLink)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink.LongLink, shortLink.LongLink)
			assert.Equal(t, testCase.expectedShortLink.Alias, shortLink.Alias)
			assert.Equal(t, testCase.expectedShortLink.CreatedAt, shortLink.CreatedAt)
			assert.Equal(t, true, shortLink.UpdatedAt.After(now))
			assert.Equal(t, true, userShortLinkRepo.IsRelationExist(testCase.user, shortLink))
		})
	}
}
