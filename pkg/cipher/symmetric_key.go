package cipher

import (
	"crypto/rand"
)

type SymmetricKey [32]byte

func GenerateSymmetricKey() (SymmetricKey, error) {
	var key SymmetricKey

	_, err := rand.Read(key[:])
	if err != nil {
		return key, err
	}

	return key, nil
}

func SymmetricKeyFromSlice(keySlice []byte) SymmetricKey {
	var key SymmetricKey

	copy(key[:], keySlice)

	return key
}
