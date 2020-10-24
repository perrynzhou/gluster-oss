package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MetaServiceConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

type ServiceConf struct {
	//service address
	Addr string `yaml:"serverAddr"`
	//service grpc port
	GrpcPort int `yaml:"grpcPort"`
	//service http port
	HttpPort     int             `yaml:"httpPort"`
	StoreSrvAddr MetaServiceConf `yaml:"storeServer"`
	//metadata server address
	MetaSrvAddr MetaServiceConf `yaml:"metaServer"`
}

func NewServiceConf(confFile string) (*ServiceConf, error) {
	b, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	c := &ServiceConf{}
	if err = yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}
	return c, nil
}
