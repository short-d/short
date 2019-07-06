package entity

import "time"

type UrlAlias struct {
	Entity
	Id          string
	OriginalUrl string
	ExpireAt    time.Time
	CreatedBy   User
}
