package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreCmd_CalcScore(t *testing.T) {
	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int
	}{
		"user already exists": {
			user:     "johnman",
			expected: 2,
		},
		"user not exists": {
			user:     "kairyu",
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := &scoreCmd{
				scoreRepo: repo,
				user:      tt.user,
				operator:  incrOperator,
			}
			actual := cmd.calcScore()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestScoreRepository_Incr(t *testing.T) {
	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int
	}{
		"user already exists": {
			user:     "johnman",
			expected: 2,
		},
		"user not exists": {
			user:     "kairyu",
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.Incr(tt.user)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUser())
		})
	}
}

func TestScoreRepository_Decr(t *testing.T) {
	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int
	}{
		"user already exists": {
			user:     "johnman",
			expected: 0,
		},
		"user not exists": {
			user:     "kairyu",
			expected: -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.Decr(tt.user)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUser())
		})
	}
}

func TestScoreRepository_LastUser(t *testing.T) {
	scores := map[string]int{}
	repo := NewScoreRepository(scores)
	repo.Incr("johnman")
	assert.Equal(t, "johnman", repo.LastUser())
	repo.Decr("god")
	assert.Equal(t, "god", repo.LastUser())
}

func TestScoreRepository_UserScore(t *testing.T) {
	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int
	}{
		"user already exists": {
			user:     "johnman",
			expected: 1,
		},
		"user not exists": {
			user:     "kairyu",
			expected: 0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.UserScore(tt.user)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
