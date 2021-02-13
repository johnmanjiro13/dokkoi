package command

import (
	"errors"

	"google.golang.org/api/customsearch/v1"
)

var ErrImageNotFound = errors.New("image was not found.")

type DokkoiCmd interface {
	Exec() (string, error)
}

type Service interface {
	GetCommand(content string) DokkoiCmd
}

type CustomSearchRepository interface {
	SearchImage(query string) (*customsearch.Result, error)
}

type ScoreRepository interface {
	LastUsername() string
	Incr(username string) (int64, error)
	Decr(username string) (int64, error)
	UserScore(username string) (int64, error)
}
