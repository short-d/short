package matcher

import "strings"

type KeywordMatcher int

const (
	MatchAll KeywordMatcher = iota
	MatchAny
)

type Keyword interface {
	IsMatch(query string, input string) bool
}

func getKeywords(query string) []string {
	return strings.Split(query, " ")
}

func NewKeyword(matcher KeywordMatcher) Keyword {
	switch matcher {
	case MatchAll:
		return ContainsAll{}
	case MatchAny:
		return ContainsAny{}
	default:
		return ContainsAll{}
	}
}
