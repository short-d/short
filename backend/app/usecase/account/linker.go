package account

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repository"
)

// Linker provides account linking service.
type Linker struct {
	keyGen             keygen.KeyGenerator
	userRepo           repository.User
	accountMappingRepo repository.AccountMapping
}

// IsAccountLinked checks whether a given external account is linked to any
// internal users already.
func (l Linker) IsAccountLinked(ssoUser entity.SSOUser) (bool, error) {
	return l.accountMappingRepo.IsSSOUserExist(ssoUser)
}

// EnsureAndLinkAccount creates an internal account when there is no internal
// account sharing the same email as the given external account and link them
// together afterwards.
func (l Linker) EnsureAndLinkAccount(ssoUser entity.SSOUser) error {
	isAccountLinked, err := l.IsAccountLinked(ssoUser)
	if err != nil {
		return err
	}

	if isAccountLinked {
		return nil
	}

	user, err := l.ensureUserExist(ssoUser)
	if err != nil {
		return err
	}
	return l.accountMappingRepo.CreateMapping(ssoUser, user)
}

func (l Linker) ensureUserExist(ssoUser entity.SSOUser) (entity.User, error) {
	isEmailExist, err := l.userRepo.IsEmailExist(ssoUser.Email)
	if err != nil {
		return entity.User{}, err
	}
	userID, err := l.generateUnassignedUserID()
	if err != nil {
		return entity.User{}, err
	}

	if isEmailExist {
		err = l.assignUserID(ssoUser.Email, userID)
		return entity.User{ID: userID}, err
	}
	return l.createUser(userID, ssoUser.Name, ssoUser.Email)
}

func (l Linker) generateUnassignedUserID() (string, error) {
	newKey, err := l.keyGen.NewKey()
	return string(newKey), err
}

func (l Linker) createUser(id string, name string, email string) (entity.User, error) {
	user := entity.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	err := l.userRepo.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (l Linker) assignUserID(userEmail string, userID string) error {
	return l.userRepo.UpdateUserID(userEmail, userID)
}

// NewLinker creates a new account linking service.
func NewLinker(
	keyGen keygen.KeyGenerator,
	userRepo repository.User,
	accountMappingRepo repository.AccountMapping,
) Linker {
	return Linker{
		keyGen:             keyGen,
		userRepo:           userRepo,
		accountMappingRepo: accountMappingRepo,
	}
}
