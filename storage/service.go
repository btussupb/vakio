package db

import (
	"database/sql"
	"time"
)

// type storage struct {
// 	db *sql.DB
// }

// func NewStorage(db *sql.DB) *storage {
// 	return &storage{
// 		db: db,
// 	}
// }

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

func Exec(userInput User) error {
	id := GenerateId()

	db := openDb()

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, number TEXT, Time time.Time)")
	if err != nil {
		//
	}
	statement.Exec()

	if _, err := db.Prepare("INSERT INTO users (id, Name, Number, Time) VALUES (?, ?, ?, ?)"); err != nil {
		//
	}
	statement.Exec(id, userInput.Name, userInput.Number, time.Now())
	return nil
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		//
	}

	defer db.Close()

	return db
}

func GenerateId() int {
	return 0
}
