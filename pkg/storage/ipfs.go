package storage

import (
	"bytes"
	"io/ioutil"
	"os"

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

func (i *Ipfs) GetData(cid string) ([]byte, error) {
	temp := "./temporary"

	if _, err := os.Stat(temp); os.IsNotExist(err) {
		_ = os.Mkdir(temp, os.ModePerm)
	}

	if err := i.Get(cid, temp); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(temp + "/" + cid)
	if err != nil {
		return nil, err
	}

	_ = os.Remove(temp)

	return data, nil
}
