package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/johnmanjiro13/dokkoi/google"
)

type service struct {
	customSearchRepo google.CustomSearchRepository
	scoreRepo        ScoreRepository
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

func NewService(customSearchRepo google.CustomSearchRepository, scoreRepo ScoreRepository) Service {
	return &service{
		customSearchRepo: customSearchRepo,
		scoreRepo:        scoreRepo,
	}
}

type DokkoiCmd interface {
	SendMessage(s *discordgo.Session, channelID string) error
}

func (s *service) GetCommand(content string) DokkoiCmd {
	if strings.HasSuffix(content, IncrOperator) {
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      content[:len(content)-2],
			operator:  IncrOperator,
		}
	} else if strings.HasSuffix(content, DecrOperator) {
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      content[:len(content)-2],
			operator:  DecrOperator,
		}
	}

	// replace full-width whitespace to half size whitespace
	replacedContent := strings.Replace(content, "　", " ", -1)
	cmd := strings.Split(replacedContent, " ")
	if len(cmd) <= 1 || cmd[0] != "dokkoi" {
		return nil
	}
	switch cmd[1] {
	case "help":
		// return helpCmd only in the case of  'dokkoi help' was sent as of now.
		if len(cmd) == 2 {
			return &helpCmd{}
		}
		return nil
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
