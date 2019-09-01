package account

import (
	"short/app/entity"
	"short/app/usecase/repo"
	"short/app/usecase/service"

	"short/fw"
)

var _ service.Account = (*Repo)(nil)

type Repo struct {
	userRepo repo.User
	timer    fw.Timer
}

func (r Repo) IsAccountExist(email string) (bool, error) {
	return r.userRepo.IsEmailExist(email)
}

func (r Repo) CreateAccount(email string, name string) error {
	now := r.timer.Now()
	user := entity.User{
		Email:     email,
		Name:      name,
		CreatedAt: &now,
	}
	return r.userRepo.Create(user)
}

func NewRepoService(userRepo repo.User, timer fw.Timer) Repo {
	return Repo{
		userRepo: userRepo,
		timer:    timer,
	}
}
