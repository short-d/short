package changelog

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/keygen"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

type ChangeLog interface {
	CreateChange(title string, summaryMarkdown string) (entity.Change, error)
	GetChangeLog() ([]entity.Change, error)
}

type Persist struct {
	keyGen        keygen.KeyGenerator
	timer         fw.Timer
	changeLogRepo repository.ChangeLog
}

func (p Persist) CreateChange(title string, summaryMarkdown string) (entity.Change, error) {
	now := p.timer.Now()
	key, err := p.keyGen.NewKey()
	if err != nil {
		return entity.Change{}, err
	}
	newChange := entity.Change{ID: string(key), Title: title, SummaryMarkdown: summaryMarkdown, ReleasedAt: &now}
	change, err := p.changeLogRepo.CreateChange(newChange)
	if err != nil {
		return entity.Change{}, err
	}
	return change, nil
}

func (p Persist) GetChangeLog() ([]entity.Change, error) {
	changeLog, err := p.changeLogRepo.GetChangeLog()
	if err != nil {
		return nil, err
	}
	return changeLog, nil
}

func NewPersist(keyGen keygen.KeyGenerator, timer fw.Timer, changeLog repository.ChangeLog) Persist {
	return Persist{keyGen, timer, changeLog}
}
