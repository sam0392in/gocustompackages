package slackalert

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

func SendAlert(slacktoken string, channelId string, AppName string, message []string, level string, Icon string) error {

	// Color of slack alert
	var color string
	switch level {
	case "critical", "error":
		color = "#ff0000"
	case "warn":
		color = "#ff8c00"
	case "info":
		color = "#008000"
	}

	// Create slackclient
	slackClient := slack.New(slacktoken)

	// Insert "\n" between the message to be sent
	attachmentMessage := strings.Join(message, "\n")

	attachment := slack.Attachment{
		Color: color,
		Text:  attachmentMessage,
	}
	// Send slack message
	_, _, err := slackClient.PostMessage(channelId, slack.MsgOptionAttachments(attachment), slack.MsgOptionIconEmoji(Icon), slack.MsgOptionUsername(AppName))

	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Println("Message successfully sent to channel ", channelId)
	return nil
}
