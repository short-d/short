package repository

import (
	"time"

	"github.com/short-d/short/app/entity"
)

// Change accesses changelog information from storage, such as database.
type Changelog interface {
	GetChangeLog() ([]entity.Change, error)
	CreateOne(id string, title string, summaryMarkdown string, releasedAt time.Time) (entity.Change, error)
}
