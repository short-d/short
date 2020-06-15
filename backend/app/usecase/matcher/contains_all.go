package matcher

import "strings"

// ContainsAll checks whether the input contains all the elements in the list.
func ContainsAll(words []string, input string) bool {
	for _, word := range words {
		if !strings.Contains(input, word) {
			return false
		}
	}
	return true
}
