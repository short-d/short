package graphql

import (
	"short/app/adapter/db"
	"short/app/adapter/recaptcha"
	"short/app/entity"
	"short/app/usecase/auth"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

func TestGraphQlAPI(t *testing.T) {
	sqlDB, _, err := mdtest.NewSQLStub()
	mdtest.Equal(t, nil, err)
	defer sqlDB.Close()

	urls := map[string]entity.URL{}
	retriever := url.NewRetrieverFake(urls)

	urlRepo := db.NewURLSql(sqlDB)
	urlRelationRepo := db.NewUserURLRelationSQL(sqlDB)
	keyGen := keygen.NewFake([]string{})
	creator := url.NewCreatorPersist(urlRepo, urlRelationRepo, &keyGen)

	s := recaptcha.NewFake()
	verifier := requester.NewVerifier(s)
	authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)

	logger := mdtest.NewLoggerFake()
	tracer := mdtest.NewTracerFake()
	graphqlAPI := NewShort(&logger, &tracer, retriever, creator, verifier, authenticator)
	mdtest.Equal(t, true, mdtest.IsGraphQlAPIValid(graphqlAPI))
}
