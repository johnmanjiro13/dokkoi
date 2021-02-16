package command

import "math/rand"

type selectCmd struct {
	elements []string
}

func (c *selectCmd) Exec() (string, error) {
	result := c.elements[rand.Intn(len(c.elements))]
	return result, nil
}
