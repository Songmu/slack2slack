package slack2slack

import (
	"net/http"
	"strconv"
	"time"
)

type Payload struct {
	Token       string
	TeamID      string
	TeamDomain  string
	ChannelID   string
	ChannelName string
	Timestamp   time.Time
	UserID      string
	UserName    string
	Text        string
	TriggerWord string
}

// https://api.slack.com/custom-integrations/outgoing-webhooks
func assignPayloadMessage(w http.ResponseWriter, r *http.Request) {
	p := &Payload{
		Token:       r.FormValue("token"),
		TeamID:      r.FormValue("team_id"),
		TeamDomain:  r.FormValue("team_domain"),
		ChannelID:   r.FormValue("channel_id"),
		ChannelName: r.FormValue("channel_name"),
		UserID:      r.FormValue("user_id"),
		UserName:    r.FormValue("user_name"),
		Text:        r.FormValue("text"),
		TriggerWord: r.FormValue("trigger_word"),
	}
	ts, err := strconv.ParseFloat(r.FormValue("timestamp"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	epSec := int64(ts)
	epNsec := int64((ts - float64(epSec)) * float64(time.Nanosecond))
	p.Timestamp = time.Unix(epSec, epNsec)
	handle(p)
	return
}

func handle(p *Payload) error {
	return nil
}
