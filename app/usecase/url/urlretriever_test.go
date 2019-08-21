package url

import (
	"short/app/entity"
	"short/app/usecase/repo"
	"short/mdtest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type urlMap = map[string]entity.Url

func TestUrlRetriever_GetUrlAfter(t *testing.T) {
	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name        string
		urls        urlMap
		alias       string
		expiringAt  time.Time
		hasErr      bool
		expectedUrl entity.Url
	}{
		{
			name:        "alias not found",
			urls:        urlMap{},
			alias:       "220uFicCJj",
			expiringAt:  now,
			hasErr:      true,
			expectedUrl: entity.Url{},
		},
		{
			name: "url expired",
			urls: urlMap{
				"220uFicCJj": entity.Url{
					Alias:    "220uFicCJj",
					ExpireAt: &before,
				},
			},
			alias:       "220uFicCJj",
			expiringAt:  now,
			hasErr:      true,
			expectedUrl: entity.Url{},
		},
		{
			name: "url never expire",
			urls: urlMap{
				"220uFicCJj": entity.Url{
					Alias:    "220uFicCJj",
					ExpireAt: nil,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: now,
			hasErr:     false,
			expectedUrl: entity.Url{
				Alias:    "220uFicCJj",
				ExpireAt: nil,
			},
		},
		{
			name: "unexpired url found",
			urls: urlMap{
				"220uFicCJj": entity.Url{
					Alias:    "220uFicCJj",
					ExpireAt: &after,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: now,
			hasErr:     false,
			expectedUrl: entity.Url{
				Alias:    "220uFicCJj",
				ExpireAt: &after,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fakeRepo := repo.NewUrlFake(testCase.urls)
			retriever := NewRetrieverPersist(&fakeRepo)
			fakeTrace := mdtest.FakeTracer.BeginTrace("GetUrlAfter")
			url, err := retriever.GetAfter(fakeTrace, testCase.alias, testCase.expiringAt)

			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedUrl, url)
		})
	}
}

func TestUrlRetriever_GetUrl(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        urlMap
		alias       string
		hasErr      bool
		expectedUrl entity.Url
	}{
		{
			name:        "alias not found",
			urls:        urlMap{},
			alias:       "220uFicCJj",
			hasErr:      true,
			expectedUrl: entity.Url{},
		},
		{
			name: "valid url found",
			urls: urlMap{
				"220uFicCJj": entity.Url{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedUrl: entity.Url{
				Alias:    "220uFicCJj",
				ExpireAt: &now,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fakeRepo := repo.NewUrlFake(testCase.urls)
			retriever := NewRetrieverPersist(&fakeRepo)
			fakeTrace := mdtest.FakeTracer.BeginTrace("GetUrl")
			url, err := retriever.Get(fakeTrace, testCase.alias)

			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedUrl, url)
		})
	}
}
