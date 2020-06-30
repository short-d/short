package emotic

import (
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type Feedback struct {
	timer timer.Timer
	keyGen keygen.KeyGenerator
	emoticInstrumentation instrumentation.Emotic
	feedbackRepo repository.Feedback
	notifiers    []Notifier
}

func (f Feedback) ReceiveFeedback(input entity.FeedbackInput) (entity.Feedback, error) {
	key, err :=  f.keyGen.NewKey()
	if err != nil {
		return entity.Feedback{}, err
	}
	feedbackID := string(key)
	input.FeedbackID = &feedbackID
	now := f.timer.Now()
	input.ReceivedAt = &now

	fb, err := f.feedbackRepo.CreateFeedback(input)
	if err != nil {
		return entity.Feedback{}, err
	}
	f.emoticInstrumentation.FeedbackReceived(fb)

	for _, notifier := range f.notifiers {
		notifier := notifier
		go func() {
			err := notifier.NotifyFeedbackReceived(fb)
			if err != nil {
				f.emoticInstrumentation.FeedbackNotifyFailed(err)
			}
		}()
	}
	return fb, nil
}

func NewFeedback(
	timer timer.Timer,
	feedbackRepo repository.Feedback,
	keyGen keygen.KeyGenerator,
	notifiers []Notifier,
) Feedback {
	return Feedback{
		timer: timer,
		feedbackRepo: feedbackRepo,
		keyGen: keyGen,
		notifiers:    notifiers,
	}
}
