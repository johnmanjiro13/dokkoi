package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.db", "REDIS_DB")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")

	viper.SetDefault("redis.host", "redis:6379")
	viper.SetDefault("redis.db", "1")
	viper.SetDefault("redis.password", "")
}

func Open(ctx context.Context, host string, db int, password string) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     host,
		DB:       db,
		Password: password,
	})
	ping := cli.Ping(ctx)
	if _, err := ping.Result(); err != nil {
		return nil, pkgerrors.Wrap(err, "redis ping failed")
	}
	return cli, nil
}
