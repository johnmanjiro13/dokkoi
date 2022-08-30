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
		"help": {
			content:  "dokkoi help",
			expected: &helpCmd{},
		},
		"help and other": {
			content:  "dokkoi help image",
			expected: &helpCmd{target: "image"},
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
		"lgtm": {
			content: "dokkoi lgtm mt-03",
			expected: &lgtmCmd{
				customSearchRepo: nil,
				query:            "mt-03",
			},
		},
		"LGTM": {
			content: "dokkoi LGTM yzf-r25",
			expected: &lgtmCmd{
				customSearchRepo: nil,
				query:            "yzf-r25",
			},
		},
		"score incr": {
			content: "John Doe++",
			expected: &scoreCmd{
				scoreRepo: nil,
				user:      "JohnDoe",
				operator:  incrOperator,
			},
		},
		"score decr": {
			content: "Jane Doe --",
			expected: &scoreCmd{
				scoreRepo: nil,
				user:      "JaneDoe",
				operator:  decrOperator,
			},
		},
		"score user": {
			content: "dokkoi score john",
			expected: &scoreCmd{
				scoreRepo: nil,
				user:      "john",
				operator:  noOperator,
			},
		},
		"select": {
			content: "dokkoi select john,man,jiro",
			expected: &selectCmd{
				elements: []string{"john", "man", "jiro"},
			},
		},
		"select with whitespaces": {
			content: "dokkoi select john, man, jiro",
			expected: &selectCmd{
				elements: []string{"john", "man", "jiro"},
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
		"include full-width whitespace": {
			content: "dokkoi　image　yamaha　sr400",
			expected: &imageCmd{
				customSearchRepo: nil,
				query:            "yamaha sr400",
			},
		},
		"without dokkoi prefix": {
			content:  "bot image yamaha",
			expected: nil,
		},
	}

	svc := service{
		customSearchRepo: nil,
		scoreRepo:        nil,
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := svc.GetCommand(tt.content)
			assert.Equal(t, tt.expected, cmd)
		})
	}
}
