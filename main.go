package main

import (
	"os"
	//"github.com/go-yaml/yaml"
)

var (
	dome  *DomeEventPage
	slack *SlackParameter
)

func init() {
	dome = &DomeEventPage{
		url: "https://www.kyoceradome-osaka.jp/events/?yearId=%YEAR%&monthId=%MONTH%", //yamlに変更
	}
	slack = &SlackParameter{
		token:   os.Getenv("SLACK_BOT_TOKEN"),             //yamlに変更
		channel: os.Getenv("DOME_CHANNEL"),                //yamlに変更
		url:     "https://slack.com/api/chat.postMessage", //yamlに変更
	}
}

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))

}
