package usecase

import (
	"tinyURL/app/entity"
	"tinyURL/app/repo"
)

func NewUrlRetrieverFake(urls map[string]entity.Url) UrlRetriever {
	repoFake := repo.NewUrlFake(urls)
	return NewUrlRetrieverRepo(repoFake)
}
