package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
	"short/app/usecase/repo"
	"short/app/usecase/url"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

type urlMap = map[string]entity.URL

func TestAuthQuery_URL(t *testing.T) {
	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name        string
		user        entity.User
		alias       string
		expireAfter *scalar.Time
		urls        urlMap
		hasErr      bool
		expectedURL *URL
	}{
		{
			name:        "alias not found with no expireAfter",
			alias:       "220uFicCJj",
			expireAfter: nil,
			urls:        urlMap{},
			hasErr:      true,
		},
		{
			name:  "alias not found with expireAfter",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			urls:   urlMap{},
			hasErr: true,
		},
		{
			name:  "alias expired",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			urls: urlMap{
				"220uFicCJj": entity.URL{
					ExpireAt: &before,
				},
			},
			hasErr: true,
		},
		{
			name:  "url found",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			urls: urlMap{
				"220uFicCJj": entity.URL{
					ExpireAt: &after,
				},
			},
			hasErr: false,
			expectedURL: &URL{
				url: entity.URL{
					ExpireAt: &after,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sqlDB, _, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer sqlDB.Close()

			fakeRepo := repo.NewURLFake(testCase.urls)
			retrieverFake := url.NewRetrieverPersist(&fakeRepo)
			query := newAuthQuery(&testCase.user, retrieverFake)

			urlArgs := &URLArgs{
				Alias:       testCase.alias,
				ExpireAfter: testCase.expireAfter,
			}

			u, err := query.URL(urlArgs)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, testCase.expectedURL, u)
		})
	}
}
