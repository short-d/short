package provider

import (
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/slack"
	"github.com/short-d/short/backend/app/usecase/emotic"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type FeedbackSlackWebHook string

func NewFeedback(
	timer timer.Timer,
	feedbackRepo repository.Feedback,
	keyGen keygen.KeyGenerator,
	emoticNotifierFactory slack.EmoticNotifierFactory,
	slackWebHook FeedbackSlackWebHook,
	) emotic.Feedback {
	var notifiers []emotic.Notifier
	if len(slackWebHook) > 0 {
		slackNotifier := emoticNotifierFactory.NewEmoticNotifier(string(slackWebHook))
		notifiers = append(notifiers, slackNotifier)
	}
	return emotic.NewFeedback(timer, feedbackRepo, keyGen, notifiers)
}
