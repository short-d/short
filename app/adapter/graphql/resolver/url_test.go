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
	urlResolver := URL{url: entity.URL{ExpireAt: &timeAfter}}

	expected := scalar.Time{Time: *urlResolver.url.ExpireAt}.Time
	got := urlResolver.ExpireAt().Time
	mdtest.Equal(t, got, expected, "*urlResolver.OriginalURL() = %v; want %v", expected, got)

	urlResolver.url.ExpireAt = nil
	var expectedTime *scalar.Time = nil
	mdtest.Equal(t, expectedTime, urlResolver.ExpireAt(), "*urlResolver.OriginalURL() = %v; want %v", got, nil)
}
