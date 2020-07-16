// +build !integration all

package shortlink

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/fw/ptr"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

func TestShortLinkCreatorPersist_CreateShortLink(t *testing.T) {
	t.Parallel()

	now := time.Now()
	utc := now.UTC()

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		availableKeys      []keygen.Key
		user               entity.User
		shortLinkArgs      entity.ShortLinkInput
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		blockedLongLinks   map[string]bool
		isPublic           bool
		// TODO(issue#803): Check error types in tests.
		expHasErr         bool
		expectedShortLink entity.ShortLink
	}{
		{
			name: "alias exists",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				LongLink:    ptr.String("https://www.google.com"),
				CustomAlias: ptr.String("220uFicCJj"),
			},
			isPublic:  false,
			expHasErr: true,
		},
		{
			name: "alias too long",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				LongLink:    ptr.String("https://www.google.com"),
				CustomAlias: ptr.String("an-alias-cannot-be-used-to-specify-default-arguments"),
			},
			expHasErr: true,
		},
		{
			name:       "alias contains invalid URL fragment character",
			shortLinks: shortLinks{},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				LongLink:    ptr.String("https://www.google.com"),
				CustomAlias: ptr.String("cant-have#chr"),
			},
			expHasErr: true,
		},
		{
			name:       "create alias successfully",
			shortLinks: shortLinks{},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				CustomAlias: ptr.String("220uFicCJj"),
				LongLink:    ptr.String("https://www.google.com"),
				ExpireAt:    &now,
			},
			isPublic:  false,
			expHasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "220uFicCJj",
				LongLink:  "https://www.google.com",
				ExpireAt:  &now,
				CreatedAt: &utc,
			},
		},
		{
			name:       "automatically generate alias if null alias provided",
			shortLinks: shortLinks{},
			availableKeys: []keygen.Key{
				"test",
			},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				LongLink: ptr.String("https://www.google.com"),
			},
			expHasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "test",
				LongLink:  "https://www.google.com",
				CreatedAt: &utc,
			},
		},
		{
			name:       "automatically generate alias if empty string alias provided",
			shortLinks: shortLinks{},
			availableKeys: []keygen.Key{
				"test",
			},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				LongLink:    ptr.String("https://www.google.com"),
				CustomAlias: ptr.String(""),
			},
			expHasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "test",
				LongLink:  "https://www.google.com",
				CreatedAt: &utc,
			},
		},
		{
			name:          "no available key",
			shortLinks:    shortLinks{},
			availableKeys: []keygen.Key{},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				LongLink:    ptr.String("https://www.google.com"),
				CustomAlias: ptr.String(""),
			},
			expHasErr: true,
		},
		{
			name:       "long link is invalid",
			shortLinks: shortLinks{},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				CustomAlias: ptr.String("220uFicCJj"),
				LongLink:    ptr.String("aaaaaaaaaaaaaaaaaaa"),
				ExpireAt:    &now,
			},
			isPublic:  false,
			expHasErr: true,
		},
		{
			name:       "reject malicious long link",
			shortLinks: shortLinks{},
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLinkArgs: entity.ShortLinkInput{
				CustomAlias: ptr.String("220uFicCJj"),
				LongLink:    ptr.String("http://malware.wicar.org/data/ms14_064_ole_not_xp.html"),
				ExpireAt:    &now,
			},
			blockedLongLinks: map[string]bool{
				"http://malware.wicar.org/data/ms14_064_ole_not_xp.html": true,
			},
			isPublic:  false,
			expHasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			blacklist := risk.NewBlackListFake(testCase.blockedLongLinks)
			shortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks, nil)
			userShortLinkRepo := repository.NewUserShortLinkRepoFake(
				testCase.relationUsers,
				testCase.relationShortLinks,
			)
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)
			longLinkValidator := validator.NewLongLink()
			aliasValidator := validator.NewCustomAlias()
			tm := timer.NewStub(now)
			riskDetector := risk.NewDetector(blacklist)

			creator := NewCreatorPersist(
				&shortLinkRepo,
				&userShortLinkRepo,
				keyGen,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)

			if testCase.shortLinkArgs.CustomAlias != nil {
				isExist, err := userShortLinkRepo.HasMapping(testCase.user, *testCase.shortLinkArgs.CustomAlias)
				assert.Equal(t, nil, err)
				assert.Equal(t, false, isExist)
			}

			shortLink, err := creator.CreateShortLink(testCase.shortLinkArgs, testCase.user, testCase.isPublic)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)

				_, err = shortLinkRepo.GetShortLinkByAlias(testCase.expectedShortLink.Alias)
				assert.NotEqual(t, nil, err)

				isExist, err := userShortLinkRepo.HasMapping(testCase.user, testCase.expectedShortLink.Alias)
				assert.Equal(t, nil, err)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink, shortLink)

			savedShortLink, err := shortLinkRepo.GetShortLinkByAlias(testCase.expectedShortLink.Alias)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink, savedShortLink)

			isExist, err := userShortLinkRepo.HasMapping(testCase.user, testCase.expectedShortLink.Alias)
			assert.Equal(t, nil, err)
			assert.Equal(t, true, isExist)
		})
	}
}
