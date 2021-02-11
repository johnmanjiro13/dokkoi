package redis

import (
	"log"
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func init() {
	viper.Set("redis.host", "localhost:6379")
	viper.Set("redis.db", 10)
}

func TestMain(m *testing.M) {
	cli, err := openTest()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cli.Close(); err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}()
	defer cli.FlushDB()

	code := m.Run()
	os.Exit(code)
}

func openTest() (*redis.Client, error) {
	return Open(viper.GetString("redis.host"), viper.GetInt("redis.db"))
}
