package db

import "time"

type (
	User struct {
		City        string
		Name        string
		Number      string
		CreatedTime time.Time
	}
	Storage interface {
		PostUser(user User) error
	}
)
