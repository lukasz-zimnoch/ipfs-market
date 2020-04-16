package main

import (
	"os"

	"github.com/ipfs/go-log"
	"github.com/urfave/cli"
)

const defaultWorkdirPath = "./workdir"

var (
	logger      = log.Logger("main")
	workdirPath string
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "workdir,w",
			Value:       defaultWorkdirPath,
			Destination: &workdirPath,
			Usage:       "full path to the working directory",
		},
	}

	app.Commands = []cli.Command{
		// TODO: add commands.
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
