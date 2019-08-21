package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
	"short/app/usecase/url"
	"short/mdtest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type urlMap = map[string]entity.URL

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
			retrieverFake := url.NewRetrieverFake(testCase.urls)
			query := NewQuery(mdtest.FakeLogger, mdtest.FakeTracer, retrieverFake)

			urlArgs := &URLArgs{
				Alias:       testCase.alias,
				ExpireAfter: testCase.expireAfter,
			}

			u, err := query.URL(urlArgs)

			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, testCase.expectedURL, u)
		})
	}
}
