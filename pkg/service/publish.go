package service

import (
	"fmt"
	"github.com/lukasz-zimnoch/ipfs-market/configs"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/cipher"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/file"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/storage"
	"math/big"
)

func Publish(
	config *configs.Config,
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

	ipfsStorage := storage.NewIpfs(config.Storage.URL)

	uploader := file.NewUploader(aesGcmCipher, ipfsStorage)

	fileCid, err := uploader.Upload(filePath)
	if err != nil {
		return fmt.Errorf("could not upload file: [%v]", err)
	}

	ethPrivateKeyBytes, err := chain.DecodeEthPrivateKey(config.Ethereum.PrivateKey)
	if err != nil {
		return fmt.Errorf("could not decode ethereum private key: [%v]", err)
	}

	eciesCipher := cipher.NewEcies(ethPrivateKeyBytes)

	ethereum, err := chain.NewEthereumClient(
		config.Ethereum.URL,
		ethPrivateKeyBytes,
		config.Ethereum.IpfsMarketContract,
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
