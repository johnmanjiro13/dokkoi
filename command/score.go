package command

import (
	"fmt"

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
	user := c.user
	if user == "" {
		user = c.scoreRepo.LastUser()
	}

	var score int
	if c.operator == IncrOperator {
		score = c.scoreRepo.Incr(user)
	} else if c.operator == DecrOperator {
		score = c.scoreRepo.Decr(user)
	}

	var message string
	if score == 1 || score == -1 {
		message = fmt.Sprintf("%s has %d point", user, score)
	} else {
		message = fmt.Sprintf("%s has %d points", user, score)
	}

	_, err := s.ChannelMessageSend(channelID, message)
	return err
}

type scoreRepository struct {
	scores   map[string]int
	lastUser string
}

type ScoreRepository interface {
	LastUser() string
	Incr(user string) int
	Decr(user string) int
}

func NewScoreRepository(scores map[string]int) ScoreRepository {
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
