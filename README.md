# go-notify

go-notify is a small Go interface that allows types to implement notifications.


## Slack 
```
client := slack.New(slackToken)
svc := notify.NewSlackService(client)
```

## SES

```
client := ses.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
svc := notify.NewSESService(sesClient)
```

## Usage

```
slackInputParams := &notify.SlackNotifyInput{
    Channel: "mychannel",
    PostMessageParams: slack.PostMessageParameters{
        Username: "mybot",
        Markdown: true,
    },
    Message: result.String(),
}

if err := svc.Notify(slackInputParams); err != nil {
    log.Fatalf("could not call notify on slack service: %v", err)
}
```


```
sesInputParams := &notify.SESNotifyInput{
    From:    "foo@example.com",
    Subject: "subject",
    To:      []string{"baz@example.com"},
    Message: msg.String(),
}

if err := sesSVC.Notify(sesInputParams); err != nil {
    log.Fatalf("count not call notify on ses service: %v", err)
}
```


