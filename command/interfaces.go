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
	LastUser() string
	Incr(user string) int
	Decr(user string) int
	UserScore(user string) int
}
