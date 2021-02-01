package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/johnmanjiro13/dokkoi/command"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var token string
	flag.StringVar(&token, "token", "", "bot token")
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

	cmdService := command.NewService()
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
