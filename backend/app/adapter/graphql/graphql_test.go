// +build !integration all

package graphql

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/url"
	"github.com/short-d/short/app/usecase/validator"
)

func TestGraphQlAPI(t *testing.T) {
	now := time.Now()
	sqlDB, _, err := mdtest.NewSQLStub()
	mdtest.Equal(t, nil, err)
	defer sqlDB.Close()

	urlRepo := db.NewURLSql(sqlDB)
	urlRelationRepo := db.NewUserURLRelationSQL(sqlDB)
	retriever := url.NewRetrieverPersist(urlRepo, urlRelationRepo)
	keyFetcher := service.NewKeyFetcherFake([]service.Key{})
	keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
	mdtest.Equal(t, nil, err)
	longLinkValidator := validator.NewLongLink()
	customAliasValidator := validator.NewCustomAlias()
	creator := url.NewCreatorPersist(
		urlRepo,
		urlRelationRepo,
		keyGen,
		longLinkValidator,
		customAliasValidator,
	)

	s := service.NewReCaptchaFake(service.VerifyResponse{})
	verifier := requester.NewVerifier(s)
	auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)

	logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
	tracer := mdtest.NewTracerFake()

	timerFake := mdtest.NewTimerFake(now)
	changeLogRepo := db.NewChangeLogSQL(sqlDB)
	changeLog := changelog.NewPersist(keyGen, timerFake, changeLogRepo)
	graphqlAPI := NewShort(&logger, &tracer, retriever, creator, changeLog, verifier, auth)
	mdtest.Equal(t, true, mdtest.IsGraphQlAPIValid(graphqlAPI))
}
