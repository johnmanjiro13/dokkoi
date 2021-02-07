package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type helpCmd struct {
	target string
}

var descriptions = map[string]string{
	"help":  "dokkoi help - Displays all of the help commands that this bot knows about.\ndokkoi help <query> - Displays all help commands that match <query>.",
	"echo":  "dokkoi echo <text> - Reply back with <text>",
	"image": "dokkoi image <query> - Queries Google Images for <query> and returns a top result.",
	"score": "dokkoi score <name> - Display the score for a name.",
	"++":    "<name>++ - Increment score for a name",
	"--":    "<name>-- - Decrement score for a name",
}

func (c *helpCmd) SendMessage(s *discordgo.Session, channelID string) error {
	var desc string
	if c.target == "" {
		desc = strings.Join(values(descriptions), "\n")
	} else {
		desc = descriptions[c.target]
	}
	_, err := s.ChannelMessageSend(channelID, desc)
	return err
}

func values(m map[string]string) (s []string) {
	for _, v := range m {
		s = append(s, v)
	}
	return
}
