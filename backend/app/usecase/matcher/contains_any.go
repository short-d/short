package matcher

import "strings"

var _ Keyword = (*ContainsAny)(nil)

type ContainsAny struct {
}

func (c ContainsAny) IsMatch(query string, input string) bool {
	if input == "" {
		return false
	} else if query == "" {
		return false
	}

	words := getKeywords(query)
	for _, word := range words {
		if strings.Contains(input, word) {
			return true
		}
	}
	return false
}
