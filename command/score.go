package command

import (
	"context"
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

func (c *scoreCmd) Exec(ctx context.Context) (string, error) {
	if c.user == "" {
		c.user = c.scoreRepo.LastUsername()
	}

	var (
		score int64
		err   error
	)
	if c.operator == incrOperator || c.operator == decrOperator {
		score, err = c.calcScore(ctx)
		if err != nil {
			return "", err
		}
	} else if c.operator == noOperator {
		score, err = c.scoreRepo.UserScore(ctx, c.user)
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

func (c *scoreCmd) calcScore(ctx context.Context) (int64, error) {
	if c.operator == incrOperator {
		return c.scoreRepo.Incr(ctx, c.user)
	} else {
		return c.scoreRepo.Decr(ctx, c.user)
	}
}

func (c *scoreCmd) SendType() string {
	return "Message"
}
