package sqldb

const (
	sqlTrue  int = 1
	sqlFalse int = 0
)

// SQLBool converts Go lang bool to the corresponding SQL data type representing boolean.
func SQLBool(value bool) int {
	if value {
		return sqlTrue
	}
	return sqlFalse
}
