package main

import (
	"os"

	"github.com/ipfs/go-log"
	"github.com/lukasz-zimnoch/ipfs-market/cmd"
	"github.com/urfave/cli"
)

const (
	defaultConfigPath = "./configs/config.yml"
)

var logger = log.Logger("im-main")

func main() {
	_ = log.SetLogLevel("*", "DEBUG")

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Value: defaultConfigPath,
			Usage: "path to the config YAML file",
		},
	}

	app.Commands = []cli.Command{
		cmd.UploadCommand,
		cmd.PurchaseCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
