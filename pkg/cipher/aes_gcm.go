package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

type SymmetricKey [32]byte

type AesGcm struct {
	cipher.AEAD
	key SymmetricKey
}

func NewAesGcm() (*AesGcm, error) {
	var key SymmetricKey

	_, err := rand.Read(key[:])
	if err != nil {
		return nil, fmt.Errorf("could not generate symmetric key: [%v]", err)
	}

	aesInstance, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("could not create AES instance: [%v]", err)
	}

	aesGcmInstance, err := cipher.NewGCM(aesInstance)
	if err != nil {
		return nil, fmt.Errorf("could not create AES-GCM instance: [%v]", err)
	}

	return &AesGcm{aesGcmInstance, key}, nil
}

func (ag *AesGcm) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, ag.NonceSize())
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("could not generate nonce: [%v]", err)
	}

	return ag.Seal(nil, nonce, data, nil), nil
}
