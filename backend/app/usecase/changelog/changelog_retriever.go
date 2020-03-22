package changelog

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ Retriever = (*RetrieverPersist)(nil)

type Retriever interface {
	GetChangelog() ([]entity.Changelog, error)
}

type RetrieverPersist struct {
	changelogRepo repository.Changelog
}

func (c RetrieverPersist) GetChangelog() ([]entity.Changelog, error) {
	changelog, err := c.changelogRepo.GetComplete()
	if err != nil {
		return nil, err
	}
	return changelog, nil
}

func NewRetrieverPersist(changelogRepo repository.Changelog) RetrieverPersist {
	return RetrieverPersist{changelogRepo}
}
