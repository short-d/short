package matcher

import "strings"

// ContainsAny checks whether the input contains any element of the list.
func ContainsAny(words []string, input string) bool {
	for _, word := range words {
		if strings.Contains(input, word) {
			return true
		}
	}
	return false
}
