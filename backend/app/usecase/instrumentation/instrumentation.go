package instrumentation

import (
	"fmt"
	"strings"

	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
)

// Instrumentation measures the internal operation of the system.
type Instrumentation struct {
	logger                          logger.Logger
	timer                           timer.Timer
	metrics                         metrics.Metrics
	analytics                       analytics.Analytics
	ctxCh                           chan ctx.ExecutionContext
	redirectingAliasToLongLinkCh    chan ctx.ExecutionContext
	redirectedAliasToLongLinkCh     chan ctx.ExecutionContext
	longLinkRetrievalSucceedCh      chan ctx.ExecutionContext
	longLinkRetrievalFailedCh       chan ctx.ExecutionContext
	featureToggleRetrievalSucceedCh chan ctx.ExecutionContext
	featureToggleRetrievalFailedCh  chan ctx.ExecutionContext
	searchSucceedCh                 chan ctx.ExecutionContext
	searchFailedCh                  chan ctx.ExecutionContext
	madeFeatureDecisionCh           chan ctx.ExecutionContext
	trackCh                         chan ctx.ExecutionContext
}

// RedirectingAliasToLongLink tracks RedirectingAliasToLongLink event.
func (i Instrumentation) RedirectingAliasToLongLink(alias string) {
	go func() {
		c := <-i.redirectingAliasToLongLinkCh
		userID := i.getUserID(nil)
		props := map[string]string{
			"request-id": c.RequestID,
			"alias":      alias,
		}
		i.analytics.Track("RedirectingAliasToLongLink", props, userID, c)
	}()
}

// RedirectedAliasToLongLink tracks RedirectedAliasToLongLink event.
func (i Instrumentation) RedirectedAliasToLongLink(shortLink entity.ShortLink) {
	go func() {
		c := <-i.redirectedAliasToLongLinkCh
		userID := i.getUserID(nil)
		props := map[string]string{
			"request-id": c.RequestID,
			"alias":      shortLink.Alias,
			"long-link":  shortLink.LongLink,
		}
		i.analytics.Track("RedirectedAliasToLongLink", props, userID, c)
	}()
}

// LongLinkRetrievalSucceed tracks the successes when retrieving long links.
func (i Instrumentation) LongLinkRetrievalSucceed() {
	go func() {
		c := <-i.longLinkRetrievalSucceedCh
		i.metrics.Count("long-link-retrieval-succeed", 1, 1, c)
	}()
}

// LongLinkRetrievalFailed tracks the failures when retrieving long links.
func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	go func() {
		c := <-i.longLinkRetrievalFailedCh
		i.logger.Error(err)
		i.metrics.Count("long-link-retrieval-failed", 1, 1, c)
	}()
}

// FeatureToggleRetrievalSucceed tracks the successes when retrieving the status
// of the feature toggle.
func (i Instrumentation) FeatureToggleRetrievalSucceed() {
	go func() {
		c := <-i.featureToggleRetrievalSucceedCh
		i.metrics.Count("feature-toggle-retrieval-succeed", 1, 1, c)
	}()
}

// FeatureToggleRetrievalFailed tracks the failures when retrieving the status
// of the feature toggle.
func (i Instrumentation) FeatureToggleRetrievalFailed(err error) {
	go func() {
		c := <-i.featureToggleRetrievalFailedCh
		i.logger.Error(err)
		i.metrics.Count("feature-toggle-retrieval-failed", 1, 1, c)
	}()
}

// SearchSucceed tracks the successes when searching the short resources.
func (i Instrumentation) SearchSucceed(user *entity.User, keywords string, resources []string) {
	go func() {
		c := <-i.searchSucceedCh
		i.metrics.Count("search-succeed", 1, 1, c)
		userID := i.getUserID(user)
		props := map[string]string{
			"keywords":  keywords,
			"resources": strings.Join(resources, ","),
		}
		i.analytics.Track("Search", props, userID, c)
	}()
}

// SearchFailed tracks the failures when searching the short resources.
func (i Instrumentation) SearchFailed(err error) {
	go func() {
		c := <-i.searchFailedCh
		i.logger.Error(err)
		i.metrics.Count("search-failed", 1, 1, c)
	}()
}

// MadeFeatureDecision tracks MadeFeatureDecision event.
func (i Instrumentation) MadeFeatureDecision(
	featureID string,
	isEnabled bool,
) {
	go func() {
		c := <-i.madeFeatureDecisionCh
		userID := i.getUserID(nil)
		isEnabledStr := fmt.Sprintf("%v", isEnabled)
		props := map[string]string{
			"request-id": c.RequestID,
			"feature-id": featureID,
			"is-enabled": isEnabledStr,
		}
		i.analytics.Track("MadeFeatureDecision", props, userID, c)
	}()
}

// Track records events happened in the system.
func (i Instrumentation) Track(event string) {
	go func() {
		c := <-i.trackCh
		userID := i.getUserID(nil)
		props := map[string]string{}
		i.analytics.Track(event, props, userID, c)
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
		return "anonymous"
	}
	return user.ID
}

// NewInstrumentation initializes instrumentation code.
func NewInstrumentation(
	logger logger.Logger,
	timer timer.Timer,
	metrics metrics.Metrics,
	analytics analytics.Analytics,
	ctxCh chan ctx.ExecutionContext,
) Instrumentation {
	redirectingAliasToLongLinkCh := make(chan ctx.ExecutionContext)
	redirectedAliasToLongLinkCh := make(chan ctx.ExecutionContext)
	longLinkRetrievalSucceedCh := make(chan ctx.ExecutionContext)
	longLinkRetrievalFailedCh := make(chan ctx.ExecutionContext)
	featureToggleRetrievalSucceedCh := make(chan ctx.ExecutionContext)
	featureToggleRetrievalFailedCh := make(chan ctx.ExecutionContext)
	searchSucceedCh := make(chan ctx.ExecutionContext)
	searchFailedCh := make(chan ctx.ExecutionContext)
	madeFeatureDecisionCh := make(chan ctx.ExecutionContext)
	trackCh := make(chan ctx.ExecutionContext)
	ssoSignIn := make(chan ctx.ExecutionContext)

	ins := &Instrumentation{
		logger:                          logger,
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
		searchSucceedCh:                 searchSucceedCh,
		searchFailedCh:                  searchFailedCh,
		madeFeatureDecisionCh:           madeFeatureDecisionCh,
		trackCh:                         trackCh,
	}
	go func() {
		c := <-ctxCh
		go func() { redirectingAliasToLongLinkCh <- c }()
		go func() { redirectedAliasToLongLinkCh <- c }()
		go func() { longLinkRetrievalSucceedCh <- c }()
		go func() { longLinkRetrievalFailedCh <- c }()
		go func() { featureToggleRetrievalSucceedCh <- c }()
		go func() { featureToggleRetrievalFailedCh <- c }()
		go func() { searchSucceedCh <- c }()
		go func() { searchFailedCh <- c }()
		go func() { madeFeatureDecisionCh <- c }()
		go func() { trackCh <- c }()
		go func() { ssoSignIn <- c }()
		close(ctxCh)
	}()
	return *ins
}
