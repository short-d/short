package instrumentation

import (
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/keygen"
)

type Factory struct {
	serverEnv fw.ServerEnv
	keyGen    keygen.KeyGenerator
	logger    fw.Logger
	tracer    fw.Tracer
	timer     fw.Timer
	metrics   fw.Metrics
}

func (f Factory) NewHTTPRequest(req *http.Request) Instrumentation {
	requestID, err := f.keyGen.NewKey()
	if err != nil {
		f.logger.Error(err)
	}

	ctx := fw.ExecutionContext{
		RequestID:      string(requestID),
		RequestStartAt: f.timer.Now(),
	}

	return Instrumentation{
		logger:  f.logger,
		tracer:  f.tracer,
		timer:   f.timer,
		metrics: f.metrics,
		ctx:     ctx,
	}
}

func NewFactory(
	serverEnv fw.ServerEnv,
	logger fw.Logger,
	tracer fw.Tracer,
	timer fw.Timer,
	metrics fw.Metrics,
	keyGen keygen.KeyGenerator,
) Factory {
	return Factory{
		serverEnv: serverEnv,
		logger:    logger,
		tracer:    tracer,
		timer:     timer,
		metrics:   metrics,
		keyGen:    keyGen,
	}
}
