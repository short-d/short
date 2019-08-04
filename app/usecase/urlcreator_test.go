package usecase

import (
	"short/app/entity"
	"short/app/repo"
	"short/modern/mdtest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUrlCreatorPersist_CreateUrl(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        urlMap
		alias       string
		hasErr      bool
		expectedUrl entity.Url
	}{
		{
			name: "alias exists",
			urls: urlMap{
				"220uFicCJj": entity.Url{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias:       "220uFicCJj",
			hasErr:      true,
			expectedUrl: entity.Url{},
		},
		{
			name:   "create alias successfully",
			urls:   urlMap{},
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
			retriever := NewUrlRetrieverPersist(fakeRepo)
			fakeTrace := mdtest.FakeTracer.BeginTrace("GetUrl")
			url, err := retriever.GetUrl(fakeTrace, testCase.alias)

			if testCase.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedUrl, url)
			}
		})
	}
}
