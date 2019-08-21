package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
)

var _ Creator = (*FakeCreator)(nil)

type FakeCreator struct {
	urls   map[string]entity.Url
	keyGen keygen.Fake
}

func (f FakeCreator) Create(url entity.Url) (entity.Url, error) {
	randomAlias := f.keyGen.NewKey()
	return f.CreateWithCustomAlias(url, randomAlias)
}

func (f FakeCreator) CreateWithCustomAlias(url entity.Url, alias string) (entity.Url, error) {
	url.Alias = alias

	_, ok := f.urls[alias]
	if ok {
		return entity.Url{}, ErrAliasExist("usecase: url alias already exist")
	}

	f.urls[alias] = url
	return url, nil
}

func NewCreatorFake(urls map[string]entity.Url, availableAlias []string) FakeCreator {
	return FakeCreator{
		urls:   urls,
		keyGen: keygen.NewFake(availableAlias),
	}
}
