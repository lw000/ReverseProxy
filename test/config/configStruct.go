package config

type Config struct {
	Count       int64  `json:"count"`
	Method      string `json:"method"`
	Millisecond int64  `json:"millisecond"`
	PostURL     string `json:"postUrl"`
	URL         string `json:"url"`
}
