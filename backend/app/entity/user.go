package entity

import "time"

type User struct {
	Name           string
	Email          string
	LastSignedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
