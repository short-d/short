package usecase

import uuid "github.com/satori/go.uuid"

type KeyGenerator interface {
	NewKey() string
}

type KeyGenInMemory struct {
}

func (KeyGenInMemory) NewKey() string {
	return uuid.NewV4().String()
}

func NewKeyGenInMemory() KeyGenerator {
	return KeyGenInMemory{}
}
