package db

// SQLBool converts Go lang bool to the corresponding SQL data type representing boolean.
const (
	true  int = 1
	false int = 0
)

func SQLBool(value bool) int {
	if value {
		return true
	}
	return false
}
