package wezario

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type OpenWeatherClient struct {
	redis      *redis.Client
	httpClient *httpClient
}

func newOpenWeatherClient(cfg *Config) *OpenWeatherClient {
	return &OpenWeatherClient{
		redis: redis.NewClient(&redis.Options{
			Addr:     cfg.RedisURL.String(),
			Password: cfg.RedisPassword,
			DB:       cfg.RedisDB,
		}),
		httpClient: newHTTPClient(cfg),
	}
}

func (r *OpenWeatherClient) getOrSetWeatherData(city string) (string, error) {
	weatherdata, err := r.redis.Get(city).Result()
	if err == redis.Nil {
		result, err := r.httpClient.requestWeather(city)
		if err != nil {
			return " ", err
		}

		// TTL hardcoded to 15 minutes, consider to make it as ENVvar
		duration := (60 * time.Second) * 15
		ttl := time.Duration(duration)
		r.redis.Set(city, result, ttl)

		weatherdata = result

	} else if err != nil {
		return " ", err
	}

	return weatherdata, nil
}
