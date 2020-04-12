package wezario

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/urfave/cli/v2"
)

// NewRedisClient creates a new redis connection client
func NewRedisClient(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL.String(),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
}

func getOrSetWeatherData(cxt *cli.Context, city string) error {
	weatherdata, err := redisClient.Get(city).Result()
	if err == redis.Nil {
		result, err := HTTPClient.requestWeather(city)
		if err != nil {
			return err
		}

		// TTL hardcoded to 15 minutes, consider to make it as ENVvar
		duration := (60 * time.Second) * 15
		ttl := time.Duration(duration)
		redisClient.Set(city, result, ttl)

		weatherdata = result

	} else if err != nil {
		return err
	}

	fmt.Println(weatherdata)

	return nil
}
