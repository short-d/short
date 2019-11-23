package validator

import "regexp"

const (
	customAliasMaxLength = 50
)

type CustomAlias struct {
	uriPattern *regexp.Regexp
}

func (c CustomAlias) IsValid(alias *string) bool {
	if alias == nil {
		return true
	}

	if *alias == "" {
		return true
	}

	if len(*alias) >= customAliasMaxLength {
		return false
	}
	return true
}

func NewCustomAlias() CustomAlias {
	return CustomAlias{}
}
