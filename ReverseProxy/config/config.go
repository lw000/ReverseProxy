package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Background []string `json:"background"`
	Debug      int64    `json:"debug"`
	Port       int64    `json:"port"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}
