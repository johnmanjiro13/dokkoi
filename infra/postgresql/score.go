package postgresql

import (
	"database/sql"

	pkgerrors "github.com/pkg/errors"
)

const (
	incr = "incr"
	decr = "decr"
)

type User struct {
	Name  string
	Score int
}

type scoreRepository struct {
	db       *sql.DB
	lastUser *User
}

func NewScoreRepository(db *sql.DB) *scoreRepository {
	return &scoreRepository{db: db}
}

func (r *scoreRepository) LastUsername() string {
	return r.lastUser.Name
}

func (r *scoreRepository) Incr(username string) (int, error) {
	user, err := r.findUser(username)
	if pkgerrors.Is(err, sql.ErrNoRows) {
		// user does not exist
		user, err = r.insertUser(username, incr)
		if err != nil {
			return 0, err
		}
		r.lastUser = user
		return 1, nil
	}
	if err != nil {
		return 0, err
	}
	// increment user's score
	user.Score++
	err = r.updateScore(user.Name, user.Score)
	if err != nil {
		return 0, err
	}
	r.lastUser = user
	return user.Score, nil
}

func (r *scoreRepository) findUser(username string) (*User, error) {
	var user User
	row := r.db.QueryRow(`SELECT name, score FROM users WHERE name = $1`, username)

	if err := row.Scan(&user.Name, &user.Score); err != nil {
		return nil, pkgerrors.Wrap(err, "find user failed")
	}
	return &user, nil
}

func (r *scoreRepository) insertUser(username, operator string) (*User, error) {
	var score int
	if operator == incr {
		score = 1
	} else if operator == decr {
		score = -1
	}
	q, err := r.db.Prepare(`INSERT INTO users (name, score) VALUES ($1, $2)`)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "insert prepare failed")
	}
	_, err = q.Exec(username, score)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "insert user failed")
	}
	return &User{Name: username, Score: score}, nil
}

func (r *scoreRepository) updateScore(username string, score int) error {
	q, err := r.db.Prepare(`UPDATE users SET score = $1 WHERE name = $2`)
	if err != nil {
		return pkgerrors.Wrap(err, "update prepare failed")
	}
	_, err = q.Exec(score, username)
	if err != nil {
		return pkgerrors.Wrap(err, "update score failed")
	}
	return nil
}
