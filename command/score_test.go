package command

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/johnmanjiro13/dokkoi/command/mock_command"
)

func TestScoreCmd_Exec(t *testing.T) {
	tests := map[string]struct {
		user     string
		lastUser string
		operator string
		score    int64
		expected string
	}{
		"increment result 1": {
			user:     "johnman",
			lastUser: "kairyu",
			operator: incrOperator,
			score:    1,
			expected: "johnman has 1 point",
		},
		"decrement result 2": {
			user:     "johnman",
			lastUser: "kairyu",
			operator: decrOperator,
			score:    2,
			expected: "johnman has 2 points",
		},
		"no operator": {
			user:     "johnman",
			lastUser: "kairyu",
			operator: noOperator,
			score:    2,
			expected: "johnman has 2 points",
		},
		"last user": {
			user:     "",
			lastUser: "kairyu",
			operator: incrOperator,
			score:    2,
			expected: "kairyu has 2 points",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockScoreRepo := mock_command.NewMockScoreRepository(ctrl)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := &scoreCmd{
				scoreRepo: mockScoreRepo,
				user:      tt.user,
				operator:  tt.operator,
			}
			switch tt.operator {
			case incrOperator:
				if tt.user != "" {
					mockScoreRepo.EXPECT().Incr(gomock.Any(), tt.user).Return(tt.score, nil)
				} else {
					mockScoreRepo.EXPECT().LastUsername().Return(tt.lastUser)
					mockScoreRepo.EXPECT().Incr(gomock.Any(), tt.lastUser).Return(tt.score, nil)
				}
			case decrOperator:
				mockScoreRepo.EXPECT().Decr(gomock.Any(), tt.user).Return(tt.score, nil)
			case noOperator:
				mockScoreRepo.EXPECT().UserScore(gomock.Any(), tt.user).Return(tt.score, nil)
			}
			actual, err := cmd.Exec(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestScoreCmd_CalcScore(t *testing.T) {
	tests := map[string]struct {
		user     string
		operator string
	}{
		"incr": {
			user:     "johnman",
			operator: incrOperator,
		},
		"decr": {
			user:     "kairyu",
			operator: decrOperator,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockScoreRepo := mock_command.NewMockScoreRepository(ctrl)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := &scoreCmd{
				scoreRepo: mockScoreRepo,
				user:      tt.user,
				operator:  tt.operator,
			}
			if tt.operator == incrOperator {
				mockScoreRepo.EXPECT().Incr(gomock.Any(), tt.user)
			} else if tt.operator == decrOperator {
				mockScoreRepo.EXPECT().Decr(gomock.Any(), tt.user)
			}
			cmd.calcScore(context.Background())
		})
	}
}

func TestScoreCmd_SendType(t *testing.T) {
	cmd := &scoreCmd{}
	assert.Equal(t, "Message", cmd.SendType())
}
