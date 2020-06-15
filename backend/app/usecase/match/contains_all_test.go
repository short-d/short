package match

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
			expected: false,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			containsAll, err := NewKeyword(ContainsAllKeywords)

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expected, containsAll.IsMatch(testCase.words, testCase.input))
		})
	}
}
