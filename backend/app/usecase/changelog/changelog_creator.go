package changelog

import (
	"time"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ Creator = (*CreatorPersist)(nil)

type Creator interface {
	CreateChange(id string, title string, summaryMarkdown string) (entity.Changelog, error)
}

type CreatorPersist struct {
	changelogRepo repository.Changelog
}

func (c CreatorPersist) CreateChange(id string, title string, summaryMarkdown string) (entity.Changelog, error) {
	currentTime := time.Now()
	change, err := c.changelogRepo.CreateOne(id, title, summaryMarkdown, currentTime)
	if err != nil {
		return entity.Changelog{}, err
	}
	return change, nil
}

func NewCreatorPersist(changelog repository.Changelog) CreatorPersist {
	return CreatorPersist{changelog}
}
