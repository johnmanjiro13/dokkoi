package command

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_command "github.com/johnmanjiro13/dokkoi/command/mock_score"
)

func TestScoreCmd_Exec(t *testing.T) {
	tests := map[string]struct {
		user     string
		lastUser string
		operator string
		score    int
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
					mockScoreRepo.EXPECT().Incr(tt.user).Return(tt.score)
				} else {
					mockScoreRepo.EXPECT().LastUser().Return(tt.lastUser)
					mockScoreRepo.EXPECT().Incr(tt.lastUser).Return(tt.score)
				}
			case decrOperator:
				mockScoreRepo.EXPECT().Decr(tt.user).Return(tt.score)
			case noOperator:
				mockScoreRepo.EXPECT().UserScore(tt.user).Return(tt.score)
			}
			actual, err := cmd.Exec()
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
				mockScoreRepo.EXPECT().Incr(tt.user)
			} else if tt.operator == decrOperator {
				mockScoreRepo.EXPECT().Decr(tt.user)
			}
			cmd.calcScore()
		})
	}
}
