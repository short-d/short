package search

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

type shortLinks = map[string]entity.ShortLink

func TestSearch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		Query              Query
		filter             Filter
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		expectedResult     Result
	}{
		{
			name: "search without user",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "http google",
			},
			filter: Filter{
				MaxResults: 2,
				OrderedResources: []OrderedResource{
					{
						Resource: ShortLink,
						Order:    order.ByCreatedTimeASC,
					},
				},
			},
			relationUsers: []entity.User{
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
			},
			relationShortLinks: []entity.ShortLink{
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
			},
			expectedResult: Result{},
		},
		{
			name: "search without query",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				OrderedResources: []OrderedResource{
					{
						Resource: ShortLink,
						Order:    order.ByCreatedTimeASC,
					},
				},
			},
			relationUsers: []entity.User{
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
			},
			relationShortLinks: []entity.ShortLink{
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
			},
			expectedResult: Result{
				shortLinks: []entity.ShortLink{
					{
						Alias:    "facebook",
						LongLink: "https://facebook.com",
					},
					{
						Alias:    "short",
						LongLink: "https://short-d.com",
					},
				},
			},
		},
		{
			name: "no order given",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "google",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				OrderedResources: []OrderedResource{
					{
						Resource: ShortLink,
					},
				},
			},
			relationUsers: []entity.User{
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
			},
			relationShortLinks: []entity.ShortLink{
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
			},
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
			},
		},
		{
			name: "valid search",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "http google",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				OrderedResources: []OrderedResource{
					{
						Resource: ShortLink,
						Order:    order.ByCreatedTimeASC,
					},
				},
			},
			relationUsers: []entity.User{
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
			},
			relationShortLinks: []entity.ShortLink{
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
			},
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
			name: "query no match",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "non existent",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				OrderedResources: []OrderedResource{
					{
						Resource: ShortLink,
						Order:    order.ByCreatedTimeASC,
					},
				},
			},
			relationUsers: []entity.User{
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
			},
			relationShortLinks: []entity.ShortLink{
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
			},
			expectedResult: Result{
				shortLinks: nil,
				users:      nil,
			},
		},
		{
			name: "less matches than maxResults",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "short",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			filter: Filter{
				MaxResults: 2,
				OrderedResources: []OrderedResource{
					{
						Resource: ShortLink,
						Order:    order.ByCreatedTimeASC,
					},
				},
			},
			relationUsers: []entity.User{
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
			},
			relationShortLinks: []entity.ShortLink{
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
			},
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

			entryRepo := logger.NewEntryRepoFake()
			lg, err := logger.NewFake(logger.LogOff, &entryRepo)
			assert.Equal(t, nil, err)

			search := NewSearch(lg, &shortLinkRepo, &userShortLinkRepo, timeout)

			result, err := search.Search(testCase.Query, testCase.filter)

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}
