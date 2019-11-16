package repo

import "short/app/entity"

var _ UserURLRelation = (*UserURLRelationFake)(nil)

// UserURLRelationFake represents in memory implementation of User-URL relationship accessor.
type UserURLRelationFake struct {
}

// CreateRelation creates many to many relationship between User and URL.
func (u UserURLRelationFake) CreateRelation(user entity.User, url entity.URL) error {
	return nil
}

// NewUserURLRepoFake creates UserURLFake
func NewUserURLRepoFake() UserURLRelationFake {
	return UserURLRelationFake{}
}
