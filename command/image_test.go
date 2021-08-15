package command

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/customsearch/v1"

	"github.com/johnmanjiro13/dokkoi/command/mock_command"
)

func TestImageCmd_Exec(t *testing.T) {
	tests := map[string]struct {
		query    string
		image    *customsearch.Result
		err      error
		expected string
	}{
		"normal success": {
			query:    "dragapult",
			image:    &customsearch.Result{Link: "https://example.com/dragapult.jpeg"},
			err:      nil,
			expected: "https://example.com/dragapult.jpeg",
		},
		"image not found": {
			query:    "not found",
			image:    nil,
			err:      ErrImageNotFound,
			expected: "image was not found",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCustomSearchRepo := mock_command.NewMockCustomSearchRepository(ctrl)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCustomSearchRepo.EXPECT().SearchImage(gomock.Any(), tt.query).Return(tt.image, tt.err)
			cmd := &imageCmd{
				customSearchRepo: mockCustomSearchRepo,
				query:            tt.query,
			}
			actual, err := cmd.Exec(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestImageCmd_searchImage(t *testing.T) {
	tests := map[string]struct {
		query    string
		image    *customsearch.Result
		expected string
	}{
		"normal success": {
			query:    "dragapult",
			image:    &customsearch.Result{Link: "https://example.com/dragapult.jpeg"},
			expected: "https://example.com/dragapult.jpeg",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCustomSearchRepo := mock_command.NewMockCustomSearchRepository(ctrl)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCustomSearchRepo.EXPECT().SearchImage(gomock.Any(), tt.query).Return(tt.image, nil)
			cmd := &imageCmd{
				customSearchRepo: mockCustomSearchRepo,
				query:            tt.query,
			}
			actual, err := cmd.searchImage(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestImageCmd_SendType(t *testing.T) {
	cmd := &imageCmd{}
	assert.Equal(t, "Message", cmd.SendType())
}
