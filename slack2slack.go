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

type Web struct {
	HandlePayload func(*Payload) error
}

// https://api.slack.com/custom-integrations/outgoing-webhooks
func (we *Web) Handle(w http.ResponseWriter, r *http.Request) {
	p, err := request2payload(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	hn := defaultHandlePayload
	if we.HandlePayload != nil {
		hn = we.HandlePayload
	}
	err = hn(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func request2payload(r *http.Request) (*Payload, error) {
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
		return nil, err
	}
	epSec := int64(ts)
	epNsec := int64((ts - float64(epSec)) * float64(time.Nanosecond))
	p.Timestamp = time.Unix(epSec, epNsec)
	return p, nil
}

func defaultHandlePayload(p *Payload) error {
	return nil
}
