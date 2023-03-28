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

func NewStorage(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) PostUser(userInput User) error {
	fmt.Println(userInput.Name, userInput.Number, time.Now())

	s.init()

	if _, err := s.db.Exec("INSERT INTO users (Name, Number, CreatedTime) VALUES ($1, $2, $3)", userInput.Name, userInput.Number, time.Now()); err != nil {
		fmt.Println("Exec")
	}

	user, err := s.getUser()
	if err != nil {
		fmt.Println("getUser")
	}
	fmt.Println(user)

	return nil
}

func (s *storage) getUser() (*User, error) {
	row, err := s.db.Query("SELECT * FROM таблица ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println("Query")
	}
	defer row.Close()

	user := &User{}

	if err := row.Scan(&user.Name, &user.Number, &user.CreatedTime); err != nil {
		fmt.Println("Scan")
	}
	return user, nil
}

func (s *storage) init() error {
	q := ("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, number TEXT, Time time.Time)")

	_, err := s.db.Exec(q)
	if err != nil {
		fmt.Println("init bd")
	}
	return nil
}
