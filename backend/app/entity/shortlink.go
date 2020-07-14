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

// ShortLinkInput represents possible ShortLink attributes for a short link.
type ShortLinkInput struct {
	LongLink    *string
	CustomAlias *string
	ExpireAt    *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// GetLongLink fetches LongLink for ShortLinkInput with default value.
func (s *ShortLinkInput) GetLongLink(defaultVal string) string {
	if s.LongLink == nil {
		return defaultVal
	}
	return *s.LongLink
}

// GetCustomAlias fetches CustomAlias for ShortLinkInput with default value.
func (s *ShortLinkInput) GetCustomAlias(defaultVal string) string {
	if s.CustomAlias == nil {
		return defaultVal
	}
	return *s.CustomAlias
}
