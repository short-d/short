package instrumentation

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/observability"
)

var _ observability.Observability = (*Instrumentation)(nil)

// Instrumentation specifics how the system should be observed.
type Instrumentation struct {
	logger  fw.Logger
	tracer  fw.Tracer
	timer   fw.Timer
	metrics fw.Metrics
	ctx     fw.ExecutionContext
}

// LongLinkRetrievalFailed tracks the failures when retrieving long links.
func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	i.logger.Error(err)
	i.metrics.Count("long-link-retrieval-failed", 1, 1, i.ctx)
}

// LongLinkRetrievalSucceed tracks the successes when retrieving long links.
func (i Instrumentation) LongLinkRetrievalSucceed() {
	i.metrics.Count("long-link-retrieval-succeed", 1, 1, i.ctx)
}
