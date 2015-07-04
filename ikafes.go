package ikafes

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
	"time"
)

const (
	jsonUrl = "http://s3-ap-northeast-1.amazonaws.com/splatoon-data.nintendo.net/"
)

func Run() int {
	c := cli.NewCLI("ikafes", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"status": func() (cli.Command, error) {
			return &Status{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}

func Now() int64 {
	return time.Now().Unix() / 1000
}
