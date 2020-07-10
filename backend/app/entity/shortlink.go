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

// ShortLinkInput represents possible ShortLink attributes
type ShortLinkInput struct {
	LongLink    *string
	CustomAlias *string
	ExpireAt    *time.Time
}

// GetLongLink fetches LongLink for ShortLinkInput, with empty string as the default value.
func (s *ShortLinkInput) GetLongLink() string {
	if s.LongLink == nil {
		return ""
	}
	return *s.LongLink
}

// GetCustomAlias fetches CustomAlias for ShortLinkInput, with empty string as the default value.
func (s *ShortLinkInput) GetCustomAlias() string {
	if s.CustomAlias == nil {
		return ""
	}
	return *s.CustomAlias
}

// GetShortLink creates a ShortLink entity instance using ShortLinkInput.
func (s *ShortLinkInput) GetShortLink() ShortLink {
	return ShortLink{
		Alias:    s.GetCustomAlias(),
		LongLink: s.GetLongLink(),
		ExpireAt: s.ExpireAt,
	}
}
