package wezario

import (
	"github.com/sirupsen/logrus"
)

// Start starts new app which parses command line arguments
// to provide weather information
func Start(cfg *Config, logger *logrus.Logger) error {
	openWeatherClient := newOpenWeatherClient(cfg)
	teleBot, err := newTeleBot(cfg, logger, openWeatherClient)
	if err != nil {
		return err
	}

	if err := teleBot.processMessage(); err != nil {
		return err
	}

	return nil
}
