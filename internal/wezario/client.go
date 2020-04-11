package wezario

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type client struct {
	logger *logrus.Logger
	cfg    *Config
}

func NewClient(cfg *Config) *client {
	return &client{
		cfg:    cfg,
		logger: logrus.New(),
	}
}

func (c *client) requestWeather(cxt *cli.Context, city string) error {

	q := c.cfg.OpenweathermapURL.Query()
	q.Set("q", city)
	q.Set("appid", c.cfg.OpenweathermapAPIKey)
	q.Set("units", "metric") // TODO add units flag
	c.cfg.OpenweathermapURL.RawQuery = q.Encode()

	resp, err := http.Get(c.cfg.OpenweathermapURL.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	fmt.Printf(
		"Temp\t\t%v\nFeels like\t%v\nThere is mostly %v (%v)\n",
		res.Main.Temp,
		res.Main.FeelsLike,
		res.Weather[0].Main,
		res.Weather[0].Description,
	)

	return nil
}
