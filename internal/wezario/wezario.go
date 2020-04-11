package wezario

import (
	"github.com/urfave/cli/v2"
)

var HTTPClient *client

// Start starts new app which parses command line arguments
// to provide weather information
func Start(cfg *Config) *cli.App {
	var city string
	HTTPClient = NewClient(cfg)

	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "city",
				Value:       "Moscow",
				Usage:       "city to show weather information for",
				Destination: &city,
			},
		},
		Action: func(c *cli.Context) error {
			return HTTPClient.requestWeather(c, city)
		},
	}
}
