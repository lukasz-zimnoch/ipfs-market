package cipher

import (
	"fmt"
	"strings"

	eciesgo "github.com/ecies/go"
)

type Ecies struct {
	privateKey *eciesgo.PrivateKey
}

func NewEcies(privateKeyHex string) (*Ecies, error) {
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = privateKeyHex[2:]
	}

	privateKey, err := eciesgo.NewPrivateKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: [%v]", err)
	}

	return &Ecies{privateKey}, nil
}

func (r *Ecies) Encrypt(data []byte) ([]byte, error) {
	return eciesgo.Encrypt(r.privateKey.PublicKey, data)
}

func (r *Ecies) Dencrypt(data []byte) ([]byte, error) {
	return eciesgo.Decrypt(r.privateKey, data)
}
