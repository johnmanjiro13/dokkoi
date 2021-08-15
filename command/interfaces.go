package command

import (
	"context"
	"errors"
	"io"

	"google.golang.org/api/customsearch/v1"
)

var ErrImageNotFound = errors.New("image was not found.")

type DokkoiCmd interface {
	ExecString(ctx context.Context) (string, error)
	ExecFile(ctx context.Context) (io.Reader, error)
	SendType() string
}

type noExecStringCmd struct{}

func (c *noExecStringCmd) ExecString(ctx context.Context) (string, error) {
	panic("not implemented")
}

type noExecFileCmd struct{}

func (c *noExecFileCmd) ExecFile(ctx context.Context) (io.Reader, error) {
	panic("not implemented")
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

type CustomSearchRepository interface {
	SearchImage(ctx context.Context, query string) (*customsearch.Result, error)
	LGTM(ctx context.Context, query string) (io.Reader, error)
}

type ScoreRepository interface {
	LastUsername() string
	Incr(ctx context.Context, username string) (int64, error)
	Decr(ctx context.Context, username string) (int64, error)
	UserScore(ctx context.Context, username string) (int64, error)
}
