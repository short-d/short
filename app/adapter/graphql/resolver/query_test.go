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
			retrieverFake := url.NewRetrieverFake(testCase.urls)
			query := NewQuery(mdtest.FakeLogger, mdtest.FakeTracer, retrieverFake)

			urlArgs := &UrlArgs{
				Alias:       testCase.alias,
				ExpireAfter: testCase.expireAfter,
			}

			u, err := query.Url(urlArgs)

			if testCase.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, testCase.expectedUrl, u)
			}
		})
	}
}
