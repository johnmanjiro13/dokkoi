package command

import "github.com/bwmarrin/discordgo"

type echoCmd struct {
	message string
}

func (c *echoCmd) SendMessage(s *discordgo.Session, channelID string) error {
	_, err := s.ChannelMessageSend(channelID, c.message)
	return err
}
