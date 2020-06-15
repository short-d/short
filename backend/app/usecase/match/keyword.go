package match

import (
	"fmt"
)

// Algorithm represents keyword matching type.
type Algorithm int

const (
	// ContainsAllKeywords represents ContainsAll type
	ContainsAllKeywords Algorithm = iota
	// ContainsAnyKeyword represents ContainsAny type
	ContainsAnyKeyword
)

// Keyword matches a list of words against an input.
type Keyword interface {
	IsMatch(words []string, input string) bool
}

// NewKeyword creates Keyword.
func NewKeyword(algorithm Algorithm) (Keyword, error) {
	switch algorithm {
	case ContainsAllKeywords:
		return new(ContainsAll), nil
	case ContainsAnyKeyword:
		return new(ContainsAny), nil
	default:
		return nil, fmt.Errorf("match algorithm %d not recognized", algorithm)
	}
}
