package url

import (
	"fmt"
	"short/app/entity"
	"short/app/usecase/keygen"
)

var _ Creator = (*FakeCreator)(nil)

// FakeCreator represents in-memory url creator
type FakeCreator struct {
	urls   map[string]entity.URL
	keyGen keygen.Fake
}

// CreateURL persists a new url with a generated alias in the repository.
func (f FakeCreator) CreateURL(url entity.URL, user entity.User) (entity.URL, error) {
	key, err := f.keyGen.NewKey()
	if err != nil {
		return entity.URL{}, err
	}
	randomAlias := string(key)
	return f.CreateURLWithCustomAlias(url, randomAlias, user)
}

// CreateURLWithCustomAlias persists a new url with a custom alias in
// the repository.
func (f FakeCreator) CreateURLWithCustomAlias(
	url entity.URL,
	alias string,
	user entity.User,
) (entity.URL, error) {
	url.Alias = alias

	fmt.Println(f.urls)

	_, ok := f.urls[alias]
	if ok {
		return entity.URL{}, ErrAliasExist("usecase: url alias already exist")
	}

	f.urls[alias] = url
	return url, nil
}

// NewCreatorFake creates in-memory url creator
func NewCreatorFake(
	urls map[string]entity.URL,
	availableAlias []string,
) FakeCreator {
	return FakeCreator{
		urls:   urls,
		keyGen: keygen.NewFake(availableAlias),
	}
}
