package validator

import "regexp"

const longLinkMaxLength = 200

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

func NewLongLink() LongLink {
	uriPattern := regexp.MustCompile(`^[a-zA-Z]+://.+$`)
	return LongLink{
		uriPattern: uriPattern,
	}
}
