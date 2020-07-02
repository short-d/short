package filesystem

import "io/ioutil"

var _ FileSystem = (*Local)(nil)

// Local saves and reads files from local disk.
type Local struct {
}

// ReadFile reads the content of the given file from local disk.
func (l Local) ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

// NewLocal creates Local file system.
func NewLocal() Local {
	return Local{}
}
