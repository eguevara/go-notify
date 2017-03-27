package notify

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

var _ Notifier = &sesService{}

type sesService struct {
	client sesiface.SESAPI
}

// NewSESService returns a new sesService instance.
func NewSESService(svc sesiface.SESAPI) Notifier {
	return &sesService{
		client: svc,
	}
}

// Notify implements the SES behavior of posting a message.
func (s *sesService) Notify(params interface{}) error {
	sesParams := params.(*SESNotifyInput)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice(sesParams.To),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data:    &sesParams.Message,
					Charset: aws.String("utf-8"),
				},
			},
			Subject: &ses.Content{
				Data:    &sesParams.Subject,
				Charset: aws.String("utf-8"),
			},
		},
		Source: aws.String(sesParams.From),
	}

	resp, err := s.client.SendEmail(input)
	if err != nil {
		return err
	}

	if sesParams.Debug {
		log.Printf("message sent on aws ses: %v", resp.MessageId)
	}

	return nil
}

// SESNotifyInput stores the input parameters to send notification.
type SESNotifyInput struct {
	From    string
	Message string
	To      []string
	Subject string
	Debug   bool
}
