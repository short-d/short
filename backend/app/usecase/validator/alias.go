package validator

import (
	"regexp"
	"strings"

	"github.com/short-d/short/backend/app/entity"
)

const (
	customAliasMaxLength = 50
)

var forbiddenCharacters = map[rune]entity.Empty{
	'#': {},
}

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

	if c.hasForbiddenCharacter(*alias) {
		return false, HasFragmentCharacter
	}

	return true, Valid
}

// hasForbiddenCharacter returns whether the alias contains the one of forbidden character which starts
// fragment identifiers in URLs which starts fragment identifiers in URLs.
func (c CustomAlias) hasForbiddenCharacter(alias string) bool {
	for ch := range forbiddenCharacters {
		if strings.ContainsRune(alias, ch) {
			return true
		}
	}
	return false
}

// NewCustomAlias creates custom alias validator.
func NewCustomAlias() CustomAlias {
	return CustomAlias{}
}
