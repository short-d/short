package usecase

import (
	"tinyURL/app/entity"
	"tinyURL/app/repo"
	"tinyURL/modern/mdtest"
)

func NewUrlRetrieverFake(urls map[string]entity.Url) UrlRetriever {
	repoFake := repo.NewUrlFake(urls)
	tracerFake := mdtest.FakeTracer
	return NewUrlRetrieverRepo(tracerFake, repoFake)
}
