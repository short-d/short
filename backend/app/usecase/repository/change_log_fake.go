package repository

import (
	"errors"

	"github.com/short-d/short/app/entity"
)

var _ ChangeLog = (*ChangeLogFake)(nil)

// UserFake represents in memory implementation of user repository.
type ChangeLogFake struct {
	changeLog []entity.Change
}

func (c ChangeLogFake) GetChangeLog() ([]entity.Change, error) {
	return c.changeLog, nil
}

func (c *ChangeLogFake) CreateChange(newChange entity.Change) (entity.Change, error) {
	for _, change := range c.changeLog {
		if change.ID == newChange.ID {
			return entity.Change{}, errors.New("change exists")
		}
	}

	c.changeLog = append(c.changeLog, newChange)
	return newChange, nil
}

// NewUserFake create in memory user repository implementation.
func NewChangeLogFake(changeLog []entity.Change) ChangeLogFake {
	return ChangeLogFake{
		changeLog: changeLog,
	}
}
