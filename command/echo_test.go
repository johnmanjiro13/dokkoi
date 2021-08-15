package command

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoCmd_ExecString(t *testing.T) {
	expected := "test"
	cmd := &echoCmd{message: expected}
	actual, err := cmd.ExecString(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, actual)
}

func TestEchoCmd_SendType(t *testing.T) {
	cmd := &echoCmd{}
	assert.Equal(t, "Message", cmd.SendType())
}
