package repo

type UserURLRelation interface {
	CreateRelation(userEmail string, urlAlias string) error
}
