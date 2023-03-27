package db

import "time"

type (
	User struct {
		Sity        string
		Name        string
		Number      string
		CreatedTime time.Time
	}
	Storage interface {
		PostUser(user User) error
	}
)
