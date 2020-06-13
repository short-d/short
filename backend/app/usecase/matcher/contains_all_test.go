package matcher

import (
	"testing"

	"github.com/short-d/app/fw/assert"
)

func TestContainsAll_IsMatch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		query    string
		input    string
		expected bool
	}{
		{
			name:     "empty query and empty input",
			query:    "",
			input:    "",
			expected: false,
		},
		{
			name:     "empty query",
			query:    "",
			input:    "a",
			expected: false,
		},
		{
			name:     "empty input",
			query:    "a",
			input:    "",
			expected: false,
		},
		{
			name:     "complete match",
			query:    "a",
			input:    "a",
			expected: true,
		},
		{
			name:     "match",
			query:    "a",
			input:    "aaaa",
			expected: true,
		},
		{
			name:     "no match",
			query:    "a b c",
			input:    "xyz",
			expected: false,
		},
		{
			name:     "in between match",
			query:    "a b c",
			input:    "xcz",
			expected: false,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			matchAll := NewKeyword(MatchAll)

			assert.Equal(t, testCase.expected, matchAll.IsMatch(testCase.query, testCase.input))
		})
	}
}
