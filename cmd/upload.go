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

func Upload(c *cli.Context) error {
	workdir := c.GlobalString("workdir")
	file := c.String("file")
	filePath := workdir + "/" + file

	//TODO: pass storage url from config.
	return process.Upload("", filePath)
}
