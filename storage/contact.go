package db

import "time"

type (
	User struct {
		Name   string
		Number string
		Time   time.Time
	}
	Storage interface {
		Exec(user User) error
	}
)
