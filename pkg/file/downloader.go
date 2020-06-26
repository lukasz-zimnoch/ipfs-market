package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Downloader struct {
	cipher  Cipher
	storage Storage
	workdir string
}

func NewDownloader(cipher Cipher, storage Storage, workdir string) *Downloader {
	return &Downloader{
		cipher:  cipher,
		storage: storage,
	}
}

func (d *Downloader) Download(cid string) error {
	logger.Infof("starting downloading CID [%v]", cid)

	data, err := d.storage.GetData(cid)
	if err != nil {
		return fmt.Errorf(
			"could not get CID [%v] from storage: [%v]",
			cid,
			err,
		)
	}

	decryptedData, err := d.cipher.Decrypt(data)
	if err != nil {
		return fmt.Errorf(
			"could not decrypt CID [%v] data: [%v]",
			cid,
			err,
		)
	}

	err = ioutil.WriteFile(d.workdir+"/"+cid, decryptedData, os.ModePerm)
	if err != nil {
		return fmt.Errorf(
			"could not write file for CID [%v]: [%v]",
			cid,
			err,
		)
	}

	logger.Infof("CID [%v] has been downloaded successfully", cid)

	return nil
}
