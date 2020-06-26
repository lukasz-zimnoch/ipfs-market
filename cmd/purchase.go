package cmd

import (
	"github.com/lukasz-zimnoch/ipfs-market/configs"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/service"
	"github.com/urfave/cli"
)

var PurchaseCommand = cli.Command{
	Name:   "purchase",
	Action: Purchase,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "file-cid,fc",
			Usage: "file CID",
		},
	},
}

func Purchase(c *cli.Context) error {
	config, err := configs.ReadConfig(c.GlobalString("config"))
	if err != nil {
		return err
	}

	fileCid := c.String("file-cid")

	return service.Purchase(config, fileCid)
}
