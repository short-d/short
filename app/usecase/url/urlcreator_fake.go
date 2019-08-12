package url

import "short/app/entity"

type FakeCreator struct {
	urls map[string]entity.Url
}

func (f FakeCreator) Create(url entity.Url) (entity.Url, error) {
	panic("implement me")
}

func (f FakeCreator) CreateWithCustomAlias(url entity.Url, alias string) (entity.Url, error) {
	panic("implement me")
}

func NewCreatorFake(urls map[string]entity.Url) Creator {
	return FakeCreator{urls: urls}
}
