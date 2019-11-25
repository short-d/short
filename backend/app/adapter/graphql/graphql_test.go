package graphql

import (
	"short/app/adapter/db"
	"short/app/adapter/recaptcha"
	"short/app/usecase/auth"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/app/usecase/validator"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestGraphQlAPI(t *testing.T) {
	sqlDB, _, err := mdtest.NewSQLStub()
	mdtest.Equal(t, nil, err)
	defer sqlDB.Close()

	urlRepo := db.NewURLSql(sqlDB)
	retriever := url.NewRetrieverPersist(urlRepo)
	urlRelationRepo := db.NewUserURLRelationSQL(sqlDB)
	keyGen := keygen.NewFake([]string{})
	longLinkValidator := validator.NewLongLink()
	customAliasValidator := validator.NewCustomAlias()
	creator := url.NewCreatorPersist(
		urlRepo,
		urlRelationRepo,
		&keyGen,
		longLinkValidator,
		customAliasValidator,
	)

	s := recaptcha.NewFake()
	verifier := requester.NewVerifier(s)
	authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)

	logger := mdtest.NewLoggerFake()
	tracer := mdtest.NewTracerFake()
	graphqlAPI := NewShort(&logger, &tracer, retriever, creator, verifier, authenticator)
	mdtest.Equal(t, true, mdtest.IsGraphQlAPIValid(graphqlAPI))
}
