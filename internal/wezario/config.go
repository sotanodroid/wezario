package wezario

import "net/url"

type Config struct {
	OpenweathermapURL    *url.URL `env:"SERVICE_URL"`
	OpenweathermapAPIKey string   `env:"SERVICE_KEY"`
	Loglevel             string   `env:"SERVICE_LOG_LEVEL"`
}

// NewConfig return new instance of service config
func NewConfig() *Config {
	return &Config{}
}
