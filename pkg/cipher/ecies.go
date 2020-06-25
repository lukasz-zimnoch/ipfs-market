package cipher

import (
	eciesgo "github.com/ecies/go"
)

type Ecies struct {
	privateKey *eciesgo.PrivateKey
}

func NewEcies(privateKeyBytes []byte) *Ecies {
	privateKey := eciesgo.NewPrivateKeyFromBytes(privateKeyBytes)

	return &Ecies{privateKey}
}

func (r *Ecies) Encrypt(data []byte) ([]byte, error) {
	return eciesgo.Encrypt(r.privateKey.PublicKey, data)
}

func (r *Ecies) Decrypt(data []byte) ([]byte, error) {
	return eciesgo.Decrypt(r.privateKey, data)
}
