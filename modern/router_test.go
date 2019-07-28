package modern

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUriMatcher(t *testing.T) {
	testCases := []struct {
		uri    string
		hasErr bool
	}{
		{
			uri:    "",
			hasErr: true,
		},
		{
			uri:    "/",
			hasErr: false,
		},
		{
			uri:    "/users",
			hasErr: false,
		},
		{
			uri:    "/users/:userId",
			hasErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			_, err := NewUriMatcher(testCase.uri)

			if testCase.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUriMatcher_Params(t *testing.T) {
	testCases := []struct {
		uri       string
		expParams []string
	}{
		{
			uri:       "/",
			expParams: []string{},
		},
		{
			uri:       "/users",
			expParams: []string{},
		},
		{
			uri:       "/users/:userId",
			expParams: []string{"userId"},
		},
		{
			uri:       "/users/:userId/articles",
			expParams: []string{"userId"},
		},
		{
			uri:       "/users/:userId/articles/:articleId",
			expParams: []string{"userId", "articleId"},
		},
		{
			uri:       "/users/:userId/articles/:articleId/doSomething",
			expParams: []string{"userId", "articleId"},
		},
	}

	for _, testCase := range testCases {

		t.Run("", func(t *testing.T) {
			urlMatcher, err := NewUriMatcher(testCase.uri)

			if err != nil {
				return
			}

			gotParams := urlMatcher.Params()
			assert.Equal(t, testCase.expParams, gotParams)
		})
	}
}

func TestUriMatcher_Match(t *testing.T) {
	uriTemplate := "/users/:userId/articles/:articleId"
	matcher, err := NewUriMatcher(uriTemplate)

	assert.Nil(t, err)

	testCases := []struct {
		path       string
		expIsMatch bool
		expParams  Params
	}{
		{
			path:       "/",
			expIsMatch: false,
		},
		{
			path:       "/users",
			expIsMatch: false,
		},
		{
			path:       "/users/fr4esw1rdf",
			expIsMatch: false,
		},
		{
			path:       "/users/fr4esw1rdf/articles",
			expIsMatch: false,
		},
		{
			path:       "/users/fr4esw1rdf/articles/1dsd2DwxS",
			expIsMatch: true,
			expParams: Params{
				"articleId": "1dsd2DwxS",
				"userId":    "fr4esw1rdf",
			},
		},
		{
			path:       "/users/fr4esw1rdf/articles/1dsd2DwxS/",
			expIsMatch: true,
			expParams: Params{
				"articleId": "1dsd2DwxS",
				"userId":    "fr4esw1rdf",
			},
		},
		{
			path:       "/users/fr4esw1rdf/articles/1dsd2DwxS/doSomething",
			expIsMatch: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("  %s", testCase.path), func(t *testing.T) {
			isMatch, gotParams := matcher.IsMatch(testCase.path)
			assert.Equal(t, testCase.expIsMatch, isMatch)

			if isMatch {
				assert.Equal(t, testCase.expParams, gotParams)
			}
		})
	}
}
