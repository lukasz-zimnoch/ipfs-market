package cipher

import (
	"fmt"
	eciesgo "github.com/ecies/go"
)

type Ecies struct {
	privateKey *eciesgo.PrivateKey
	publicKey  *eciesgo.PublicKey
}

func NewEcies(privateKeyBytes []byte) *Ecies {
	privateKey := eciesgo.NewPrivateKeyFromBytes(privateKeyBytes)

	return &Ecies{privateKey, privateKey.PublicKey}
}

func NewEciesFromPublicKey(publicKeyBytes []byte) (*Ecies, error) {
	publicKey, err := eciesgo.NewPublicKeyFromBytes(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not create ECIES instance: [%v]", err)
	}

	return &Ecies{nil, publicKey}, nil
}

func (r *Ecies) Encrypt(data []byte) ([]byte, error) {
	return eciesgo.Encrypt(r.publicKey, data)
}

func (r *Ecies) Decrypt(data []byte) ([]byte, error) {
	if r.privateKey == nil {
		return nil, fmt.Errorf("encrypt-only ECIES instance")
	}

	return eciesgo.Decrypt(r.privateKey, data)
}
