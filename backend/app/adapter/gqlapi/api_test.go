// +build !integration all

package gqlapi

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/requester"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/url"
	"github.com/short-d/short/backend/app/usecase/validator"
)

func TestGraphQlAPI(t *testing.T) {
	t.Parallel()
	now := time.Now()
	blockedURLs := map[string]bool{}
	blacklist := risk.NewBlackListFake(blockedURLs)

	urlRepo := repository.NewShortLinkFake(map[string]entity.ShortLink{})
	urlRelationRepo := repository.NewUserShortLinkRepoFake([]entity.User{}, []entity.ShortLink{})
	retriever := url.NewRetrieverPersist(&urlRepo, &urlRelationRepo)
	keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
	keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
	assert.Equal(t, nil, err)

	longLinkValidator := validator.NewLongLink()
	customAliasValidator := validator.NewCustomAlias()
	tm := timer.NewStub(now)
	riskDetector := risk.NewDetector(blacklist)

	creator := url.NewCreatorPersist(
		&urlRepo,
		&urlRelationRepo,
		keyGen,
		longLinkValidator,
		customAliasValidator,
		tm,
		riskDetector,
	)

	s := requester.NewReCaptchaFake(requester.VerifyResponse{})
	verifier := requester.NewVerifier(s)
	auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)

	entryRepo := logger.NewEntryRepoFake()
	lg, err := logger.NewFake(logger.LogOff, &entryRepo)
	assert.Equal(t, nil, err)

	changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
	userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
	changeLog := changelog.NewPersist(keyGen, tm, &changeLogRepo, &userChangeLogRepo)
	r := resolver.NewResolver(lg, retriever, creator, changeLog, verifier, auth)
	graphqlAPI := NewShort(r)
	assert.Equal(t, true, graphql.IsGraphQlAPIValid(graphqlAPI))
}
