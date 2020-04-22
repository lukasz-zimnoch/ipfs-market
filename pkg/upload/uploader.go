package upload

import (
	"fmt"
	"io/ioutil"

	"github.com/ipfs/go-log"
)

var logger = log.Logger("im-uploader")

type Cipher interface {
	Encrypt(data []byte) ([]byte, error)
}

type Storage interface {
	Store(data []byte) (string, error)
}

type Uploader struct {
	cipher  Cipher
	storage Storage
}

func NewUploader(cipher Cipher, storage Storage) *Uploader {
	return &Uploader{
		cipher:  cipher,
		storage: storage,
	}
}

func (u *Uploader) Upload(filePath string) error {
	logger.Infof("starting uploading file [%v]", filePath)

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read file [%v]: [%v]", filePath, err)
	}

	// TODO: check if file size does not exceed 1048576 bytes.

	logger.Infof(
		"file [%v] has been read successfully; size: [%v] bytes",
		filePath,
		len(fileBytes),
	)

	fileEncryptedBytes, err := u.cipher.Encrypt(fileBytes)
	if err != nil {
		return fmt.Errorf("could not encrypt file [%v]: [%v]", filePath, err)
	}

	logger.Infof(
		"file [%v] has been encrypted successfully; "+
			"size after encryption: [%v] bytes",
		filePath,
		len(fileEncryptedBytes),
	)

	cid, err := u.storage.Store(fileEncryptedBytes)
	if err != nil {
		return fmt.Errorf("could not store file [%v]: [%v]", filePath, err)
	}

	logger.Infof(
		"file [%v] has been stored successfully; "+
			"content identifier: [%v]",
		filePath,
		cid,
	)

	// TODO: store CID and encrypted key on-chain.

	return nil
}
