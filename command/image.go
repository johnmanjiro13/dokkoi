package command

import (
	pkgerrors "github.com/pkg/errors"
)

type imageCmd struct {
	customSearchRepo CustomSearchRepository
	query            string
}

func (c *imageCmd) Exec() (string, error) {
	url, err := c.searchImage()
	if err != nil {
		if pkgerrors.Is(err, ErrImageNotFound) {
			return "image was not found", nil
		}
		return "", err
	}
	return url, nil
}

func (c *imageCmd) searchImage() (string, error) {
	image, err := c.customSearchRepo.SearchImage(c.query)
	if err != nil {
		return "", pkgerrors.Wrap(err, "image search failed")
	}
	return image.Link, nil
}
