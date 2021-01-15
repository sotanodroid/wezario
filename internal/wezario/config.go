package wezario

import (
	"net/url"
)

// Config ...
type Config struct {
	OpenweathermapURL    *url.URL `env:"WEZARIO_OPENWEATHER_URL"`
	OpenweathermapAPIKey string   `env:"WEZARIO_OPENWEATHER_KEY"`
	Loglevel             string   `env:"WEZARIO_LOG_LEVEL"`
	RedisURL             *url.URL `env:"WEZARIO_REDIS_URL"`
	RedisPassword        string   `env:"WEZARIO_REDIS_PASSWORD"`
	RedisDB              int      `env:"WEZARIO_REDIS_DB"`
	TelegramToken        string   `env:"WEZARIO_TELEGRAM_TOKEN"`
}

// NewConfig returns new instance of service config
func NewConfig() *Config {
	return &Config{}
}
