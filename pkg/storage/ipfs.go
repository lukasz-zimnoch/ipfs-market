package storage

import (
	"bytes"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type Ipfs struct {
	*ipfs.Shell
}

func NewIpfs(url string) *Ipfs {
	return &Ipfs{ipfs.NewShell(url)}
}

func (i *Ipfs) Store(data []byte) (string, error) {
	return i.Add(bytes.NewReader(data))
}
