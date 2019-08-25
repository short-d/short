package repo

var _ UserURL = (*UserURLFake)(nil)

// UserURLFake represents in memory implementation of User-URL relationship accessor.
type UserURLFake struct {
}

// CreateRelation creates many to many relationship between User and URL.
func (u UserURLFake) CreateRelation(userEmail string, urlAlias string) error {
	return nil
}

// NewUserURLRepoFake creates UserURLFake
func NewUserURLRepoFake() UserURLFake {
	return UserURLFake{}
}
