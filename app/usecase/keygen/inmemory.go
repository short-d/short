package keygen

import uuid "github.com/satori/go.uuid"

type InMemory struct {
}

func (InMemory) NewKey() string {
	return uuid.NewV4().String()
}

func NewInMemory() KeyGenerator {
	return InMemory{}
}
