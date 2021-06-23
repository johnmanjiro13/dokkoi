package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScoreRepository_Incr(t *testing.T) {
	cli, err := openTest()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	ctx := context.Background()
	status := cli.Set(ctx, "score:johnman", 1, 5*time.Minute)
	if status.Err() != nil {
		t.Fatal(status.Err())
	}
	defer cli.FlushDB(ctx)
	repo := NewScoreRepository(cli)

	tests := map[string]struct {
		username string
		expected int64
	}{
		"user already exists": {
			username: "johnman",
			expected: 2,
		},
		"user not exists": {
			username: "kairyu",
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := repo.Incr(ctx, tt.username)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.username, repo.LastUsername())
		})
	}
}

func TestScoreRepository_Decr(t *testing.T) {
	cli, err := openTest()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	ctx := context.Background()
	status := cli.Set(ctx, "score:johnman", 1, 5*time.Minute)
	if status.Err() != nil {
		t.Fatal(status.Err())
	}
	defer cli.FlushDB(ctx)
	repo := NewScoreRepository(cli)

	tests := map[string]struct {
		username string
		expected int64
	}{
		"user already exists": {
			username: "johnman",
			expected: 0,
		},
		"user not exists": {
			username: "kairyu",
			expected: -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := repo.Decr(ctx, tt.username)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.username, repo.LastUsername())
		})
	}
}

func TestScoreRepository_LastUser(t *testing.T) {
	ctx := context.Background()
	cli, err := openTest()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	defer cli.FlushDB(ctx)
	repo := NewScoreRepository(cli)
	if _, err := repo.Incr(ctx, "johnman"); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "johnman", repo.LastUsername())
	if _, err := repo.Decr(ctx, "god"); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "god", repo.LastUsername())
}

func TestScoreRepository_UserScore(t *testing.T) {
	cli, err := openTest()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()
	status := cli.Set(ctx, "score:johnman", 1, 5*time.Minute)
	if status.Err() != nil {
		t.Fatal(status.Err())
	}
	defer cli.FlushDB(ctx)
	repo := NewScoreRepository(cli)

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
			actual, err := repo.UserScore(ctx, tt.user)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, "", repo.LastUsername())
		})
	}
}
