package match

import "strings"

var _ Keyword = (*ContainsAny)(nil)

// ContainsAny checks if an input contains any element of a list.
type ContainsAny struct {
}

// IsMatch checks if the input contains any word.
func (c ContainsAny) IsMatch(words []string, input string) bool {
	if len(words) == 0 {
		return true
	}
	for _, word := range words {
		if strings.Contains(input, word) {
			return true
		}
	}
	return false
}
