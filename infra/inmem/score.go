package inmem

import "github.com/johnmanjiro13/dokkoi/command"

type scoreRepository struct {
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

func (r *scoreRepository) Incr(username string) (int64, error) {
	score := r.scores[username]
	score++
	r.scores[username] = score
	r.lastUsername = username
	return score, nil
}

func (r *scoreRepository) Decr(user string) (int64, error) {
	score := r.scores[user]
	score--
	r.scores[user] = score
	r.lastUsername = user
	return score, nil
}

func (r *scoreRepository) UserScore(user string) (int64, error) {
	return r.scores[user], nil
}
