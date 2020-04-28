package main

import (
	"log"
	"os"
)

var (
	slackAPIToken = os.Getenv("SLACK_BOT_TOKEN")
	slackChannel  = os.Getenv("DOME_CHANNEL")
)

func main() {
	sche, err := GetSchedule()
	if err != nil {
		log.Fatal(err)
	}
	if err = PostMsg(sche); err != nil {
		log.Fatal(err)
	}
}
