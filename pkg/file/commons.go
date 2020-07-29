package file

import "math/big"
import "github.com/ipfs/go-log"

var logger = log.Logger("im-file")

type Cipher interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

type CipherFactory func(publicKey []byte) (Cipher, error)

type Storage interface {
	Store(data []byte) (string, error)
	GetData(cid string) ([]byte, error)
}

type Chain interface {
	Publish(cid string, accessKey []byte, price *big.Int) (string, error)
	Purchase(cid string, publicKey []byte) (string, error)
	AnswerPurchase(cid string, purchaser string, accessKey []byte) (string, error)
	HasPurchased(cid string) (bool, error)
	GetAccessKey(cid string) ([]byte, error)
	IsAuthor(cid string) (bool, error)
	DeriveAddress(publicKeyBytes []byte) (string, error)

	SubscribePurchaseAnsweredEvent(
		onEvent func(cid, purchaser string, accessKey []byte),
		onError func(error),
	) error
	SubscribePurchaseCreatedEvent(
		onEvent func(cid string, publicKey []byte),
		onError func(error),
	) error
}
