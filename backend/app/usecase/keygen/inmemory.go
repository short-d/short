package keygen

import uuid "github.com/satori/go.uuid"

var _ KeyGenerator = (*InMemory)(nil)

type InMemory struct {
}

func (InMemory) NewKey() string {
	return uuid.NewV4().String()
}

func NewInMemory() KeyGenerator {
	return InMemory{}
}
