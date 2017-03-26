# go-notify

go-notify is a small Go interface that allows types to implement notifications.


## Slack 
```
client := slack.New(slackToken)
svc := notify.NewSlackService(client)
```

## SES

wip


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

