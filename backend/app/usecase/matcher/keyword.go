package matcher

// KeywordType represents the type of keyword matching.
type KeywordType int

const (
	ContainsAllKeywords KeywordType = iota
	ContainsAnyKeyword
)

// Keyword matches the slice of words against the input.
type Keyword interface {
	IsMatch(words []string, input string) bool
}

// NewKeyword creates Keyword.
func NewKeyword(keywordType KeywordType) Keyword {
	switch keywordType {
	case ContainsAllKeywords:
		return ContainsAll{}
	case ContainsAnyKeyword:
		return ContainsAny{}
	default:
		return ContainsAll{}
	}
}
