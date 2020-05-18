package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukasz-zimnoch/ipfs-market/pkg/chain/gen/contracts"
	"math/big"
	"time"
)

type EthereumClient struct {
	client             *ethclient.Client
	ipfsMarketContract *contracts.IpfsMarket
}

func NewEthereumClient(url string) (*EthereumClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	// TODO 2: Take address from config.
	ipfsMarketContract, err := contracts.NewIpfsMarket(
		common.HexToAddress("0x97B08a01aA799861e2a736a1e2E40E6E548Da647"),
		client,
	)
	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		client:             client,
		ipfsMarketContract: ipfsMarketContract,
	}, nil
}

func (ec *EthereumClient) Publish(hash string, accessKey []byte) {
	// TODO 3: Take price from config.
	price := big.NewInt(100000000000000000)

	// TODO 4: Do we need to use this context?
	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelCtx()

	opts := &bind.TransactOpts{Context: ctx}

	// TODO 5: Check transaction outcome.
	ec.ipfsMarketContract.Publish(opts, hash, accessKey, price)
}
