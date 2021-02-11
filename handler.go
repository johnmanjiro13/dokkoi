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
	if m.Author.Bot {
		return
	}

	cmd := h.commandService.GetCommand(m.Content)
	if cmd == nil {
		return
	}
	res, err := cmd.Exec()
	if err != nil {
		log.Printf("an error occurred in command execution. err: %v", err)
		return
	}
	if _, err := s.ChannelMessageSend(m.ChannelID, res); err != nil {
		log.Printf("an error occurred in sending message. err: %v", err)
	}
}
