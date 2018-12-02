package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/2matz/soracom-ltem-button-go-handler/oneclick"
	"github.com/2matz/soracom-ltem-button-go-handler/slack"
	"github.com/aws/aws-lambda-go/lambda"
)

var logger log.Logger

// HandleRequest ...
func HandleRequest(ctx context.Context, event oneclick.Event) (result string, err error) {
	var slackWebhookURL *url.URL
	attributes := event.GetPlacementAttributes()
	u := attributes.(map[string]interface{})["slack_webhook_url"].(string)
	if u == "" {
		result = ""
		err = fmt.Errorf("%s", "Slack Webhook URL is not defined.")
		return
	}
	slackWebhookURL, err = url.ParseRequestURI(u)
	if err != nil {
		result = ""
		err = fmt.Errorf("%s", "Slack Webhook URL is incorrect format.")
		return
	}

	clickType := event.GetClickType()
	userName := event.GetProjectName()
	if userName == "" {
		userName = "SORACOM LTE-M Button"
	}
	deviceID := event.GetDeviceID()

	slackWebHookClient := slack.NewSlackWebhookClient(
		ctx,
		slackWebhookURL.String(),
		fmt.Sprintf("%s is %s clicked.", deviceID, clickType),
		userName,
		"",
		"",
		"",
	)
	result, err = slackWebHookClient.Post()

	log.Printf("result: %v, error: %v", result, err)

	return
}

func main() {
	lambda.Start(HandleRequest)
}
