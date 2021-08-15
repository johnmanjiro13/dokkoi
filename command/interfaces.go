package command

import (
	"context"
	"errors"

	"google.golang.org/api/customsearch/v1"
)

var ErrImageNotFound = errors.New("image was not found.")

type DokkoiCmd interface {
	Exec(ctx context.Context) (string, error)
	SendType() string
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

type CustomSearchRepository interface {
	SearchImage(ctx context.Context, query string) (*customsearch.Result, error)
}

type ScoreRepository interface {
	LastUsername() string
	Incr(ctx context.Context, username string) (int64, error)
	Decr(ctx context.Context, username string) (int64, error)
	UserScore(ctx context.Context, username string) (int64, error)
}
