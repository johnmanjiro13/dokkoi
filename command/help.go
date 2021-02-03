package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type helpCmd struct {
}

var (
	descriptions = []string{
		"dokkoi help - Displays all of the help commands that this bot knows about.",
		"dokkoi echo <text> - Reply back with <text>",
		"dokkoi echo <query> - Queries Google Images for <query> and returns a top result.",
		"<name>++ - Increment score for a name",
		"<name>-- - Decrement score for a name",
	}
)

func (c *helpCmd) SendMessage(s *discordgo.Session, channelID string) error {
	_, err := s.ChannelMessageSend(channelID, strings.Join(descriptions, "\n"))
	return err
}
