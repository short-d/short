package repository

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
)

var _ ChangeLog = (*ChangeLogFake)(nil)

// ChangeLogFake represents in memory implementation of ChangeLog repository
type ChangeLogFake struct {
	changeLog []entity.Change
}

// GetChangeLog fetches full ChangeLog from memory
func (c ChangeLogFake) GetChangeLog() ([]entity.Change, error) {
	return c.changeLog, nil
}

// CreateChange creates and persists new Change in the repository
func (c *ChangeLogFake) CreateChange(newChange entity.Change) (entity.Change, error) {
	for _, change := range c.changeLog {
		if change.ID == newChange.ID {
			return entity.Change{}, errors.New("change exists")
		}
	}

	c.changeLog = append(c.changeLog, newChange)
	return newChange, nil
}

// DeleteChange removes a change based on a given id
func (c *ChangeLogFake) DeleteChange(id string) error {
	for idx, change := range c.changeLog {
		if change.ID == id {
			return c.removeChangeAt(idx)
		}
	}

	return nil
}

func (c *ChangeLogFake) removeChangeAt(idx int) error {
	if idx < 0 || idx >= len(c.changeLog) {
		return errors.New("index not in range for removing change")
	}

	c.changeLog = append(c.changeLog[:idx], c.changeLog[idx+1:]...)
	return nil
}

func (c *ChangeLogFake) UpdateChange(newChange entity.Change) (entity.Change, error) {
	for idx, change := range c.changeLog {
		if change.ID == newChange.ID {
			c.changeLog[idx].Title = newChange.Title
			c.changeLog[idx].SummaryMarkdown = newChange.SummaryMarkdown
			break
		}
	}

	return newChange, nil
}

// NewChangeLogFake creates ChangeLogFake
func NewChangeLogFake(changeLog []entity.Change) ChangeLogFake {
	return ChangeLogFake{
		changeLog: changeLog,
	}
}
