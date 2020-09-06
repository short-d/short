package instrumentation

import (
	"fmt"
	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/usecase/keygen"
)

type SSOInstrumentation struct {
	event     string
	keyGen    keygen.KeyGenerator
	logger    logger.Logger
	timer     timer.Timer
	metrics   metrics.Metrics
	analytics analytics.Analytics
	ssoSignIn chan ctx.ExecutionContext
}

func (i SSOInstrumentation) StartingSSO() {
}

func (i SSOInstrumentation) UserAlreadySignedIn() {

}

func (i SSOInstrumentation) RetrievingAuthToken() {

}

func (i SSOInstrumentation) SignInFailed() {

}

func (i SSOInstrumentation) SignInSucceeded() {

}

func (i SSOInstrumentation) VerifyingAuthToken(token string) {
	i.logger.Info("Checking if auth token exists")
}

func (i SSOInstrumentation) RedirectingUserToSignInLink(signInLink string) {

}

func NewSSO(providerName string) SSOInstrumentation {
	return SSOInstrumentation{
		event: fmt.Sprintf("%s-SSO", providerName),
	}
}
