package search

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type shortLinks = map[string]entity.ShortLink

func TestSearch(t *testing.T) {
	t.Parallel()

	shortLinksMap := make(shortLinks)
	shortLinksMap["git-google"] = entity.ShortLink{
		Alias:    "git-google",
		LongLink: "http://github.com/google",
	}
	shortLinksMap["google"] = entity.ShortLink{
		Alias:    "google",
		LongLink: "https://google.com",
	}

	users := []entity.User{
		{
			ID:    "alpha",
			Email: "alpha@example.com",
		},
		{
			ID:    "alpha",
			Email: "alpha@example.com",
		},
	}

	userShortLinks := []entity.ShortLink{
		{
			Alias:    "git-google",
			LongLink: "http://github.com/google",
		},
		{
			Alias:    "google",
			LongLink: "https://google.com",
		},
	}

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		query              Query
		filter             Filter
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		expectedResult     Result
	}{
		{
			name:       "valid Search",
			shortLinks: shortLinksMap,
			query: Query{
				keywords: "http google",
				user:     entity.User{},
			},
			filter: Filter{
				maxResults: 2,
				resources:  []Resource{ShortLink},
				orders:     []Order{CreatedTimeASC},
			},
			relationUsers:      users,
			relationShortLinks: userShortLinks,
			expectedResult: Result{
				shortLinks: []entity.ShortLink{
					{
						Alias:    "git-google",
						LongLink: "http://github.com/google",
					},
					{
						Alias:    "google",
						LongLink: "https://google.com",
					},
				},
				users: nil,
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			userShortLinkRepo := repository.NewUserShortLinkRepoFake(testCase.relationUsers, testCase.relationShortLinks)
			shortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks)

			search := NewSearch(&shortLinkRepo, &userShortLinkRepo)
			result := search.Search(testCase.query, testCase.filter)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}
