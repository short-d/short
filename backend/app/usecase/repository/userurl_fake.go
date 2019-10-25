package repository

import "short/app/entity"

var _ UserURLRelation = (*UserURLRelationFake)(nil)

// UserURLRelationFake represents in memory implementation of User-URL relationship accessor.
type UserURLRelationFake struct {
	userURLRelations map[string]string
}

// CreateRelation creates many to many relationship between User and URL.
func (u UserURLRelationFake) CreateRelation(user entity.User, url entity.URL) error {
	u.userURLRelations[url.Alias] = user.Email
	return nil
}

// FindAliasesByUser fetches all the aliases corresponding to the given user.
func (u UserURLRelationFake) FindAliasesByUser(user entity.User) ([]string, error) {
	var aliases []string
	for alias := range u.userURLRelations {
		if u.userURLRelations[alias] == user.Email {
			aliases = append(aliases, alias)
		}
	}
	return aliases, nil
}

// NewUserURLRepoFake creates UserURLFake
func NewUserURLRepoFake(userURLRelations map[string]string) UserURLRelationFake {
	if userURLRelations == nil {
		userURLRelations = make(map[string]string)
	}

	return UserURLRelationFake{
		userURLRelations: userURLRelations,
	}
}
