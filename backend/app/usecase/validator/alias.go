package validator

import (
	"regexp"
	"strings"
)

const (
	customAliasMaxLength = 50
)

// CustomAlias represents format validator for custom alias
type CustomAlias struct {
	uriPattern *regexp.Regexp
}

// IsValid checks whether the given alias has valid format.
func (c CustomAlias) IsValid(alias *string) (bool, Violation) {
	if alias == nil {
		return true, Valid
	}

	if *alias == "" {
		return true, Valid
	}

	if len(*alias) >= customAliasMaxLength {
		return false, AliasTooLong
	}

	if c.HasFragmentCharacter(*alias) {
		return false, HasFragmentCharacter
	}

	return true, Valid
}

// HasFragmentCharacter returns whether the alias contains the '#' character which starts fragment identifiers in URLs
func (c CustomAlias) HasFragmentCharacter(alias string) bool {
	return strings.ContainsRune(alias, rune('#'))
}

// NewCustomAlias creates custom alias validator.
func NewCustomAlias() CustomAlias {
	return CustomAlias{}
}
