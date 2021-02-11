package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreRepository_Incr(t *testing.T) {
	scores := map[string]int64{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int64
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
			actual, err := repo.Incr(tt.user)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUsername())
		})
	}
}

func TestScoreRepository_Decr(t *testing.T) {
	scores := map[string]int64{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int64
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
			actual, err := repo.Decr(tt.user)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUsername())
		})
	}
}

func TestScoreRepository_LastUser(t *testing.T) {
	scores := map[string]int64{}
	repo := NewScoreRepository(scores)
	repo.Incr("johnman")
	assert.Equal(t, "johnman", repo.LastUsername())
	repo.Decr("god")
	assert.Equal(t, "god", repo.LastUsername())
}

func TestScoreRepository_UserScore(t *testing.T) {
	scores := map[string]int64{
		"johnman": 1,
		"625":     2,
	}
	repo := NewScoreRepository(scores)

	tests := map[string]struct {
		user     string
		expected int64
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
			actual, err := repo.UserScore(tt.user)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
