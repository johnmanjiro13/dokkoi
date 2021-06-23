package command

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoCmd_Exec(t *testing.T) {
	expected := "test"
	cmd := &echoCmd{message: expected}
	actual, err := cmd.Exec(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, actual)
}
