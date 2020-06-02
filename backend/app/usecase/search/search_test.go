package search

import (
	"fmt"
	"testing"

	"github.com/short-d/short/backend/app/entity"
)

type shortLinks = map[string]entity.ShortLink

func TestMerging(t *testing.T) {
	results := []Result{
		{
			shortLinks: []entity.ShortLink{
				{Alias: "google"},
				{Alias: "facebook"},
			},
			users: nil,
		},
		{
			shortLinks: nil,
			users: []entity.User{
				{ID: "test"},
			},
		},
	}
	result := mergeResults(results)
	fmt.Println(result)
}

//func TestSearch(t *testing.T) {
//	t.Parallel()
//
//	shortLinksMap := make(shortLinks)
//	shortLinksMap["git-google"] = entity.ShortLink{
//		Alias:    "git-google",
//		LongLink: "http://github.com/google",
//	}
//	shortLinksMap["google"] = entity.ShortLink{
//		Alias:    "google",
//		LongLink: "https://google.com",
//	}
//
//	users := []entity.User{
//		{
//			ID:    "alpha",
//			Email: "alpha@example.com",
//		},
//		{
//			ID:    "alpha",
//			Email: "alpha@example.com",
//		},
//	}
//
//	userShortLinks := []entity.ShortLink{
//		{
//			Alias:    "git-google",
//			LongLink: "http://github.com/google",
//		},
//		{
//			Alias:    "google",
//			LongLink: "https://google.com",
//		},
//	}
//
//	testCases := []struct {
//		name               string
//		shortLinks         shortLinks
//		Query              Query
//		filter             Filter
//		relationUsers      []entity.User
//		relationShortLinks []entity.ShortLink
//		expectedResult     Result
//	}{
//		{
//			name:       "valid Search",
//			shortLinks: shortLinksMap,
//			Query: Query{
//				Query: "http google",
//				User:  entity.User{},
//			},
//			filter: Filter{
//				MaxResults: 2,
//				Resources:  []Resource{ShortLink},
//				Orders:     []OrderBy{CreatedTimeASC},
//			},
//			relationUsers:      users,
//			relationShortLinks: userShortLinks,
//			expectedResult: Result{
//				shortLinks: []entity.ShortLink{
//					{
//						Alias:    "git-google",
//						LongLink: "http://github.com/google",
//					},
//					{
//						Alias:    "google",
//						LongLink: "https://google.com",
//					},
//				},
//				users: nil,
//			},
//		},
//	}
//
//	for _, testCase := range testCases {
//		testCase := testCase
//		t.Run(testCase.name, func(t *testing.T) {
//			t.Parallel()
//			userShortLinkRepo := repository.NewUserShortLinkRepoFake(testCase.relationUsers, testCase.relationShortLinks)
//			shortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks)
//
//			search := NewSearch(&shortLinkRepo, &userShortLinkRepo)
//			result := search.Search(testCase.Query, testCase.filter)
//			assert.Equal(t, testCase.expectedResult, result)
//		})
//	}
//}
