package entity

import (
	"time"

	"github.com/short-d/short/backend/app/entity/metatag"
)

// ShortLink represents a short link.
type ShortLink struct {
	Alias         string
	LongLink      string
	ExpireAt      *time.Time
	CreatedBy     *User
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	OpenGraphTags metatag.OpenGraph
	TwitterTags   metatag.Twitter
}
