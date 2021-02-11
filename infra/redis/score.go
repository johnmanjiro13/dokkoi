package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	pkgerrors "github.com/pkg/errors"

	"github.com/johnmanjiro13/dokkoi/command"
)

type scoreRepository struct {
	cli          *redis.Client
	lastUsername string
}

func NewScoreRepository(cli *redis.Client) command.ScoreRepository {
	return &scoreRepository{cli: cli}
}

func (r *scoreRepository) Incr(username string) (int64, error) {
	score, err := r.cli.Incr(fmt.Sprintf("score:%s", username)).Result()
	if err != nil {
		return 0, pkgerrors.Wrap(err, "incr score failed")
	}
	r.lastUsername = username
	return score, nil
}

func (r *scoreRepository) Decr(username string) (int64, error) {
	score, err := r.cli.Decr(fmt.Sprintf("score:%s", username)).Result()
	if err != nil {
		return 0, pkgerrors.Wrap(err, "decr score failed")
	}
	r.lastUsername = username
	return score, nil
}

func (r *scoreRepository) UserScore(username string) (int64, error) {
	value, err := r.cli.Get(fmt.Sprintf("score:%s", username)).Result()
	if pkgerrors.Is(err, redis.Nil) {
		// user does not exist
		return 0, nil
	}
	if err != nil {
		return 0, pkgerrors.Wrap(err, "get score failed")
	}
	score, err := strconv.Atoi(value)
	if err != nil {
		return 0, pkgerrors.Wrap(err, "invalid score value")
	}
	return int64(score), nil
}

func (r *scoreRepository) LastUsername() string {
	return r.lastUsername
}
