package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
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
	fmt.Println(userInput.City, userInput.Name, userInput.Number, time.Now())

	s.init()

	if _, err := s.db.Exec("INSERT INTO users (city, name, number, createdTime) VALUES ($1, $2, $3, $4)", userInput.City, userInput.Name, userInput.Number, time.Now()); err != nil {
		fmt.Println("Exec", err)
	}

	user, err := s.getUser()
	if err != nil {
		fmt.Println("getUser", err)
	}
	fmt.Println(user, "end code")

	s.sendMail(user)

	return nil
}

func (s *storage) getUser() (*User, error) {
	row, err := s.db.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println("Query", err)
		return nil, err
	}
	defer row.Close()

	user := &User{}
	for row.Next() {
		if err := row.Scan(&user.Id, &user.City, &user.Name, &user.Number, &user.CreatedTime); err != nil {
			fmt.Println("Scan", err)
			return nil, err
		}
	}
	return user, nil
}

func (s *storage) sendMail(data *User) error {
	auth := smtp.PlainAuth("", "karagandavakio@gmail.com", "ajnbbhiqrjfctfuv", "smtp.gmail.com")

	to := []string{"Neveryun1212@gmail.com"}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	body := string(jsonData)

	msg := []byte(body)

	err = smtp.SendMail("smtp.gmail.com:587", auth, "bjjbaha@mail.ru", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *storage) init() error {
	q := ("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, city TEXT, name TEXT, number TEXT, createdTime DATE)")

	_, err := s.db.Exec(q)
	if err != nil {
		fmt.Println("init bd", err)
	}

	return nil
}
