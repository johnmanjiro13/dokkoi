package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"

	"github.com/johnmanjiro13/dokkoi/command"
	"github.com/johnmanjiro13/dokkoi/infra/google"
	"github.com/johnmanjiro13/dokkoi/infra/inmem"
)

func init() {
	viper.BindEnv("discord.token", "DISCORD_TOKEN")

	viper.SetDefault("discord.token", "")

	viper.AutomaticEnv()
}

func main() {
	dg, err := discordgo.New("Bot " + viper.GetString("discord.token"))
	if err != nil {
		log.Fatalf("creating discord session is fail. err: %v", err)
	}

	csService, err := customsearch.NewService(context.Background(), option.WithAPIKey(viper.GetString("customsearch.api.key")))
	if err != nil {
		log.Fatalf("creating customsearch service is fail. err: %v", err)
	}
	csRepo := google.NewCustomSearchRepository(csService, viper.GetString("customsearch.engine.id"))
	scoreRepo := inmem.NewScoreRepository(map[string]int64{})
	cmdService := command.NewService(csRepo, scoreRepo)
	handler := newHandler(cmdService)

	dg.AddHandler(handler.onMessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("opening discord connection is fail. err: %v", err)
	}
	defer dg.Close()

	log.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}
