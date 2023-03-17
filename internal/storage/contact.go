package db

import "time"

type (
	User struct {
		Name        string
		Number      string
		CreatedTime time.Time
	}
	Storage interface {
		PostUser(user User) error
	}
)
