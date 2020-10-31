package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type CommonBackendConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}
type StoreBackendConf struct {
	CommonBackendConf
	Volume  string `yaml:"volume"`
}
type ServiceBackendConfig struct {
	BackendAddr string `yaml:"addr"`
	GrpcPort    int    `yaml:"grpc_port"`
	//service http port
	HttpPort int `yaml:"http_port"`
}
type ServerConfig struct {
	//service address
	ServiceBackend ServiceBackendConfig `yaml:"serviceBackend"`
	StoreBackend  StoreBackendConf      `yaml:"storageBackend"`
	//metadata server address
	MetaBacked CommonBackendConf `yaml:"metaBackend"`
}

func NewServerConf(confFile string) (*ServerConfig, error) {
	b, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	c := &ServerConfig{}
	if err = yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}
	return c, nil
}
