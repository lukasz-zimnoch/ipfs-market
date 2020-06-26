package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Ethereum Ethereum `yaml:"ethereum"`
	Storage  Storage  `yaml:"storage"`
}

type Ethereum struct {
	URL                string `yaml:"url"`
	PrivateKey         string `yaml:"privateKey"`
	IpfsMarketContract string `yaml:"ipfsMarketContract"`
}

type Storage struct {
	URL     string `yaml:"url"`
	Workdir string `yaml:"workdir"`
}

func ReadConfig(configPath string) (*Config, error) {
	var config Config

	configYaml, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
