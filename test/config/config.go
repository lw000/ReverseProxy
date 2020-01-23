package config

import (
	"encoding/json"
	"io/ioutil"
)

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}
