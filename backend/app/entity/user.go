package entity

import "time"

type User struct {
	ID             string
	Name           string
	Email          string
	LastSignedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
