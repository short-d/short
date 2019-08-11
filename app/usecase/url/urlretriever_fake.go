package url

import (
	"short/app/adapter/repo"
	"short/app/entity"
)

func NewRetrieverFake(urls map[string]entity.Url) Retriever {
	repoFake := repo.NewUrlFake(urls)
	return NewRetrieverPersist(repoFake)
}
