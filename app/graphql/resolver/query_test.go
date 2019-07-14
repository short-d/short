package resolver

import (
	"testing"
	"time"
	"tinyURL/app/entity"
	"tinyURL/app/graphql/scalar"
	"tinyURL/app/usecase"
	"tinyURL/modern/mdtest"

	"github.com/stretchr/testify/assert"
)

type urlMap = map[string]entity.Url

func TestQuery_Url(t *testing.T) {
	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name        string
		alias       string
		expireAfter *scalar.Time
		urls        urlMap
		hasErr      bool
		expectedUrl *Url
	}{
		{
			name:        "alias not found with no expireAfter",
			alias:       "220uFicCJj",
			expireAfter: nil,
			urls:        urlMap{},
			hasErr:      true,
			expectedUrl: nil,
		},
		{
			name:  "alias not found with expireAfter",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			urls:        urlMap{},
			hasErr:      true,
			expectedUrl: nil,
		},
		{
			name:  "alias expired",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			urls: urlMap{
				"220uFicCJj": entity.Url{
					ExpireAt: &before,
				},
			},
			hasErr:      true,
			expectedUrl: nil,
		},
		{
			name:  "url found",
			alias: "220uFicCJj",
			expireAfter: &scalar.Time{
				Time: now,
			},
			urls: urlMap{
				"220uFicCJj": entity.Url{
					ExpireAt: &after,
				},
			},
			hasErr: false,
			expectedUrl: &Url{
				url: entity.Url{
					ExpireAt: &after,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			retrieverFake := usecase.NewUrlRetrieverFake(testCase.urls)
			query := NewQuery(mdtest.FakeLogger, mdtest.FakeTracer, retrieverFake)

			urlArgs := &UrlArgs{
				Alias:       testCase.alias,
				ExpireAfter: testCase.expireAfter,
			}

			url, err := query.Url(urlArgs)

			if testCase.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, testCase.expectedUrl, url)
			}
		})
	}
}
