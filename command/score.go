package command

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

const (
	IncrOperator = "++"
	DecrOperator = "--"
)

type scoreCmd struct {
	scoreRepo ScoreRepository
	user      string
	operator  string
}

func (c *scoreCmd) SendMessage(s *discordgo.Session, channelID string) error {
	var score int
	if c.operator == IncrOperator {
		score = c.scoreRepo.Incr(c.user)
	} else if c.operator == DecrOperator {
		score = c.scoreRepo.Decr(c.user)
	}
	_, err := s.ChannelMessageSend(channelID, strconv.Itoa(score))
	return err
}

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
