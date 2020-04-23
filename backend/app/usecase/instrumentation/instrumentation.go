package instrumentation

import (
	"fmt"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

// Instrumentation measures the internal operation of the system.
type Instrumentation struct {
	logger                          fw.Logger
	tracer                          fw.Tracer
	timer                           fw.Timer
	metrics                         fw.Metrics
	analytics                       fw.Analytics
	ctx                             fw.ExecutionContext
	ctxCh                           chan fw.ExecutionContext
	redirectingAliasToLongLinkCh    chan struct{}
	redirectedAliasToLongLinkCh     chan struct{}
	longLinkRetrievalSucceedCh      chan struct{}
	longLinkRetrievalFailedCh       chan struct{}
	featureToggleRetrievalSucceedCh chan struct{}
	featureToggleRetrievalFailedCh  chan struct{}
}

// RedirectingAliasToLongLink tracks RedirectingAliasToLongLink event.
func (i Instrumentation) RedirectingAliasToLongLink(user *entity.User) {
	go func() {
		<-i.redirectingAliasToLongLinkCh
		userID := i.getUserID(user)
		i.analytics.Track("RedirectingAliasToLongLink", map[string]string{}, userID, i.ctx)
	}()
}

// RedirectedAliasToLongLink tracks RedirectedAliasToLongLink event.
func (i Instrumentation) RedirectedAliasToLongLink(user *entity.User) {
	go func() {
		<-i.redirectedAliasToLongLinkCh
		userID := i.getUserID(user)
		i.analytics.Track("RedirectedAliasToLongLink", map[string]string{}, userID, i.ctx)
	}()
}

// LongLinkRetrievalSucceed tracks the successes when retrieving long links.
func (i Instrumentation) LongLinkRetrievalSucceed() {
	go func() {
		<-i.longLinkRetrievalSucceedCh
		i.metrics.Count("long-link-retrieval-succeed", 1, 1, i.ctx)
	}()
}

// LongLinkRetrievalFailed tracks the failures when retrieving long links.
func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	go func() {
		<-i.longLinkRetrievalFailedCh
		i.logger.Error(err)
		i.metrics.Count("long-link-retrieval-failed", 1, 1, i.ctx)
	}()
}

// FeatureToggleRetrievalSucceed tracks the successes when retrieving the status
// of the feature toggle.
func (i Instrumentation) FeatureToggleRetrievalSucceed() {
	go func() {
		<-i.featureToggleRetrievalSucceedCh
		i.metrics.Count("feature-toggle-retrieval-succeed", 1, 1, i.ctx)
	}()
}

// FeatureToggleRetrievalFailed tracks the failures when retrieving the status
// of the feature toggle.
func (i Instrumentation) FeatureToggleRetrievalFailed(err error) {
	go func() {
		<-i.featureToggleRetrievalFailedCh
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

// Done closes all the channels to prevent memory leak.
func (i Instrumentation) Done() {
	close(i.redirectingAliasToLongLinkCh)
	close(i.redirectedAliasToLongLinkCh)
	close(i.longLinkRetrievalSucceedCh)
	close(i.longLinkRetrievalFailedCh)
	close(i.featureToggleRetrievalSucceedCh)
	close(i.featureToggleRetrievalFailedCh)
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
	ctxCh chan fw.ExecutionContext,
) Instrumentation {
	redirectingAliasToLongLinkCh := make(chan struct{})
	redirectedAliasToLongLinkCh := make(chan struct{})
	longLinkRetrievalSucceedCh := make(chan struct{})
	longLinkRetrievalFailedCh := make(chan struct{})
	featureToggleRetrievalSucceedCh := make(chan struct{})
	featureToggleRetrievalFailedCh := make(chan struct{})

	ins := Instrumentation{
		logger:                          logger,
		tracer:                          tracer,
		timer:                           timer,
		metrics:                         metrics,
		analytics:                       analytics,
		ctxCh:                           ctxCh,
		redirectingAliasToLongLinkCh:    redirectingAliasToLongLinkCh,
		redirectedAliasToLongLinkCh:     redirectedAliasToLongLinkCh,
		longLinkRetrievalSucceedCh:      longLinkRetrievalSucceedCh,
		longLinkRetrievalFailedCh:       longLinkRetrievalFailedCh,
		featureToggleRetrievalSucceedCh: featureToggleRetrievalSucceedCh,
		featureToggleRetrievalFailedCh:  featureToggleRetrievalFailedCh,
	}
	go func() {
		ctx := <-ctxCh
		ins.ctx = ctx
		redirectingAliasToLongLinkCh <- struct{}{}
		redirectedAliasToLongLinkCh <- struct{}{}
		longLinkRetrievalSucceedCh <- struct{}{}
		longLinkRetrievalFailedCh <- struct{}{}
		featureToggleRetrievalSucceedCh <- struct{}{}
		featureToggleRetrievalFailedCh <- struct{}{}
		defer close(ctxCh)
	}()
	return ins
}
