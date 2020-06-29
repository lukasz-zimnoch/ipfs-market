package file

import (
	"fmt"
)

type Purchaser struct {
	cipher Cipher
	chain  Chain
}

func NewPurchaser(cipher Cipher, chain Chain) *Purchaser {
	return &Purchaser{
		cipher: cipher,
		chain:  chain,
	}
}

func (p *Purchaser) Purchase(
	cid string,
	publicKey []byte,
) ([]byte, error) {
	logger.Infof("starting purchasing CID [%v]", cid)

	hasPurchased, err := p.chain.HasPurchased(cid)
	if err != nil {
		return nil, fmt.Errorf(
			"could not check existing purchase of CID [%v]: [%v]",
			cid,
			err,
		)
	}

	if hasPurchased {
		logger.Infof(
			"CID [%v] purchase already exists; checking author's response",
			cid,
		)
	} else {
		transactionHash, err := p.chain.Purchase(cid, publicKey)
		if err != nil {
			return nil, fmt.Errorf("could not purchase CID [%v]: [%v]", cid, err)
		}

		logger.Infof(
			"CID [%v] has been purchased successfully; "+
				"transaction hash on-chain: [%v]; waiting for author's response",
			cid,
			transactionHash,
		)
	}

	accessKey, exists := p.checkAccessKey(cid)

	if !exists {
		accessKey, err = p.waitForAccessKey(cid, publicKey)
		if err != nil {
			return nil, fmt.Errorf(
				"could not get access key for CID [%v]: [%v]",
				cid,
				err,
			)
		}
	}

	return p.cipher.Decrypt(accessKey)
}

func (p *Purchaser) checkAccessKey(cid string) ([]byte, bool) {
	accessKey, err := p.chain.GetAccessKey(cid)
	if err != nil {
		return nil, false
	}

	if len(accessKey) == 0 {
		return accessKey, false
	}

	return accessKey, true
}

func (p *Purchaser) waitForAccessKey(
	cid string,
	publicKey []byte,
) ([]byte, error) {
	address, err := p.chain.DeriveAddress(publicKey)
	if err != nil {
		return nil, err
	}

	accessKeyChannel := make(chan []byte)
	errorChannel := make(chan error)

	err = p.chain.SubscribePurchaseAnsweredEvent(
		func(eventCid, purchaser string, accessKey []byte) {
			if cid == eventCid && address == purchaser {
				accessKeyChannel <- accessKey
			}
		},
		func(err error) {
			errorChannel <- err
		},
	)
	if err != nil {
		return nil, err
	}

	select {
	case accessKey := <-accessKeyChannel:
		return accessKey, nil
	case err := <-errorChannel:
		return nil, err
	}
}
