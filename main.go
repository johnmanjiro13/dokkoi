package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token = flag.String("t", "", "bot token")
)

func main() {
	flag.Parse()

	dg, err := discordgo.New("Bot " + *token)
	if err != nil {
		fmt.Println("creating discord session is fail, ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(onMessageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("opening discord connection is fail, ", err)
		return
	}
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "ping" {
		sendMessage(s, m.ChannelID, "pong")
	}
}

func sendMessage(s *discordgo.Session, channelID, message string) {
	if _, err := s.ChannelMessageSend(channelID, message); err != nil {
		fmt.Println("error sending message, ", err)
		return
	}
	fmt.Println(">>> ", message)
}
