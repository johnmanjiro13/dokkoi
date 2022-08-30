package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"

	"github.com/johnmanjiro13/dokkoi/pkg/command"
)

type scoreRepository struct {
	cli          *redis.Client
	lastUsername string
}

func NewScoreRepository(cli *redis.Client) command.ScoreRepository {
	return &scoreRepository{cli: cli}
}

func (r *scoreRepository) Incr(ctx context.Context, username string) (int64, error) {
	score, err := r.cli.Incr(ctx, fmt.Sprintf("score:%s", username)).Result()
	if err != nil {
		return 0, pkgerrors.Wrap(err, "incr score failed")
	}
	r.lastUsername = username
	return score, nil
}

func (r *scoreRepository) Decr(ctx context.Context, username string) (int64, error) {
	score, err := r.cli.Decr(ctx, fmt.Sprintf("score:%s", username)).Result()
	if err != nil {
		return 0, pkgerrors.Wrap(err, "decr score failed")
	}
	r.lastUsername = username
	return score, nil
}

func (r *scoreRepository) UserScore(ctx context.Context, username string) (int64, error) {
	value, err := r.cli.Get(ctx, fmt.Sprintf("score:%s", username)).Result()
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
