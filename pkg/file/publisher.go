package file

import "fmt"

type Chain interface {
	Publish(cid string, accessKey []byte)
}

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

func (p *Publisher) Publish(cid string, accessKey []byte) error {
	encryptedAccessKey, err := p.cipher.Encrypt(accessKey[:])
	if err != nil {
		return fmt.Errorf("could not encrypt key: [%v]", err)
	}

	p.chain.Publish(cid, encryptedAccessKey)

	return nil
}
