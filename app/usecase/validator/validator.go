package validator

import (
	"regexp"
)

type Validator interface {
	IsValid(text *string) bool
}

const (
	longLinkMaxLength    = 200
	customAliasMaxLength = 50
)

type LongLink struct {
	uriPattern *regexp.Regexp
}

func (l LongLink) IsValid(longLink *string) bool {
	if longLink == nil {
		return false
	}

	if *longLink == "" {
		return false
	}

	if len(*longLink) >= longLinkMaxLength {
		return false
	}

	if !l.uriPattern.MatchString(*longLink) {
		return false
	}

	return true
}

func NewLongLink() Validator {
	uriPattern := regexp.MustCompile(`^[a-zA-Z]+://.+$`)
	return LongLink{
		uriPattern: uriPattern,
	}
}

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

func NewCustomAlias() Validator {
	return CustomAlias{}
}
