package changelog

import (
	"time"

	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ ChangeLog = (*Persist)(nil)

// ErrUnauthorizedAction represents unauthorized action error
type ErrUnauthorizedAction string

func (e ErrUnauthorizedAction) Error() string {
	return string(e)
}

// ChangeLog retrieves change log and create changes.
type ChangeLog interface {
	CreateChange(title string, summaryMarkdown *string, user entity.User) (entity.Change, error)
	GetChangeLog() ([]entity.Change, error)
	GetLastViewedAt(user entity.User) (*time.Time, error)
	ViewChangeLog(user entity.User) (time.Time, error)
	DeleteChange(id string, user entity.User) error
	UpdateChange(id string, title string, summaryMarkdown *string, user entity.User) (entity.Change, error)
}

// Persist retrieves change log from and saves changes to persistent data store.
type Persist struct {
	keyGen            keygen.KeyGenerator
	timer             timer.Timer
	changeLogRepo     repository.ChangeLog
	userChangeLogRepo repository.UserChangeLog
	authorizer        authorizer.Authorizer
}

// CreateChange creates a new change in the data store.
func (p Persist) CreateChange(title string, summaryMarkdown *string, user entity.User) (entity.Change, error) {
	canCreateChange, err := p.authorizer.CanCreateChange(user)
	if err != nil {
		return entity.Change{}, err
	}

	if !canCreateChange {
		return entity.Change{}, ErrUnauthorizedAction("Unauthorized action: create change logs")
	}

	now := p.timer.Now().UTC()
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
func (p Persist) GetLastViewedAt(user entity.User) (*time.Time, error) {
	lastViewedAt, err := p.userChangeLogRepo.GetLastViewedAt(user)
	if err == nil {
		return &lastViewedAt, nil
	}

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case repository.ErrEntryNotFound:
		return nil, nil
	}

	return nil, err
}

// ViewChangeLog records the time when the user viewed the change log
func (p Persist) ViewChangeLog(user entity.User) (time.Time, error) {
	now := p.timer.Now().UTC()
	lastViewedAt, err := p.userChangeLogRepo.UpdateLastViewedAt(user, now)
	if err == nil {
		return lastViewedAt, nil
	}

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case repository.ErrEntryNotFound:
		err = p.userChangeLogRepo.CreateRelation(user, now)
		if err != nil {
			return time.Time{}, err
		}

		return now, nil
	}

	return time.Time{}, err
}

// DeleteChange removes the change with given id
func (p Persist) DeleteChange(id string, user entity.User) error {
	canDeleteChange, err := p.authorizer.CanDeleteChange(user)
	if err != nil {
		return err
	}

	if !canDeleteChange {
		return ErrUnauthorizedAction("Unauthorized action: delete change logs")
	}

	return p.changeLogRepo.DeleteChange(id)
}

// UpdateChange updates an existing change with given id in data store.
func (p Persist) UpdateChange(
	id string,
	title string,
	summaryMarkdown *string,
	user entity.User,
) (entity.Change, error) {
	canUpdateChange, err := p.authorizer.CanUpdateChange(user)
	if err != nil {
		return entity.Change{}, err
	}

	if !canUpdateChange {
		return entity.Change{}, ErrUnauthorizedAction("Unauthorized action: update change logs")
	}

	newChange := entity.Change{
		ID:              id,
		Title:           title,
		SummaryMarkdown: summaryMarkdown,
	}
	return p.changeLogRepo.UpdateChange(newChange)
}

// NewPersist creates Persist
func NewPersist(
	keyGen keygen.KeyGenerator,
	timer timer.Timer,
	changeLog repository.ChangeLog,
	userChangeLog repository.UserChangeLog,
	authorizer authorizer.Authorizer,
) Persist {
	return Persist{
		keyGen:            keyGen,
		timer:             timer,
		changeLogRepo:     changeLog,
		userChangeLogRepo: userChangeLog,
		authorizer:        authorizer,
	}
}
