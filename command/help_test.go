package command

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelpCmd_Exec(t *testing.T) {
	fullDesc := `<name>++ - Increment score for a name
<name>-- - Decrement score for a name
dokkoi echo <text> - Reply back with <text>
dokkoi help - Displays all of the help commands that this bot knows about.
dokkoi help <query> - Displays all help commands that match <query>.
dokkoi image <query> - Queries Google Images for <query> and returns a top result.
dokkoi score <name> - Display the score for a name.
dokkoi select <element1>,<element2>,... - Choose one of the elements in your list randomly.`

	tests := map[string]struct {
		target   string
		expected string
	}{
		"full description": {
			target:   "",
			expected: fullDesc,
		},
		"image description": {
			target:   "image",
			expected: "dokkoi image <query> - Queries Google Images for <query> and returns a top result.",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := &helpCmd{target: tt.target}
			actual, err := cmd.Exec(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestValues(t *testing.T) {
	fixture := map[string]string{
		"++": "<name>++ - Increment score for a name",
		"--": "<name>-- - Decrement score for a name",
	}
	expected := []string{
		"<name>++ - Increment score for a name",
		"<name>-- - Decrement score for a name",
	}
	assert.Equal(t, expected, values(fixture))
}
