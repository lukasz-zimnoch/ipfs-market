package cmd

import (
	"github.com/lukasz-zimnoch/ipfs-market/configs"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/service"
	"github.com/urfave/cli"
)

var WatchCommand = cli.Command{
	Name:   "watch",
	Action: Watch,
}

func Watch(c *cli.Context) error {
	config, err := configs.ReadConfig(c.GlobalString("config"))
	if err != nil {
		return err
	}

	return service.Watch(config)
}
