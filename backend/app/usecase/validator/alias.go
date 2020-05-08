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
func (c CustomAlias) IsValid(alias *string) bool {
	if alias == nil {
		return true
	}

	if *alias == "" {
		return true
	}

	if strings.ContainsRune(*alias, rune('#')) {
		return false
	}

	if len(*alias) >= customAliasMaxLength {
		return false
	}
	return true
}

// NewCustomAlias creates custom alias validator.
func NewCustomAlias() CustomAlias {
	return CustomAlias{}
}
