package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/johnmanjiro13/dokkoi/pkg/command"
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
	if cmd.SendType() == "Message" {
		if err := sendMessage(s, m.ChannelID, cmd); err != nil {
			log.Println(err)
		}
	} else if cmd.SendType() == "File" {
		if err := sendFile(s, m.ChannelID, cmd); err != nil {
			log.Println(err)
		}
	}
}

func sendMessage(s *discordgo.Session, channelID string, cmd command.DokkoiCmd) error {
	res, err := cmd.ExecString(context.Background())
	if err != nil {
		return fmt.Errorf("an error occurred in message command execution. err: %+v", err)
	}
	if _, err := s.ChannelMessageSend(channelID, res); err != nil {
		return fmt.Errorf("an error occurred in sending message. err: %+v", err)
	}
	return nil
}

func sendFile(s *discordgo.Session, channelID string, cmd command.DokkoiCmd) error {
	res, err := cmd.ExecFile(context.Background())
	if err != nil {
		return fmt.Errorf("an error occurred in file command execution. err: %+v", err)
	}
	if res == nil {
		return fmt.Errorf("result is not found.")
	}
	if _, err := s.ChannelFileSend(channelID, "lgtm.jpg", res); err != nil {
		return fmt.Errorf("an error occurred in sending file. err: %+v", err)
	}
	return nil
}
