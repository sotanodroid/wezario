package wezario

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/go-redis/redis/v7"
)

var httpClient *client
var redisClient *redis.Client
var teleBot *tgbotapi.BotAPI

// Start starts new app which parses command line arguments
// to provide weather information
func Start(cfg *Config) error {
	var err error

	httpClient = newHTTPClient(cfg)
	redisClient = newRedisClient(cfg)
	teleBot, err = newTeleBot(cfg)
	if err != nil {
		return err
	}

	if err := processMessage(); err != nil {
		return err
	}

	return nil
}
