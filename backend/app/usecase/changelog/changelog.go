package changelog

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/keygen"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

type ChangeLog interface {
	CreateChange(title string, summaryMarkdown string) (entity.Change, error)
	GetChangelog() ([]entity.Change, error)
}

type Persist struct {
	keyGen        keygen.KeyGenerator
	timer 		  fw.Timer
	changelogRepo repository.Changelog
}

func (p Persist) CreateChange(title string, summaryMarkdown string) (entity.Change, error) {
	now := p.timer.Now()
	key, err := p.keyGen.NewKey()
	if err != nil {
		return entity.Change{}, nil
	}
	newChange := entity.Change{ID: string(key), Title: title, SummaryMarkdown: summaryMarkdown, ReleasedAt: &now}
	change, err := p.changelogRepo.CreateOne(newChange)
	if err != nil {
		return entity.Change{}, err
	}
	return change, nil
}

func (p Persist) GetChangelog() ([]entity.Change, error) {
	changelog, err := p.changelogRepo.GetChangeLog()
	if err != nil {
		return nil, err
	}
	return changelog, nil
}

func NewPersist(keyGen keygen.KeyGenerator, timer fw.Timer, changelog repository.Changelog) Persist {
	return Persist{keyGen, timer,changelog}
}
