package db

func SQLBool(value bool) int {
	if value {
		return 1
	}
	return 0
}
