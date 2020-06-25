package cmd

import (
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

// TODO: refactoring:
//  - config extractor
//  - services as objects?
//  - use config object instead of separate params
func Purchase(c *cli.Context) error {
	ethNodeUrl := c.GlobalString("eth-node")
	ethPrivateKey := c.GlobalString("eth-private-key")
	ipfsMarketContractAddress := c.GlobalString("ipfs-market-contract")
	storageUrl := c.GlobalString("storage")

	fileCid := c.String("file-cid")

	return service.Purchase(
		ethNodeUrl,
		ethPrivateKey,
		ipfsMarketContractAddress,
		storageUrl,
		fileCid,
	)
}