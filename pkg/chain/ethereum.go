package chain

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain/ethereum_gen/contracts"
	"math/big"
	"strings"
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
		publicKey,
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

	return crypto.FromECDSAPub(&privateKey.PublicKey)[1:], nil
}
