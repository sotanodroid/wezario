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

type result struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
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
	q.Set("units", "metric") // TODO Добавить флаг выбора метрической системы
	c.cfg.OpenweathermapURL.RawQuery = q.Encode()

	resp, err := http.Get(c.cfg.OpenweathermapURL.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var res result
	json.NewDecoder(resp.Body).Decode(&res)

	weather, err := json.Marshal(res)
	if err != nil {
		return err
	}

	// FIXME Сделать человекочитаемый вывод в терминал, после публикации как бота -- убрать
	fmt.Println(string(weather))

	return nil
}
