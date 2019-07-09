package entity

import "time"

type Url struct {
	Alias       string
	OriginalUrl string
	ExpireAt    *time.Time
	CreatedBy   *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
