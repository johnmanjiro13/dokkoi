package command

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/johnmanjiro13/dokkoi/command/mock_command"
	"github.com/stretchr/testify/assert"
)

func TestLGTMCommand_ExecFile(t *testing.T) {
	tests := map[string]struct {
		query    string
		image    io.Reader
		err      error
		expected io.Reader
	}{
		"normal success": {
			query:    "dragapult",
			image:    bytes.NewBufferString("dragapult"),
			err:      nil,
			expected: bytes.NewBufferString("dragapult"),
		},
		"image not found": {
			query:    "not found",
			image:    nil,
			err:      ErrImageNotFound,
			expected: nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCustomSearchRepo := mock_command.NewMockCustomSearchRepository(ctrl)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCustomSearchRepo.EXPECT().LGTM(gomock.Any(), tt.query).Return(tt.image, tt.err)
			cmd := &lgtmCmd{
				customSearchRepo: mockCustomSearchRepo,
				query:            tt.query,
			}
			actual, err := cmd.ExecFile(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLGTMCommand_SendType(t *testing.T) {
	cmd := &lgtmCmd{}
	assert.Equal(t, "File", cmd.SendType())
}
