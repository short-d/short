// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
)

func TestShortLink_Alias(t *testing.T) {
	t.Parallel()
	shortLinkResolver := ShortLink{shortLink: entity.ShortLink{Alias: "TestAlias"}}

	expected := shortLinkResolver.shortLink.Alias
	got := *shortLinkResolver.Alias()
	assert.Equal(t, got, expected, "*shortLinkTest.Alias() = %v; want %v", expected, got)
}

func TestShortLink_LongLink(t *testing.T) {
	t.Parallel()
	shortLinkResolver := ShortLink{shortLink: entity.ShortLink{LongLink: "TestLongLink"}}

	expected := shortLinkResolver.shortLink.LongLink
	got := *shortLinkResolver.LongLink()
	assert.Equal(t, got, expected, "*shortLinkResolver.LongLink() = %v; want %v", expected, got)
}

func TestShortLink_ExpireAt(t *testing.T) {
	t.Parallel()
	timeAfter := time.Now().Add(5 * time.Second)
	testCases := []struct {
		shortLink ShortLink
		expected  *scalar.Time
	}{
		{
			shortLink: ShortLink{shortLink: entity.ShortLink{ExpireAt: &timeAfter}},
			expected:  &scalar.Time{Time: timeAfter},
		},
		{
			shortLink: ShortLink{shortLink: entity.ShortLink{ExpireAt: nil}},
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		assert.Equal(t, testCase.expected, testCase.shortLink.ExpireAt())
	}
}
