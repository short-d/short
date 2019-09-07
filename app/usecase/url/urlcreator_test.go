package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestURLCreatorPersist_CreateUrl(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        urlMap
		alias       string
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
			alias:     "220uFicCJj",
			userEmail: "alpha@example.com",
			url:       entity.URL{},
			hasErr:    true,
		},
		{
			name:      "create alias successfully",
			urls:      urlMap{},
			alias:     "220uFicCJj",
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
			urlRepo := repo.NewURLFake(testCase.urls)
			userURLRepo := repo.NewUserURLRepoFake()
			keyGen := keygen.NewFake([]string{
				testCase.alias,
			})

			creator := NewCreatorPersist(&urlRepo, &userURLRepo, &keyGen)
			url, err := creator.Create(testCase.url, testCase.userEmail)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, url)
		})
	}
}
