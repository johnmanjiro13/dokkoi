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
		return err
	}
	_, err = s.ChannelMessageSend(channelID, url)
	return err
}

func (c *imageCmd) searchImage() (string, error) {
	image, err := c.customSearchRepo.SearchImage(c.query)
	if err != nil {
		if pkgerrors.Is(err, ErrImageNotFound) {
			return "image was not found", nil
		}
		return "", pkgerrors.Wrap(err, "image search failed")
	}
	return image.Link, nil
}
