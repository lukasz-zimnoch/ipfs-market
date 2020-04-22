package process

import (
	"fmt"

	"github.com/lukasz-zimnoch/ipfs-market/pkg/cipher"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/storage"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/upload"
)

func Upload(storageUrl string, filePath string) error {
	aesGcmCipher, err := cipher.NewAesGcm()
	if err != nil {
		return fmt.Errorf("could not create cipher: [%v]", err)
	}

	ipfsStorage := storage.NewIpfs(storageUrl)

	uploader := upload.NewUploader(aesGcmCipher, ipfsStorage)

	return uploader.Upload(filePath)
}
