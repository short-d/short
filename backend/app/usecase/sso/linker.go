package sso

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// AccountLinker maps external user accounts to Short user accounts.
type AccountLinker struct {
	keyGen   keygen.KeyGenerator
	userRepo repository.User
	ssoMap   repository.SSOMap
}

// IsAccountLinked checks whether a given external account is linked to any
// internal users already.
func (a AccountLinker) IsAccountLinked(ssoUser entity.SSOUser) (bool, error) {
	return a.ssoMap.IsSSOUserExist(ssoUser.ID)
}

// GetShortUser fetches the internal user linked to the given external user.
func (a AccountLinker) GetShortUser(ssoUser entity.SSOUser) (entity.User, error) {
	id, err := a.ssoMap.GetShortUserID(ssoUser.ID)
	if err != nil {
		return entity.User{}, err
	}
	return a.userRepo.GetUserByID(id)
}

// CreateAndLinkAccount creates an internal account when there is no internal
// account sharing the same email as the given external account and link them
// together afterwards.
func (a AccountLinker) CreateAndLinkAccount(ssoUser entity.SSOUser) error {
	if len(ssoUser.Email) < 1 {
		userID, err := a.createAccount(ssoUser)
		if err != nil {
			return err
		}
		return a.ssoMap.CreateMapping(ssoUser.ID, userID)
	}

	user, err := a.userRepo.GetUserByEmail(ssoUser.Email)
	if err == nil {
		return a.linkExistingAccount(ssoUser, user)
	}

	var errNotFound repository.ErrEntryNotFound
	if !errors.As(err, &errNotFound) {
		return err
	}

	userID, err := a.createAccount(ssoUser)
	if err != nil {
		return err
	}
	return a.ssoMap.CreateMapping(ssoUser.ID, userID)
}

func (a AccountLinker) linkExistingAccount(ssoUser entity.SSOUser, user entity.User) error {
	if len(user.ID) > 0 {
		return a.ssoMap.CreateMapping(ssoUser.ID, user.ID)
	}

	newID, err := a.generateUnassignedUserID()
	if err != nil {
		return err
	}

	err = a.userRepo.UpdateUserID(ssoUser.Email, newID)
	if err != nil {
		return err
	}
	return a.ssoMap.CreateMapping(ssoUser.ID, newID)
}

func (a AccountLinker) createAccount(ssoUser entity.SSOUser) (string, error) {
	userID, err := a.generateUnassignedUserID()
	if err != nil {
		return "", err
	}
	err = a.createUser(userID, ssoUser.Name, ssoUser.Email)
	return userID, err
}

func (a AccountLinker) generateUnassignedUserID() (string, error) {
	newKey, err := a.keyGen.NewKey()
	return string(newKey), err
}

func (a AccountLinker) createUser(id string, name string, email string) error {
	user := entity.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return a.userRepo.CreateUser(user)
}

// AccountLinkerFactory creates AccountLinker.
type AccountLinkerFactory struct {
	keyGen   keygen.KeyGenerator
	userRepo repository.User
}

// NewAccountLinker creates a new account linker.
func (a AccountLinkerFactory) NewAccountLinker(
	ssoMap repository.SSOMap,
) AccountLinker {
	return AccountLinker{
		keyGen:   a.keyGen,
		userRepo: a.userRepo,
		ssoMap:   ssoMap,
	}
}

// NewAccountLinkerFactory creates AccountLinkerFactory.
func NewAccountLinkerFactory(
	keyGen keygen.KeyGenerator,
	userRepo repository.User,
) AccountLinkerFactory {
	return AccountLinkerFactory{
		keyGen:   keyGen,
		userRepo: userRepo,
	}
}
