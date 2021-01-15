package wezario

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	log *logrus.Logger
	api *tgbotapi.BotAPI
	ow  *OpenWeatherClient
}

func newTeleBot(cfg *Config, logger *logrus.Logger, ow *OpenWeatherClient) (*Bot, error) {
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		log: logger,
		api: botApi,
		ow:  ow,
	}, nil
}

func (b *Bot) processMessage() error {
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updatesChan, err := b.api.GetUpdatesChan(ucfg)
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
			_, err := b.api.Send(msg)
			if err != nil {
				b.log.Errorf("Error sending msg: %v", err)
			}
			continue
		}

		city := update.Message.Text
		weatherInfo, err := b.ow.getOrSetWeatherData(city)
		if err != nil || weatherInfo == " " {
			result := fmt.Sprintf(
				"Не удается получить информацию по городу %s",
				city,
			)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
			_, err := b.api.Send(msg)
			if err != nil {
				b.log.Errorf("Error sending msg: %v", err)
			}
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherInfo)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err = b.api.Send(msg)
		if err != nil {
			b.log.Errorf("Error sending msg: %v", err)
		}
	}

	return nil
}
