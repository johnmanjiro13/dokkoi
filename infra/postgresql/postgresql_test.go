package postgresql

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	db, err := OpenDB()
	if err != nil {
		log.Fatalf("db opening failed. err: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("db closing failed. err: %v", err)
		}
	}()

	// run tests
	code := m.Run()
	os.Exit(code)
}
