package slack2slack

import (
	"fmt"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

type App struct {
	WebhookURL string
	Channel    string
	UserName   string
	IconURL    string
	IconEmoji  string
}

func pstr(s string) *string {
	return &s
}

func (app *App) forward(p *Payload) error {
	att := slack.Attachment{
		Title:     pstr(fmt.Sprintf("%s#%s@%s", p.TeamDomain, p.ChannelName, p.UserName)),
		TitleLink: pstr(fmt.Sprintf("https://%s.slack.com/archives/%s/", p.TeamDomain, p.ChannelID)),
		Text:      pstr(p.Text),
	}
	slackPayload := slack.Payload{
		Username:    app.UserName,
		IconUrl:     app.IconURL,
		IconEmoji:   app.IconEmoji,
		Attachments: []slack.Attachment{att},
	}
	errs := slack.Send(app.WebhookURL, "", slackPayload)
	if len(errs) > 0 {
		return fmt.Errorf("error occured while sending payload to slack incomming hook: %#v", errs)
	}
	return nil
}
