package file

import (
	"context"
	"fmt"
	"time"
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

	accessKey, err := p.pollAccessKey(cid)
	if err != nil {
		return nil, fmt.Errorf(
			"could not get access key for CID [%v]: [%v]",
			cid,
			err,
		)
	}

	return p.cipher.Decrypt(accessKey)
}

func (p *Purchaser) pollAccessKey(cid string) ([]byte, error) {
	//TODO: configurable poll parameters.
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancelCtx()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logger.Infof(
				"checking author's response for CID [%v] purchase",
				cid,
			)

			if accessKey, exists := p.checkAccessKey(cid); exists {
				logger.Infof(
					"author responded for CID [%v] purchase",
					cid,
				)
				return accessKey, nil
			}

			logger.Infof(
				"author didn't respond for CID [%v] purchase yet; "+
					"next check will be performed after a while",
				cid,
			)
		case <-ctx.Done():
			return nil, fmt.Errorf(
				"author didn't respond for CID [%v] purchase "+
					"at given time window; please try again later",
				cid,
			)
		}
	}
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
