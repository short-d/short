package match

import "strings"

var _ Keyword = (*ContainsAll)(nil)

// ContainsAll checks if an input contains all the elements of a list.
type ContainsAll struct {
}

// IsMatch checks if the input contains all the words.
func (c ContainsAll) IsMatch(words []string, input string) bool {
	for _, word := range words {
		if !strings.Contains(input, word) {
			return false
		}
	}
	return true
}
