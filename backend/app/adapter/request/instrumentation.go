package request

import (
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/instrumentation"
	"github.com/short-d/short/app/usecase/keygen"
)

// InstrumentationFactory initializes instrumentation code.
type InstrumentationFactory struct {
	serverEnv fw.ServerEnv
	keyGen    keygen.KeyGenerator
	logger    fw.Logger
	tracer    fw.Tracer
	timer     fw.Timer
	metrics   fw.Metrics
	analytics fw.Analytics
	client    Client
}

// NewHTTPRequest creates and initializes Instrumentation tied to the given HTTP
// request.
func (f InstrumentationFactory) NewHTTP(req *http.Request) instrumentation.Instrumentation {
	requestID, err := f.keyGen.NewKey()
	if err != nil {
		f.logger.Error(err)
	}

	location, err := f.client.GetLocation(req)
	if err != nil {
		f.logger.Error(err)
	}

	ctx := fw.ExecutionContext{
		RequestID:      string(requestID),
		RequestStartAt: f.timer.Now(),
		Location:       location,
	}

	return instrumentation.NewInstrumentation(
		f.logger,
		f.tracer,
		f.timer,
		f.metrics,
		f.analytics,
		ctx,
	)
}

// NewFactory creates Instrumentation factory.
func NewInstrumentationFactory(
	serverEnv fw.ServerEnv,
	logger fw.Logger,
	tracer fw.Tracer,
	timer fw.Timer,
	metrics fw.Metrics,
	analytics fw.Analytics,
	keyGen keygen.KeyGenerator,
	client Client,
) InstrumentationFactory {
	return InstrumentationFactory{
		serverEnv: serverEnv,
		logger:    logger,
		tracer:    tracer,
		timer:     timer,
		metrics:   metrics,
		analytics: analytics,
		keyGen:    keyGen,
		client:    client,
	}
}
