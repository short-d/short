package repo

type UserURL interface {
	CreateRelation(userEmail string, urlAlias string) error
}
