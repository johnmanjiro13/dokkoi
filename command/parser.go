package command

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var (
	commandRegExp = regexp.MustCompile(`^dokkoi\s(.+)\s(.+)$`)
)

type DokkoiCmd interface {
	SendMessage(s *discordgo.Session, channelID string) error
}

func Parse(content string) DokkoiCmd {
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
