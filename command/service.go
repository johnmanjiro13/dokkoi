package command

import (
	"regexp"

	"github.com/bwmarrin/discordgo"

	"github.com/johnmanjiro13/dokkoi/google"
)

var commandRegExp = regexp.MustCompile(`^dokkoi\s(.+)\s(.+)$`)

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
	cmd := commandRegExp.FindStringSubmatch(content)
	if len(cmd) != 3 {
		return nil
	}
	switch cmd[1] {
	case "echo":
		return &echoCmd{message: cmd[2]}
	case "image":
		return &imageCmd{
			customSearchRepo: s.customSearchRepo,
			query:            cmd[2],
		}
	default:
		return nil
	}
}
