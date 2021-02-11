package redis

import (
	"github.com/go-redis/redis"
	pkgerrors "github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.db", "REDIS_DB")

	viper.SetDefault("redis.host", "redis:6379")
	viper.SetDefault("redis.db", "1")
}

func Open(host string, db int) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    host,
		DB:      db,
	})
	ping := cli.Ping()
	if _, err := ping.Result(); err != nil {
		return nil, pkgerrors.Wrap(err, "redis ping failed")
	}
	return cli, nil
}
