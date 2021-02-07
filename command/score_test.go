package command

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_command "github.com/johnmanjiro13/dokkoi/command/mock_score"
)

func TestScoreCmd_CalcScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_command.NewMockScoreRepository(ctrl)

	tests := map[string]struct {
		user     string
		operator string
	}{
		"incr": {
			user:     "johnman",
			operator: incrOperator,
		},
		"user not exists": {
			user:     "kairyu",
			operator: decrOperator,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := &scoreCmd{
				scoreRepo: repo,
				user:      tt.user,
				operator:  incrOperator,
			}
			if cmd.operator == incrOperator {
				repo.EXPECT().Incr(cmd.user)
			} else if cmd.operator == decrOperator {
				repo.EXPECT().Decr(cmd.user)
			}
			cmd.calcScore()
		})
	}
}
