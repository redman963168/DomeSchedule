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

//Run 実処理
func (c *CLI) Run(args []string) int {
	//引数設定
	flg := flag.NewFlagSet("cli", flag.ContinueOnError)
	var (
		verOpt = flg.Bool("version", false, "Show version info")
	)
	flg.Parse(args[1:])

	//バージョン表示
	if *verOpt {
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
