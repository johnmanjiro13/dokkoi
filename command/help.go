package command

import (
	"context"
	"sort"
	"strings"
)

type helpCmd struct {
	target string
}

var descriptions = map[string]string{
	"help":   "dokkoi help - Displays all of the help commands that this bot knows about.\ndokkoi help <query> - Displays all help commands that match <query>.",
	"echo":   "dokkoi echo <text> - Reply back with <text>",
	"image":  "dokkoi image <query> - Queries Google Images for <query> and returns a top result.",
	"score":  "dokkoi score <name> - Display the score for a name.",
	"select": "dokkoi select <element1>,<element2>,... - Choose one of the elements in your list randomly.",
	"++":     "<name>++ - Increment score for a name",
	"--":     "<name>-- - Decrement score for a name",
}

func (c *helpCmd) Exec(ctx context.Context) (desc string, err error) {
	if c.target == "" {
		desc = strings.Join(values(descriptions), "\n")
	} else {
		desc = descriptions[c.target]
	}
	return
}

func values(m map[string]string) (s []string) {
	for _, v := range m {
		s = append(s, v)
	}
	sort.Strings(s)
	return
}
