package usecase

import (
	"short/app/entity"
	"short/app/repo"
)

func NewUrlRetrieverFake(urls map[string]entity.Url) UrlRetriever {
	repoFake := repo.NewUrlFake(urls)
	return NewUrlRetrieverPersist(repoFake)
}
