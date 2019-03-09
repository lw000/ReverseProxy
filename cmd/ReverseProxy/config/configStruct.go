package config

type Config struct {
	BackgroudURL []string `json:"backgroudUrl"`
	Debug        int64    `json:"debug"`
	Port         int64    `json:"port"`
}
