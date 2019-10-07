package url

import (
	"short/app/entity"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

type fakeUrlMap = map[string]entity.URL

func TestFakeUrlRetriever_GetUrlAfter(t *testing.T) {
	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name        string
		urls        fakeUrlMap
		alias       string
		expiringAt  time.Time
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name:        "alias not found",
			urls:        fakeUrlMap{},
			alias:       "220uFicCJj",
			expiringAt:  now,
			hasErr:      true,
			expectedURL: entity.URL{},
		},
		{
			name: "url expired",
			urls: fakeUrlMap{
				"220uFicCJj": entity.URL{
					Alias:    "220uFicCJj",
					ExpireAt: &before,
				},
			},
			alias:       "220uFicCJj",
			expiringAt:  now,
			hasErr:      true,
			expectedURL: entity.URL{},
		},
		{
			name: "url never expire",
			urls: fakeUrlMap{
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
			urls: fakeUrlMap{
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
			fakeTrace := tracer.BeginTrace("GetUrlAfter")
			fakeRetriever := NewRetrieverFake(testCase.urls)

			url, err := fakeRetriever.GetAfter(fakeTrace, testCase.alias, testCase.expiringAt)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, url)
		})
	}
}

func TestFakeUrlRetriever_GetUrl(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name        string
		urls        fakeUrlMap
		alias       string
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name:        "alias not found",
			urls:        fakeUrlMap{},
			alias:       "220uFicCJj",
			hasErr:      true,
			expectedURL: entity.URL{},
		},
		{
			name: "valid url found",
			urls: fakeUrlMap{
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
			fakeTrace := tracer.BeginTrace("GetUrl")
			fakeRetriever := NewRetrieverFake(testCase.urls)

			url, err := fakeRetriever.Get(fakeTrace, testCase.alias)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, url)
		})
	}
}
