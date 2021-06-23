package command

import (
	"context"
	"math/rand"
)

type selectCmd struct {
	elements []string
}

func (c *selectCmd) Exec(ctx context.Context) (string, error) {
	result := c.elements[rand.Intn(len(c.elements))]
	return result, nil
}
