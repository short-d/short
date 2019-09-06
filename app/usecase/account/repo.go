package account

import (
	"short/app/entity"
	"short/app/usecase/repo"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

var _ service.Account = (*RepoService)(nil)

type RepoService struct {
	userRepo repo.User
	timer    fw.Timer
}

func (r RepoService) IsAccountExist(email string) (bool, error) {
	return r.userRepo.IsEmailExist(email)
}

func (r RepoService) CreateAccount(email string, name string) error {
	now := r.timer.Now()
	user := entity.User{
		Email:     email,
		Name:      name,
		CreatedAt: &now,
	}
	return r.userRepo.Create(user)
}

func NewRepoService(userRepo repo.User, timer fw.Timer) RepoService {
	return RepoService{
		userRepo: userRepo,
		timer:    timer,
	}
}
