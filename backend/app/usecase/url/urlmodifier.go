package url

import (
	"short/app/entity"
	"short/app/usecase/repository"
	"short/app/usecase/validator"
)

var _ Modifier = (*ModifierPersist)(nil)

// Modifier represents URL modifier
type Modifier interface {
	UpdateURL(oldAlias string, newAlias string, user entity.User) (entity.URL, error)
}

// ModifierPersist represents URL modifier that modifies URL from persistent
// storage, such as database
type ModifierPersist struct {
	urlRepo             repository.URL
	userURLRelationRepo repository.UserURLRelation
	aliasValidator      validator.CustomAlias
}

// UpdateURL update URL matching oldAlias with newAlias
func (r ModifierPersist) UpdateURL(oldAlias string, newAlias string, user entity.User) (entity.URL, error) {
	url, err := r.getURL(oldAlias)
	if err != nil {
		return entity.URL{}, err
	}

	url.Alias = newAlias

	err = r.urlRepo.Update(url)
	if err != nil {
		return entity.URL{}, err
	}

	err = r.userURLRelationRepo.UpdateRelation(user, url)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func (r ModifierPersist) getURL(alias string) (entity.URL, error) {
	url, err := r.urlRepo.GetByAlias(alias)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

// NewModifierPersist creates persistent URL modifier
func NewModifierPersist(
	urlRepo repository.URL,
	userURLRelationRepo repository.UserURLRelation,
	aliasValidator validator.CustomAlias,
) ModifierPersist {
	return ModifierPersist{
		urlRepo:             urlRepo,
		userURLRelationRepo: userURLRelationRepo,
		aliasValidator:      aliasValidator,
	}
}
