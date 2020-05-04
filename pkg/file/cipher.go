package file

type Cipher interface {
	Encrypt(data []byte) ([]byte, error)
}
