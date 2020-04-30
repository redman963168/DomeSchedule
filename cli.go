package main

import (
	"fmt"
	"io"
	"os"
)

//SlackのAPIトークン、チャンネル名
var (
	slackAPIToken = os.Getenv("SLACK_BOT_TOKEN")
	slackChannel  = os.Getenv("DOME_CHANNEL")
)

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeError
)

//CLI stdout,stderr
type CLI struct {
	outStream, errStream io.Writer
}

//Run 実処理
func (c *CLI) Run(args []string) int {
	sche, err := GetSchedule()
	if err != nil {
		fmt.Fprintln(c.errStream, os.ErrInvalid, err)
		return ExitCodeError
	}
	if err = PostMsg(sche); err != nil {
		fmt.Fprintln(c.errStream, os.ErrInvalid, err)
		return ExitCodeError
	}
	return ExitCodeOK
}
