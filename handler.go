package main

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/johnmanjiro13/dokkoi/command"
)

type handler struct {
	commandService command.Service
}

func newHandler(commandService command.Service) *handler {
	return &handler{
		commandService: commandService,
	}
}

func (h *handler) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Username == "dokkoi" {
		return
	}

	cmd := h.commandService.GetCommand(m.Content)
	if cmd == nil {
		return
	}
	err := cmd.SendMessage(s, m.ChannelID)
	if err != nil {
		log.Printf("an error occurred in sending message. err: %s", err)
	}
}
