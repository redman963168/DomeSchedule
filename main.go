package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type yml struct {
	Dome  DomeParameter  `yaml:"dome_params"`
	Slack SlackParameter `yaml:"slack_params"`
}

var y *yml

func init() {
	//yaml読み込み
	buf, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(buf, &y)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))

}
