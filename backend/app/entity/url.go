package entity

import "time"

// URL represents a short link.
type URL struct {
	Alias       string
	OriginalURL string
	ExpireAt    *time.Time
	CreatedBy   *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	MetaTags
}
