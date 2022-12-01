package inmem

import (
	"context"
	"sync"

	"github.com/johnmanjiro13/dokkoi/pkg/command"
)

type scoreRepository struct {
	mu           sync.Mutex
	scores       map[string]int64
	lastUsername string
}

func NewScoreRepository(scores map[string]int64) command.ScoreRepository {
	return &scoreRepository{
		scores: scores,
	}
}

func (r *scoreRepository) LastUsername() string {
	return r.lastUsername
}

func (r *scoreRepository) Incr(ctx context.Context, username string) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	score := r.scores[username]
	score++
	r.scores[username] = score
	r.lastUsername = username
	return score, nil
}

func (r *scoreRepository) Decr(ctx context.Context, user string) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	score := r.scores[user]
	score--
	r.scores[user] = score
	r.lastUsername = user
	return score, nil
}

func (r *scoreRepository) UserScore(ctx context.Context, user string) (int64, error) {
	return r.scores[user], nil
}
