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
	sche, err := GetSchedule()
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}

	//スケジュールPOST
	if err = PostMsg(sche); err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}

	return ExitCodeOK
}
