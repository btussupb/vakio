package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type storage struct {
	db *sql.DB
}

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

func (s *storage) PostUser(userInput User) error {
	// id := GenerateId()

	// db := s.openDb()

	// statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, number TEXT, Time time.Time)")
	// if err != nil {
	// 	//
	// }
	fmt.Println(userInput.Name, userInput.Number, time.Now())

	s.init()

	if _, err := s.db.Exec("INSERT INTO users (Name, Number, CreatedTime) VALUES ($1, $2, $3)", userInput.Name, userInput.Number, time.Now()); err != nil {
		//
	}

	user, err := s.getUser()
	if err != nil {
		//
	}
	fmt.Println(user)

	return nil
}

// func (s *storage) openDb() *sql.DB {
// 	db, err := sql.Open("sqlite3", "store.db")
// 	if err != nil {
// 		//
// 	}

// 	defer db.Close()

// 	return db
// }

func (s *storage) getUser() (*User, error) {
	row, err := s.db.Query("SELECT * FROM таблица ORDER BY id DESC LIMIT 1")
	if err != nil {
		//
	}
	defer row.Close()

	user := &User{}

	if err := row.Scan(&user.Name, &user.Number, &user.CreatedTime); err != nil {
		//
	}
	return user, nil
}

func (s *storage) init() error {
	q := ("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, number TEXT, Time time.Time)")

	_, err := s.db.Exec(q)
	if err != nil {
		//
	}
	return nil
}
