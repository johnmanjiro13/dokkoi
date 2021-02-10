package postgresql

import (
	"database/sql"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	db, err := OpenDB()
	if err != nil {
		log.Fatalf("db opening failed. err: %v", err)
	}
	defer db.Close()

	setup(db)
	// run tests
	m.Run()
	teardown(db)
}

func setup(db *sql.DB) {
	const createUsers = `
	CREATE TABLE IF NOT EXISTS users (
		id serial not null,
		name varchar(255) unique not null,
		score integer not null default 0,
		primary key (id)
	);`

	if _, err := db.Exec(createUsers); err != nil {
		log.Fatal(err)
	}
}

func teardown(db *sql.DB) {
	if _, err := db.Exec(`TRUNCATE TABLE users RESTART IDENTITY;`); err != nil {
		log.Fatal(err)
	}
}
