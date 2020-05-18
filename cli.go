package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var version = "0.0.1"

//引数
var (
	verOpt  = flag.Bool("version", false, "Show version info")
  )

//SlackのAPIトークン、チャンネル名
var (
	slackAPIToken = os.Getenv("SLACK_BOT_TOKEN")
	slackChannel  = os.Getenv("DOME_CHANNEL")
)

// 終了コード 0,1...
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
	//version
	flag.Parse()
	if *verOpt  {
		fmt.Fprintln(c.outStream, version)
		return ExitCodeOK
	}
	//スケジュール取得
	sche, err := GetSchedule()
	if err != nil {
		fmt.Fprintln(c.errStream, os.ErrInvalid, err)
		return ExitCodeError
	}
	//スケジュールPOST
	if err = PostMsg(sche); err != nil {
		fmt.Fprintln(c.errStream, os.ErrInvalid, err)
		return ExitCodeError
	}
	return ExitCodeOK
}
