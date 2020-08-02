package main

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sotanodroid/wezario/internal/wezario"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}
}

func main() {
	cfg := wezario.NewConfig()
	if err := envdecode.Decode(cfg); err != nil {
		logger.Error(err)
	}
	logger.Info("Application starts")

	if err := wezario.Start(cfg); err != nil {
		logger.Fatal(err)
	}
}
