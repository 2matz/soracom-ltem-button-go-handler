package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Slack ...
type Slack struct {
	Context context.Context
	Params  params
	URL     string
}

type params struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

// NewSlackWebhookClient ...
// See also: https://api.slack.com/incoming-webhooks
func NewSlackWebhookClient(ctx context.Context, url, text, username, iconEmoji, iconURL, channel string) *Slack {
	// TODO: params validation
	p := params{
		Text:      text,
		Username:  username,
		IconEmoji: iconEmoji,
		IconURL:   iconURL,
		Channel:   channel,
	}

	return &Slack{
		Context: ctx,
		Params:  p,
		URL:     url,
	}
}

// Post execute webhook request.
func (s *Slack) Post() (result string, err error) {
	params, err := json.Marshal(s.Params)
	chkErr(err)
	body := bytes.NewBuffer(params)

	req, err := http.NewRequest(http.MethodPut, s.URL, body)
	req = req.WithContext(s.Context)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	chkErr(err)

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	chkErr(err)

	result = fmt.Sprintf("HTTP Status: %s, Response: %s", resp.Status, string(respBody))

	return
}

func chkErr(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)

		return
	}
}
