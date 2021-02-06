package command

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/customsearch/v1"

	mock_command "github.com/johnmanjiro13/dokkoi/command/mock_google"
)

func TestImageCmd_searchImage(t *testing.T) {
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
		"not found": {
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
			mockCustomSearchRepo.EXPECT().SearchImage(tt.query).Return(tt.image, tt.err)
			cmd := &imageCmd{
				customSearchRepo: mockCustomSearchRepo,
				query:            tt.query,
			}
			actual, err := cmd.searchImage()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
