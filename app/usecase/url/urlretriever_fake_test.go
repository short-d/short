package url

import (
	"short/app/entity"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

type fakeURLMap = map[string]entity.URL

func TestFakeURLRetriever_GetURLAfter(t *testing.T) {
	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name        string
		urls        fakeURLMap
		alias       string
		expiringAt  time.Time
		hasErr      bool
		expectedURL entity.URL
		expectedErr string
	}{
		{
			name:        "alias not found",
			urls:        fakeURLMap{},
			alias:       "220uFicCJj",
			expiringAt:  now,
			hasErr:      true,
			expectedURL: entity.URL{},
			expectedErr: "url not found",
		},
		{
			name: "url expired",
			urls: fakeURLMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &before,
				},
			},
			alias:       "220uFicCJj",
			expiringAt:  now,
			hasErr:      true,
			expectedURL: entity.URL{},
			expectedErr: "url expired",
		},
		{
			name: "url never expires",
			urls: fakeURLMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: nil,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: now,
			hasErr:     false,
			expectedURL: entity.URL{
				Alias:    "220uFicCJj",
				ExpireAt: nil,
			},
		},
		{
			name: "unexpired url found",
			urls: fakeURLMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &after,
				},
			},
			alias:      "220uFicCJj",
			expiringAt: now,
			hasErr:     false,
			expectedURL: entity.URL{
				Alias:    "220uFicCJj",
				ExpireAt: &after,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tracer := mdtest.NewTracerFake()
			fakeTrace := tracer.BeginTrace("GetURLAfter")
			fakeRetriever := NewRetrieverFake(testCase.urls)

			gotURL, err := fakeRetriever.GetAfter(fakeTrace, testCase.alias, testCase.expiringAt)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				mdtest.Equal(t, testCase.expectedErr, err.Error())
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, gotURL)
		})
	}
}

func TestFakeURLRetriever_GetURL(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        fakeURLMap
		alias       string
		hasErr      bool
		expectedURL entity.URL
		expectedErr string
	}{
		{
			name:        "alias not found",
			urls:        fakeURLMap{},
			alias:       "220uFicCJj",
			hasErr:      true,
			expectedURL: entity.URL{},
			expectedErr: "url not found",
		},
		{
			name: "valid url found",
			urls: fakeURLMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &now,
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedURL: entity.URL{
				Alias:    "220uFicCJj",
				ExpireAt: &now,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tracer := mdtest.NewTracerFake()
			fakeTrace := tracer.BeginTrace("GetURL")
			fakeRetriever := NewRetrieverFake(testCase.urls)

			gotURL, err := fakeRetriever.Get(fakeTrace, testCase.alias)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				mdtest.Equal(t, testCase.expectedErr, err.Error())
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, gotURL)
		})
	}
}
