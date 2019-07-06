package entity

import "time"

type User struct {
	Entity
	Name           string
	Email          string
	LastLoggedInAt *time.Time
}
