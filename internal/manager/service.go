package manager

import (
	"fmt"

	wayToDb "github.com/btussupb/vakio/internal/storage"
)

type service struct {
	wayToDb wayToDb.Storage
}

// func NewUserService(wayToDb wayToDb.Storage) *service {
// 	return &service{wayToDb: wayToDb}
// }

func (s *service) PostUser(userInput User) error {
	fmt.Println("manager")
	return s.wayToDb.PostUser(userInput.toStorageModel())
}
