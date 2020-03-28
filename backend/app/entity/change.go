package entity

import "time"

// Change represents a single change in change log
type Change struct {
	ID              string
	Title           string
	SummaryMarkdown *string
	ReleasedAt      *time.Time
}
