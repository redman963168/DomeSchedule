package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var version = "0.0.1"

// 終了コード 0,1...
const (
	ExitCodeOK = iota
	ExitCodeError
)

//CLI stdout,stderr
type CLI struct {
	outStream, errStream io.Writer
}

//DomeEventPage ドームのURL
type DomeEventPage struct {
	url string
}

//SlackParameter slackAPIに必要な情報
type SlackParameter struct{
	token string
	channel string
	url string
}

//Run 実処理
func (c *CLI) Run(args []string) int {
	//引数設定
	var showVersion bool
	flg := flag.NewFlagSet("cli", flag.ContinueOnError)
	flg.SetOutput(c.errStream)
	flg.BoolVar(&showVersion, "version", false, "Show version info")
	if err := flg.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.errStream, os.ErrInvalid, err)
		return ExitCodeError
	}

	//バージョン表示
	if showVersion {
		fmt.Fprintln(c.outStream, version)
		return ExitCodeOK
	}

	//スケジュール取得
	sche, err := dome.GetSchedule()
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}

	//スケジュールPOST
	if err = slack.PostMsg(sche); err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}

	return ExitCodeOK
}
