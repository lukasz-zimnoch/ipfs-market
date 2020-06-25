package file

import (
	"fmt"
	"io/ioutil"
)

const maxFileByteSize = 1048576

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

func (u *Uploader) Upload(filePath string) (string, error) {
	logger.Infof("starting uploading file [%v]", filePath)

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("could not read file [%v]: [%v]", filePath, err)
	}

	if len(fileBytes) > maxFileByteSize {
		return "", fmt.Errorf(
			"file [%v] exceeds maximum file size of [%v] bytes",
			filePath,
			maxFileByteSize,
		)
	}

	logger.Infof(
		"file [%v] has been read successfully; size: [%v] bytes",
		filePath,
		len(fileBytes),
	)

	fileEncryptedBytes, err := u.cipher.Encrypt(fileBytes)
	if err != nil {
		return "", fmt.Errorf("could not encrypt file [%v]: [%v]", filePath, err)
	}

	logger.Infof(
		"file [%v] has been encrypted successfully; "+
			"size after encryption: [%v] bytes",
		filePath,
		len(fileEncryptedBytes),
	)

	fileCid, err := u.storage.Store(fileEncryptedBytes)
	if err != nil {
		return "", fmt.Errorf("could not store file [%v]: [%v]", filePath, err)
	}

	logger.Infof(
		"file [%v] has been stored successfully; "+
			"content identifier: [%v]",
		filePath,
		fileCid,
	)

	return fileCid, nil
}
