package request

import (
	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/usecase/keygen"
)

type SSOInstrumentationFactory struct {
	ssoProviderName string
	keyGen          keygen.KeyGenerator
	logger          logger.Logger
	timer           timer.Timer
	metrics         metrics.Metrics
	analytics       analytics.Analytics
	client          Client
}

// NewInstrumentationcFactory creates Instrumentation factory.
func NewSSOInstrumentationFactory(
	provider string) SSOInstrumentationFactory {
	return SSOInstrumentationFactory{}
}
