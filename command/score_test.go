package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreRepository_Incr(t *testing.T) {
	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)
	repo.lastUser = "625"

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
		"not given user name": {
			user:     "",
			expected: 3,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.Incr(tt.user)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestScoreRepository_Decr(t *testing.T) {
	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)
	repo.lastUser = "625"

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
		"not given user name": {
			user:     "",
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.Decr(tt.user)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
