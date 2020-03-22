package entity

import "time"

type Change struct {
	ID              string
	Title           string
	SummaryMarkdown string
	ReleasedAt      *time.Time
}
