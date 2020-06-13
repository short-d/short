package matcher

import "strings"

var _ Keyword = (*ContainsAll)(nil)

type ContainsAll struct {
}

func (c ContainsAll) IsMatch(query string, input string) bool {
	if input == "" {
		return false
	} else if query == "" {
		return false
	}

	words := getKeywords(query)
	for _, word := range words {
		if !strings.Contains(input, word) {
			return false
		}
	}
	return true
}
