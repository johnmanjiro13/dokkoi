package command

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/customsearch/v1"
)

type service struct {
	customSearchRepo CustomSearchRepository
	scoreRepo        ScoreRepository
}

var ErrImageNotFound = errors.New("image was not found.")

type CustomSearchRepository interface {
	SearchImage(query string) (*customsearch.Result, error)
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

func NewService(customSearchRepo CustomSearchRepository, scoreRepo ScoreRepository) Service {
	return &service{
		customSearchRepo: customSearchRepo,
		scoreRepo:        scoreRepo,
	}
}

type DokkoiCmd interface {
	SendMessage(s *discordgo.Session, channelID string) error
}

func (s *service) GetCommand(content string) DokkoiCmd {
	// replace full-width whitespace to half size whitespace
	replacedContent := strings.Replace(content, "　", " ", -1)
	cmd := strings.Split(replacedContent, " ")
	switch {
	case len(cmd) >= 2 && cmd[1] == "help":
		return &helpCmd{target: strings.Join(cmd[2:], " ")}
	case len(cmd) >= 3 && cmd[1] == "echo":
		return &echoCmd{message: strings.Join(cmd[2:], " ")}
	case len(cmd) >= 3 && cmd[1] == "image":
		return &imageCmd{
			customSearchRepo: s.customSearchRepo,
			query:            strings.Join(cmd[2:], " "),
		}
	case len(cmd) >= 3 && cmd[1] == "score":
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      strings.Join(cmd[2:], " "),
			operator:  noOperator,
		}
	case strings.HasSuffix(content, incrOperator):
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      content[:len(content)-2],
			operator:  incrOperator,
		}
	case strings.HasSuffix(content, decrOperator):
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      content[:len(content)-2],
			operator:  decrOperator,
		}
	default:
		return nil
	}
}
