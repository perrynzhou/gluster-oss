package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MetaServiceConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

type ServerConfig struct {
	//service address
	Addr string `yaml:"serverAddr"`
	//service grpc port
	GrpcPort int `yaml:"grpcPort"`
	//service http port
	HttpPort     int             `yaml:"httpPort"`
	StoreBackend MetaServiceConf `yaml:"storageBackend"`
	//metadata server address
	MetaBacked MetaServiceConf `yaml:"metaBackend"`
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
