package main

import (
	"os"
	//"fmt"
)

var (
	slackAPIToken = os.Getenv("SLACK_BOT_TOKEN")
	slackChannel  = os.Getenv("DOME_CHANNEL")
)

func main() {
	sche := GetSchedule()
	if err := PostMsg(sche); err != nil {
		panic(err)
	}
}
