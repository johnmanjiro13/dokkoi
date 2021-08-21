package command

import (
	"context"

	pkgerrors "github.com/pkg/errors"
)

type imageCmd struct {
	noExecFileCmd
	customSearchRepo CustomSearchRepository
	query            string
}

func (c *imageCmd) ExecString(ctx context.Context) (string, error) {
	url, err := c.searchImage(ctx)
	if err != nil {
		if pkgerrors.Is(err, ErrImageNotFound) {
			return "image was not found", nil
		}
		return "", err
	}
	return url, nil
}

func (c *imageCmd) searchImage(ctx context.Context) (string, error) {
	image, err := c.customSearchRepo.SearchImage(ctx, c.query)
	if err != nil {
		return "", pkgerrors.Wrap(err, "image search failed")
	}
	return image.Link, nil
}

func (c *imageCmd) SendType() string {
	return "Message"
}
