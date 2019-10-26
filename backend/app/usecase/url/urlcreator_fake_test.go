package url

import (
	"short/app/entity"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestURLFakeCreator_CreateURL(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        urlMap
		alias       []string
		userEmail   string
		url         entity.URL
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name: "alias exists",
			urls: urlMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias:     []string{"220uFicCJj"},
			userEmail: "alpha@example.com",
			url:       entity.URL{},
			hasErr:    true,
		},
		{
			name:      "create alias successfully",
			urls:      urlMap{},
			alias:     []string{"220uFicCJj"},
			userEmail: "alpha@example.com",
			url: entity.URL{
				Alias:    "220uFicCJj",
				ExpireAt: &now,
			},
			hasErr: false,
			expectedURL: entity.URL{
				Alias:    "220uFicCJj",
				ExpireAt: &now,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fakeCreator := NewCreatorFake(testCase.urls, testCase.alias)
			gotURL, err := fakeCreator.CreateURL(testCase.url, testCase.userEmail)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, gotURL)
		})
	}
}
