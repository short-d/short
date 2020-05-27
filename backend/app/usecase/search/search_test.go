package search

import (
	"testing"

	"github.com/short-d/short/backend/app/entity"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		query          Query
		filter         Filter
		expectedResult Result
	}{
		{
			name: "valid search",
			query: Query{
				keywords: "http google",
				user:     entity.User{},
			},
			filter: Filter{
				maxResults: 2,
				resources:  []Resource{ShortLink},
				orders:     []Order{CreatedTimeASC},
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
				users: nil,
			},
		},
	}

	_ = testCases
}
