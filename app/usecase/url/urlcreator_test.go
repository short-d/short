package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
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
		url         entity.Url
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
			url:    entity.Url{},
			alias:  "220uFicCJj",
			hasErr: true,
		},
		{
			name:  "create alias successfully",
			urls:  urlMap{},
			alias: "220uFicCJj",
			url: entity.Url{
				Alias:    "220uFicCJj",
				ExpireAt: &now,
			},
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
			fakeKeyGen := keygen.NewFake([]string{
				testCase.alias,
			})
			creator := NewCreatorPersist(&fakeRepo, &fakeKeyGen)
			url, err := creator.Create(testCase.url)

			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedUrl, url)
		})
	}
}
