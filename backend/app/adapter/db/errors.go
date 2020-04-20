package db

import "errors"

var errRecordNotFound = errors.New("user does not exist")

// IsErrRecordNotFound checks whether the error is ErrRecordNotFound.
func IsErrRecordNotFound(err error) bool {
	return errors.Is(err, errRecordNotFound)
}
