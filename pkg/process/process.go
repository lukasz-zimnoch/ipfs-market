package process

import (
	"fmt"

	"github.com/lukasz-zimnoch/ipfs-market/pkg/cipher"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/storage"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/upload"
)

func Upload(storageUrl string, filePath string) error {
	key, err := cipher.GenerateSymmetricKey()
	if err != nil {
		return fmt.Errorf("could not generate symmetric key: [%v]", err)
	}

	aesGcmCipher, err := cipher.NewAesGcm(key)
	if err != nil {
		return fmt.Errorf("could not create cipher: [%v]", err)
	}

	ipfsStorage := storage.NewIpfs(storageUrl)

	uploader := upload.NewUploader(aesGcmCipher, ipfsStorage)

	_, err = uploader.Upload(filePath)
	if err != nil {
		return fmt.Errorf("could not upload file: [%v]", err)
	}

	return nil
}
