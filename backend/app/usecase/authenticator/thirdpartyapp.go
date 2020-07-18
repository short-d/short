package authenticator

import (
	"errors"
	"fmt"

	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator/payload"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// ThirdPartyApp authenticates the identify of a third application.
type ThirdPartyApp struct {
	authorizer authorizer.Authorizer
	tokenizer  crypto.Tokenizer
	keyGen     keygen.KeyGenerator
	timer      timer.Timer
	apiKeyRepo repository.APIKey
	appRepo    repository.App
}

// GetApp retrieves app information based on the credential provided.
func (t ThirdPartyApp) GetApp(cred Credential) (entity.App, error) {
	if cred.APIKey == nil {
		return entity.App{}, nil
	}
	apiKeyStr := *cred.APIKey
	tokenPayload, err := t.tokenizer.Decode(apiKeyStr)
	if err != nil {
		return entity.App{}, err
	}

	apiKeyPayload, err := payload.NewAPIKey(tokenPayload)
	if err != nil {
		return entity.App{}, errors.New("invalid api key")
	}
	apiKey, err := t.apiKeyRepo.GetAPIKey(apiKeyPayload.AppID, apiKeyPayload.Key)
	if err != nil {
		return entity.App{}, errors.New("invalid api key")
	}
	if apiKey.IsDisabled {
		return entity.App{}, errors.New("invalid api key")
	}
	return t.appRepo.GetAppByID(apiKey.AppID)
}

// GenerateAPIKey generates a new API key for the given app.
func (t ThirdPartyApp) GenerateAPIKey(user entity.User, app entity.App) (string, error) {
	canGenerateKey, err := t.authorizer.CanGenerateAPIKey(user)
	if err != nil {
		return "", err
	}

	if !canGenerateKey {
		return "", fmt.Errorf("user(%s) cannot generate api key for app(%s)", user.ID, app.ID)
	}

	var entryNotFound repository.ErrEntryNotFound

	_, err = t.appRepo.GetAppByID(app.ID)
	if err != nil {
		return "", err
	}

	key, err := t.newKey()
	if err != nil {
		return "", err
	}

	_, err = t.apiKeyRepo.GetAPIKey(app.ID, key)
	if err == nil {
		return "", fmt.Errorf("key(%s) already exists for app(%s)", key, app.ID)
	}
	if !errors.As(err, &entryNotFound) {
		return "", err
	}

	isDisabled := false
	now := t.timer.Now()
	in := entity.APIKeyInput{
		AppID:      &app.ID,
		Key:        &key,
		IsDisabled: &isDisabled,
		CreatedAt:  &now,
	}
	apiKey, err := t.apiKeyRepo.CreateAPIKey(in)
	if err != nil {
		return "", err
	}

	apiKeyPayload := payload.APIKey{
		AppID: apiKey.AppID,
		Key:   apiKey.Key,
	}
	tokenPayload := apiKeyPayload.NewTokenPayload()
	return t.tokenizer.Encode(tokenPayload)
}

func (t ThirdPartyApp) newKey() (string, error) {
	key, err := t.keyGen.NewKey()
	return string(key), err
}

// NewThirdPartyApp creates ThirdPartyApp authenticator.
func NewThirdPartyApp(
	authorizer authorizer.Authorizer,
	tokenizer crypto.Tokenizer,
	keyGen keygen.KeyGenerator,
	timer timer.Timer,
	apiKeyRepo repository.APIKey,
	appRepo repository.App,
) ThirdPartyApp {
	return ThirdPartyApp{
		authorizer: authorizer,
		tokenizer:  tokenizer,
		keyGen:     keyGen,
		timer:      timer,
		apiKeyRepo: apiKeyRepo,
		appRepo:    appRepo,
	}
}
