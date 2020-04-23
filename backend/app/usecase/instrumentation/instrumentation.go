package instrumentation

import (
	"fmt"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

// Instrumentation measures the internal operation of the system.
type Instrumentation struct {
	logger    fw.Logger
	tracer    fw.Tracer
	timer     fw.Timer
	metrics   fw.Metrics
	analytics fw.Analytics
	ctx       fw.ExecutionContext
}

// RedirectingAliasToLongLink tracks RedirectingAliasToLongLink event.
func (i Instrumentation) RedirectingAliasToLongLink(user *entity.User) {
	go func() {
		userID := i.getUserID(user)
		i.analytics.Track("RedirectingAliasToLongLink", map[string]string{}, userID, i.ctx)
	}()
}

// RedirectedAliasToLongLink tracks RedirectedAliasToLongLink event.
func (i Instrumentation) RedirectedAliasToLongLink(user *entity.User) {
	go func() {
		userID := i.getUserID(user)
		i.analytics.Track("RedirectedAliasToLongLink", map[string]string{}, userID, i.ctx)
	}()
}

// LongLinkRetrievalSucceed tracks the successes when retrieving long links.
func (i Instrumentation) LongLinkRetrievalSucceed() {
	go func() {
		i.metrics.Count("long-link-retrieval-succeed", 1, 1, i.ctx)
	}()
}

// LongLinkRetrievalFailed tracks the failures when retrieving long links.
func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	go func() {
		i.logger.Error(err)
		i.metrics.Count("long-link-retrieval-failed", 1, 1, i.ctx)
	}()
}

// FeatureToggleRetrievalSucceed tracks the successes when retrieving the status
// of the feature toggle.
func (i Instrumentation) FeatureToggleRetrievalSucceed() {
	go func() {
		i.metrics.Count("feature-toggle-retrieval-succeed", 1, 1, i.ctx)
	}()
}

// FeatureToggleRetrievalSucceed tracks the failures when retrieving the status
// of the feature toggle.
func (i Instrumentation) FeatureToggleRetrievalFailed(err error) {
	go func() {
		i.logger.Error(err)
		i.metrics.Count("feature-toggle-retrieval-failed", 1, 1, i.ctx)
	}()
}

// MadeFeatureDecision tracks MadeFeatureDecision event.
func (i Instrumentation) MadeFeatureDecision(
	featureID string,
	isEnabled bool,
) {
	go func() {
		userID := i.getUserID(nil)
		isEnabledStr := fmt.Sprintf("%v", isEnabled)
		props := map[string]string{
			"feature-id": featureID,
			"is-enabled": isEnabledStr,
		}
		i.analytics.Track("MadeFeatureDecision", props, userID, i.ctx)
	}()
}

func (i Instrumentation) getUserID(user *entity.User) string {
	if user == nil {
		return i.ctx.RequestID
	}
	return user.Email
}

// NewInstrumentation initializes instrumentation code.
func NewInstrumentation(logger fw.Logger,
	tracer fw.Tracer,
	timer fw.Timer,
	metrics fw.Metrics,
	analytics fw.Analytics,
	ctx fw.ExecutionContext,
) Instrumentation {
	return Instrumentation{
		logger:    logger,
		tracer:    tracer,
		timer:     timer,
		metrics:   metrics,
		analytics: analytics,
		ctx:       ctx,
	}
}
