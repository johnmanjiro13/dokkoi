package command

import (
	"fmt"
)

const (
	noOperator   = "none"
	incrOperator = "++"
	decrOperator = "--"
)

type scoreCmd struct {
	scoreRepo ScoreRepository
	user      string
	operator  string
}

func (c *scoreCmd) Exec() (string, error) {
	if c.user == "" {
		c.user = c.scoreRepo.LastUsername()
	}

	var (
		score int64
		err   error
	)
	if c.operator == incrOperator || c.operator == decrOperator {
		score, err = c.calcScore()
		if err != nil {
			return "", err
		}
	} else if c.operator == noOperator {
		score, err = c.scoreRepo.UserScore(c.user)
		if err != nil {
			return "", err
		}
	}

	var message string
	if score == 1 || score == -1 {
		message = fmt.Sprintf("%s has %d point", c.user, score)
	} else {
		message = fmt.Sprintf("%s has %d points", c.user, score)
	}
	return message, nil
}

func (c *scoreCmd) calcScore() (int64, error) {
	if c.operator == incrOperator {
		return c.scoreRepo.Incr(c.user)
	} else {
		return c.scoreRepo.Decr(c.user)
	}
}
