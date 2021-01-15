package main

import (
	"github.com/joeshaw/envdecode"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/sotanodroid/wezario/internal/wezario"
)

func main() {
	logger := logrus.New()
	cfg := wezario.NewConfig()
	if err := envdecode.Decode(cfg); err != nil {
		logger.Error(err)
	}

	level, err := logrus.ParseLevel(cfg.Loglevel)
	if err != nil {
		logger.Error(err)
	}

	logger.SetLevel(level)
	logger.Info("Application starts")

	if err := wezario.Start(cfg, logger); err != nil {
		logger.Fatal(err)
	}
}
