package entity

import "time"

type Url struct {
	Entity
	Id          string
	OriginalUrl string
	ExpireAt    *time.Time
	CreatedBy   User
}
