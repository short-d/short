package changelog

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/repository"
)

// ChangeLog retrieves change log and create changes.
type ChangeLog interface {
	CreateChange(title string, summaryMarkdown string) (entity.Change, error)
	GetChangeLog() ([]entity.Change, error)
}

// Persist retrieves change log from and saves changes to persistent data store.
type Persist struct {
	keyGen        keygen.KeyGenerator
	timer         fw.Timer
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
		ReleasedAt:      &now,
	}
	return p.changeLogRepo.CreateChange(newChange)
}

// GetChangeLog retrieves full ChangeLog from persistent data store.
func (p Persist) GetChangeLog() ([]entity.Change, error) {
	return p.changeLogRepo.GetChangeLog()
}

// NewPersist creates Persist
func NewPersist(
	keyGen keygen.KeyGenerator,
	timer fw.Timer,
	changeLog repository.ChangeLog,
) Persist {
	return Persist{
		keyGen:        keyGen,
		timer:         timer,
		changeLogRepo: changeLog,
	}
}
