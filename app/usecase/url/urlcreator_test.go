package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestURLCreatorPersist_CreateUrl(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        urlMap
		alias       string
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
			url:    entity.URL{},
			alias:  "220uFicCJj",
			hasErr: true,
		},
		{
			name:  "create alias successfully",
			urls:  urlMap{},
			alias: "220uFicCJj",
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
			fakeRepo := repo.NewURLFake(testCase.urls)
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
			assert.Equal(t, testCase.expectedURL, url)
		})
	}
}
