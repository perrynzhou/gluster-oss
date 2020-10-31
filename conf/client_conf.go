package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)


type StorageGatewayClient struct {
	CommonConf  CommonBackendConf `yaml:"storageGateway"`
}
func NewClientConf(confFile string) (*StorageGatewayClient, error) {
	b, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	c := &StorageGatewayClient{}
	if err = yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}
	return c, nil
}
