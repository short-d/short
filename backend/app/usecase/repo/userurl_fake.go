package repo

var _ UserURLRelation = (*UserURLRelationFake)(nil)

// UserURLFake represents in memory implementation of User-URL relationship accessor.
type UserURLRelationFake struct {
}

// CreateRelation creates many to many relationship between User and URL.
func (u UserURLRelationFake) CreateRelation(userEmail string, urlAlias string) error {
	return nil
}

// NewUserURLRepoFake creates UserURLFake
func NewUserURLRepoFake() UserURLRelationFake {
	return UserURLRelationFake{}
}
