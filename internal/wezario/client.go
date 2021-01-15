package wezario

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const notFound = 404

type httpClient struct {
	cfg *Config
}

func newHTTPClient(cfg *Config) *httpClient {
	return &httpClient{
		cfg: cfg,
	}
}

func (c *httpClient) requestWeather(city string) (string, error) {

	q := c.cfg.OpenweathermapURL.Query()
	q.Set("q", city)
	q.Set("appid", c.cfg.OpenweathermapAPIKey)
	q.Set("units", "metric")
	c.cfg.OpenweathermapURL.RawQuery = q.Encode()

	resp, err := http.Get(c.cfg.OpenweathermapURL.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == notFound {
		return "", errors.New("City not found")
	}

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
