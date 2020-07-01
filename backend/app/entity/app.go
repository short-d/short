package entity

import "time"

// App represents an application owned by third party developer.
type App struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
