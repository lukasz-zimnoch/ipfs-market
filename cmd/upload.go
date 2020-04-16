package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/ipfs/go-log"
	"github.com/urfave/cli"
)

var logger = log.Logger("im-cmd")

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

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not get file [%v]: [%v]", filePath, err)
	}

	logger.Infof("file has [%v] bytes", len(fileBytes))

	return nil
}
