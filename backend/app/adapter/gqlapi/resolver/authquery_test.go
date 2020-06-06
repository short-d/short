// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

type shortLinkMap = map[string]entity.ShortLink

func TestAuthQuery_ShortLink(t *testing.T) {
	t.Parallel()
	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name              string
		user              entity.User
		alias             string
		expireAfter       *scalar.Time
		shortLinks        shortLinkMap
		hasErr            bool
		expectedShortLink *ShortLink
	}{
		{
			name:        "alias not found with no expireAfter",
			alias:       "220uFicCJj",
			expireAfter: nil,
			shortLinks:  shortLinkMap{},
			hasErr:      true,
		},
		{
			name:  "alias not found with expireAfter",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			shortLinks: shortLinkMap{},
			hasErr:     true,
		},
		{
			name:  "alias expired",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			shortLinks: shortLinkMap{
				"220uFicCJj": entity.ShortLink{
					ExpireAt: &before,
				},
			},
			hasErr: true,
		},
		{
			name:  "shortlink found",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			shortLinks: shortLinkMap{
				"220uFicCJj": entity.ShortLink{
					ExpireAt: &after,
				},
			},
			hasErr: false,
			expectedShortLink: &ShortLink{
				shortLink: entity.ShortLink{
					ExpireAt: &after,
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			fakeShortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks)
			fakeUserShortLinkRepo := repository.NewUserShortLinkRepoFake(nil, nil)
			retrieverFake := shortlink.NewRetrieverPersist(&fakeShortLinkRepo, &fakeUserShortLinkRepo)

			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			timerFake := timer.NewStub(now)
			changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			changeLog := changelog.NewPersist(keyGen, timerFake, &changeLogRepo, &userChangeLogRepo)

			tokenizer := crypto.NewTokenizerFake()
			auth := authenticator.NewAuthenticator(tokenizer, timerFake, time.Hour)

			authToken, err := auth.GenerateToken(testCase.user)
			assert.Equal(t, nil, err)

			query := newAuthQuery(&authToken, auth, changeLog, retrieverFake)

			shortLinkArgs := &ShortLinkArgs{
				Alias:       testCase.alias,
				ExpireAfter: testCase.expireAfter,
			}

			s, err := query.ShortLink(shortLinkArgs)

			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, testCase.expectedShortLink, s)
		})
	}
}
