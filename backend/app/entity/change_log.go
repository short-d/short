package entity

import "time"

type Changelog struct {
	ID              string
	Title           string
	SummaryMarkdown string
	ReleasedAt      *time.Time
}
