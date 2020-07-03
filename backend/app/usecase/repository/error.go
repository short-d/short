package repository

// ErrEntryNotFound represents no entry found in the repository error
var _ error = (*ErrEntryNotFound)(nil)

// ErrEntryNotFound represents table entry not found error.
type ErrEntryNotFound string

// Error coverts ErrEntryNotFound into human readable message to easy debugging.
func (e ErrEntryNotFound) Error() string {
	return string(e)
}

var _ error = (*ErrEntryExists)(nil)

// ErrEntryExists represents table entry already exists error.
type ErrEntryExists string

// Error coverts ErrEntryExists into human readable message to easy debugging.
func (e ErrEntryExists) Error() string {
	return string(e)
}
