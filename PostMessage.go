package main

import (
	"github.com/slack-go/slack"
)

//PostMsg Slackにメッセージを送る
func PostMsg(schedules []Schedule) error {
	msg := "【本日の京セラドームの予定】"
	api := slack.New(slackAPIToken)
	for _, sche := range schedules {
		msg = msg + "\n" + sche.title
		msg = msg + "\n" + sche.dateInfo

	}
	_, _, err := api.PostMessage(
		slackChannel,
		slack.MsgOptionText(msg, false))
	return err
}
