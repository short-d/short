package account

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

// Provider providers user account service.
type Provider struct {
	userRepo repository.User
	timer    fw.Timer
}

// IsAccountExist checks whether an user account exist.
func (r Provider) IsAccountExist(email string) (bool, error) {
	return r.userRepo.IsEmailExist(email)
}

// CreateAccount creates an user account.
func (r Provider) CreateAccount(email string, name string) error {
	now := r.timer.Now()
	user := entity.User{
		Email:     email,
		Name:      name,
		CreatedAt: &now,
	}
	return r.userRepo.CreateUser(user)
}

// NewProvider creates user account service provider.
func NewProvider(userRepo repository.User, timer fw.Timer) Provider {
	return Provider{
		userRepo: userRepo,
		timer:    timer,
	}
}
