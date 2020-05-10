package repository

// ErrEntryNotFound represents no entry found in the repository error
var _ error = (*ErrEntryNotFound)(nil)

// ErrEntryNotFound represents table entry not found error.
type ErrEntryNotFound string

func (e ErrEntryNotFound) Error() string {
	return string(e)
}
