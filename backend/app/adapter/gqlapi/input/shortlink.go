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

// CreateShortLinkInput converts GraphQL ShortLinkInput into consumable entity for use cases.
func (s ShortLinkInput) CreateShortLinkInput() entity.ShortLinkInput {
	return entity.ShortLinkInput{
		LongLink:    s.LongLink,
		CustomAlias: s.CustomAlias,
		ExpireAt:    s.ExpireAt,
	}
}
