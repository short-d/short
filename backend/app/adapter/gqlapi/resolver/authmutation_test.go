// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/url"
	"github.com/short-d/short/backend/app/usecase/validator"
)

func TestUpdateURL(t *testing.T) {
	now := time.Now().UTC()
	newAlias := "myNewAlias"
	newLongLink := "https://www.short-d.com"
	urls := urlMap{
		"mySimpleAlias": entity.ShortLink{
			Alias:    "mySimpleAlias",
			LongLink: "https://www.google.com/",
		},
	}
	testCases := []struct {
		name               string
		args               *UpdateURLArgs
		user               entity.User
		urls               urlMap
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		expectedShortLink  *URL
		shouldError        bool
	}{
		{
			name: "empty update returns empty response",
			args: &UpdateURLArgs{
				OldAlias: "mySimpleAlias",
				URL:      URLInput{},
			},
			user: entity.User{
				ID:    "1",
				Email: "short@gmail.com",
			},
			urls: urls,
			relationUsers: []entity.User{
				entity.User{
					ID:    "1",
					Email: "short@gmail.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				entity.ShortLink{
					Alias:    "mySimpleAlias",
					LongLink: "https://www.google.com/",
				},
			},
			expectedShortLink: nil,
			shouldError:       false,
		},
		{
			name: "update short link alias",
			args: &UpdateURLArgs{
				OldAlias: "mySimpleAlias",
				URL: URLInput{
					CustomAlias: &newAlias,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "short@gmail.com",
			},
			urls: urls,
			relationUsers: []entity.User{
				entity.User{
					ID:    "1",
					Email: "short@gmail.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				entity.ShortLink{
					Alias:    "mySimpleAlias",
					LongLink: "https://www.google.com/",
				},
			},
			expectedShortLink: &URL{
				url: entity.ShortLink{
					Alias:    newAlias,
					LongLink: "https://www.google.com/",
				},
			},
			shouldError: false,
		},
		{
			name: "update only long link",
			args: &UpdateURLArgs{
				OldAlias: "mySimpleAlias",
				URL: URLInput{
					OriginalURL: &newLongLink,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "short@gmail.com",
			},
			urls: urls,
			relationUsers: []entity.User{
				entity.User{
					ID:    "1",
					Email: "short@gmail.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				entity.ShortLink{
					Alias:    "mySimpleAlias",
					LongLink: "https://www.google.com/",
				},
			},
			expectedShortLink: &URL{
				url: entity.ShortLink{
					Alias:    "mySimpleAlias",
					LongLink: newLongLink,
				},
			},
			shouldError: false,
		},
		{
			name: "update both alias and long link",
			args: &UpdateURLArgs{
				OldAlias: "mySimpleAlias",
				URL: URLInput{
					CustomAlias: &newAlias,
					OriginalURL: &newLongLink,
				},
			},
			user: entity.User{
				ID:    "1",
				Email: "short@gmail.com",
			},
			urls: urls,
			relationUsers: []entity.User{
				entity.User{
					ID:    "1",
					Email: "short@gmail.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				entity.ShortLink{
					Alias:    "mySimpleAlias",
					LongLink: "https://www.google.com/",
				},
			},
			expectedShortLink: &URL{
				url: entity.ShortLink{
					Alias:    newAlias,
					LongLink: newLongLink,
				},
			},
			shouldError: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			blockedHash := map[string]bool{
				"http://malware.wicar.org/data/ms14_064_ole_not_xp.html": false,
			}
			blacklist := risk.NewBlackListFake(blockedHash)
			shortLinkRepo := repository.NewShortLinkFake(testCase.urls)
			userShortLinkRepo := repository.NewUserShortLinkRepoFake(
				testCase.relationUsers,
				testCase.relationShortLinks,
			)

			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			longLinkValidator := validator.NewLongLink()
			aliasValidator := validator.NewCustomAlias()
			riskDetector := risk.NewDetector(blacklist)

			tm := timer.NewStub(now)
			changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			changeLog := changelog.NewPersist(keyGen, tm, &changeLogRepo, &userChangeLogRepo)

			tokenizer := crypto.NewTokenizerFake()
			auth := authenticator.NewAuthenticator(tokenizer, tm, time.Hour)

			authToken, err := auth.GenerateToken(testCase.user)
			assert.Equal(t, nil, err)

			creator := url.NewCreatorPersist(
				&shortLinkRepo,
				&userShortLinkRepo,
				keyGen,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)
			updater := url.NewUpdaterPersist(
				&shortLinkRepo,
				&userShortLinkRepo,
				keyGen,
				longLinkValidator,
				aliasValidator,
				tm,
				riskDetector,
			)
			authMutation := newAuthMutation(
				&authToken,
				auth,
				changeLog,
				creator,
				updater,
			)
			shortLink, err := authMutation.UpdateURL(testCase.args)
			assert.Equal(t, nil, err)
			if shortLink == nil {
				return
			}
			assert.Equal(t, testCase.expectedShortLink.url.Alias, shortLink.url.Alias)
			assert.Equal(t, testCase.expectedShortLink.url.LongLink, shortLink.url.LongLink)
			assert.Equal(t, true, shortLink.url.UpdatedAt.After(now))
		})
	}
}
