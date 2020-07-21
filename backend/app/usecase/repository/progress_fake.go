package repository

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
)

var _ Progress = (*ProgressFake)(nil)

type ProgressFake struct {
	progresses []entity.Progress
}

// GetProgress finds a Progress in the progress table given a Progress ID
func (p ProgressFake) GetProgress(progressID string) (entity.Progress, error) {
	for _, progress := range p.progresses {
		if progress.ID == progressID {
			return progress, nil
		}
	}
	return entity.Progress{}, errors.New("progress not found")
}

func NewProcessFake(progresses []entity.Progress) ProgressFake {
	return ProgressFake{progresses: progress}
}
