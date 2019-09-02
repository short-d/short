package inject

import (
	"short/app/usecase/account"
	"short/app/usecase/repo"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

func RepoAccount(userRepo repo.User, timer fw.Timer) service.Account {
	return account.NewRepoService(userRepo, timer)
}
