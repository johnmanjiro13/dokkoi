package postgresql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreRepository_Incr(t *testing.T) {
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	q, err := db.Prepare(`INSERT INTO users (name, score) VALUES ($1, $2);`)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := q.Exec("johnman", 1); err != nil {
		t.Fatal(err)
	}

	repo := NewScoreRepository(db)

	tests := map[string]struct {
		user string

		expected int
	}{
		"User already exists": {
			user:     "johnman",
			expected: 2,
		},
		"User not exists": {
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
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	q, err := db.Prepare(`INSERT INTO users (name, score) VALUES ($1, $2);`)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := q.Exec("johnman", 1); err != nil {
		t.Fatal(err)
	}

	repo := NewScoreRepository(db)

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
			actual, err := repo.Decr(tt.user)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.user, repo.LastUsername())
		})
	}
}
