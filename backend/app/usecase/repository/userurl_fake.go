package repository

import (
	"database/sql"
	"errors"
	"short/app/entity"
)

var _ UserURLRelation = (*UserURLRelationFake)(nil)

// UserURLRelationFake represents in memory implementation of User-URL relationship accessor.
type UserURLRelationFake struct {
	users []entity.User
	urls  []entity.URL
}

func (u *UserURLRelationFake) NewTransaction() (*sql.Tx, error) {
	panic("implement me")
}

func (u *UserURLRelationFake) CreateRelationWithTransaction(tx *sql.Tx, user entity.User, url entity.URL) error {
	panic("implement me")
}

// CreateRelation creates many to many relationship between User and URL.
func (u *UserURLRelationFake) CreateRelation(user entity.User, url entity.URL) error {
	if u.IsRelationExist(user, url) {
		return errors.New("relationship exists")
	}
	u.users = append(u.users, user)
	u.urls = append(u.urls, url)
	return nil
}

// FindAliasesByUser fetches the aliases of all the URLs created by the given user.
func (u UserURLRelationFake) FindAliasesByUser(user entity.User) ([]string, error) {
	var aliases []string
	for idx, currUser := range u.users {
		if currUser.ID != user.ID {
			continue
		}
		aliases = append(aliases, u.urls[idx].Alias)
	}
	return aliases, nil
}

// IsRelationExist checks whether the an URL is own by a given user.
func (u UserURLRelationFake) IsRelationExist(user entity.User, url entity.URL) bool {
	for idx, currUser := range u.users {
		if currUser.ID != user.ID {
			continue
		}

		if u.urls[idx].Alias == url.Alias {
			return true
		}
	}
	return false
}

// NewUserURLRepoFake creates UserURLFake
func NewUserURLRepoFake(users []entity.User, urls []entity.URL) UserURLRelationFake {
	return UserURLRelationFake{
		users: users,
		urls:  urls,
	}
}
