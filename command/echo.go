package command

import "context"

type echoCmd struct {
	message string
}

func (c *echoCmd) Exec(ctx context.Context) (string, error) {
	return c.message, nil
}

func (c *echoCmd) SendType() string {
	return "Message"
}
