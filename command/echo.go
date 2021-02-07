package command

type echoCmd struct {
	message string
}

func (c *echoCmd) Exec() (string, error) {
	return c.message, nil
}
