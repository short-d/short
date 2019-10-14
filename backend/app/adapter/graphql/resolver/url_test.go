package resolver

import (
	"github.com/byliuyang/app/mdtest"
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
	"testing"
	"time"
)

func TestURL_Alias(t *testing.T) {
	urlResolver := URL{url: entity.URL{Alias: "TestAlias"}}

	expected := urlResolver.url.Alias
	got := *urlResolver.Alias()
	mdtest.Equal(t, got, expected, "*urlTest.Alias() = %v; want %v", expected, got)
}

func TestURL_OriginalURL(t *testing.T) {
	urlResolver := URL{url: entity.URL{OriginalURL: "TestOriginalUrl"}}

	expected := urlResolver.url.OriginalURL
	got := *urlResolver.OriginalURL()
	mdtest.Equal(t, got, expected, "*urlResolver.OriginalURL() = %v; want %v", expected, got)
}

func TestURL_ExpireAt(t *testing.T) {
	timeAfter := time.Now().Add(5 * time.Second)
	testCases := []struct {
		url    URL
		expected *scalar.Time
	}{
		{
			url:    URL{url: entity.URL{ExpireAt: &timeAfter}},
			expected: &scalar.Time{Time: timeAfter},
		},
		{
			url: URL{url: entity.URL{ExpireAt: nil}},
		},
	}
	for _, testCase := range testCases {
		mdtest.Equal(t, testCase.expected, testCase.url.ExpireAt())
	}
}
