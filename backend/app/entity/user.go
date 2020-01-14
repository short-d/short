package entity

import "time"

// User contains basic user information such as, user ID, name, and email.
type User struct {
	ID             string
	Name           string
	Email          string
	LastSignedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
