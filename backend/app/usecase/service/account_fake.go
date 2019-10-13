package service

var _ Account = (*AccountFake)(nil)

type AccountFake struct {
	isAccountExist    bool
	isAccountExistErr error
	createAccountErr  error
}

func (a AccountFake) IsAccountExist(email string) (bool, error) {
	return a.isAccountExist, a.isAccountExistErr
}

func (a AccountFake) CreateAccount(email string, name string) error {
	return a.createAccountErr
}

func NewAccountFake(
	isAccountExist bool,
	isAccountExistErr error,
	createAccountErr error,
) AccountFake {
	return AccountFake{
		isAccountExist:    isAccountExist,
		isAccountExistErr: isAccountExistErr,
		createAccountErr:  createAccountErr,
	}
}
