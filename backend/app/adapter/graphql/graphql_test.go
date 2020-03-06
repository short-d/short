// +build !integration all

package graphql

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/url"
	"github.com/short-d/short/app/usecase/validator"
)

func TestGraphQlAPI(t *testing.T) {
	sqlDB, _, err := mdtest.NewSQLStub()
	mdtest.Equal(t, nil, err)
	defer sqlDB.Close()

	urlRepo := db.NewURLSql(sqlDB)
	retriever := url.NewRetrieverPersist(urlRepo)
	urlRelationRepo := db.NewUserURLRelationSQL(sqlDB)
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
	authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)

	logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
	tracer := mdtest.NewTracerFake()
	graphqlAPI := NewShort(&logger, &tracer, retriever, creator, verifier, authenticator)
	mdtest.Equal(t, true, mdtest.IsGraphQlAPIValid(graphqlAPI))
}
