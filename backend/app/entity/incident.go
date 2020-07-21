package entity

import "time"

// Incident represents an incident
type Incident struct {
	ID        string
	Title     string
	CreatedAt *time.Time
}
