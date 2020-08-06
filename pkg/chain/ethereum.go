package chain

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain/ethereum_gen/contracts"
	"math/big"
	"strings"
	"sync"
)

var logger = log.Logger("im-eth")

type EthereumClient struct {
	client             *ethclient.Client
	transactOpts       *bind.TransactOpts
	callOpts           *bind.CallOpts
	ipfsMarketContract *contracts.IpfsMarket
}

func NewEthereumClient(
	url string,
	privateKeyBytes []byte,
	ipfsMarketContractAddress string,
) (*EthereumClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	transactOpts := bind.NewKeyedTransactor(privateKey)

	callOpts := &bind.CallOpts{
		From: transactOpts.From,
	}

	ipfsMarketContract, err := contracts.NewIpfsMarket(
		common.HexToAddress(ipfsMarketContractAddress),
		client,
	)
	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		client:             client,
		transactOpts:       transactOpts,
		callOpts:           callOpts,
		ipfsMarketContract: ipfsMarketContract,
	}, nil
}

func (ec *EthereumClient) Publish(
	cid string,
	accessKey []byte,
	price *big.Int,
) (string, error) {
	transaction, err := ec.ipfsMarketContract.Publish(
		ec.transactOpts,
		cid,
		accessKey,
		price,
	)
	if err != nil {
		return "", err
	}

	return transaction.Hash().Hex(), nil
}

func (ec *EthereumClient) Purchase(
	cid string,
	publicKey []byte,
) (string, error) {
	price, err := ec.ipfsMarketContract.GetPrice(ec.callOpts, cid)
	if err != nil {
		return "", err
	}

	// Copy the original transactOpts.
	tempTransactOpts := new(bind.TransactOpts)
	*tempTransactOpts = *ec.transactOpts
	tempTransactOpts.Value = price

	transaction, err := ec.ipfsMarketContract.Purchase(
		tempTransactOpts,
		cid,
		RemovePublicKeyPrefix(publicKey),
	)
	if err != nil {
		return "", err
	}

	return transaction.Hash().Hex(), nil
}

func (ec *EthereumClient) AnswerPurchase(
	cid string,
	purchaser string,
	accessKey []byte,
) (string, error) {
	transaction, err := ec.ipfsMarketContract.AnswerPurchase(
		ec.transactOpts,
		cid,
		common.HexToAddress(purchaser),
		accessKey,
	)
	if err != nil {
		return "", err
	}

	return transaction.Hash().Hex(), nil
}

func (ec *EthereumClient) HasPurchased(cid string) (bool, error) {
	return ec.ipfsMarketContract.HasPurchased(ec.callOpts, cid)
}

func (ec *EthereumClient) GetAccessKey(cid string) ([]byte, error) {
	return ec.ipfsMarketContract.GetAccessKey(ec.callOpts, cid)
}

func (ec *EthereumClient) IsAuthor(cid string) (bool, error) {
	return ec.ipfsMarketContract.IsAuthor(ec.callOpts, cid)
}

func (ec *EthereumClient) DeriveAddress(publicKeyBytes []byte) (string, error) {
	publicKey, err := crypto.UnmarshalPubkey(AddPublicKeyPrefix(publicKeyBytes))
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*publicKey).Hex(), nil
}

func (ec *EthereumClient) WithdrawPayments() (string, error) {
	transaction, err := ec.ipfsMarketContract.WithdrawPayments(
		ec.transactOpts,
		ec.transactOpts.From,
	)

	if err != nil {
		return "", err
	}

	return transaction.Hash().Hex(), nil
}

func (ec *EthereumClient) SubscribePurchaseAnsweredEvent(
	onEvent func(cid, purchaser string, accessKey []byte),
	onError func(error),
) error {
	eventChannel := make(chan *contracts.IpfsMarketPurchaseAnswered)

	eventSubscription, err := ec.ipfsMarketContract.WatchPurchaseAnswered(
		nil,
		eventChannel,
	)
	if err != nil {
		close(eventChannel)
		return fmt.Errorf(
			"error creating watch for PurchaseAnswered events: [%v]",
			err,
		)
	}

	var mutex sync.Mutex

	go func() {
		for {
			select {
			case event := <-eventChannel:
				mutex.Lock()

				onEvent(
					event.Cid,
					event.Purchaser.Hex(),
					event.AccessKey,
				)

				mutex.Unlock()
			case err := <-eventSubscription.Err():
				onError(err)
				return
			}
		}
	}()

	return nil
}

func (ec *EthereumClient) SubscribePurchaseCreatedEvent(
	onEvent func(cid string, publicKey []byte),
	onError func(error),
) error {
	eventChannel := make(chan *contracts.IpfsMarketPurchaseCreated)

	eventSubscription, err := ec.ipfsMarketContract.WatchPurchaseCreated(
		nil,
		eventChannel,
	)
	if err != nil {
		close(eventChannel)
		return fmt.Errorf(
			"error creating watch for PurchaseCreated events: [%v]",
			err,
		)
	}

	var mutex sync.Mutex

	go func() {
		for {
			select {
			case event := <-eventChannel:
				mutex.Lock()

				onEvent(
					event.Cid,
					event.PublicKey,
				)

				mutex.Unlock()
			case err := <-eventSubscription.Err():
				onError(err)
				return
			}
		}
	}()

	return nil
}

func DecodeEthPrivateKey(privateKeyHex string) ([]byte, error) {
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = privateKeyHex[2:]
	}

	return hex.DecodeString(privateKeyHex)
}

func DeriveEthPublicKey(privateKeyHex string) ([]byte, error) {
	privateKeyBytes, err := DecodeEthPrivateKey(privateKeyHex)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return crypto.FromECDSAPub(&privateKey.PublicKey), nil
}

func RemovePublicKeyPrefix(publicKey []byte) []byte {
	if len(publicKey) == 65 && publicKey[0] == 0x04 {
		return publicKey[1:]
	}

	return publicKey
}

func AddPublicKeyPrefix(publicKey []byte) []byte {
	if len(publicKey) == 64 {
		return append([]byte{0x04}, publicKey...)
	}

	return publicKey
}
