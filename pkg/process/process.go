package process

import (
	"fmt"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/cipher"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/file"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/storage"
	"math/big"
)

func Upload(
	ethNodeUrl string,
	ethPrivateKey string,
	ipfsMarketContractAddress string,
	storageUrl string,
	filePath string,
	filePrice *big.Int,
) error {
	key, err := cipher.GenerateSymmetricKey()
	if err != nil {
		return fmt.Errorf("could not generate symmetric key: [%v]", err)
	}

	aesGcmCipher, err := cipher.NewAesGcm(key)
	if err != nil {
		return fmt.Errorf("could not create AES-GCM cipher: [%v]", err)
	}

	ipfsStorage := storage.NewIpfs(storageUrl)

	uploader := file.NewUploader(aesGcmCipher, ipfsStorage)

	fileCid, err := uploader.Upload(filePath)
	if err != nil {
		return fmt.Errorf("could not upload file: [%v]", err)
	}

	ethPrivateKeyBytes, err := chain.DecodeEthPrivateKey(ethPrivateKey)
	if err != nil {
		return fmt.Errorf("could not decode ethereum private key: [%v]", err)
	}

	eciesCipher := cipher.NewEcies(ethPrivateKeyBytes)

	ethereum, err := chain.NewEthereumClient(
		ethNodeUrl,
		ethPrivateKeyBytes,
		ipfsMarketContractAddress,
	)
	if err != nil {
		return fmt.Errorf("could not create ethereum client: [%v]", err)
	}

	publisher := file.NewPublisher(eciesCipher, ethereum)

	err = publisher.Publish(fileCid, key[:], filePrice)
	if err != nil {
		return fmt.Errorf("could not publish file: [%v]", err)
	}

	return nil
}
