package request

import (
	"net/http"

	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"

	"github.com/short-d/short/app/usecase/instrumentation"
	"github.com/short-d/short/app/usecase/keygen"
)

// InstrumentationFactory initializes instrumentation code.
type InstrumentationFactory struct {
	runtime   env.Runtime
	keyGen    keygen.KeyGenerator
	logger    logger.Logger
	timer     timer.Timer
	metrics   metrics.Metrics
	analytics analytics.Analytics
	client    Client
}

// NewHTTP creates and initializes Instrumentation tied to the given HTTP
// request.
func (f InstrumentationFactory) NewHTTP(req *http.Request) instrumentation.Instrumentation {
	ctxCh := make(chan ctx.ExecutionContext)

	go func() {
		requestID, err := f.keyGen.NewKey()
		if err != nil {
			f.logger.Error(err)
		}

		location, err := f.client.GetLocation(req)
		if err != nil {
			f.logger.Error(err)
		}

		c := ctx.ExecutionContext{
			RequestID:      string(requestID),
			RequestStartAt: f.timer.Now(),
			Location:       location,
		}
		ctxCh <- c
	}()

	return instrumentation.NewInstrumentation(
		f.logger,
		f.timer,
		f.metrics,
		f.analytics,
		ctxCh,
	)
}

// NewRequest creates and initializes Instrumentation for a given user request.
func (f InstrumentationFactory) NewRequest() instrumentation.Instrumentation {
	ctxCh := make(chan ctx.ExecutionContext)

	go func() {
		requestID, err := f.keyGen.NewKey()
		if err != nil {
			f.logger.Error(err)
		}

		c := ctx.ExecutionContext{
			RequestID:      string(requestID),
			RequestStartAt: f.timer.Now(),
		}
		ctxCh <- c
	}()

	return instrumentation.NewInstrumentation(
		f.logger,
		f.timer,
		f.metrics,
		f.analytics,
		ctxCh,
	)
}

// NewInstrumentationFactory creates Instrumentation factory.
func NewInstrumentationFactory(
	runtime env.Runtime,
	logger logger.Logger,
	timer timer.Timer,
	metrics metrics.Metrics,
	analytics analytics.Analytics,
	keyGen keygen.KeyGenerator,
	client Client,
) InstrumentationFactory {
	return InstrumentationFactory{
		runtime:   runtime,
		logger:    logger,
		timer:     timer,
		metrics:   metrics,
		analytics: analytics,
		keyGen:    keyGen,
		client:    client,
	}
}
