package command

import (
	"github.com/bwmarrin/discordgo"
	pkgerrors "github.com/pkg/errors"
)

type imageCmd struct {
	customSearchRepo CustomSearchRepository
	query            string
}

func (c *imageCmd) SendMessage(s *discordgo.Session, channelID string) error {
	url, err := c.searchImage()
	if err != nil {
		if pkgerrors.Is(err, ErrImageNotFound) {
			return c.sendMessage(s, channelID, "image was not found")
		}
		return err
	}
	return c.sendMessage(s, channelID, url)
}

func (c *imageCmd) searchImage() (string, error) {
	image, err := c.customSearchRepo.SearchImage(c.query)
	if err != nil {
		return "", pkgerrors.Wrap(err, "image search failed")
	}
	return image.Link, nil
}

func (c *imageCmd) sendMessage(s *discordgo.Session, channelID, message string) error {
	_, err := s.ChannelMessageSend(channelID, message)
	return err
}
