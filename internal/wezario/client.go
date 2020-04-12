package wezario

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type client struct {
	cfg *Config
}

func NewHTTPClient(cfg *Config) *client {
	return &client{
		cfg: cfg,
	}
}

func (c *client) requestWeather(city, units string) (string, error) {

	q := c.cfg.OpenweathermapURL.Query()
	q.Set("q", city)
	q.Set("appid", c.cfg.OpenweathermapAPIKey)
	q.Set("units", units)
	c.cfg.OpenweathermapURL.RawQuery = q.Encode()

	resp, err := http.Get(c.cfg.OpenweathermapURL.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Temp\t\t%v\nFeels like\t%v\nThere is mostly %v (%v)\n",
		res.Main.Temp,
		res.Main.FeelsLike,
		res.Weather[0].Main,
		res.Weather[0].Description,
	), nil
}
