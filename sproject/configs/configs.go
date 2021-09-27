package configs

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   int
	Name   string
	Server string
}

func LoadConfig(c string) (*Config, error) {
	f, err := ioutil.ReadFile(c)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	err = yaml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}
	return config, err
}
