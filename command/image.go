package command

import (
	"errors"

	"github.com/bwmarrin/discordgo"
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
	result, err := c.customSearchRepo.SearchImage(c.query)
	if err != nil {
		return "", err
	}
	if len(result.Items) <= 0 {
		return "", errors.New("image not found")
	}
	item := result.Items[0]
	return item.Link, nil
}
