package command

import (
	"strings"
)

type service struct {
	customSearchRepo CustomSearchRepository
	scoreRepo        ScoreRepository
}

func NewService(customSearchRepo CustomSearchRepository, scoreRepo ScoreRepository) Service {
	return &service{
		customSearchRepo: customSearchRepo,
		scoreRepo:        scoreRepo,
	}
}

func (s *service) GetCommand(content string) DokkoiCmd {
	// replace full-width whitespace to half size whitespace
	replacedContent := strings.Replace(content, "　", " ", -1)
	cmd := strings.Split(replacedContent, " ")
	if cmd[0] == "dokkoi" {
		return s.withPrefixCommand(cmd)
	} else {
		return s.withoutPrefixCommand(content)
	}
}

func (s *service) withPrefixCommand(cmd []string) DokkoiCmd {
	switch {
	case len(cmd) >= 2 && cmd[1] == "help":
		return &helpCmd{
			target: strings.Join(cmd[2:], " "),
		}
	case len(cmd) >= 3 && cmd[1] == "echo":
		return &echoCmd{
			message: strings.Join(cmd[2:], " "),
		}
	case len(cmd) >= 3 && cmd[1] == "image":
		return &imageCmd{
			customSearchRepo: s.customSearchRepo,
			query:            strings.Join(cmd[2:], " "),
		}
	case len(cmd) >= 3 && (cmd[1] == "lgtm" || cmd[1] == "LGTM"):
		return &lgtmCmd{
			customSearchRepo: s.customSearchRepo,
			query:            strings.Join(cmd[2:], " "),
		}
	case len(cmd) >= 3 && cmd[1] == "score":
		user := strings.Join(cmd[2:], "")
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      user,
			operator:  noOperator,
		}
	case len(cmd) >= 3 && cmd[1] == "select":
		joined := strings.Join(cmd[2:], "")
		elements := strings.Split(joined, ",")
		return &selectCmd{
			elements: elements,
		}
	default:
		return nil
	}
}

func (s *service) withoutPrefixCommand(content string) DokkoiCmd {
	switch {
	case strings.HasSuffix(content, incrOperator):
		user := strings.Replace(content[:len(content)-2], " ", "", -1)
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      user,
			operator:  incrOperator,
		}
	case strings.HasSuffix(content, decrOperator):
		user := strings.Replace(content[:len(content)-2], " ", "", -1)
		return &scoreCmd{
			scoreRepo: s.scoreRepo,
			user:      user,
			operator:  decrOperator,
		}
	default:
		return nil
	}
}
