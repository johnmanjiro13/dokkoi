package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	goredis "github.com/go-redis/redis/v8"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"

	"github.com/johnmanjiro13/dokkoi/pkg/command"
	"github.com/johnmanjiro13/dokkoi/pkg/infra/google"
	"github.com/johnmanjiro13/dokkoi/pkg/infra/inmem"
	"github.com/johnmanjiro13/dokkoi/pkg/infra/redis"
)

func init() {
	viper.BindEnv("discord.token", "DISCORD_TOKEN")

	viper.SetDefault("discord.token", "")

	viper.AutomaticEnv()
}

var (
	scoreDBType = flag.StringP("scoredb-type", "s", "redis", "database type for score command")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// setup discord session
	dg, err := discordgo.New("Bot " + viper.GetString("discord.token"))
	if err != nil {
		log.Fatalf("creating discord session is fail. err: %v", err)
	}

	// setup custom search service and repository
	csService, err := customsearch.NewService(context.Background(), option.WithAPIKey(viper.GetString("customsearch.api.key")))
	if err != nil {
		log.Fatalf("creating customsearch service is fail. err: %v", err)
	}
	csRepo := google.NewCustomSearchRepository(csService, viper.GetString("customsearch.engine.id"))

	// setup score repository
	var (
		scoreRepo command.ScoreRepository
		redisCli  *goredis.Client
	)
	switch *scoreDBType {
	case "inmem":
		scoreRepo = inmem.NewScoreRepository(map[string]int64{})
	case "redis":
		redisCli, err = redis.Open(context.Background(), viper.GetString("redis.host"), viper.GetInt("redis.db"), viper.GetString("redis.password"))
		if err != nil {
			log.Fatalf("connecting redis is fail. err: %v", err)
		}
		scoreRepo = redis.NewScoreRepository(redisCli)
	default:
		log.Fatalf("invalid type for score database: %s", *scoreDBType)
	}

	// setup handler
	cmdService := command.NewService(csRepo, scoreRepo)
	handler := newHandler(cmdService)

	dg.AddHandler(handler.onMessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("opening discord connection is fail. err: %v", err)
	}

	log.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// close connections clearly
	dg.Close()
	if redisCli != nil {
		redisCli.Close()
	}
}
