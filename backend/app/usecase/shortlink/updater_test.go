// +build !integration all

package shortlink

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

func TestShortLinkUpdaterPersist_UpdateShortLink(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()

	// ShortLink creator and updater methods require string pointers
	// It is not possible to reference a literal string directly in Go
	// To make the test suite as clean as possible, they are instead declared below and
	// the variables are passed into the test case by reference, which is a legal semantic in Go.
	// https://golang.org/ref/spec#Address_operators

	// Set of aliases representing different classes of aliases to be tested.
	alias := "short-d"
	fragmentAlias := "#http-bin"

	// Set of long links representing different classes of long links to be tested.
	longLink := "https://httpbin.org/get?p1=v1"
	invalidLongLink := "aaaaaaaaaaaaaaaaaaa"
	maliciousLongLink := "http://malware.wicar.org/data/ms14_064_ole_not_xp.html"

	testCases := []struct {
		name               string
		alias              string
		shortlinks         shortLinks
		user               entity.User
		updateArgs         entity.ShortLinkInput
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		blockedLongLinks   map[string]bool
		expectedHasErr     bool
		expectedShortLink  entity.ShortLink
	}{
		{
			name:  "successfully update existing long link",
			alias: "boGp9w35",
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
			updateArgs: entity.ShortLinkInput{
				LongLink: &longLink,
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
			expectedHasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:    "boGp9w35",
				LongLink: "https://httpbin.org/get?p1=v1",
			},
		},
		{
			name:  "alias doesn't exist",
			alias: "eBJRJJty",
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
			updateArgs: entity.ShortLinkInput{
				LongLink: &longLink,
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
			expectedHasErr:    true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "alias already exist",
			alias: "git",
			shortlinks: shortLinks{
				"git": entity.ShortLink{
					Alias:     "git",
					LongLink:  "https://github.com/short-d",
					UpdatedAt: &now,
				},
				"short-d": entity.ShortLink{
					Alias:     "short-d",
					LongLink:  "http://short-d.com/",
					UpdatedAt: &now,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "gopher@golang.org",
			},
			updateArgs: entity.ShortLinkInput{
				CustomAlias: &alias,
			},
			relationUsers: []entity.User{
				{ID: "1"},
				{ID: "1"},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:     "git",
					LongLink:  "https://github.com/short-d",
					UpdatedAt: &now,
				},
				{
					Alias:     "short-d",
					LongLink:  "http://short-d.com/",
					UpdatedAt: &now,
				},
			},
			expectedHasErr: true,
		},
		{
			name:  "long link is invalid",
			alias: "boGp9w35",
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
			updateArgs: entity.ShortLinkInput{
				LongLink: &invalidLongLink,
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
			expectedHasErr:    true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "alias contains hash tag",
			alias: "boGp9w35",
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
			updateArgs: entity.ShortLinkInput{
				CustomAlias: &fragmentAlias,
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
			expectedHasErr:    true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "short link is not owned by the user",
			alias: "boGp9w35",
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
			updateArgs: entity.ShortLinkInput{
				LongLink: &longLink,
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
			expectedHasErr:    true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:  "reject malicious long link",
			alias: "boGp9w35",
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
			updateArgs: entity.ShortLinkInput{
				LongLink: &maliciousLongLink,
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
			blockedLongLinks: map[string]bool{
				maliciousLongLink: true,
			},
			expectedHasErr:    true,
			expectedShortLink: entity.ShortLink{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			tm := timer.NewStub(now)
			shortLinkRepo := repository.NewShortLinkFake(testCase.shortlinks)
			userShortLinkRepo := repository.NewUserShortLinkRepoFake(
				testCase.relationUsers,
				testCase.relationShortLinks,
			)

			longLinkValidator := validator.NewLongLink()
			aliasValidator := validator.NewCustomAlias()
			blacklist := risk.NewBlackListFake(testCase.blockedLongLinks)
			riskDetector := risk.NewDetector(blacklist)
			updater := NewUpdaterPersist(
				&shortLinkRepo,
				&userShortLinkRepo,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)

			shortLink, err := updater.UpdateShortLink(testCase.alias, testCase.updateArgs, testCase.user)
			if testCase.expectedHasErr {
				assert.NotEqual(t, nil, err)

				_, err = shortLinkRepo.GetShortLinkByAlias(testCase.expectedShortLink.Alias)
				assert.NotEqual(t, nil, err)

				isExist, err := userShortLinkRepo.HasMapping(testCase.user, testCase.expectedShortLink.Alias)
				assert.Equal(t, nil, err)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink.LongLink, shortLink.LongLink)
			assert.Equal(t, testCase.expectedShortLink.Alias, shortLink.Alias)
			assert.Equal(t, testCase.expectedShortLink.CreatedAt, shortLink.CreatedAt)
			assert.Equal(t, true, shortLink.UpdatedAt.After(now))
			isExist, err := userShortLinkRepo.HasMapping(testCase.user, shortLink.Alias)
			assert.Equal(t, nil, err)
			assert.Equal(t, true, isExist)
		})
	}
}
