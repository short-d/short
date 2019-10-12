package resolver

import (
	"github.com/byliuyang/app/mdtest"
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
	"testing"
	"time"
)

func TestURL_Alias(t *testing.T) {
	urlTest := URL{url: struct {
		Alias       string
		OriginalURL string
		ExpireAt    *time.Time
		CreatedBy   *entity.User
		CreatedAt   *time.Time
		UpdatedAt   *time.Time
	}{Alias: "TestAlias"}}

	expected := urlTest.url.Alias
	result := *urlTest.Alias()
	mdtest.Equal(t, result, expected, "*urlTest.Alias() of '%v' is %v", expected, result)
}

func TestURL_OriginalURL(t *testing.T) {
	urlTest := URL{url: struct {
		Alias       string
		OriginalURL string
		ExpireAt    *time.Time
		CreatedBy   *entity.User
		CreatedAt   *time.Time
		UpdatedAt   *time.Time
	}{OriginalURL: "TestOriginalUrl"}}

	expected := urlTest.url.OriginalURL
	result := *urlTest.OriginalURL()
	mdtest.Equal(t, result, expected, "*urlTest.OriginalURL() of '%v' is %v", expected, result)
}

func TestURL_ExpireAt(t *testing.T) {
	timeAfter := time.Now().Add(5 * time.Second)
	urlTest := URL{url: struct {
		Alias       string
		OriginalURL string
		ExpireAt    *time.Time
		CreatedBy   *entity.User
		CreatedAt   *time.Time
		UpdatedAt   *time.Time
	}{ExpireAt: &timeAfter}}

	expected := scalar.Time{Time: *urlTest.url.ExpireAt}.Time
	result := urlTest.ExpireAt().Time
	mdtest.Equal(t, result, expected, "*urlTest.OriginalURL() instead of '%v' is %v", expected, result)

	urlTest.url.ExpireAt = nil
	var expectedTime *scalar.Time = nil
	mdtest.Equal(t, expectedTime, urlTest.ExpireAt(), "*urlTest.OriginalURL() of '%v' is %v", result, nil)
}
