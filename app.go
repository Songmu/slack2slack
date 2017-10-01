package slack2slack

import (
	"fmt"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

type App struct {
	WebhookURL string
}

func (app *App) forward(p *Payload) error {
	slackPayload := slack.Payload{
		Text: fmt.Printf("Team: %s, Chanel: %s User: %s, Msg: %s\n",
			p.TeamDomain, p, ChannelName, p.UserName, p.Text),
	}
	errs := slack.Send(app.WebhookURL, "", slackPayload)
	if len(errs) > 0 {
		return fmt.Errorf("error occured while sending payload to slack incomming hook: %#v", err)
	}
	return nil
}
