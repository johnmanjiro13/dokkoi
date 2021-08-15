package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLGTMCommand_SendType(t *testing.T) {
	cmd := &lgtmCmd{}
	assert.Equal(t, "File", cmd.SendType())
}
