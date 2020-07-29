package service

import (
	"context"
	"fmt"
	"github.com/lukasz-zimnoch/ipfs-market/configs"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/cipher"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/file"
)

func Watch(config *configs.Config) error {
	ethPrivateKeyBytes, err := chain.DecodeEthPrivateKey(config.Ethereum.PrivateKey)
	if err != nil {
		return fmt.Errorf("could not decode ethereum private key: [%v]", err)
	}

	eciesCipher := cipher.NewEcies(ethPrivateKeyBytes)

	eciesCipherFactory := func(publicKey []byte) (file.Cipher, error) {
		return cipher.NewEciesFromPublicKey(chain.AddPublicKeyPrefix(publicKey))
	}

	ethereum, err := chain.NewEthereumClient(
		config.Ethereum.URL,
		ethPrivateKeyBytes,
		config.Ethereum.IpfsMarketContract,
	)
	if err != nil {
		return fmt.Errorf("could not create ethereum client: [%v]", err)
	}

	watcher := file.NewWatcher(eciesCipher, eciesCipherFactory, ethereum)

	err = watcher.Watch(context.TODO())
	if err != nil {
		return fmt.Errorf("could not trigger watch process: [%v]", err)
	}

	return nil
}
