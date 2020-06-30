package slack

import (
	"fmt"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/fw/slackapi"
	"github.com/short-d/short/backend/app/usecase/emotic"
	"time"
)

var _ emotic.Notifier = (*EmoticNotifier)(nil)

type EmoticNotifier struct {
	slackWebHookURL string
	slack           slackapi.Slack
}

func (e EmoticNotifier) NotifyFeedbackReceived(feedback entity.Feedback) error {
	message := newFeedbackReceivedMessage(feedback)
	return e.slack.SendMessage(e.slackWebHookURL, message)
}

func newFeedbackReceivedMessage(feedback entity.Feedback) string {
	const feedbackReceivedTemplate = `
*App ID*: %s
*Feedback ID*: %s
*Customer rating*: %d
*Customer comment*: %v
*Customer email*: %v
*Received at*: %s
`
	receivedAt := feedback.ReceivedAt.Format(time.RFC1123)
	comment := ""
	if feedback.Comment != nil {
		comment = *feedback.Comment
	}
	email := ""
	if feedback.CustomerEmail != nil {
		email = *feedback.CustomerEmail
	}
	return fmt.Sprintf(
		feedbackReceivedTemplate,
		feedback.AppID,
		feedback.FeedbackID,
		feedback.CustomerRating,
		comment,
		email,
		receivedAt,
	)
}

type EmoticNotifierFactory struct {
	slack slackapi.Slack
}

func (s EmoticNotifierFactory) NewEmoticNotifier(slackWebHookURL string) EmoticNotifier {
	return EmoticNotifier{
		slackWebHookURL: slackWebHookURL,
		slack:           s.slack,
	}
}

func NewEmoticNotifierFactory(
	slack slackapi.Slack,
) EmoticNotifierFactory {
	return EmoticNotifierFactory{
		slack: slack,
	}
}
