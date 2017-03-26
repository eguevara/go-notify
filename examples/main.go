package main

import (
	"log"
	"os"

	"github.com/eguevara/go-notify"
	"github.com/nlopes/slack"
)

func main() {

	// Get api token from environment variables.
	slackToken := os.Getenv("SLACK_API_TOKEN")
	if slackToken == "" {
		log.Fatal("slack api token is required. (hint: export SLACK_API_TOKEN)")
	}

	// Creates a nlopes.slack client using api token.
	client := slack.New(slackToken)

	// Create a SlackService that implements Notifier interface.
	svc := notify.NewSlackService(client)

	// Configure message to be sent.
	msg := &notify.Message{
		Title:    "title",
		Text:     "text",
		Endpoint: "slack.main",
	}

	// Contains the parameters for slacknotify.
	slackInputParams := &notify.SlackNotifyInput{
		Channel: "devops",
		PostMessageParams: slack.PostMessageParameters{
			Username: "mbot",
			Markdown: true,
		},
		Message: msg.String(),
		Debug:   true,
	}

	// Calls Notify method on service.
	if err := svc.Notify(slackInputParams); err != nil {
		log.Fatalf("could not call notify on slack service: %v", err)
	}
}
