package changelog

import (
	"time"

	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ ChangeLog = (*Persist)(nil)

// ChangeLog retrieves change log and create changes.
type ChangeLog interface {
	CreateChange(title string, summaryMarkdown *string) (entity.Change, error)
	GetChangeLog() ([]entity.Change, error)
	GetLastViewedAt() *time.Time
}

// Persist retrieves change log from and saves changes to persistent data store.
type Persist struct {
	keyGen        keygen.KeyGenerator
	timer         timer.Timer
	changeLogRepo repository.ChangeLog
}

// CreateChange creates a new change in the data store.
func (p Persist) CreateChange(title string, summaryMarkdown *string) (entity.Change, error) {
	now := p.timer.Now()
	key, err := p.keyGen.NewKey()
	if err != nil {
		return entity.Change{}, err
	}
	newChange := entity.Change{
		ID:              string(key),
		Title:           title,
		SummaryMarkdown: summaryMarkdown,
		ReleasedAt:      now,
	}
	return p.changeLogRepo.CreateChange(newChange)
}

// GetChangeLog retrieves full ChangeLog from persistent data store.
func (p Persist) GetChangeLog() ([]entity.Change, error) {
	return p.changeLogRepo.GetChangeLog()
}

// GetLastViewedAt retrieves the last time the user viewed the change log
// TODO(issue#613): fetch the last time the user viewed the change log from persistent storage.
func (p Persist) GetLastViewedAt() *time.Time {
	now := p.timer.Now()
	return &now
}

// NewPersist creates Persist
func NewPersist(
	keyGen keygen.KeyGenerator,
	timer timer.Timer,
	changeLog repository.ChangeLog,
) Persist {
	return Persist{
		keyGen:        keyGen,
		timer:         timer,
		changeLogRepo: changeLog,
	}
}
