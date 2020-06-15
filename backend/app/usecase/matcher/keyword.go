package matcher

import (
	"fmt"
)

// KeywordType represents keyword matching type.
type KeywordType int

const (
	// ContainsAllKeywords represents ContainsAll type
	ContainsAllKeywords KeywordType = iota
	// ContainsAnyKeyword represents ContainsAny type
	ContainsAnyKeyword
)

// Keyword matches a list of words against an input.
type Keyword interface {
	IsMatch(words []string, input string) bool
}

// NewKeyword creates Keyword.
func NewKeyword(keywordType KeywordType) (Keyword, error) {
	switch keywordType {
	case ContainsAllKeywords:
		return new(ContainsAll), nil
	case ContainsAnyKeyword:
		return new(ContainsAny), nil
	default:
		return nil, fmt.Errorf("keyword matching type %d not recognized", keywordType)
	}
}
