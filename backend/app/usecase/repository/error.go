package repository

// ErrEntryNotFound represents no entry found in the repository error
var _ error = (*ErrEntryNotFound)(nil)

type ErrEntryNotFound string

func (e ErrEntryNotFound) Error() string {
	return string(e)
}
