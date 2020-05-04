package cmd

import (
	"github.com/lukasz-zimnoch/ipfs-market/pkg/process"
	"github.com/urfave/cli"
)

var UploadCommand = cli.Command{
	Name:   "upload",
	Action: Upload,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "file,f",
			Usage: "file name",
		},
	},
}

// TODO 1: refactoring:
//  - config extractor
//  - processes as objects?
//  - use config object instead of separate params
func Upload(c *cli.Context) error {
	ethNodeUrl := c.GlobalString("eth-node")
	ethPrivateKey := c.GlobalString("eth-private-key")
	storageUrl := c.GlobalString("storage")
	workdir := c.GlobalString("workdir")
	file := c.String("file")
	filePath := workdir + "/" + file

	return process.Upload(ethNodeUrl, ethPrivateKey, storageUrl, filePath)
}
