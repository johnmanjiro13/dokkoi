package command

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var commandRegExp = regexp.MustCompile(`^dokkoi\s(.+)\s(.+)$`)

type service struct {
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

func NewService() Service {
	return &service{}
}

type DokkoiCmd interface {
	SendMessage(s *discordgo.Session, channelID string) error
}

func (s *service) GetCommand(content string) DokkoiCmd {
	cmd := commandRegExp.FindStringSubmatch(content)
	if len(cmd) != 3 {
		return nil
	}
	switch cmd[1] {
	case "echo":
		return &echoCmd{message: cmd[2]}
	default:
		return nil
	}
}
