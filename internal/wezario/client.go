package wezario

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type client struct {
	logger *logrus.Logger
	cfg    *Config
}

type result struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      int     `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   int     `json:"temp_min"`
		TempMax   int     `json:"temp_max"`
	} `json:"main"`
}

func NewClient(cfg *Config) *client {
	return &client{
		cfg:    cfg,
		logger: logrus.New(),
	}
}

func (c *client) requestWeather(city string) {

	q := c.cfg.OpenweathermapURL.Query()
	q.Set("q", city)
	q.Set("appid", c.cfg.OpenweathermapAPIKey)
	q.Set("units", "metric") // Добавить флаг выбора метрической системы

	c.cfg.OpenweathermapURL.RawQuery = q.Encode()

	resp, err := http.Get(c.cfg.OpenweathermapURL.String())
	if err != nil {
		c.logger.Fatal(err.Error())
	}

	defer resp.Body.Close()

	var res result
	json.NewDecoder(resp.Body).Decode(&res)

	weather, _ := json.Marshal(res)
 
	fmt.Println(string(weather))
}
