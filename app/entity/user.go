package entity

import "time"

type User struct {
	Name           string
	Email          string
	LastLoggedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
