package slack2slack

import (
	"os"
	"testing"
)

const webhookURLEnv = "SLACK2SLACK_WEBHOOK_URL"

func TestSend(t *testing.T) {
	u := os.Getenv(webhookURLEnv)
	if u == "" {
		t.Skipf("please set %s environment variable for testing", webhookURLEnv)
	}
	app := App{WebhookURL: u}
	p := &Payload{
		Token:       "XXXXXXXX",
		TeamID:      "T001",
		TeamDomain:  "example-songmu",
		ChannelID:   "C1235",
		ChannelName: "general",
		UserID:      "U1122",
		UserName:    "songmu-hhh",
		Text:        "テスト",
		TriggerWord: "",
	}
	err := app.forward(p)
	if err != nil {
		t.Errorf("error should be nil but: %s", err)
	}
}
