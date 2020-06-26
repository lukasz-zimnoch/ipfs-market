package service

import (
	"fmt"
	"github.com/lukasz-zimnoch/ipfs-market/configs"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/cipher"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/file"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/storage"
)

func Purchase(config *configs.Config, fileCid string) error {
	ethPrivateKeyBytes, err := chain.DecodeEthPrivateKey(config.Ethereum.PrivateKey)
	if err != nil {
		return fmt.Errorf("could not decode ethereum private key: [%v]", err)
	}

	eciesCipher := cipher.NewEcies(ethPrivateKeyBytes)

	ethPublicKeyBytes, err := chain.DeriveEthPublicKey(config.Ethereum.PrivateKey)
	if err != nil {
		return fmt.Errorf("could not derive ethereum public key: [%v]", err)
	}

	ethereum, err := chain.NewEthereumClient(
		config.Ethereum.URL,
		ethPrivateKeyBytes,
		config.Ethereum.IpfsMarketContract,
	)
	if err != nil {
		return fmt.Errorf("could not create ethereum client: [%v]", err)
	}

	purchaser := file.NewPurchaser(eciesCipher, ethereum)

	accessKey, err := purchaser.Purchase(fileCid, ethPublicKeyBytes)
	if err != nil {
		return fmt.Errorf("could not purchase file: [%v]", err)
	}

	aesGcmCipher, err := cipher.NewAesGcm(
		cipher.SymmetricKeyFromSlice(accessKey),
	)
	if err != nil {
		return fmt.Errorf("could not create AES-GCM cipher: [%v]", err)
	}

	ipfsStorage := storage.NewIpfs(config.Storage.URL)

	downloader := file.NewDownloader(aesGcmCipher, ipfsStorage, config.Storage.Workdir)

	err = downloader.Download(fileCid)
	if err != nil {
		return fmt.Errorf("could not download file: [%v]", err)
	}

	return nil
}
