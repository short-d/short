package filesystem

import "io/ioutil"

var _ FileSystem = (*Local)(nil)

type Local struct {
}

func (l Local) ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func NewLocal() Local {
	return Local{}
}
