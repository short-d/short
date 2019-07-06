package entity

import "time"

type Entity struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}
