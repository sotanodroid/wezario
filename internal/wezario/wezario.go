package wezario

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/go-redis/redis/v7"
)

var httpClient *client
var redisClient *redis.Client

// Start starts new app which parses command line arguments
// to provide weather information
func Start(cfg *Config) error {
	httpClient = newHTTPClient(cfg)
	redisClient = newRedisClient(cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return err
	}

	bot.Debug = true

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updatesChan, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		return err
	}

	for update := range updatesChan {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/start" {
			result := fmt.Sprintf(
				"%s\n\n%s\n%s\n",
				"Привет!",
				"Я могу узнать для тебя погоду в любом городе.",
				"Напиши мне название города на английском языке и я выдам результат.",
			)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
			bot.Send(msg)
			continue
		}

		city := update.Message.Text
		weatherInfo, err := getOrSetWeatherData(city)
		if err != nil || weatherInfo == " " {
			result := fmt.Sprintf(
				"Не удается получить информацию по городу %s",
				city,
			)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
			bot.Send(msg)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherInfo)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	return nil
}
