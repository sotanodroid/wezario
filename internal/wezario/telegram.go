package wezario

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func newTeleBot(cfg *Config) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(cfg.TelegramToken)
}

func processMessage() error {
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updatesChan, err := teleBot.GetUpdatesChan(ucfg)
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
			teleBot.Send(msg)
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
			teleBot.Send(msg)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherInfo)
		msg.ReplyToMessageID = update.Message.MessageID

		teleBot.Send(msg)
	}

	return nil
}
