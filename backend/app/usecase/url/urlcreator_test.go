package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestURLCreatorPersist_CreateURL(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        urlMap
		alias       string
		user        entity.User
		url         entity.URL
		expHasErr   bool
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
			alias: "220uFicCJj",
			user: entity.User{
				Email: "alpha@example.com",
			},
			url:       entity.URL{},
			expHasErr: true,
		},
		{
			name:  "create alias successfully",
			urls:  urlMap{},
			alias: "220uFicCJj",
			user: entity.User{
				Email: "alpha@example.com",
			},
			url: entity.URL{
				Alias:    "220uFicCJj",
				ExpireAt: &now,
			},
			expHasErr: false,
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

			urlsCopy := copyURLMap(testCase.urls)
			fakeCreator := NewCreatorFake(urlsCopy, []string{
				testCase.alias,
			})

			creators := []Creator{creator, fakeCreator}

			for _, c := range creators {
				url, err := c.CreateURL(testCase.url, testCase.user)
				if testCase.expHasErr {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedURL, url)
			}
		})
	}
}

func copyURLMap(srcURLMap map[string]entity.URL) map[string]entity.URL {
	urlsCopy := make(map[string]entity.URL)
	for key, val := range srcURLMap {
		urlsCopy[key] = val
	}
	return urlsCopy
}
