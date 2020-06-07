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

func TestShortLinkCreatorPersist_CreateShortLink(t *testing.T) {
	t.Parallel()

	now := time.Now()
	utc := now.UTC()

	alias := "220uFicCJj"
	longAlias := "an-alias-cannot-be-used-to-specify-default-arguments"
	invalidFragmentAlias := "cant-have#chr"
	emptyAlias := ""

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		alias              *string
		availableKeys      []keygen.Key
		user               entity.User
		shortLink          entity.ShortLink
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		isPublic           bool
		expHasErr          bool
		expectedShortLink  entity.ShortLink
	}{
		{
			name: "alias exists",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias: &alias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{},
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
			alias: &longAlias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{
				LongLink: "https://www.google.com",
			},
			expHasErr: true,
		},
		{
			name:       "alias contains invalid URL fragment character",
			shortLinks: shortLinks{},
			alias:      &invalidFragmentAlias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{
				LongLink: "https://www.google.com",
			},
			expHasErr: true,
		},
		{
			name:       "create alias successfully",
			shortLinks: shortLinks{},
			alias:      &alias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "https://www.google.com",
				ExpireAt: &now,
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
			alias: nil,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{
				LongLink: "https://www.google.com",
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
			alias: &emptyAlias,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{
				LongLink: "https://www.google.com",
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
			alias:         nil,
			user: entity.User{
				Email: "alpha@example.com",
			},
			shortLink: entity.ShortLink{
				LongLink: "https://www.google.com",
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
			shortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks)
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

			_, err = shortLinkRepo.GetShortLinkByAlias(testCase.shortLink.Alias)
			assert.NotEqual(t, nil, err)

			isExist := userShortLinkRepo.IsRelationExist(testCase.user, testCase.shortLink)
			assert.Equal(t, false, isExist)

			shortLink, err := creator.CreateShortLink(testCase.shortLink, testCase.alias, testCase.user, testCase.isPublic)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)

				_, err = shortLinkRepo.GetShortLinkByAlias(testCase.expectedShortLink.Alias)
				assert.NotEqual(t, nil, err)

				isExist := userShortLinkRepo.IsRelationExist(testCase.user, testCase.expectedShortLink)
				assert.Equal(t, false, isExist)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink, shortLink)

			savedShortLink, err := shortLinkRepo.GetShortLinkByAlias(testCase.expectedShortLink.Alias)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink, savedShortLink)

			isExist = userShortLinkRepo.IsRelationExist(testCase.user, testCase.expectedShortLink)
			assert.Equal(t, true, isExist)
		})
	}
}
