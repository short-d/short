package search

import (
	"testing"
	"time"

	"github.com/short-d/short/backend/app/usecase/search/order"

	"github.com/bmizerany/assert"
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
	shortLinksMap["short"] = entity.ShortLink{
		Alias:    "short",
		LongLink: "https://short-d.com",
	}
	shortLinksMap["facebook"] = entity.ShortLink{
		Alias:    "facebook",
		LongLink: "https://facebook.com",
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
		{
			ID:    "beta",
			Email: "beta@example.com",
		},
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
		{
			Alias:    "google",
			LongLink: "https://google.com",
		},
		{
			Alias:    "short",
			LongLink: "https://short-d.com",
		},
		{
			Alias:    "facebook",
			LongLink: "https://facebook.com",
		},
	}

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		Query              Query
		filter             Filter
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		expectedHasErr     bool
		expectedResult     Result
	}{
		{
			name:       "valid search",
			shortLinks: shortLinksMap,
			Query: Query{
				Query: "http google",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				Resources:  []Resource{ShortLink},
				Orders:     []order.By{order.ByCreatedTimeASC},
			},
			relationUsers:      users,
			relationShortLinks: userShortLinks,
			expectedResult: Result{
				shortLinks: []entity.ShortLink{
					{
						Alias:    "google",
						LongLink: "https://google.com",
					},
					{
						Alias:    "git-google",
						LongLink: "http://github.com/google",
					},
				},
				users: nil,
			},
		},
		{
			name:       "empty search",
			shortLinks: shortLinksMap,
			Query: Query{
				Query: "",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				Resources:  []Resource{ShortLink},
				Orders:     []order.By{order.ByCreatedTimeASC},
			},
			relationUsers:      users,
			relationShortLinks: userShortLinks,
			expectedResult: Result{
				shortLinks: nil,
				users:      nil,
			},
		},
		{
			name:       "one result",
			shortLinks: shortLinksMap,
			Query: Query{
				Query: "short",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				Resources:  []Resource{ShortLink},
				Orders:     []order.By{order.ByCreatedTimeASC},
			},
			relationUsers:      users,
			relationShortLinks: userShortLinks,
			expectedResult: Result{
				shortLinks: []entity.ShortLink{
					{
						Alias:    "short",
						LongLink: "https://short-d.com",
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
			timeout := time.Second

			search := NewSearch(&shortLinkRepo, &userShortLinkRepo, timeout)

			result, err := search.Search(testCase.Query, testCase.filter)
			if testCase.expectedHasErr {
				assert.NotEqual(t, nil, err)
			}

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}
