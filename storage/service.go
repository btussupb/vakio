package db

import (
	"database/sql"
	"time"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *storage {
	return &storage{
		db: db,
	}
}

// func (s *storage) open(userInput User) error {
// 	db, err := sql.Open("sqlite3", "store.db")
// 	if err != nil {
// 		//
// 	}
// 	if err = s.Exec(userInput); err != nil {
// 		//
// 	}

// 	defer db.Close()

// 	return nil
// }

func (s *storage) Exec(userInput User) error {
	id := GenerateId()

	if _, err := s.db.Exec("insert into users (id, Name, Number, Time) values ($1, $2, $3, $4)", id, userInput.Name, userInput.Number, time.Now()); err != nil {
		//
	}
	return nil
}

func GenerateId() int {
	return 0
}
