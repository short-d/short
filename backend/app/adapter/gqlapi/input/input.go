package input

import (
	"time"

	"github.com/short-d/short/backend/app/entity"
)

// ShortLinkInput represents possible ShortLink attributes
type ShortLinkInput struct {
	LongLink    *string
	CustomAlias *string
	ExpireAt    *time.Time
}

// CreateShortLinkInput will massage the GraphQL arguments into a form that use case will accept.
func (s ShortLinkInput) CreateShortLinkInput() entity.ShortLinkInput {
	var longLink, customAlias string

	if s.LongLink == nil {
		longLink = ""
	} else {
		longLink = *s.LongLink
	}

	if s.CustomAlias == nil {
		customAlias = ""
	} else {
		customAlias = *s.CustomAlias
	}

	return entity.ShortLinkInput{
		LongLink:    longLink,
		CustomAlias: customAlias,
		ExpireAt:    s.ExpireAt,
	}
}
