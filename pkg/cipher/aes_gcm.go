package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

type AesGcm struct {
	cipher.AEAD
}

func NewAesGcm(key SymmetricKey) (*AesGcm, error) {
	aesInstance, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("could not create AES instance: [%v]", err)
	}

	aesGcmInstance, err := cipher.NewGCM(aesInstance)
	if err != nil {
		return nil, fmt.Errorf("could not create AES-GCM instance: [%v]", err)
	}

	return &AesGcm{aesGcmInstance}, nil
}

func (ag *AesGcm) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, ag.NonceSize())
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("could not generate nonce: [%v]", err)
	}

	return ag.Seal(nil, nonce, data, nil), nil
}

func (ag *AesGcm) Decrypt(data []byte) ([]byte, error) {
	nonce, data := data[:ag.NonceSize()], data[ag.NonceSize():]

	return ag.Open(nil, nonce, data, nil)
}
