package chain

import "github.com/ethereum/go-ethereum/ethclient"

type EthereumClient struct {
	client *ethclient.Client
}

func NewEthereumClient(url string) (*EthereumClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		client: client,
	}, nil
}

func (ec *EthereumClient) Publish(hash string, accessKey []byte) {
	// TODO 2: Implementation.
}
