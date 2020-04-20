package db

// SQLBool converts Go lang bool to the corresponding SQL data type representing boolean.
func SQLBool(value bool) int {
	if value {
		return 1
	}
	return 0
}
