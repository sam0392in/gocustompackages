# Slack Alert

![logo](slack.png "icon")

- This package will help you to send the slack messages without populating your main code.
- Just call this package in main file and pass required parameters as shown in below sample code.

## Usage

```
package main

import (
	"fmt"

	"github.com/sam0392in/gocustompackages/slackalert"
)

func main() {
	var data []string
	slacktoken := "__YOUR__SLACK_TOKEN__"

    // message level: "critical", "error", "warn", "info"
    level := "info"
	Icon := ":kubernetes"
	AppName := "sam-test"

    // ChannelId you will get from login to slack from browser and click on channel.
    // In address bar you will see like : https://app.slack.com/client/_CLIENT_ID__/__CHANNEL_ID__
	channelId := "__CHANNEL_ID__"

    // Data to be sent as message in slack channel
	message := "test message"
	status := "test status"
	events := "test event"
	logs := "test log"

	inputs := []string{message, status, events, logs}

	for _, input := range inputs {
		if input != "" {
			data = append(data, input)
		}
	}

    // Call slackalert module and pass below required params
	err := slackalert.SendAlert(slacktoken, channelId, AppName, data, level, Icon)

	if err != nil {
		fmt.Println(err)
	}
}

```