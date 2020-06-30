package authenticator

import (
	"errors"
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type CloudAPI struct {
	timer timer.Timer
	kegGen keygen.KeyGenerator
	tokenizer  crypto.Tokenizer
	apiKeyRepo repository.ApiKey
	appRepo    repository.App
}

func (c CloudAPI) GetApp(apiKeyToken string) (entity.App, error) {
	payload, err := c.tokenizer.Decode(apiKeyToken)
	appID, ok := payload["app_id"]
	if !ok {
		return entity.App{}, errors.New("invalid api key")
	}
	key, ok := payload["key"]
	if !ok {
		return entity.App{}, errors.New("invalid api key")
	}

	apiKey, err := c.apiKeyRepo.FindApiKey(appID.(string), key.(string))
	if err != nil {
		return entity.App{}, err
	}

	if apiKey.IsDisabled {
		return entity.App{}, errors.New("invalid api key")
	}
	return c.appRepo.FindAppByID(apiKey.AppID)
}

func (c CloudAPI) GenerateApiKey(appID string) (string, error) {
	key, err := c.kegGen.NewKey()
	if err != nil {
		return "", err
	}

	keyStr := string(key)
	now := c.timer.Now()
	input := entity.ApiKeyInput{
		AppID:      &appID,
		Key:        &keyStr,
		CreatedAt:  &now,
	}
	apiKey, err := c.apiKeyRepo.CreateApiKey(input)
	if err != nil {
		return "", err
	}
	payload := crypto.TokenPayload{
		"app_id": apiKey.AppID,
		"key": apiKey.Key,
	}
	return c.tokenizer.Encode(payload)
}

func NewCloudAPI(
	timer timer.Timer,
	kegGen keygen.KeyGenerator,
	tokenizer crypto.Tokenizer,
	apiKeyRepo repository.ApiKey,
	appRepo repository.App,
) CloudAPI {
	return CloudAPI{
		timer: timer,
		kegGen: kegGen,
		tokenizer: tokenizer,
		apiKeyRepo: apiKeyRepo,
		appRepo: appRepo,
	}
}
