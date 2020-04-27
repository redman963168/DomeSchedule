package main

import (
	"os"
	//"fmt"
)

var (
	slackAPIToken = os.Getenv("SLACK_BOT_TOKEN")
	slackChannel  = "apitテスト"
)

func main() {

	sche := GetSchedule()
	//fmt.Println(sche)
	if err := PostMsg(sche); err != nil {
		panic(err)
	}
	//run()
	//
}
