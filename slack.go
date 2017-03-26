package notify

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
)

var _ SLACKAPI = &slack.Client{}

// SLACKAPI is a interface wrapper around Slack.Client
type SLACKAPI interface {
	PostMessage(string, string, slack.PostMessageParameters) (string, string, error)
	DeleteMessage(string, string) (string, string, error)
	UpdateMessage(string, string, string) (string, string, string, error)
}

var _ Notifier = &slackService{}

type slackService struct {
	client SLACKAPI
}

// NewSlackService returns a new slackService instance.
func NewSlackService(svc SLACKAPI) Notifier {
	return &slackService{
		client: svc,
	}
}

// Notify implements the Slack behavior of posting a message
func (s *slackService) Notify(params interface{}) error {
	slackParams := params.(*SlackNotifyInput)

	channel, timestamp, err := s.client.PostMessage(slackParams.Channel, slackParams.Message, slackParams.PostMessageParams)
	if err != nil {
		return fmt.Errorf("error posting message: %v", err)
	}

	if slackParams.Debug {
		log.Printf("message sent on channel %s at %v", channel, timestamp)
	}

	return nil
}

// SlackNotifyInput stores the input parameters to send notification.
type SlackNotifyInput struct {
	Channel           string
	Debug             bool
	Message           string
	PostMessageParams slack.PostMessageParameters
}
