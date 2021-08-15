package command

import (
	"context"
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
		return nil, err
	}
	return image, err
}

func (c *lgtmCmd) SendType() string {
	return "File"
}
