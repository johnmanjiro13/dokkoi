package command

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/johnmanjiro13/dokkoi/google"
)

type imageCmd struct {
	customSearchRepo google.CustomSearchRepository
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
	result, err := c.customSearchRepo.SearchImage(c.query)
	if err != nil {
		return "", err
	}
	if result.Items == nil {
		return "", errors.New("image not found")
	}
	item := result.Items[0]
	return item.Link, nil
}
