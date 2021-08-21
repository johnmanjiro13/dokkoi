package command

import (
	"context"
	"errors"
	"io"
)

type lgtmCmd struct {
	noExecStringCmd
	customSearchRepo CustomSearchRepository
	query            string
}

func (c *lgtmCmd) ExecFile(ctx context.Context) (io.Reader, error) {
	image, err := c.customSearchRepo.LGTM(ctx, c.query)
	if err != nil {
		if errors.Is(err, ErrImageNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return image, nil
}

func (c *lgtmCmd) SendType() string {
	return "File"
}
