package input

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

type LongLinkValidator struct {
	uriPattern *regexp.Regexp
}

func (l LongLinkValidator) IsValid(longLink *string) bool {
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

func NewLongLinkValidator() Validator {
	uriPattern := regexp.MustCompile(`^[a-zA-Z]+://.+$`)
	return LongLinkValidator{
		uriPattern: uriPattern,
	}
}

type CustomAliasValidator struct {
	uriPattern *regexp.Regexp
}

func (c CustomAliasValidator) IsValid(alias *string) bool {
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

func NewCustomAliasValidator() Validator {
	return CustomAliasValidator{}
}
