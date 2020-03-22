package repository

import (
	"time"

	"github.com/short-d/short/app/entity"
)

// Changelog accesses changelog information from storage, such as database.
type Changelog interface {
	GetComplete() ([]entity.Changelog, error)
	CreateOne(id string, title string, summaryMarkdown string, releasedAt time.Time) (entity.Changelog, error)
}
