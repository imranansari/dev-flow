package chat

import (
	"os"
	
	"github.com/nlopes/slack"
)

type Slack struct{}

func (s Slack) client() *slack.Client {
	return slack.New(os.Getenv("SLACK_API_TOKEN"))
}

func (s Slack) getUserID(username string) (string) {
	client := s.client()

	users, err := client.GetUsers()

	if err != nil {
		panic(err)
	}

	var userID string

	for _, slackuser := range users {
		if slackuser.Name == username {
			userID = slackuser.ID
		}
	}
	
	return userID
}

func (s Slack) DirectMessage(username string, message string) {
	userID := s.getUserID(username)

	client := s.client()
	
	_, _, channelID, err := client.OpenIMChannel(userID)

	if err != nil {
		panic(err)
	}
	
	client.PostMessage(channelID, message, slack.PostMessageParameters{})
}