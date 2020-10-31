package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ClientConf struct {
	StorageGatewayAddr string `yaml:"storageGatewayAddr"`
	StorageGatewayPort int    `yaml:"storageGatewayPort"`
}

func NewClientConf(confFile string) (*ClientConf, error) {
	b, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	c := &ClientConf{}
	if err = yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}
	return c, nil
}
