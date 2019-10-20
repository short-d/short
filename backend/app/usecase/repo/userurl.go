package repo

// UserURLRelation accesses User-URL relationship from storage, such as database.
type UserURLRelation interface {
	CreateRelation(userEmail string, urlAlias string) error
}
