package inmem

import "github.com/johnmanjiro13/dokkoi/command"

type scoreRepository struct {
	scores   map[string]int
	lastUser string
}

func NewScoreRepository(scores map[string]int) command.ScoreRepository {
	return &scoreRepository{
		scores: scores,
	}
}

func (r *scoreRepository) LastUser() string {
	return r.lastUser
}

func (r *scoreRepository) Incr(user string) int {
	score := r.scores[user]
	score++
	r.scores[user] = score
	r.lastUser = user
	return score
}

func (r *scoreRepository) Decr(user string) int {
	score := r.scores[user]
	score--
	r.scores[user] = score
	r.lastUser = user
	return score
}

func (r *scoreRepository) UserScore(user string) int {
	return r.scores[user]
}
