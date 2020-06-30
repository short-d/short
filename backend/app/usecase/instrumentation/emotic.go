package instrumentation

import (
	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
)

type Emotic struct {
	logger                          logger.Logger
	timer                           timer.Timer
	metrics                         metrics.Metrics
	analytics                       analytics.Analytics
	ctxCh                           chan ctx.ExecutionContext
	feedbackReceivedCh         chan ctx.ExecutionContext
	feedbackNotifyFailedCh          chan ctx.ExecutionContext
}

func (e Emotic) FeedbackReceived(feedback entity.Feedback) {
	go func() {
		c := <-e.feedbackReceivedCh
		userID := getUserID(nil)
		props := map[string]string{
			"customer-rating": string(feedback.CustomerRating),
			"comment": feedback.GetComment(""),
			"customer-email": feedback.GetCustomerEmail(""),
		}
		e.analytics.Track("FeedbackReceived", props, userID, c)
	}()
}

func (e Emotic) FeedbackNotifyFailed(err error) {
	go func() {
		c := <-e.feedbackNotifyFailedCh
		e.logger.Error(err)
		e.metrics.Count("feedback-notify-failed", 1, 1, c)
	}()
}

// Done closes all the channels to prevent memory leak.
func (e Emotic) Done() {
	close(e.feedbackReceivedCh)
	close(e.feedbackNotifyFailedCh)
}

// NewEmotic initializes instrumentation code for Emotic.
func NewEmotic(
	logger logger.Logger,
	timer timer.Timer,
	metrics metrics.Metrics,
	analytics analytics.Analytics,
	ctxCh chan ctx.ExecutionContext,
) Emotic {
	feedbackReceivedCh := make(chan ctx.ExecutionContext)
	feedbackNotifyFailedCh := make(chan ctx.ExecutionContext)

	ins := Emotic{
		logger:                          logger,
		timer:                           timer,
		metrics:                         metrics,
		analytics:                       analytics,
		ctxCh:                           ctxCh,
		feedbackReceivedCh:    feedbackReceivedCh,
		feedbackNotifyFailedCh:     feedbackNotifyFailedCh,
	}
	go func() {
		c := <-ctxCh
		go func() { feedbackReceivedCh <- c }()
		go func() { feedbackNotifyFailedCh <- c }()
		close(ctxCh)
	}()
	return ins
}

