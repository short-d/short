package instrumentation

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/observability"
)

var _ observability.Observability = (*Instrumentation)(nil)

// Instrumentation specifics how the system should be observed.
type Instrumentation struct {
	logger      fw.Logger
	tracer      fw.Tracer
	timer       fw.Timer
	metrics     fw.Metrics
	analytics   fw.Analytics
	geoLocation fw.GeoLocation
	ctx         fw.ExecutionContext
}

func (i Instrumentation) RedirectingAliasToLongLink(user *entity.User) {
	userID := i.getUserID(user)
	i.analytics.Track("RedirectingAliasToLongLink", map[string]string{}, userID, i.ctx)
}

func (i Instrumentation) RedirectedAliasToLongLink(user *entity.User) {
	userID := i.getUserID(user)
	i.analytics.Track("RedirectedAliasToLongLink", map[string]string{}, userID, i.ctx)
}

func (i Instrumentation) getUserID(user *entity.User) string {
	if user == nil {
		return i.ctx.RequestID
	}
	return user.Email
}

// LongLinkRetrievalSucceed tracks the successes when retrieving long links.
func (i Instrumentation) LongLinkRetrievalSucceed() {
	i.metrics.Count("long-link-retrieval-succeed", 1, 1, i.ctx)
}

// LongLinkRetrievalFailed tracks the failures when retrieving long links.
func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	i.logger.Error(err)
	i.metrics.Count("long-link-retrieval-failed", 1, 1, i.ctx)
}
