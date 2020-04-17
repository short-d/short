package instrumentation

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/observability"
)

var _ observability.Observability = (*Instrumentation)(nil)

type Instrumentation struct {
	logger  fw.Logger
	tracer  fw.Tracer
	timer   fw.Timer
	metrics fw.Metrics
	ctx     fw.ExecutionContext
}

func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	i.logger.Error(err)
	i.metrics.Count("long-link-retrieval-failed", 1, 1, i.ctx)
}

func (i Instrumentation) LongLinkRetrievalSucceed() {
	i.metrics.Count("long-link-retrieval-succeed", 1, 1, i.ctx)
}
