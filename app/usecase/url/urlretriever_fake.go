package url

import (
	"short/app/entity"
	"short/app/usecase/repo"
)

func NewRetrieverFake(urls map[string]entity.Url) Retriever {
	repoFake := repo.NewUrlFake(urls)
	return NewRetrieverPersist(repoFake)
}
