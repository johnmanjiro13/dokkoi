package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"

	"github.com/johnmanjiro13/dokkoi/command"
	"github.com/johnmanjiro13/dokkoi/google"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var token, apiKey, engineID string
	flag.StringVar(&token, "token", "", "bot token")
	flag.StringVar(&apiKey, "api_key", "", "google api key")
	flag.StringVar(&engineID, "engine_id", "", "google search engine id")
	flag.VisitAll(func(f *flag.Flag) {
		if v, ok := os.LookupEnv(strings.ToUpper(f.Name)); ok {
			f.Value.Set(v)
		}
	})
	flag.Parse()

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("creating discord session is fail. err: %s", err)
	}

	csService, err := customsearch.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("creating customsearch service is fail. err: %s", err)
	}
	csRepo := google.NewCustomSearchRepository(csService, engineID)
	scoreRepo := command.NewScoreRepository(map[string]int{})
	cmdService := command.NewService(csRepo, scoreRepo)
	handler := newHandler(cmdService)

	dg.AddHandler(handler.onMessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("opening discord connection is fail. err: %s", err)
	}
	defer dg.Close()

	log.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}
