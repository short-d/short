// +build !integration all

package shortlink

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type shortLinks = map[string]entity.ShortLink

func TestRetriever_GetShortLink(t *testing.T) {
	t.Parallel()

	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name              string
		shortLinks        shortLinks
		alias             string
		expiringAt        *time.Time
		hasErr            bool
		expectedShortLink entity.ShortLink
	}{
		{
			name:              "alias not found",
			shortLinks:        shortLinks{},
			alias:             "220uFicCJj",
			expiringAt:        &now,
			hasErr:            true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name: "short link expired",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: &before,
				},
			},
			alias:             "220uFicCJj",
			expiringAt:        &now,
			hasErr:            true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name: "short link never expire",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: nil,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: &now,
			hasErr:     false,
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				ExpireAt: nil,
			},
		},
		{
			name: "get without expiration",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: &before,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: nil,
			hasErr:     false,
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				ExpireAt: &before,
			},
		},
		{
			name: "unexpired short link found",
			shortLinks: shortLinks{
				"220uFicCJj": entity.ShortLink{
					Alias:    "220uFicCJj",
					ExpireAt: &after,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: &now,
			hasErr:     false,
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				ExpireAt: &after,
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			fakeShortLinkRepo := repository.NewShortLinkFake(nil, testCase.shortLinks)
			fakeUserShortLinkRepo := repository.NewUserShortLinkRepoFake([]entity.User{}, []entity.ShortLink{})
			retriever := NewRetrieverPersist(&fakeShortLinkRepo, &fakeUserShortLinkRepo)
			shortLink, err := retriever.GetShortLink(testCase.alias, testCase.expiringAt)

			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLink, shortLink)
		})
	}
}

func TestRetrieverPersist_GetShortLinks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		users              []entity.User
		createdShortLinks  []entity.ShortLink
		user               entity.User
		hasErr             bool
		expectedShortLinks []entity.ShortLink
	}{
		{
			name: "user created short links",
			shortLinks: shortLinks{
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://www.google.com/",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://github.com/short-d/short/",
				},
				"mozilla": entity.ShortLink{
					Alias:    "mozilla",
					LongLink: "https://www.mozilla.org/",
				},
			},
			users: []entity.User{
				{
					ID:    "12345",
					Name:  "Test User",
					Email: "test@gmail.com",
				}, {
					ID:    "12345",
					Name:  "Test User",
					Email: "test@gmail.com",
				}, {
					ID:    "12346",
					Name:  "Test User 2",
					Email: "test2@gmail.com",
				},
			},
			createdShortLinks: []entity.ShortLink{
				{
					Alias:    "google",
					LongLink: "https://www.google.com/",
				},
				{
					Alias:    "short",
					LongLink: "https://github.com/short-d/short/",
				},
				{
					Alias:    "mozilla",
					LongLink: "https://www.mozilla.org/",
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			hasErr: false,
			expectedShortLinks: []entity.ShortLink{
				{
					Alias:    "google",
					LongLink: "https://www.google.com/",
				},
				{
					Alias:    "short",
					LongLink: "https://github.com/short-d/short/",
				},
			},
		},
		{
			name: "user has no ShortLink",
			shortLinks: shortLinks{
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://www.google.com/",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://github.com/short-d/short/",
				},
				"mozilla": entity.ShortLink{
					Alias:    "mozilla",
					LongLink: "https://www.mozilla.org/",
				},
			},
			users: []entity.User{
				{
					ID:    "12345",
					Name:  "Test User",
					Email: "test@gmail.com",
				}, {
					ID:    "12345",
					Name:  "Test User",
					Email: "test@gmail.com",
				}, {
					ID:    "12345",
					Name:  "Test User",
					Email: "test@gmail.com",
				},
			},
			createdShortLinks: []entity.ShortLink{
				{
					Alias:    "google",
					LongLink: "https://www.google.com/",
				},
				{
					Alias:    "short",
					LongLink: "https://github.com/short-d/short/",
				},
				{
					Alias:    "mozilla",
					LongLink: "https://www.mozilla.org/",
				},
			},
			user: entity.User{
				ID:    "12346",
				Name:  "Test User 2",
				Email: "test2@gmail.com",
			},
			hasErr:             false,
			expectedShortLinks: []entity.ShortLink{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			fakeShortLinkRepo := repository.NewShortLinkFake(nil, testCase.shortLinks)
			fakeUserShortLinkRepo := repository.NewUserShortLinkRepoFake(testCase.users, testCase.createdShortLinks)
			retriever := NewRetrieverPersist(&fakeShortLinkRepo, &fakeUserShortLinkRepo)

			shortLinks, err := retriever.GetShortLinksByUser(testCase.user)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedShortLinks, shortLinks)
		})
	}
}
