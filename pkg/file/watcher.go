package file

import (
	"context"
	"fmt"
)

type Watcher struct {
	authorCipher           Cipher
	purchaserCipherFactory CipherFactory
	chain                  Chain
}

func NewWatcher(
	authorCipher Cipher,
	purchaserCipherFactory CipherFactory,
	chain Chain,
) *Watcher {
	return &Watcher{
		authorCipher:           authorCipher,
		purchaserCipherFactory: purchaserCipherFactory,
		chain:                  chain,
	}
}

// TODO: check past events as well.
func (w *Watcher) Watch(ctx context.Context) error {
	logger.Info("starting watch process")

	purchaseChannel := make(chan *purchase)
	errorChannel := make(chan error)

	err := w.chain.SubscribePurchaseCreatedEvent(
		func(cid string, publicKey []byte) {
			purchaseChannel <- &purchase{cid, publicKey}
		},
		func(err error) {
			errorChannel <- err
		},
	)
	if err != nil {
		return fmt.Errorf(
			"could not subscribe purchase created event: [%v]",
			err,
		)
	}

	logger.Info("subscribed to purchase events")

	for {
		select {
		case purchase := <-purchaseChannel:
			go w.handlePurchase(purchase)
		case err := <-errorChannel:
			return fmt.Errorf(
				"subscription to purchase created event failed: [%v]",
				err,
			)
		case <-ctx.Done():
			logger.Info("watch context is done; terminating process")
			return nil
		}
	}
}

// TODO: retry on error.
func (w *Watcher) handlePurchase(purchase *purchase) {
	logger.Infof(
		"detected purchase of CID [%v]; checking",
		purchase.cid,
	)

	isAuthor, err := w.chain.IsAuthor(purchase.cid)
	if err != nil {
		logger.Errorf(
			"could not determine authorship of purchased "+
				"CID [%v]: [%v]; purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	if !isAuthor {
		logger.Infof(
			"CID [%v] belongs to other author; "+
				"purchase will be ignored",
			purchase.cid,
		)
		return
	}

	authorAccessKey, err := w.chain.GetAccessKey(purchase.cid)
	if err != nil {
		logger.Errorf(
			"could not get author's access key of purchased "+
				"CID [%v]: [%v]; purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	accessKey, err := w.authorCipher.Decrypt(authorAccessKey)
	if err != nil {
		logger.Errorf(
			"could not decrypt author's access key of purchased "+
				"CID [%v]: [%v]; purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	purchaserCipher, err := w.purchaserCipherFactory(purchase.publicKey)
	if err != nil {
		logger.Errorf(
			"could not create purchaser's cipher of purchased "+
				"CID [%v]: [%v]; purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	purchaserAccessKey, err := purchaserCipher.Encrypt(accessKey)
	if err != nil {
		logger.Errorf(
			"could not encrypt purchaser's access key of purchased "+
				"CID [%v]: [%v]; purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	purchaser, err := w.chain.DeriveAddress(purchase.publicKey)
	if err != nil {
		logger.Errorf(
			"could not derive purchaser's address of purchased "+
				"CID [%v]: [%v]; purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	transactionHash, err := w.chain.AnswerPurchase(
		purchase.cid,
		purchaser,
		purchaserAccessKey,
	)
	if err != nil {
		logger.Errorf(
			"could not answer purchase of CID [%v]: [%v]; "+
				"purchase will be ignored",
			purchase.cid,
			err,
		)
		return
	}

	logger.Infof(
		"purchase of CID [%v] has been answered successfully; "+
			"transaction hash on-chain: [%v]",
		purchase.cid,
		transactionHash,
	)

	go w.withdrawPayments()
}

func (w *Watcher) withdrawPayments() {
	logger.Infof("starting payments withdrawal")

	transactionHash, err := w.chain.WithdrawPayments()
	if err != nil {
		logger.Errorf("could not perform payments withdrawal: [%v]", err)
		return
	}

	logger.Infof(
		"payments withdrawal has been completed successfully; "+
			"transaction hash on-chain: [%v]",
		transactionHash,
	)
}

type purchase struct {
	cid       string
	publicKey []byte
}
