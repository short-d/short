// +build !integration all

package gqlapi

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/sqldb"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/external"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/risk"
	"github.com/short-d/short/app/usecase/url"
	"github.com/short-d/short/app/usecase/validator"
)

func TestGraphQlAPI(t *testing.T) {
	now := time.Now()
	blockedURLs := map[string]bool{}
	blacklist := risk.NewBlackListFake(blockedURLs)
	sqlDB, _, err := mdtest.NewSQLStub()
	assert.Equal(t, nil, err)
	defer sqlDB.Close()

	urlRepo := sqldb.NewURLSql(sqlDB)
	urlRelationRepo := sqldb.NewUserURLRelationSQL(sqlDB)
	retriever := url.NewRetrieverPersist(urlRepo, urlRelationRepo)
	keyFetcher := external.NewKeyFetcherFake([]external.Key{})
	keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
	assert.Equal(t, nil, err)
	longLinkValidator := validator.NewLongLink()
	customAliasValidator := validator.NewCustomAlias()
	timer := mdtest.NewTimerFake(time.Now())
	riskDetector := risk.NewDetector(blacklist)

	creator := url.NewCreatorPersist(
		urlRepo,
		urlRelationRepo,
		keyGen,
		longLinkValidator,
		customAliasValidator,
		timer,
		riskDetector,
	)

	s := external.NewReCaptchaFake(external.VerifyResponse{})
	verifier := requester.NewVerifier(s)
	auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)

	logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
	tracer := mdtest.NewTracerFake()

	timerFake := mdtest.NewTimerFake(now)
	changeLogRepo := sqldb.NewChangeLogSQL(sqlDB)
	changeLog := changelog.NewPersist(keyGen, timerFake, changeLogRepo)
	graphqlAPI := NewShort(&logger, &tracer, retriever, creator, changeLog, verifier, auth)
	assert.Equal(t, true, mdtest.IsGraphQlAPIValid(graphqlAPI))
}
