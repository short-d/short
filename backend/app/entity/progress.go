package entity

import (
	"time"
)

// Progress represents an Progress
type Progress struct {
	ID        string
	Incident  *Incident
	Status    string
	Info      string
	CreatedAt *time.Time
}
