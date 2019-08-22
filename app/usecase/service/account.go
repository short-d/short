package service

type Account interface {
	IsAccountExist(email string) (bool, error)
	CreateAccount(email string, name string) error
}
