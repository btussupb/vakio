package db

import "time"

type (
	User struct {
		Id          int
		City        string
		Name        string
		Number      string
		CreatedTime time.Time
	}
	Storage interface {
		PostUser(user User) error
	}
)
