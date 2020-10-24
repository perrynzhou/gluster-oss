package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ClientConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
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
