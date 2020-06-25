package file

import (
	"fmt"
	"math/big"
)

type Publisher struct {
	cipher Cipher
	chain  Chain
}

func NewPublisher(cipher Cipher, chain Chain) *Publisher {
	return &Publisher{
		cipher: cipher,
		chain:  chain,
	}
}

func (p *Publisher) Publish(cid string, accessKey []byte, price *big.Int) error {
	logger.Infof("starting publishing CID [%v]", cid)

	encryptedAccessKey, err := p.cipher.Encrypt(accessKey[:])
	if err != nil {
		return fmt.Errorf("could not encrypt key: [%v]", err)
	}

	transactionHash, err := p.chain.Publish(cid, encryptedAccessKey, price)
	if err != nil {
		return fmt.Errorf("could not publish CID [%v]: [%v]", cid, err)
	}

	logger.Infof(
		"CID [%v] has been published successfully; "+
			"transaction hash on-chain: [%v]",
		cid,
		transactionHash,
	)

	return nil
}
