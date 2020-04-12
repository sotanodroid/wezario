package wezario

import (
	"github.com/go-redis/redis/v7"
	"github.com/urfave/cli/v2"
)

var HTTPClient *client
var redisClient *redis.Client

// Start starts new app which parses command line arguments
// to provide weather information
func Start(cfg *Config) *cli.App {
	var city string
	var units string

	HTTPClient = NewHTTPClient(cfg)
	redisClient = NewRedisClient(cfg)

	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "city",
				Value:       "Moscow",
				Aliases:     []string{"c"},
				Usage:       "city to show weather information for",
				Destination: &city,
			},
			&cli.StringFlag{
				Name:        "units",
				Value:       "metric",
				Aliases:     []string{"u"},
				Usage:       "Unit metric system to show. Choses 'imperial' or 'metric'.",
				Destination: &units,
			},
		},
		Action: func(c *cli.Context) error {
			return getOrSetWeatherData(c, city, units)
		},
	}
}
