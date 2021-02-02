package command

type scoreRepository struct {
	scores   map[string]int
	lastUser string
}

type ScoreRepository interface {
	Incr(user string) int
	Decr(user string) int
}

func NewScoreRepository(scores map[string]int) *scoreRepository {
	return &scoreRepository{
		scores: scores,
	}
}

func (r *scoreRepository) Incr(user string) int {
	var score int
	if user == "" {
		score = r.scores[r.lastUser]
	} else {
		score = r.scores[user]
	}
	score++
	r.scores[user] = score
	return score
}

func (r *scoreRepository) Decr(user string) int {
	var score int
	if user == "" {
		score = r.scores[r.lastUser]
	} else {
		score = r.scores[user]
	}
	score--
	r.scores[user] = score
	return score
}
