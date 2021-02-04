package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreRepository_Incr(t *testing.T) {
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

	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.Incr(tt.user)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUser())
		})
	}
}

func TestScoreRepository_Decr(t *testing.T) {
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

	scores := map[string]int{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := repo.Decr(tt.user)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUser())
		})
	}
}
