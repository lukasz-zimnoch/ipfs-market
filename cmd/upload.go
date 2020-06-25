package cmd

import (
	"github.com/lukasz-zimnoch/ipfs-market/pkg/service"
	"github.com/urfave/cli"
	"math/big"
)

const defaultFilePrice = 100000000000000000 //0.1 ETH

var UploadCommand = cli.Command{
	Name:   "upload",
	Action: Upload,
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

// TODO: refactoring:
//  - config extractor
//  - services as objects?
//  - use config object instead of separate params
func Upload(c *cli.Context) error {
	ethNodeUrl := c.GlobalString("eth-node")
	ethPrivateKey := c.GlobalString("eth-private-key")
	ipfsMarketContractAddress := c.GlobalString("ipfs-market-contract")
	storageUrl := c.GlobalString("storage")
	workdir := c.GlobalString("workdir")

	file := c.String("file")
	filePath := workdir + "/" + file
	filePrice := big.NewInt(c.Int64("price"))

	return service.Upload(
		ethNodeUrl,
		ethPrivateKey,
		ipfsMarketContractAddress,
		storageUrl,
		filePath,
		filePrice,
	)
}
