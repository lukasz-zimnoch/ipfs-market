package cmd

import (
	"fmt"
	"github.com/lukasz-zimnoch/ipfs-market/configs"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/service"
	"github.com/urfave/cli"
	"math/big"
)

const defaultFilePrice = 100000000000000000 //0.1 ETH

var PublishCommand = cli.Command{
	Name:   "publish",
	Action: Publish,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "file,f",
			Usage: "file name",
		},
		&cli.Int64Flag{
			Name:  "price,p",
			Value: defaultFilePrice,
			Usage: "file price in WEI",
		},
	},
}

func Publish(c *cli.Context) error {
	config, err := configs.ReadConfig(c.GlobalString("config"))
	if err != nil {
		return err
	}

	filePath := c.String("file")
	filePrice := big.NewInt(c.Int64("price"))

	fileCid, err := service.Publish(config, filePath, filePrice)
	if err != nil {
		return err
	}

	// pass the file CID to stdout
	fmt.Println(fileCid)

	return nil
}
