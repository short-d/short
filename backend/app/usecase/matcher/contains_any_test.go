package matcher

import (
	"testing"

	"github.com/short-d/app/fw/assert"
)

func TestContainsAny_IsMatch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		words    []string
		input    string
		expected bool
	}{
		{
			name:     "empty words and empty input",
			words:    nil,
			input:    "",
			expected: false,
		},
		{
			name:     "empty words",
			words:    nil,
			input:    "a",
			expected: false,
		},
		{
			name:     "empty input",
			words:    []string{"a"},
			input:    "",
			expected: false,
		},
		{
			name:     "complete match",
			words:    []string{"a"},
			input:    "a",
			expected: true,
		},
		{
			name:     "match",
			words:    []string{"a"},
			input:    "aaaa",
			expected: true,
		},
		{
			name:     "no match",
			words:    []string{"a", "b", "c"},
			input:    "xyz",
			expected: false,
		},
		{
			name:     "in between match",
			words:    []string{"a", "b", "c"},
			input:    "xcz",
			expected: true,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			containsAny, err := NewKeyword(ContainsAnyKeyword)

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expected, containsAny.IsMatch(testCase.words, testCase.input))
		})
	}
}
