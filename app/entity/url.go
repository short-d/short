package entity

import "time"

type Url struct {
	Entity
	Alias          string
	OriginalUrl string
	ExpireAt    *time.Time
	CreatedBy   User
}
