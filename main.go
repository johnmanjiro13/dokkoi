package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	echoCmd = "echo"
)

var (
	commandRegExp = regexp.MustCompile(`^dokkoi\s(.+)\s(.+)$`)
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
		log.Printf("creating discord session is fail. err: %v ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(onMessageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Printf("opening discord connection is fail. err: %v ", err)
		return
	}
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	log.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := commandRegExp.FindStringSubmatch(m.Content)
	if len(cmd) != 3 {
		return
	}
	switch {
	case cmd[1] == echoCmd:
		sendMessage(s, m.ChannelID, cmd[2])
	}
}

func sendMessage(s *discordgo.Session, channelID, message string) {
	if _, err := s.ChannelMessageSend(channelID, message); err != nil {
		log.Printf("error sending message. err: %v", err)
		return
	}
}
