package matcher

import (
	"testing"

	"github.com/short-d/app/fw/assert"
)

func TestContainsAll_IsMatch(t *testing.T) {
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
			expected: true,
		},
		{
			name:     "empty words",
			words:    nil,
			input:    "a",
			expected: true,
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
			name:     "all match",
			words:    []string{"a", "ab", "aa"},
			input:    "aaaba",
			expected: true,
		},
		{
			name:     "no match",
			words:    []string{"a", "b", "c"},
			input:    "xyz",
			expected: false,
		},
		{
			name:     "at least one mismatch",
			words:    []string{"a", "b", "c"},
			input:    "xcz",
			expected: false,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testCase.expected, ContainsAll(testCase.words, testCase.input))
		})
	}
}
