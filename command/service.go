package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/johnmanjiro13/dokkoi/google"
)

type service struct {
	customSearchRepo google.CustomSearchRepository
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

func NewService(customSearchRepo google.CustomSearchRepository) Service {
	return &service{
		customSearchRepo: customSearchRepo,
	}
}

type DokkoiCmd interface {
	SendMessage(s *discordgo.Session, channelID string) error
}

func (s *service) GetCommand(content string) DokkoiCmd {
	cmd := strings.Split(content, " ")
	if len(cmd) <= 1 || cmd[0] != "dokkoi" {
		return nil
	}
	switch cmd[1] {
	case "echo":
		return &echoCmd{message: strings.Join(cmd[2:], " ")}
	case "image":
		return &imageCmd{
			customSearchRepo: s.customSearchRepo,
			query:            strings.Join(cmd[2:], " "),
		}
	default:
		return nil
	}
}
