package command

import (
	"context"
	"math/rand"
)

type selectCmd struct {
	noExecFileCmd
	elements []string
}

func (c *selectCmd) ExecString(ctx context.Context) (string, error) {
	result := c.elements[rand.Intn(len(c.elements))]
	return result, nil
}

func (c *selectCmd) SendType() string {
	return "Message"
}
