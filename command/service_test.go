package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetCommand(t *testing.T) {
	tests := map[string]struct {
		content  string
		expected DokkoiCmd
	}{
		"other than dokkoi": {
			content:  "work",
			expected: nil,
		},
		"echo": {
			content:  "dokkoi echo hoge",
			expected: &echoCmd{message: "hoge"},
		},
		"image": {
			content: "dokkoi image z900rs",
			expected: &imageCmd{
				customSearchRepo: nil,
				query:            "z900rs",
			},
		},
		"dokkoi only": {
			content:  "dokkoi",
			expected: nil,
		},
		"multi target": {
			content: "dokkoi image yamaha sr400",
			expected: &imageCmd{
				customSearchRepo: nil,
				query:            "yamaha sr400",
			},
		},
	}

	svc := service{
		customSearchRepo: nil,
	}
	for desc, tt := range tests {
		t.Run(desc, func(t *testing.T) {
			cmd := svc.GetCommand(tt.content)
			assert.Equal(t, tt.expected, cmd)
		})
	}
}
