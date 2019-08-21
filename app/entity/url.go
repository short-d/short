package entity

import "time"

type URL struct {
	Alias       string
	OriginalURL string
	ExpireAt    *time.Time
	CreatedBy   *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
