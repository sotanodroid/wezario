package wezario

import (
	"time"

	"github.com/go-redis/redis/v7"
)

func newRedisClient(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL.String(),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
}

func getOrSetWeatherData(city string) (string, error) {
	weatherdata, err := redisClient.Get(city).Result()
	if err == redis.Nil {
		result, err := httpClient.requestWeather(city)
		if err != nil {
			return " ", err
		}

		// TTL hardcoded to 15 minutes, consider to make it as ENVvar
		duration := (60 * time.Second) * 15
		ttl := time.Duration(duration)
		redisClient.Set(city, result, ttl)

		weatherdata = result

	} else if err != nil {
		return " ", err
	}

	return weatherdata, nil
}
