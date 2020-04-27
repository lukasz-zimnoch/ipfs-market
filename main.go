package main

import (
	"os"

	"github.com/ipfs/go-log"
	"github.com/lukasz-zimnoch/ipfs-market/cmd"
	"github.com/urfave/cli"
)

const (
	defaultStorageUrl  = "http://localhost:5001"
	defaultWorkdirPath = "./workdir"
)

var logger = log.Logger("im-main")

func main() {
	_ = log.SetLogLevel("*", "DEBUG")

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "storage,s",
			Value: defaultStorageUrl,
			Usage: "url to the storage node",
		},
		cli.StringFlag{
			Name:  "workdir,w",
			Value: defaultWorkdirPath,
			Usage: "full path to the working directory",
		},
	}

	app.Commands = []cli.Command{
		cmd.UploadCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
