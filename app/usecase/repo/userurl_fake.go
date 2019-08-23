package repo

var _ UserURL = (*UserURLFake)(nil)

type UserURLFake struct {
}

func (u UserURLFake) CreateRelation(userEmail string, urlAlias string) error {
	return nil
}

func NewUserURLRepoFake() UserURLFake {
	return UserURLFake{}
}
