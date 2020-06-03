package entity

import "time"

// ShortLink represents a short link.
type ShortLink struct {
	Alias       string
	LongLink 	string
	ExpireAt    *time.Time
	CreatedBy   *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	MetaTags
}
